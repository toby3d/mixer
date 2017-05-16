package interactive

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"errors"
	"io"

	"github.com/pierrec/lz4"
)

type FlusherWriter interface {
	io.Writer
	Flush() error
}

// A Coder is responsible for encoding and decoding discreet JSON messages
// for sending and receiving with a consumer. The Coder is allowed to be
// stateful but it must be safe to use Encode and Decode in parallel with
// each other, however Encode and Decode themselves may not be thread safe.
type Coder struct {
	initializeReader func() (io.Reader, error)
	varUintBuffer    [binary.MaxVarintLen64]byte
	maxDecodeSize    uint64
	rBuf             *bytes.Buffer
	wBuf             *bytes.Buffer
	r                io.Reader
	w                FlusherWriter
}

// Creates a GZIP coder
func NewGzip() *Coder {
	var wBuf bytes.Buffer
	var rBuf bytes.Buffer

	return &Coder{
		maxDecodeSize:    2000000,
		initializeReader: func() (io.Reader, error) { return gzip.NewReader(rBuf) },
		wBuf:             wBuf,
		rBuf:             rBuf,
		w:                gzip.NewWriter(wBuf),
	}
}

// Creates an LZ4 coder
func NewLZ4() *Coder {
	var wBuf bytes.Buffer
	var rBuf bytes.Buffer

	return &Coder{
		maxDecodeSize: 2000000,
		wBuf:          wBuf,
		rBuf:          rBuf,
		w:             lz4.NewWriter(wBuf),
		r:             lz4.NewReader(rBuf),
	}
}

// Encode coverts a JSON string to a binary packet.
func (c *Coder) Encode(json []byte) (packet []byte, err error) {
	// Prepend the target message size to the write buffer.
	sizeLen := binary.PutUvarint(c.varUintBuffer[:], uint64(len(json)))
	c.wBuf.Write(c.varUintBuffer[:sizeLen])
	// Write the json the writer and flush it so the consumer can decode it.
	if _, err := c.w.Write(json); err != nil {
		return nil, err
	}
	if err := c.w.Flush(); err != nil {
		return nil, err
	}
	packet = c.wBuf.Bytes()
	c.wBuf.Reset()
	return packet, nil
}

// Decode coverts a binary packet to a JSON string.
func (c *Coder) Decode(packet []byte) (json []byte, err error) {
	// Look up the decompressed message size.
	size, read := binary.Uvarint(packet)
	if read <= 0 {
		return nil, errors.New("ref/code: error reading packet size")
	}
	if size > c.maxDecodeSize {
		return nil, errors.New("ref/code: packet too large, refusing to decode")
	}
	// Push the packet into the reader's input
	c.rBuf.Write(packet[read:])
	// Initialize the reader. We can't do this until we get the start of the
	// data stream since it may want the headers when it's initialized.
	if c.r == nil {
		if c.r, err = c.initializeReader(); err != nil {
			return nil, err
		}
	}
	json = make([]byte, size)
	if _, err := io.ReadFull(c.r, json); err != nil {
		return nil, err
	}
	return json, nil
}
