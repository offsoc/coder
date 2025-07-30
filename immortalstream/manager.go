package immortalstream

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"

	"cdr.dev/slog"
)

const (
	// MaxStreams is the maximum number of immortal streams per agent (32 streams = 2 GiB max)
	MaxStreams = 32
)

// Stream represents an immortal stream with metadata and connection management.
type Stream struct {
	ID                  uuid.UUID  `json:"id"`
	Name                string     `json:"name"`
	TCPPort             int        `json:"tcp_port"`
	CreatedAt           time.Time  `json:"created_at"`
	LastConnectionAt    time.Time  `json:"last_connection_at"`
	LastDisconnectionAt *time.Time `json:"last_disconnection_at,omitempty"`

	// Internal state
	pipe       *BackedPipe
	targetConn io.ReadWriteCloser
	connected  bool
	mu         sync.RWMutex

	// Connection sequence tracking
	readerSeqNum int64
	writerSeqNum int64
}

// IsConnected returns whether the stream is currently connected.
func (s *Stream) IsConnected() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.connected
}

// SetConnected updates the connection status and timestamps.
func (s *Stream) SetConnected(connected bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	if connected && !s.connected {
		s.LastConnectionAt = now
		s.LastDisconnectionAt = nil
	} else if !connected && s.connected {
		s.LastDisconnectionAt = &now
	}
	s.connected = connected
}

// UpdateSequenceNumbers updates the sequence numbers for reconnection.
func (s *Stream) UpdateSequenceNumbers(readerSeqNum, writerSeqNum int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.readerSeqNum = readerSeqNum
	s.writerSeqNum = writerSeqNum
}

// GetSequenceNumbers returns the current sequence numbers for reconnection.
func (s *Stream) GetSequenceNumbers() (int64, int64) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.readerSeqNum, s.writerSeqNum
}

// Close closes the stream and all associated connections.
func (s *Stream) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var errs []error

	if s.pipe != nil {
		if err := s.pipe.Close(); err != nil {
			errs = append(errs, err)
		}
		s.pipe = nil
	}

	if s.targetConn != nil {
		if err := s.targetConn.Close(); err != nil {
			errs = append(errs, err)
		}
		s.targetConn = nil
	}

	s.connected = false

	if len(errs) > 0 {
		return xerrors.Errorf("close stream: %v", errs)
	}
	return nil
}

// Manager manages immortal streams for an agent.
type Manager struct {
	logger slog.Logger
	dialer func(network, address string) (net.Conn, error)

	streams   map[uuid.UUID]*Stream
	streamsMu sync.RWMutex
}

// NewManager creates a new immortal stream manager.
func NewManager(logger slog.Logger, dialer func(network, address string) (net.Conn, error)) *Manager {
	return &Manager{
		logger:  logger,
		dialer:  dialer,
		streams: make(map[uuid.UUID]*Stream),
	}
}

// CreateStreamRequest represents a request to create a new immortal stream.
type CreateStreamRequest struct {
	TCPPort int `json:"tcp_port"`
}

// CreateStreamResponse represents the response after creating a stream.
type CreateStreamResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// CreateStream creates a new immortal stream.
func (m *Manager) CreateStream(ctx context.Context, req CreateStreamRequest) (*CreateStreamResponse, error) {
	m.streamsMu.Lock()
	defer m.streamsMu.Unlock()

	// Check if we're at the limit
	if len(m.streams) >= MaxStreams {
		// Try to evict a disconnected stream
		evicted := m.evictOldestDisconnectedStreamLocked()
		if !evicted {
			return nil, xerrors.New("Too many Immortal Streams")
		}
	}

	// Generate a unique ID and human-readable name
	streamID := uuid.New()
	streamName := generateStreamName()

	// Attempt to dial the local service
	targetAddr := fmt.Sprintf("localhost:%d", req.TCPPort)
	targetConn, err := m.dialer("tcp", targetAddr)
	if err != nil {
		return nil, xerrors.Errorf("The connection was refused")
	}

	// Create the stream
	stream := &Stream{
		ID:               streamID,
		Name:             streamName,
		TCPPort:          req.TCPPort,
		CreatedAt:        time.Now(),
		LastConnectionAt: time.Now(),
		targetConn:       targetConn,
		connected:        false,
	}

	m.streams[streamID] = stream

	m.logger.Info(ctx, "created immortal stream",
		slog.F("id", streamID),
		slog.F("name", streamName),
		slog.F("port", req.TCPPort))

	return &CreateStreamResponse{
		ID:   streamID,
		Name: streamName,
	}, nil
}

