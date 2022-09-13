package chunk

import (
	"bytes"
	"sort"
)

var (
	_ sort.Interface = (Bytes)(nil)
)

type Bytes [][]byte

func (b Bytes) Len() int {
	return len(b)
}

func (b Bytes) Less(i, j int) bool {
	return bytes.Compare(b[i], b[j]) == -1
}

func (b Bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
