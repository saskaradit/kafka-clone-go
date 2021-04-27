package client

import (
	"bytes"
	"errors"
)

const defaultScratchSize = 64 * 1024

var errBuffer = errors.New("buffer is too small to fit a single message")

// Simple represents an instance of client connected to a set of kafka servers
type Simple struct {
	addrs []string

	buf bytes.Buffer
}

// NewClient creates a new client for the kafka server
func NewSimple(addrs []string) *Simple {
	return &Simple{
		addrs: addrs,
	}
}

// Send sends the message to the kafka server
func (s *Simple) Send(msg []byte) error {
	_, err := s.buf.Write(msg)
	return err
}

// Receives accepts the message to the kafka server
// error in case something is wrong
// The scratch can be used to read
func (s *Simple) Receive(scratch []byte) ([]byte, error) {
	if scratch == nil {
		scratch = make([]byte, defaultScratchSize)
	}

	n, err := s.buf.Read(scratch)
	if err != nil {
		return nil, err
	}

	truncated, rest, err := cutToLast(scratch[0:n])
	if err != nil {
		return nil, err
	}

	_ = rest
	return truncated, nil
}

func cutToLast(res []byte) (truncated []byte, rest []byte, err error) {
	n := len(res)
	if n == 0 {
		return res, nil, nil
	}

	if res[n-1] == '\n' {
		return res, nil, nil
	}

	lastPos := bytes.LastIndexByte(res, '\n')
	if lastPos < 0 {
		return nil, nil, errBuffer
	}
	return res[0 : lastPos+1], res[lastPos+1:], nil

}
