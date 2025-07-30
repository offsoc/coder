package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"golang.org/x/xerrors"

	"cdr.dev/slog/sloggers/slogtest"
	"github.com/coder/coder/v2/immortalstream"
)

func TestImmortalStreamHandler(t *testing.T) {
	t.Parallel()

	t.Run("CreateStream", func(t *testing.T) {
		t.Parallel()

		logger := slogtest.Make(t, nil)
		dialer := func(network, address string) (net.Conn, error) {
			return &mockConn{}, nil
		}
		manager := immortalstream.NewManager(logger, dialer)
		defer manager.Close()

		handler := newImmortalStreamHandler(manager)

		req := immortalstream.CreateStreamRequest{TCPPort: 8080}
		body, _ := json.Marshal(req)

		httpReq := httptest.NewRequest("POST", "/api/v0/immortal-stream", bytes.NewReader(body))
		httpReq.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler.createStream(rr, httpReq)

		if rr.Code != http.StatusCreated {
			t.Fatalf("expected status 201, got %d", rr.Code)
		}

		var resp immortalstream.CreateStreamResponse
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		if err != nil {
			t.Fatalf("failed to unmarshal response: %v", err)
		}

		if resp.ID == uuid.Nil {
			t.Fatalf("expected valid stream ID")
		}
		if resp.Name == "" {
			t.Fatalf("expected non-empty stream name")
		}
	})

	t.Run("CreateStreamInvalidJSON", func(t *testing.T) {
		t.Parallel()

		logger := slogtest.Make(t, nil)
		manager := immortalstream.NewManager(logger, nil)
		defer manager.Close()

		handler := newImmortalStreamHandler(manager)

		httpReq := httptest.NewRequest("POST", "/api/v0/immortal-stream", strings.NewReader("invalid json"))
		httpReq.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler.createStream(rr, httpReq)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected status 400, got %d", rr.Code)
		}
	})

	t.Run("CreateStreamConnectionRefused", func(t *testing.T) {
		t.Parallel()

		logger := slogtest.Make(t, nil)
		dialer := func(network, address string) (net.Conn, error) {
			return nil, xerrors.New("connection refused")
		}
		manager := immortalstream.NewManager(logger, dialer)
		defer manager.Close()

		handler := newImmortalStreamHandler(manager)

		req := immortalstream.CreateStreamRequest{TCPPort: 8080}
		body, _ := json.Marshal(req)

		httpReq := httptest.NewRequest("POST", "/api/v0/immortal-stream", bytes.NewReader(body))
		httpReq.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler.createStream(rr, httpReq)

		if rr.Code != http.StatusNotFound {
			t.Fatalf("expected status 404, got %d", rr.Code)
		}
	})

	t.Run("ListStreams", func(t *testing.T) {
		t.Parallel()

		logger := slogtest.Make(t, nil)
		dialer := func(network, address string) (net.Conn, error) {
			return &mockConn{}, nil
		}
		manager := immortalstream.NewManager(logger, dialer)
		defer manager.Close()

		// Create a test stream
		createReq := immortalstream.CreateStreamRequest{TCPPort: 8080}
		_, err := manager.CreateStream(context.Background(), createReq)
		if err != nil {
			t.Fatalf("failed to create test stream: %v", err)
		}

		handler := newImmortalStreamHandler(manager)

		httpReq := httptest.NewRequest("GET", "/api/v0/immortal-stream", nil)
		rr := httptest.NewRecorder()
		handler.listStreams(rr, httpReq)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected status 200, got %d", rr.Code)
		}

		var streams []*immortalstream.Stream
		err = json.Unmarshal(rr.Body.Bytes(), &streams)
		if err != nil {
			t.Fatalf("failed to unmarshal response: %v", err)
		}

		if len(streams) != 1 {
			t.Fatalf("expected 1 stream, got %d", len(streams))
		}
		if streams[0].TCPPort != 8080 {
			t.Fatalf("expected port 8080, got %d", streams[0].TCPPort)
		}
	})

	t.Run("DeleteStream", func(t *testing.T) {
		t.Parallel()

		logger := slogtest.Make(t, nil)
		dialer := func(network, address string) (net.Conn, error) {
			return &mockConn{}, nil
		}
		manager := immortalstream.NewManager(logger, dialer)
		defer manager.Close()

		// Create a test stream
		createReq := immortalstream.CreateStreamRequest{TCPPort: 8080}
		createResp, err := manager.CreateStream(context.Background(), createReq)
		if err != nil {
			t.Fatalf("failed to create test stream: %v", err)
		}

		handler := newImmortalStreamHandler(manager)

		// Create chi router to handle URL parameters
		r := chi.NewRouter()
		r.Delete("/{id}", handler.deleteStream)

		httpReq := httptest.NewRequest("DELETE", "/"+createResp.ID.String(), nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, httpReq)

		if rr.Code != http.StatusNoContent {
			t.Fatalf("expected status 204, got %d", rr.Code)
		}

		// Verify stream was deleted
		streams := manager.ListStreams()
		if len(streams) != 0 {
			t.Fatalf("expected 0 streams after deletion, got %d", len(streams))
		}
	})

	t.Run("DeleteStreamNotFound", func(t *testing.T) {
		t.Parallel()

		logger := slogtest.Make(t, nil)
		manager := immortalstream.NewManager(logger, nil)
		defer manager.Close()

		handler := newImmortalStreamHandler(manager)

		// Create chi router to handle URL parameters
		r := chi.NewRouter()
		r.Delete("/{id}", handler.deleteStream)

		nonExistentID := uuid.New()
		httpReq := httptest.NewRequest("DELETE", "/"+nonExistentID.String(), nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, httpReq)

		if rr.Code != http.StatusNotFound {
			t.Fatalf("expected status 404, got %d", rr.Code)
		}
	})

	t.Run("ConnectToStreamMissingUpgradeHeaders", func(t *testing.T) {
		t.Parallel()

		logger := slogtest.Make(t, nil)
		dialer := func(network, address string) (net.Conn, error) {
			return &mockConn{}, nil
		}
		manager := immortalstream.NewManager(logger, dialer)
		defer manager.Close()

		// Create a test stream
		createReq := immortalstream.CreateStreamRequest{TCPPort: 8080}
		createResp, err := manager.CreateStream(context.Background(), createReq)
		if err != nil {
			t.Fatalf("failed to create test stream: %v", err)
		}

		handler := newImmortalStreamHandler(manager)

		// Create chi router to handle URL parameters
		r := chi.NewRouter()
		r.Get("/{id}", handler.connectToStream)

		httpReq := httptest.NewRequest("GET", "/"+createResp.ID.String(), nil)
		// Missing upgrade headers
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, httpReq)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected status 400, got %d", rr.Code)
		}
	})

	t.Run("ConnectToStreamInvalidSequenceNumber", func(t *testing.T) {
		t.Parallel()

		logger := slogtest.Make(t, nil)
		dialer := func(network, address string) (net.Conn, error) {
			return &mockConn{}, nil
		}
		manager := immortalstream.NewManager(logger, dialer)
		defer manager.Close()

		// Create a test stream
		createReq := immortalstream.CreateStreamRequest{TCPPort: 8080}
		createResp, err := manager.CreateStream(context.Background(), createReq)
		if err != nil {
			t.Fatalf("failed to create test stream: %v", err)
		}

		handler := newImmortalStreamHandler(manager)

		// Create chi router to handle URL parameters
		r := chi.NewRouter()
		r.Get("/{id}", handler.connectToStream)

		httpReq := httptest.NewRequest("GET", "/"+createResp.ID.String(), nil)
		httpReq.Header.Set("Upgrade", "coder-immortal-stream")
		httpReq.Header.Set("Connection", "upgrade")
		httpReq.Header.Set("x-coder-immortal-stream-sequence-num", "invalid")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, httpReq)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected status 400, got %d", rr.Code)
		}
	})

	t.Run("ParseSequenceNumber", func(t *testing.T) {
		t.Parallel()

		// Test valid sequence number
		seqNum, err := parseSequenceNumber("12345")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if seqNum != 12345 {
			t.Fatalf("expected 12345, got %d", seqNum)
		}

		// Test empty string
		seqNum, err = parseSequenceNumber("")
		if err != nil {
			t.Fatalf("expected no error for empty string, got %v", err)
		}
		if seqNum != 0 {
			t.Fatalf("expected 0 for empty string, got %d", seqNum)
		}

		// Test invalid string
		_, err = parseSequenceNumber("invalid")
		if err == nil {
			t.Fatalf("expected error for invalid string")
		}
	})

	t.Run("Routes", func(t *testing.T) {
		t.Parallel()

		logger := slogtest.Make(t, nil)
		manager := immortalstream.NewManager(logger, nil)
		defer manager.Close()

		handler := newImmortalStreamHandler(manager)
		router := handler.Routes()

		// Test that routes are properly configured
		// This is a basic test to ensure the router doesn't panic
		if router == nil {
			t.Fatalf("expected non-nil router")
		}
	})
}

// Mock connection for testing
type mockConn struct {
	closed bool
	data   []byte
	pos    int
}

func (m *mockConn) Read(b []byte) (int, error) {
	if m.closed {
		return 0, io.EOF
	}
	if m.pos >= len(m.data) {
		return 0, io.EOF
	}

	n := copy(b, m.data[m.pos:])
	m.pos += n
	return n, nil
}

func (m *mockConn) Write(b []byte) (int, error) {
	if m.closed {
		return 0, xerrors.New("connection closed")
	}
	m.data = append(m.data, b...)
	return len(b), nil
}

func (m *mockConn) Close() error {
	m.closed = true
	return nil
}

func (m *mockConn) LocalAddr() net.Addr  { return &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 0} }
func (m *mockConn) RemoteAddr() net.Addr { return &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 0} }

func (m *mockConn) SetDeadline(t time.Time) error      { return nil }
func (m *mockConn) SetReadDeadline(t time.Time) error  { return nil }
func (m *mockConn) SetWriteDeadline(t time.Time) error { return nil }
