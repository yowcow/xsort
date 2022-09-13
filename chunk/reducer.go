package chunk

import (
	"io"

	"github.com/yowcow/xsort/types"
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
		if b == nil {
			b = tmp
			c = m.chunks[i]
		} else if types.CmpBytes(tmp, b) {
			b = tmp
			c = m.chunks[i]
		}
	}

	if b == nil {
		return nil, io.EOF
	}

	c.Next()

	return b, nil
}
