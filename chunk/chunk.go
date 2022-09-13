package chunk

import (
	"bufio"
	"io"
)

type Chunk struct {
	s    *bufio.Scanner
	head []byte
}

func NewChunk(r io.Reader) (*Chunk, error) {
	s := bufio.NewScanner(r)
	if !s.Scan() {
		return nil, io.EOF
	}

	return &Chunk{
		s:    s,
		head: s.Bytes(),
	}, nil
}

func (c *Chunk) Head() []byte {
	return c.head
}

func (c *Chunk) Next() bool {
	c.head = nil

	if c.s.Scan() {
		c.head = c.s.Bytes()
		return true
	}

	return false
}
