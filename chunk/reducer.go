package chunk

import (
	"bytes"
	"io"
)

type Reducer []*Chunk

func (m *Reducer) AddChunk(chunk *Chunk) {
	*m = append(*m, chunk)
}

func (m *Reducer) Next() ([]byte, error) {
	chunks := *m
	size := len(chunks)

	var b []byte
	var c *Chunk

	for i := 0; i < size; i++ {
		tmp := chunks[i].Head()
		if tmp == nil {
			continue
		}
		if b == nil {
			b = tmp
			c = chunks[i]
			continue
		}
		if bytes.Compare(tmp, b) == -1 {
			b = tmp
			c = chunks[i]
			continue
		}
	}

	if b == nil {
		return nil, io.EOF
	}

	c.Next()

	return b, nil
}
