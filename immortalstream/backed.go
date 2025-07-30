package immortalstream

import (
	"context"
	"io"
	"sync"

	"golang.org/x/xerrors"
)

// MaxBufferSize is the maximum amount of data we buffer for replay (64 MiB)
const MaxBufferSize = 64 * 1024 * 1024

// BackedReader wraps an unreliable io.Reader and allows it to be replaced
// when it errors. It keeps track of the number of bytes read.
type BackedReader struct {
	cond        *sync.Cond
	r           io.Reader // unreliable
	sequenceNum int64
	closed      bool

	// Callback to notify when reconnection is needed
	onReconnect func(seqNum int64)
}

var _ io.ReadCloser = &BackedReader{}

// NewBackedReader creates a new BackedReader with the given initial reader.
func NewBackedReader(r io.Reader, onReconnect func(seqNum int64)) *BackedReader {
	return &BackedReader{
		cond:        sync.NewCond(&sync.Mutex{}),
		r:           r,
		sequenceNum: 0,
		closed:      false,
		onReconnect: onReconnect,
	}
}

// Read reads from the underlying reader. If the reader errors, it triggers
// a reconnection and waits for a new reader to be provided.
func (b *BackedReader) Read(buf []byte) (int, error) {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	for {
		if b.closed {
			return 0, io.EOF
		}

		if b.r == nil {
			b.cond.Wait()
			continue
		}

		n, err := b.r.Read(buf)
		if n > 0 {
			b.sequenceNum += int64(n)
		}

		if err != nil && err != io.EOF {
			// Reader errored, trigger reconnection
			b.r = nil
			if b.onReconnect != nil {
				go b.onReconnect(b.sequenceNum)
			}
			continue
		}

		return n, err
	}
}

// Reconnect replaces the reader with a new one, starting from the given sequence number.
func (b *BackedReader) Reconnect(newR io.Reader, replaySeqNum int64) error {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	if b.closed {
		return xerrors.New("reader is closed")
	}

	// Validate that we're not going backwards in sequence
	if replaySeqNum > b.sequenceNum {
		return xerrors.Errorf("replay sequence number %d is ahead of current %d", replaySeqNum, b.sequenceNum)
	}

	b.r = newR
	b.cond.Broadcast()
	return nil
}

// Close closes the reader.
func (b *BackedReader) Close() error {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	b.closed = true
	b.r = nil
	b.cond.Broadcast()
	return nil
}

// SequenceNum returns the current sequence number (bytes read).
func (b *BackedReader) SequenceNum() int64 {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()
	return b.sequenceNum
}

// BackedWriter wraps an unreliable io.Writer and keeps a ring buffer
// of the most recent data for replay when the writer is replaced.
type BackedWriter struct {
	cond   *sync.Cond
	w      io.Writer // unreliable
	closed bool

	// ring buffer
	buffer       []byte
	oldestIndex  int
	newestIndex  int
	oldestSeqNum int64
	newestSeqNum int64
	bufferFull   bool

	// Callback to notify when reconnection is needed
	onReconnect func(seqNum int64)
}

var _ io.WriteCloser = &BackedWriter{}

// NewBackedWriter creates a new BackedWriter with the given initial writer.
func NewBackedWriter(w io.Writer, onReconnect func(seqNum int64)) *BackedWriter {
	return &BackedWriter{
		cond:         sync.NewCond(&sync.Mutex{}),
		w:            w,
		closed:       false,
		buffer:       make([]byte, MaxBufferSize),
		oldestIndex:  0,
		newestIndex:  0,
		oldestSeqNum: 0,
		newestSeqNum: 0,
		bufferFull:   false,
		onReconnect:  onReconnect,
	}
}

// Write writes data to the underlying writer and stores it in the ring buffer.
func (b *BackedWriter) Write(data []byte) (int, error) {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	for {
		if b.closed {
			return 0, xerrors.New("writer is closed")
		}

		if b.w == nil {
			b.cond.Wait()
			continue
		}

		// Try to write to the underlying writer
		n, err := b.w.Write(data)

		// Always store data in ring buffer before returning
		b.writeToBuffer(data[:n])

		if err != nil {
			// Writer errored, trigger reconnection
			b.w = nil
			if b.onReconnect != nil {
				go b.onReconnect(b.newestSeqNum)
			}
			continue
		}

		return n, err
	}
}

// writeToBuffer writes data to the ring buffer.
func (b *BackedWriter) writeToBuffer(data []byte) {
	for _, byte := range data {
		b.buffer[b.newestIndex] = byte
		b.newestIndex = (b.newestIndex + 1) % MaxBufferSize
		b.newestSeqNum++

		if b.newestIndex == b.oldestIndex && b.bufferFull {
			// Buffer wrapped around, advance oldest pointer
			b.oldestIndex = (b.oldestIndex + 1) % MaxBufferSize
			b.oldestSeqNum++
		} else if b.newestIndex == b.oldestIndex {
			// Buffer just became full
			b.bufferFull = true
		}
	}
}

