package chunk

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSort(t *testing.T) {
	cases := []struct {
		title    string
		input    Bytes
		expected Bytes
	}{
		{
			"simple list",
			Bytes{
				[]byte("3"),
				[]byte("2"),
				[]byte("4"),
				[]byte("1"),
			},
			Bytes{
				[]byte("1"),
				[]byte("2"),
				[]byte("3"),
				[]byte("4"),
			},
		},
		{
			"with empty lines",
			Bytes{
				[]byte("3"),
				[]byte(""),
				[]byte("2"),
				[]byte("4"),
				[]byte(""),
				[]byte("1"),
			},
			Bytes{
				[]byte(""),
				[]byte(""),
				[]byte("1"),
				[]byte("2"),
				[]byte("3"),
				[]byte("4"),
			},
		},
		{
			"with common prefix",
			Bytes{
				[]byte("3"),
				[]byte("2"),
				[]byte("32"),
				[]byte("32"),
				[]byte("321"),
				[]byte("321"),
				[]byte("3211"),
				[]byte("4"),
				[]byte("31"),
				[]byte("1"),
			},
			Bytes{
				[]byte("1"),
				[]byte("2"),
				[]byte("3"),
				[]byte("31"),
				[]byte("32"),
				[]byte("32"),
				[]byte("321"),
				[]byte("321"),
				[]byte("3211"),
				[]byte("4"),
			},
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			data := c.input
			sort.Sort(data)

			if diff := cmp.Diff(c.expected, data); diff != "" {
				t.Error("got diff:", diff)
			}
		})
	}
}
