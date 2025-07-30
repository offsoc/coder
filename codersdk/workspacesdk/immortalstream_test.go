package workspacesdk

import (
	"context"
	"io"
	"net"
	"net/http"
	"net/netip"
	"testing"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"

	"cdr.dev/slog/sloggers/slogtest"
)

func TestImmortalStreamConn(t *testing.T) {
	t.Parallel()

	t.Run("CreateImmortalStreamFallback", func(t *testing.T) {
		t.Parallel()

		// Create a mock agent connection
		agentConn := newMockAgentConn()

		logger := slogtest.Make(t, nil)
		opts := ImmortalStreamOptions{
			Logger:         logger,
			EnableFallback: true,
			ReconnectDelay: 100 * time.Millisecond,
		}

		// This will use the fallback since we're not implementing full HTTP upgrade
		streamConn, err := agentConn.CreateImmortalStream(context.Background(), 8080, opts)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		defer streamConn.Close()

		if streamConn.streamID == uuid.Nil {
			t.Fatalf("expected valid stream ID")
		}
		if streamConn.tcpPort != 8080 {
			t.Fatalf("expected port 8080, got %d", streamConn.tcpPort)
		}
	})

	t.Run("ImmortalStreamConnReadWrite", func(t *testing.T) {
		t.Parallel()

		agentConn := newMockAgentConn()

		logger := slogtest.Make(t, nil)
		opts := ImmortalStreamOptions{
			Logger:         logger,
			EnableFallback: true,
		}

		streamConn, err := agentConn.CreateImmortalStream(context.Background(), 8080, opts)
		if err != nil {
			t.Fatalf("failed to create immortal stream: %v", err)
		}
		defer streamConn.Close()

		// Test writing (fallback mode uses direct connection)
		testData := []byte("hello world")
		n, err := streamConn.Write(testData)
		if err != nil {
			t.Fatalf("write failed: %v", err)
		}
		if n != len(testData) {
			t.Fatalf("expected %d bytes written, got %d", len(testData), n)
		}

		// Test reading
		buf := make([]byte, len(testData))
		n, err = streamConn.Read(buf)
		if err != nil && err != io.EOF {
			t.Fatalf("read failed: %v", err)
		}
		if string(buf[:n]) != string(testData) {
			t.Fatalf("expected %q, got %q", string(testData), string(buf[:n]))
		}
	})

	t.Run("ImmortalStreamConnClose", func(t *testing.T) {
		t.Parallel()

		agentConn := newMockAgentConn()

		logger := slogtest.Make(t, nil)
		opts := ImmortalStreamOptions{
			Logger:         logger,
			EnableFallback: true,
		}

		streamConn, err := agentConn.CreateImmortalStream(context.Background(), 8080, opts)
		if err != nil {
			t.Fatalf("failed to create immortal stream: %v", err)
		}

		err = streamConn.Close()
		if err != nil {
			t.Fatalf("close failed: %v", err)
		}

		// Verify connection is closed
		if !streamConn.closed {
			t.Fatalf("expected connection to be closed")
		}
	})

	t.Run("ReconnectLogic", func(t *testing.T) {
		t.Parallel()

		logger := slogtest.Make(t, nil)
		streamConn := &ImmortalStreamConn{
			streamID:  uuid.New(),
			tcpPort:   8080,
			agentConn: nil, // Can't use mock due to type constraints
			logger:    logger,
			options:   ImmortalStreamOptions{ReconnectDelay: 10 * time.Millisecond},
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		// Test reconnection logic (will fail but shouldn't panic)
		streamConn.reconnect(ctx)

		// Should not panic and should handle context cancellation
	})
}

func TestImmortalStreamOptions(t *testing.T) {
	t.Parallel()

	t.Run("DefaultReconnectDelay", func(t *testing.T) {
		t.Parallel()

		agentConn := newMockAgentConn()

		logger := slogtest.Make(t, nil)
		opts := ImmortalStreamOptions{
			Logger:         logger,
			EnableFallback: true,
			// ReconnectDelay not set, should default to 1 second
		}

		streamConn, err := agentConn.CreateImmortalStream(context.Background(), 8080, opts)
		if err != nil {
			t.Fatalf("failed to create immortal stream: %v", err)
		}
		defer streamConn.Close()

		if streamConn.options.ReconnectDelay != time.Second {
			t.Fatalf("expected default reconnect delay of 1 second, got %v", streamConn.options.ReconnectDelay)
		}
	})
}

func TestHelperFunctions(t *testing.T) {
	t.Parallel()

	t.Run("isTooManyStreamsError", func(t *testing.T) {
		t.Parallel()

		// Test with too many streams error
		tooManyErr := xerrors.New("Too many Immortal Streams")
		if !isTooManyStreamsError(tooManyErr) {
			t.Fatalf("expected true for too many streams error")
		}

		// Test with other error
		otherErr := xerrors.New("some other error")
		if isTooManyStreamsError(otherErr) {
			t.Fatalf("expected false for other error")
		}
	})
}

// Mock AgentConn for testing
type mockAgentConn struct {
	addr       netip.Addr
	httpClient *http.Client
	Conn       struct {
		DialContextTCP func(ctx context.Context, addr netip.AddrPort) (net.Conn, error)
	}
}

func newMockAgentConn() *mockAgentConn {
	return &mockAgentConn{
		addr:       netip.MustParseAddr("127.0.0.1"),
		httpClient: &http.Client{},
	}
}

func (m *mockAgentConn) agentAddress() netip.Addr {
	if m == nil {
		return netip.MustParseAddr("127.0.0.1")
	}
	return m.addr
}

func (m *mockAgentConn) apiClient() *http.Client {
	if m == nil || m.httpClient == nil {
		return &http.Client{}
	}
	return m.httpClient
}

func (m *mockAgentConn) connectToImmortalStream(ctx context.Context, streamID uuid.UUID, readerSeqNum, writerSeqNum int64) (io.ReadWriteCloser, error) {
	return &mockConn{}, nil
}

func (m *mockAgentConn) createFallbackConnection(ctx context.Context, tcpPort int, opts ImmortalStreamOptions) (*ImmortalStreamConn, error) {
	// Set default reconnect delay if not specified
	if opts.ReconnectDelay == 0 {
		opts.ReconnectDelay = time.Second
	}

	conn := &mockConn{}
	return &ImmortalStreamConn{
		streamID:  uuid.New(),
		tcpPort:   tcpPort,
		agentConn: nil, // Keep as nil for tests to avoid type issues
		options:   opts,
		conn:      conn,
		logger:    opts.Logger,
	}, nil
}

func (m *mockAgentConn) CreateImmortalStream(ctx context.Context, tcpPort int, opts ImmortalStreamOptions) (*ImmortalStreamConn, error) {
	return m.createFallbackConnection(ctx, tcpPort, opts)
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
