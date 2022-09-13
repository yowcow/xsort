package chunk

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReducer(t *testing.T) {
	cases := []struct {
		title    string
		inputs   []string
		expected []string
	}{
		{
			"no chunk",
			[]string{},
			nil,
		},
		{
			"one chunk",
			[]string{`001
002
002
003
`},
			[]string{
				"001",
				"002",
				"002",
				"003",
			},
		},
		{
			"two chunks",
			[]string{`001
002
002
003`, `001
003
004`},
			[]string{
				"001",
				"001",
				"002",
				"002",
				"003",
				"003",
				"004",
			},
		},
		{
			"three chunks",
			[]string{`001
002
005`, `001
003
007`, `002
006
008`,
			},
			[]string{
				"001",
				"001",
				"002",
				"002",
				"003",
				"005",
				"006",
				"007",
				"008",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			var chunks []*Chunk
			for _, input := range c.inputs {
				r := strings.NewReader(input)
				chunk, err := NewChunk(r)
				if err != nil {
					t.Fatal("expected no error but got:", err)
				}
				chunks = append(chunks, chunk)
			}

			var actual []string
			m := NewReducer(chunks)
			for {
				output, err := m.Next()
				if errors.Is(err, io.EOF) {
					break
				}
				if err != nil {
					t.Error("expected no error but got:", err)
				}

				actual = append(actual, string(output))
			}

			if diff := cmp.Diff(c.expected, actual); diff != "" {
				t.Error("expected no diff but got:", diff)
			}
		})
	}
}
