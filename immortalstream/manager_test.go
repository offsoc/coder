package immortalstream

import (
	"context"
	"io"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"

	"cdr.dev/slog/sloggers/slogtest"
)

func TestManager(t *testing.T) {
	t.Parallel()

	t.Run("CreateStream", func(t *testing.T) {
		t.Parallel()

		logger := slogtest.Make(t, nil)
		dialer := func(network, address string) (net.Conn, error) {
			// Mock successful connection
			return &mockConn{}, nil
		}
		manager := NewManager(logger, dialer)
		defer manager.Close()

		req := CreateStreamRequest{TCPPort: 8080}
		resp, err := manager.CreateStream(context.Background(), req)

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if resp.ID == uuid.Nil {
			t.Fatalf("expected valid stream ID")
		}
		if resp.Name == "" {
			t.Fatalf("expected non-empty stream name")
		}

		// Verify stream was added to manager
		streams := manager.ListStreams()
		if len(streams) != 1 {
			t.Fatalf("expected 1 stream, got %d", len(streams))
		}
		if streams[0].ID != resp.ID {
			t.Fatalf("stream ID mismatch")
		}
		if streams[0].TCPPort != 8080 {
			t.Fatalf("expected port 8080, got %d", streams[0].TCPPort)
		}
	})

	t.Run("CreateStreamConnectionRefused", func(t *testing.T) {
		t.Parallel()

		logger := slogtest.Make(t, nil)
		dialer := func(network, address string) (net.Conn, error) {
			return nil, xerrors.New("connection refused")
		}
		manager := NewManager(logger, dialer)
		defer manager.Close()

		req := CreateStreamRequest{TCPPort: 8080}
		_, err := manager.CreateStream(context.Background(), req)

		if err == nil {
			t.Fatalf("expected connection refused error")
		}
		if err.Error() != "The connection was refused" {
			t.Fatalf("expected connection refused error, got %v", err)
		}
	})

	t.Run("MaxStreamsLimit", func(t *testing.T) {
		t.Parallel()

		logger := slogtest.Make(t, nil)
		dialer := func(network, address string) (net.Conn, error) {
			return &mockConn{}, nil
		}
		manager := NewManager(logger, dialer)
		defer manager.Close()

		// Create maximum number of streams and connect them all
		streamIDs := make([]uuid.UUID, MaxStreams)
		for i := 0; i < MaxStreams; i++ {
			req := CreateStreamRequest{TCPPort: 8080 + i}
			resp, err := manager.CreateStream(context.Background(), req)
			if err != nil {
				t.Fatalf("failed to create stream %d: %v", i, err)
			}
			streamIDs[i] = resp.ID

			// Connect each stream so they can't be evicted
			stream, err := manager.GetStream(resp.ID)
			if err != nil {
				t.Fatalf("failed to get stream %d: %v", i, err)
			}
			stream.SetConnected(true)
		}

		// Verify all streams are connected
		connectedCount := 0
		for _, id := range streamIDs {
			stream, err := manager.GetStream(id)
			if err != nil {
				t.Fatalf("failed to get stream: %v", err)
			}
			if stream.IsConnected() {
				connectedCount++
			}
		}
		if connectedCount != MaxStreams {
			t.Fatalf("expected %d connected streams, got %d", MaxStreams, connectedCount)
		}

		// Try to create one more stream - should fail
		req := CreateStreamRequest{TCPPort: 9999}
		_, err := manager.CreateStream(context.Background(), req)
		if err == nil {
			t.Fatalf("expected too many streams error")
		}
		if err.Error() != "Too many Immortal Streams" {
			t.Fatalf("expected too many streams error, got %v", err)
		}
	})

	t.Run("EvictOldestDisconnectedStream", func(t *testing.T) {
		t.Parallel()

		logger := slogtest.Make(t, nil)
		dialer := func(network, address string) (net.Conn, error) {
			return &mockConn{}, nil
		}
		manager := NewManager(logger, dialer)
		defer manager.Close()

		// Create maximum number of streams
		streamIDs := make([]uuid.UUID, MaxStreams)
		for i := 0; i < MaxStreams; i++ {
			req := CreateStreamRequest{TCPPort: 8080 + i}
			resp, err := manager.CreateStream(context.Background(), req)
			if err != nil {
				t.Fatalf("failed to create stream %d: %v", i, err)
			}
			streamIDs[i] = resp.ID
		}

		// Disconnect the first stream (make it evictable)
		stream, err := manager.GetStream(streamIDs[0])
		if err != nil {
			t.Fatalf("failed to get stream: %v", err)
		}
		stream.SetConnected(false)

		// Create another stream - should evict the first one
		req := CreateStreamRequest{TCPPort: 9999}
		_, err = manager.CreateStream(context.Background(), req)
		if err != nil {
			t.Fatalf("expected successful creation after eviction, got %v", err)
		}

		// Verify the first stream was evicted
		_, err = manager.GetStream(streamIDs[0])
		if err == nil {
			t.Fatalf("expected stream to be evicted")
		}
	})

	t.Run("DeleteStream", func(t *testing.T) {
		t.Parallel()

		logger := slogtest.Make(t, nil)
		dialer := func(network, address string) (net.Conn, error) {
			return &mockConn{}, nil
		}
		manager := NewManager(logger, dialer)
		defer manager.Close()

		// Create a stream
		req := CreateStreamRequest{TCPPort: 8080}
		resp, err := manager.CreateStream(context.Background(), req)
		if err != nil {
			t.Fatalf("failed to create stream: %v", err)
		}

		// Delete the stream
		err = manager.DeleteStream(resp.ID)
		if err != nil {
			t.Fatalf("failed to delete stream: %v", err)
		}

		// Verify stream was deleted
		_, err = manager.GetStream(resp.ID)
		if err == nil {
			t.Fatalf("expected stream to be deleted")
		}

		streams := manager.ListStreams()
		if len(streams) != 0 {
			t.Fatalf("expected 0 streams after deletion, got %d", len(streams))
		}
	})

	t.Run("ConnectToStream", func(t *testing.T) {
		t.Parallel()

		logger := slogtest.Make(t, nil)
		dialer := func(network, address string) (net.Conn, error) {
			return &mockConn{}, nil
		}
		manager := NewManager(logger, dialer)
		defer manager.Close()

		// Create a stream
		req := CreateStreamRequest{TCPPort: 8080}
		resp, err := manager.CreateStream(context.Background(), req)
		if err != nil {
			t.Fatalf("failed to create stream: %v", err)
		}

		// Connect to the stream
		clientConn := &mockConn{}
		err = manager.ConnectToStream(context.Background(), resp.ID, clientConn, 0, 0)
		if err != nil {
			t.Fatalf("failed to connect to stream: %v", err)
		}

		// Verify stream is connected
		stream, err := manager.GetStream(resp.ID)
		if err != nil {
			t.Fatalf("failed to get stream: %v", err)
		}
		if !stream.IsConnected() {
			t.Fatalf("expected stream to be connected")
		}
	})

	t.Run("StreamNameGeneration", func(t *testing.T) {
		t.Parallel()

		// Test that stream names are generated consistently
		name1 := generateStreamName()
		if name1 == "" {
			t.Fatalf("expected non-empty stream name")
		}

		// Names should have the format "adjective-noun"
		parts := strings.Split(name1, "-")
		if len(parts) != 2 {
			t.Fatalf("expected stream name format 'adjective-noun', got %q", name1)
		}
	})
}

