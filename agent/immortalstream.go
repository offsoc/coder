package agent

import (
	"encoding/json"
	"net"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"golang.org/x/xerrors"

	"github.com/coder/coder/v2/coderd/httpapi"
	"github.com/coder/coder/v2/codersdk"
	"github.com/coder/coder/v2/immortalstream"
)

// immortalStreamHandler handles HTTP requests for immortal streams.
type immortalStreamHandler struct {
	manager *immortalstream.Manager
}

// newImmortalStreamHandler creates a new immortal stream handler.
func newImmortalStreamHandler(manager *immortalstream.Manager) *immortalStreamHandler {
	return &immortalStreamHandler{
		manager: manager,
	}
}

// Routes sets up the immortal stream routes.
func (h *immortalStreamHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.createStream)
	r.Get("/", h.listStreams)
	r.Get("/{id}", h.connectToStream)
	r.Delete("/{id}", h.deleteStream)

	return r
}

// createStream handles POST /api/v0/immortal-stream
func (h *immortalStreamHandler) createStream(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req immortalstream.CreateStreamRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpapi.Write(ctx, w, http.StatusBadRequest, codersdk.Response{
			Message: "Invalid request body",
			Detail:  err.Error(),
		})
		return
	}

	resp, err := h.manager.CreateStream(ctx, req)
	if err != nil {
		errMsg := err.Error()
		if errMsg == "Too many Immortal Streams" {
			httpapi.Write(ctx, w, http.StatusServiceUnavailable, codersdk.Response{
				Message: "Too many Immortal Streams",
			})
			return
		}
		if errMsg == "The connection was refused" {
			httpapi.Write(ctx, w, http.StatusNotFound, codersdk.Response{
				Message: "The connection was refused",
			})
			return
		}
		httpapi.Write(ctx, w, http.StatusInternalServerError, codersdk.Response{
			Message: "Failed to create immortal stream",
			Detail:  err.Error(),
		})
		return
	}

	httpapi.Write(ctx, w, http.StatusCreated, resp)
}

// listStreams handles GET /api/v0/immortal-stream
func (h *immortalStreamHandler) listStreams(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	streams := h.manager.ListStreams()
	httpapi.Write(ctx, w, http.StatusOK, streams)
}

// connectToStream handles GET /api/v0/immortal-stream/{id}
func (h *immortalStreamHandler) connectToStream(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse stream ID from URL
	idStr := chi.URLParam(r, "id")
	streamID, err := uuid.Parse(idStr)
	if err != nil {
		httpapi.Write(ctx, w, http.StatusBadRequest, codersdk.Response{
			Message: "Invalid stream ID",
			Detail:  err.Error(),
		})
		return
	}

	// Check if the stream exists
	stream, err := h.manager.GetStream(streamID)
	if err != nil {
		httpapi.Write(ctx, w, http.StatusNotFound, codersdk.Response{
			Message: "Stream not found",
		})
		return
	}

	// Check for upgrade headers
	if r.Header.Get("Upgrade") != "coder-immortal-stream" || r.Header.Get("Connection") != "upgrade" {
		httpapi.Write(ctx, w, http.StatusBadRequest, codersdk.Response{
			Message: "Missing required upgrade headers",
			Detail:  "Expected Upgrade: coder-immortal-stream and Connection: upgrade",
		})
		return
	}

	// Parse sequence numbers from headers
	readerSeqNum, err := parseSequenceNumber(r.Header.Get("x-coder-immortal-stream-sequence-num"))
	if err != nil {
		httpapi.Write(ctx, w, http.StatusBadRequest, codersdk.Response{
			Message: "Invalid reader sequence number",
			Detail:  err.Error(),
		})
		return
	}

	// Get current sequence numbers for the response header
	_, currentWriterSeq := stream.GetSequenceNumbers()

	// Upgrade the connection
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		httpapi.Write(ctx, w, http.StatusInternalServerError, codersdk.Response{
			Message: "Connection hijacking not supported",
		})
		return
	}

	// Send upgrade response
	w.Header().Set("Upgrade", "coder-immortal-stream")
	w.Header().Set("Connection", "upgrade")
	w.Header().Set("x-coder-immortal-stream-sequence-num", strconv.FormatInt(currentWriterSeq, 10))
	w.WriteHeader(http.StatusSwitchingProtocols)

	// Hijack the connection
	rawConn, _, err := hijacker.Hijack()
	if err != nil {
		httpapi.Write(ctx, w, http.StatusInternalServerError, codersdk.Response{
			Message: "Failed to hijack connection",
			Detail:  err.Error(),
		})
		return
	}

	// Connect to the immortal stream
	err = h.manager.ConnectToStream(ctx, streamID, rawConn, readerSeqNum, currentWriterSeq)
	if err != nil {
		_ = rawConn.Close()
		return
	}

	// The connection is now managed by the immortal stream
}

// deleteStream handles DELETE /api/v0/immortal-stream/{id}
func (h *immortalStreamHandler) deleteStream(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse stream ID from URL
	idStr := chi.URLParam(r, "id")
	streamID, err := uuid.Parse(idStr)
	if err != nil {
		httpapi.Write(ctx, w, http.StatusBadRequest, codersdk.Response{
			Message: "Invalid stream ID",
			Detail:  err.Error(),
		})
		return
	}

	err = h.manager.DeleteStream(streamID)
	if err != nil {
		if err.Error() == "stream not found" {
			httpapi.Write(ctx, w, http.StatusNotFound, codersdk.Response{
				Message: "Stream not found",
			})
			return
		}
		httpapi.Write(ctx, w, http.StatusInternalServerError, codersdk.Response{
			Message: "Failed to delete stream",
			Detail:  err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// parseSequenceNumber parses a sequence number from a header value.
func parseSequenceNumber(value string) (int64, error) {
	if value == "" {
		return 0, nil
	}

	seqNum, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, xerrors.Errorf("parse sequence number: %w", err)
	}

	return seqNum, nil
}

// httpUpgradeConn wraps a net.Conn to implement io.ReadWriteCloser for HTTP upgrade connections.
type httpUpgradeConn struct {
	net.Conn
}

func (c *httpUpgradeConn) Read(b []byte) (int, error) {
	return c.Conn.Read(b)
}

func (c *httpUpgradeConn) Write(b []byte) (int, error) {
	return c.Conn.Write(b)
}

func (c *httpUpgradeConn) Close() error {
	return c.Conn.Close()
}