// evictOldestDisconnectedStreamLocked evicts the oldest disconnected stream.
// Must be called with streamsMu write lock held.
func (m *Manager) evictOldestDisconnectedStreamLocked() bool {
	var oldestStream *Stream
	var oldestTime time.Time = time.Now()

	for _, stream := range m.streams {
		if !stream.IsConnected() && (oldestStream == nil || stream.LastConnectionAt.Before(oldestTime)) {
			oldestStream = stream
			oldestTime = stream.LastConnectionAt
		}
	}

	if oldestStream != nil {
		delete(m.streams, oldestStream.ID)
		_ = oldestStream.Close()
		return true
	}

	return false
}

// GetStream retrieves a stream by ID.
func (m *Manager) GetStream(id uuid.UUID) (*Stream, error) {
	m.streamsMu.RLock()
	defer m.streamsMu.RUnlock()

	stream, exists := m.streams[id]
	if !exists {
		return nil, xerrors.New("stream not found")
	}

	return stream, nil
}

// ListStreams returns all current streams.
func (m *Manager) ListStreams() []*Stream {
	m.streamsMu.RLock()
	defer m.streamsMu.RUnlock()

	streams := make([]*Stream, 0, len(m.streams))
	for _, stream := range m.streams {
		streams = append(streams, stream)
	}

	return streams
}

// DeleteStream deletes a stream by ID.
func (m *Manager) DeleteStream(id uuid.UUID) error {
	m.streamsMu.Lock()
	defer m.streamsMu.Unlock()

	stream, exists := m.streams[id]
	if !exists {
		return xerrors.New("stream not found")
	}

	delete(m.streams, id)
	return stream.Close()
}

// ConnectToStream establishes a connection to an existing stream.
func (m *Manager) ConnectToStream(ctx context.Context, id uuid.UUID, conn io.ReadWriteCloser, readerSeqNum, writerSeqNum int64) error {
	stream, err := m.GetStream(id)
	if err != nil {
		return err
	}

	stream.mu.Lock()
	defer stream.mu.Unlock()

	// If this is the first connection, create the pipe
	if stream.pipe == nil {
		onReconnect := func(rSeqNum, wSeqNum int64) {
			m.logger.Debug(ctx, "stream needs reconnection",
				slog.F("stream_id", id),
				slog.F("reader_seq", rSeqNum),
				slog.F("writer_seq", wSeqNum))
			stream.UpdateSequenceNumbers(rSeqNum, wSeqNum)
			stream.SetConnected(false)
		}

		stream.pipe = NewBackedPipe(stream.targetConn, conn, onReconnect)
	} else {
		// Replace the unreliable connection
		err = stream.pipe.ReplaceUnreliable(conn, readerSeqNum, writerSeqNum)
		if err != nil {
			return xerrors.Errorf("replace unreliable connection: %w", err)
		}
	}

	stream.SetConnected(true)
	stream.UpdateSequenceNumbers(readerSeqNum, writerSeqNum)

	m.logger.Info(ctx, "connected to immortal stream",
		slog.F("id", id),
		slog.F("name", stream.Name),
		slog.F("reader_seq", readerSeqNum),
		slog.F("writer_seq", writerSeqNum))

	return nil
}

// Close closes all streams and the manager.
func (m *Manager) Close() error {
	m.streamsMu.Lock()
	defer m.streamsMu.Unlock()

	var errs []error
	for id, stream := range m.streams {
		if err := stream.Close(); err != nil {
			errs = append(errs, xerrors.Errorf("close stream %s: %w", id, err))
		}
	}

	m.streams = make(map[uuid.UUID]*Stream)

	if len(errs) > 0 {
		return xerrors.Errorf("close manager: %v", errs)
	}
	return nil
}

// generateStreamName generates a human-readable name for a stream (like Docker names).
func generateStreamName() string {
	adjectives := []string{
		"happy", "clever", "brave", "swift", "gentle", "bright", "calm", "eager",
		"kind", "wise", "bold", "quick", "quiet", "strong", "warm", "cool",
		"fluffy", "sleek", "mighty", "tiny", "vast", "sharp", "smooth", "rough",
	}

	names := []string{
		"poincare", "euler", "newton", "gauss", "riemann", "hilbert", "turing", "godel",
		"feynman", "curie", "darwin", "mendel", "watson", "crick", "bohr", "planck",
		"hawking", "einstein", "tesla", "faraday", "pascal", "fermat", "cauchy", "fourier",
	}

	// Simple pseudo-random selection based on current time
	now := time.Now().UnixNano()
	adjIdx := int(now) % len(adjectives)
	nameIdx := int(now/1000) % len(names)

	return fmt.Sprintf("%s-%s", adjectives[adjIdx], names[nameIdx])
}
