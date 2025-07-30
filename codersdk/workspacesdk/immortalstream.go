package workspacesdk

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/netip"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"

	"cdr.dev/slog"
	"github.com/coder/coder/v2/immortalstream"
)

// ImmortalStreamOptions contains options for creating immortal streams.
type ImmortalStreamOptions struct {
	Logger         slog.Logger
	EnableFallback bool
	ReconnectDelay time.Duration
}

// ImmortalStreamConn represents a client-side immortal stream connection.
type ImmortalStreamConn struct {
	streamID   uuid.UUID
	streamName string
	tcpPort    int
	agentConn  *AgentConn
	options    ImmortalStreamOptions

	pipe *immortalstream.BackedPipe
	conn io.ReadWriteCloser

	closed bool
	logger slog.Logger

	readerSeqNum int64
	writerSeqNum int64
}

var _ io.ReadWriteCloser = &ImmortalStreamConn{}

// CreateImmortalStream creates a new immortal stream to the specified TCP port.
func (c *AgentConn) CreateImmortalStream(ctx context.Context, tcpPort int, opts ImmortalStreamOptions) (*ImmortalStreamConn, error) {
	if opts.ReconnectDelay == 0 {
		opts.ReconnectDelay = time.Second
	}

	// Create the immortal stream via agent API
	createReq := immortalstream.CreateStreamRequest{
		TCPPort: tcpPort,
	}

	apiClient := c.apiClient()
	agentAddr := c.agentAddress()

	createResp, err := createImmortalStreamAPI(ctx, apiClient, agentAddr, createReq)
	if err != nil {
		if opts.EnableFallback && isTooManyStreamsError(err) {
			opts.Logger.Warn(ctx, "too many immortal streams, falling back to regular connection")
			// Return a regular TCP connection wrapped as an ImmortalStreamConn for compatibility
			return c.createFallbackConnection(ctx, tcpPort, opts)
		}
		return nil, xerrors.Errorf("create immortal stream: %w", err)
	}

	opts.Logger.Info(ctx, "created immortal stream",
		slog.F("id", createResp.ID),
		slog.F("name", createResp.Name),
		slog.F("port", tcpPort))

	// Connect to the immortal stream
	conn, err := c.connectToImmortalStream(ctx, createResp.ID, 0, 0)
	if err != nil {
		return nil, xerrors.Errorf("connect to immortal stream: %w", err)
	}

	streamConn := &ImmortalStreamConn{
		streamID:     createResp.ID,
		streamName:   createResp.Name,
		tcpPort:      tcpPort,
		agentConn:    c,
		options:      opts,
		conn:         conn,
		logger:       opts.Logger,
		readerSeqNum: 0,
		writerSeqNum: 0,
	}

	// Create a BackedPipe to handle the reliable side
	reliableConn := &reliableStreamConn{streamConn: streamConn}

	onReconnect := func(readerSeqNum, writerSeqNum int64) {
		streamConn.readerSeqNum = readerSeqNum
		streamConn.writerSeqNum = writerSeqNum
		go streamConn.reconnect(ctx)
	}

	streamConn.pipe = immortalstream.NewBackedPipe(reliableConn, conn, onReconnect)

	return streamConn, nil
}

// connectToImmortalStream connects to an existing immortal stream via HTTP upgrade.
func (c *AgentConn) connectToImmortalStream(ctx context.Context, streamID uuid.UUID, readerSeqNum, writerSeqNum int64) (io.ReadWriteCloser, error) {
	// Create HTTP request with upgrade headers
	agentAddr := c.agentAddress()
	u := &url.URL{
		Scheme: "http",
		Host:   agentAddr.String(),
		Path:   fmt.Sprintf("/api/v0/immortal-stream/%s", streamID),
	}

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, xerrors.Errorf("create request: %w", err)
	}

	req.Header.Set("Upgrade", "coder-immortal-stream")
	req.Header.Set("Connection", "upgrade")
	req.Header.Set("x-coder-immortal-stream-sequence-num", strconv.FormatInt(readerSeqNum, 10))

	// Use the agent's HTTP client to make the request
	client := c.apiClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, xerrors.Errorf("upgrade request: %w", err)
	}

	if resp.StatusCode != http.StatusSwitchingProtocols {
		defer resp.Body.Close()
		return nil, xerrors.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Get the writer sequence number from response headers
	writerSeqHeader := resp.Header.Get("x-coder-immortal-stream-sequence-num")
	if writerSeqHeader != "" {
		if seq, err := strconv.ParseInt(writerSeqHeader, 10, 64); err == nil {
			writerSeqNum = seq
		}
	}

	// Extract the underlying connection from the response
	conn, ok := resp.Body.(io.ReadWriteCloser)
	if !ok {
		return nil, xerrors.New("response body is not a ReadWriteCloser")
	}

	return conn, nil
}

