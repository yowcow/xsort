package types

import "sort"

var _ sort.Interface = (Bytes)(nil)

type Bytes [][]byte

func (b Bytes) Len() int {
	return len(b)
}

func (b Bytes) Less(i, j int) bool {
	x := b[i]
	xlen := len(x)
	y := b[j]
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

func (b Bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
