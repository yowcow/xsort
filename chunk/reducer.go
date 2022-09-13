package chunk

import (
	"bytes"
	"io"
)

type Reducer struct {
	chunks []*Chunk
	count  int
}

func NewReducer(chunks []*Chunk) *Reducer {
	return &Reducer{
		chunks: chunks,
		count:  len(chunks),
	}
}

func (m *Reducer) Next() ([]byte, error) {
	var b []byte
	var c *Chunk

	for i := 0; i < m.count; i++ {
		tmp := m.chunks[i].Head()
		if tmp == nil {
			continue
		}
		if b == nil {
			b = tmp
			c = m.chunks[i]
			continue
		}
		if bytes.Compare(tmp, b) == -1 {
			b = tmp
			c = m.chunks[i]
			continue
		}
	}

	if b == nil {
		return nil, io.EOF
	}

	c.Next()

	return b, nil
}