// Reconnect replaces the writer with a new one and replays data from the given sequence number.
func (b *BackedWriter) Reconnect(newW io.Writer, replaySeqNum int64) error {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	if b.closed {
		return xerrors.New("writer is closed")
	}

	// Check if replay sequence number is available in buffer
	if replaySeqNum < b.oldestSeqNum {
		return xerrors.Errorf("replay sequence number %d is older than oldest buffered %d", replaySeqNum, b.oldestSeqNum)
	}

	if replaySeqNum > b.newestSeqNum {
		return xerrors.Errorf("replay sequence number %d is ahead of newest buffered %d", replaySeqNum, b.newestSeqNum)
	}

	// Replay data from replaySeqNum to current
	replayBytes := b.newestSeqNum - replaySeqNum
	if replayBytes > 0 {
		replayData := make([]byte, replayBytes)
		b.readFromBuffer(replaySeqNum, replayData)

		_, err := newW.Write(replayData)
		if err != nil {
			return xerrors.Errorf("replay write failed: %w", err)
		}
	}

	b.w = newW
	b.cond.Broadcast()
	return nil
}

// readFromBuffer reads data from the ring buffer starting at the given sequence number.
func (b *BackedWriter) readFromBuffer(startSeqNum int64, buf []byte) {
	bytesToSkip := startSeqNum - b.oldestSeqNum
	startIndex := (b.oldestIndex + int(bytesToSkip)) % MaxBufferSize

	for i := 0; i < len(buf); i++ {
		readIndex := (startIndex + i) % MaxBufferSize
		buf[i] = b.buffer[readIndex]
	}
}

// Close closes the writer.
func (b *BackedWriter) Close() error {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	b.closed = true
	b.w = nil
	b.cond.Broadcast()
	return nil
}

// SequenceNum returns the current sequence number (bytes written).
func (b *BackedWriter) SequenceNum() int64 {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()
	return b.newestSeqNum
}

// BackedPipe connects a reliable io.ReadWriter with a series of unreliable
// io.ReadWriteCloser instances, similar to net.Pipe but with reconnection capability.
type BackedPipe struct {
	reliable   io.ReadWriter
	unreliable io.ReadWriteCloser

	reader *BackedReader
	writer *BackedWriter

	ctx    context.Context
	cancel context.CancelFunc

	onReconnect func(readerSeqNum, writerSeqNum int64)

	closed bool
	mu     sync.Mutex
}

// NewBackedPipe creates a new BackedPipe connecting a reliable and unreliable connection.
func NewBackedPipe(reliable io.ReadWriter, unreliable io.ReadWriteCloser, onReconnect func(readerSeqNum, writerSeqNum int64)) *BackedPipe {
	ctx, cancel := context.WithCancel(context.Background())

	bp := &BackedPipe{
		reliable:    reliable,
		unreliable:  unreliable,
		ctx:         ctx,
		cancel:      cancel,
		onReconnect: onReconnect,
	}

	// Create backed reader/writer with reconnection callbacks
	bp.reader = NewBackedReader(unreliable, func(seqNum int64) {
		if bp.onReconnect != nil {
			bp.onReconnect(seqNum, bp.writer.SequenceNum())
		}
	})

	bp.writer = NewBackedWriter(unreliable, func(seqNum int64) {
		if bp.onReconnect != nil {
			bp.onReconnect(bp.reader.SequenceNum(), seqNum)
		}
	})

	// Start copying data between reliable and unreliable sides
	go bp.copyReliableToUnreliable()
	go bp.copyUnreliableToReliable()

	return bp
}

// copyReliableToUnreliable copies data from reliable reader to backed writer.
func (bp *BackedPipe) copyReliableToUnreliable() {
	defer bp.Close()
	_, _ = io.Copy(bp.writer, bp.reliable)
}

// copyUnreliableToReliable copies data from backed reader to reliable writer.
func (bp *BackedPipe) copyUnreliableToReliable() {
	defer bp.Close()
	_, _ = io.Copy(bp.reliable, bp.reader)
}

// ReplaceUnreliable replaces the unreliable connection with a new one.
func (bp *BackedPipe) ReplaceUnreliable(newUnreliable io.ReadWriteCloser, readerSeqNum, writerSeqNum int64) error {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	if bp.closed {
		return xerrors.New("pipe is closed")
	}

	// Close old unreliable connection
	if bp.unreliable != nil {
		_ = bp.unreliable.Close()
	}

	bp.unreliable = newUnreliable

	// Reconnect reader and writer
	err := bp.reader.Reconnect(newUnreliable, readerSeqNum)
	if err != nil {
		return xerrors.Errorf("reconnect reader: %w", err)
	}

	err = bp.writer.Reconnect(newUnreliable, writerSeqNum)
	if err != nil {
		return xerrors.Errorf("reconnect writer: %w", err)
	}

	return nil
}

// Close closes the pipe and all underlying connections.
func (bp *BackedPipe) Close() error {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	if bp.closed {
		return nil
	}

	bp.closed = true
	bp.cancel()

	if bp.reader != nil {
		_ = bp.reader.Close()
	}
	if bp.writer != nil {
		_ = bp.writer.Close()
	}
	if bp.unreliable != nil {
		_ = bp.unreliable.Close()
	}

	return nil
}

// ReaderSequenceNum returns the current reader sequence number.
func (bp *BackedPipe) ReaderSequenceNum() int64 {
	if bp.reader != nil {
		return bp.reader.SequenceNum()
	}
	return 0
}

// WriterSequenceNum returns the current writer sequence number.
func (bp *BackedPipe) WriterSequenceNum() int64 {
	if bp.writer != nil {
		return bp.writer.SequenceNum()
	}
	return 0
}
