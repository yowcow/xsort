package xsort

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSort(t *testing.T) {
	cases := []struct {
		title     string
		chunkSize int64
		input     string
		expected  string
	}{
		{
			"shuffled",
			10,
			`002
005
003
001
004
`,
			`001
002
003
004
005
`,
		},
		{
			"with blank lines",
			10,
			`
002
005

003
001

004
`,
			`


001
002
003
004
005
`,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			dir, err := os.MkdirTemp("", "xsort-test")
			if err != nil {
				t.Fatal("failed creating tmp dir:", err)
			}
			defer os.RemoveAll(dir)

			var buf bytes.Buffer
			opt := SortOptions{
				Input:     strings.NewReader(c.input),
				Output:    &buf,
				ChunkSize: c.chunkSize,
				TmpDir:    dir,
			}
			if err := Sort(&opt); err != nil {
				t.Error("unexpected error:", err)
			}
			if d := cmp.Diff(c.expected, buf.String()); d != "" {
				t.Error("unexpected diff:", d)
			}
		})
	}
}
