package types

import "sort"

var _ sort.Interface = (Bytes)(nil)

type Bytes [][]byte

func (b Bytes) Len() int {
	return len(b)
}

func (b Bytes) Less(i, j int) bool {
	return CmpBytes(b[i], b[j])
}

func (b Bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func CmpBytes(x, y []byte) bool {
	xlen := len(x)
	ylen := len(y)

	minlen := xlen
	if ylen < xlen {
		minlen = ylen
	}

	for i := 0; i < minlen; i++ {
		if x[i] < y[i] {
			return true
		} else if x[i] > y[i] {
			return false
		}
	}

	return xlen < ylen
}