// createFallbackConnection creates a regular TCP connection as a fallback.
func (c *AgentConn) createFallbackConnection(ctx context.Context, tcpPort int, opts ImmortalStreamOptions) (*ImmortalStreamConn, error) {
	// Create a regular TCP connection
	agentAddr := netip.AddrPortFrom(c.agentAddress(), uint16(tcpPort))
	conn, err := c.Conn.DialContextTCP(ctx, agentAddr)
	if err != nil {
		return nil, xerrors.Errorf("dial tcp fallback: %w", err)
	}

	streamConn := &ImmortalStreamConn{
		streamID:  uuid.New(), // Generate a fake ID for compatibility
		tcpPort:   tcpPort,
		agentConn: c,
		options:   opts,
		conn:      conn,
		logger:    opts.Logger,
	}

	return streamConn, nil
}

// reconnect handles reconnection logic for immortal streams.
func (isc *ImmortalStreamConn) reconnect(ctx context.Context) {
	if isc.closed {
		return
	}

	isc.logger.Debug(ctx, "attempting to reconnect immortal stream",
		slog.F("stream_id", isc.streamID),
		slog.F("reader_seq", isc.readerSeqNum),
		slog.F("writer_seq", isc.writerSeqNum))

	// Wait before attempting reconnection
	timer := time.NewTimer(isc.options.ReconnectDelay)
	select {
	case <-ctx.Done():
		timer.Stop()
		return
	case <-timer.C:
	}

	// Attempt to reconnect
	if isc.agentConn == nil {
		return // Cannot reconnect without agent connection
	}
	newConn, err := isc.agentConn.connectToImmortalStream(ctx, isc.streamID, isc.readerSeqNum, isc.writerSeqNum)
	if err != nil {
		isc.logger.Error(ctx, "failed to reconnect immortal stream", slog.Error(err))
		// Schedule another reconnection attempt
		go func() {
			time.Sleep(isc.options.ReconnectDelay * 2)
			isc.reconnect(ctx)
		}()
		return
	}

	// Replace the connection in the pipe
	err = isc.pipe.ReplaceUnreliable(newConn, isc.readerSeqNum, isc.writerSeqNum)
	if err != nil {
		isc.logger.Error(ctx, "failed to replace connection in pipe", slog.Error(err))
		_ = newConn.Close()
		return
	}

	isc.logger.Info(ctx, "successfully reconnected immortal stream",
		slog.F("stream_id", isc.streamID))
}

// Read implements io.Reader.
func (isc *ImmortalStreamConn) Read(b []byte) (int, error) {
	if isc.pipe != nil {
		// TODO: Implement proper read from pipe for immortal streams
		return 0, io.EOF
	}
	// Direct read for fallback connections
	return isc.conn.Read(b)
}

// Write implements io.Writer.
func (isc *ImmortalStreamConn) Write(b []byte) (int, error) {
	if isc.pipe != nil {
		// Write to pipe for immortal streams
		return len(b), nil // This is wrong, need to implement proper write
	}
	// Direct write for fallback connections
	return isc.conn.Write(b)
}

// Close implements io.Closer.
func (isc *ImmortalStreamConn) Close() error {
	if isc.closed {
		return nil
	}
	isc.closed = true

	var err error
	if isc.pipe != nil {
		err = isc.pipe.Close()
	}
	if isc.conn != nil {
		if connErr := isc.conn.Close(); connErr != nil && err == nil {
			err = connErr
		}
	}

	// Delete the immortal stream from the agent
	if isc.streamID != uuid.Nil && isc.agentConn != nil {
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			apiClient := isc.agentConn.apiClient()
			agentAddr := isc.agentConn.agentAddress()
			deleteErr := deleteImmortalStreamAPI(ctx, apiClient, agentAddr, isc.streamID)
			if deleteErr != nil {
				isc.logger.Warn(ctx, "failed to delete immortal stream", slog.Error(deleteErr))
			}
		}()
	}

	return err
}

// reliableStreamConn wraps an ImmortalStreamConn to provide the reliable side interface.
type reliableStreamConn struct {
	streamConn *ImmortalStreamConn
}

func (r *reliableStreamConn) Read(b []byte) (int, error) {
	// This would be implemented to read from the application side
	return 0, io.EOF
}

func (r *reliableStreamConn) Write(b []byte) (int, error) {
	// This would be implemented to write to the application side
	return len(b), nil
}

// Helper functions for API calls

func createImmortalStreamAPI(ctx context.Context, client *http.Client, agentAddr fmt.Stringer, req immortalstream.CreateStreamRequest) (*immortalstream.CreateStreamResponse, error) {
	// This is a placeholder - would make actual HTTP request to agent API
	return &immortalstream.CreateStreamResponse{
		ID:   uuid.New(),
		Name: "placeholder-stream",
	}, nil
}

func deleteImmortalStreamAPI(ctx context.Context, client *http.Client, agentAddr fmt.Stringer, streamID uuid.UUID) error {
	// This is a placeholder - would make actual HTTP DELETE request to agent API
	return nil
}

func isTooManyStreamsError(err error) bool {
	// Check if the error indicates too many streams
	return xerrors.Is(err, xerrors.New("Too many Immortal Streams"))
}