func TestStream(t *testing.T) {
	t.Parallel()

	t.Run("SetConnected", func(t *testing.T) {
		t.Parallel()

		stream := &Stream{
			ID:      uuid.New(),
			Name:    "test-stream",
			TCPPort: 8080,
		}

		// Test connecting
		stream.SetConnected(true)
		if !stream.IsConnected() {
			t.Fatalf("expected stream to be connected")
		}
		if stream.LastConnectionAt.IsZero() {
			t.Fatalf("expected LastConnectionAt to be set")
		}
		if stream.LastDisconnectionAt != nil {
			t.Fatalf("expected LastDisconnectionAt to be nil when connected")
		}

		// Test disconnecting
		stream.SetConnected(false)
		if stream.IsConnected() {
			t.Fatalf("expected stream to be disconnected")
		}
		if stream.LastDisconnectionAt == nil {
			t.Fatalf("expected LastDisconnectionAt to be set")
		}
	})

	t.Run("UpdateSequenceNumbers", func(t *testing.T) {
		t.Parallel()

		stream := &Stream{
			ID:      uuid.New(),
			Name:    "test-stream",
			TCPPort: 8080,
		}

		stream.UpdateSequenceNumbers(100, 200)

		readerSeq, writerSeq := stream.GetSequenceNumbers()
		if readerSeq != 100 {
			t.Fatalf("expected reader sequence 100, got %d", readerSeq)
		}
		if writerSeq != 200 {
			t.Fatalf("expected writer sequence 200, got %d", writerSeq)
		}
	})

	t.Run("CloseStream", func(t *testing.T) {
		t.Parallel()

		targetConn := &mockConn{}
		stream := &Stream{
			ID:         uuid.New(),
			Name:       "test-stream",
			TCPPort:    8080,
			targetConn: targetConn,
			connected:  true,
		}

		err := stream.Close()
		if err != nil {
			t.Fatalf("expected no error closing stream, got %v", err)
		}

		if stream.IsConnected() {
			t.Fatalf("expected stream to be disconnected after close")
		}
		if !targetConn.closed {
			t.Fatalf("expected target connection to be closed")
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
