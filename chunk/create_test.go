package chunk

import (
	"os"
	"testing"
)

func TestCreateChunkFiles(t *testing.T) {
	cases := []struct {
		title              string
		inputFile          string
		chunkSize          int
		expectedChunkCount int
	}{
		{
			"chunksize-32",
			"create-test-30.csv",
			32,
			8,
		},
		{
			"chunksize-128",
			"create-test-30.csv",
			128,
			2,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			tmpdir, err := os.MkdirTemp("", c.title)
			if err != nil {
				t.Fatal("failed creating tmp dir:", err)
			}
			defer os.RemoveAll(tmpdir)

			r, err := os.Open(c.inputFile)
			if err != nil {
				t.Fatal("failed reading data:", err)
			}

			chunks, err := CreateChunkFiles(r, c.chunkSize, tmpdir)
			if err != nil {
				t.Error("expected no error but got:", err)
			}
			if chunkCount := len(chunks); chunkCount != c.expectedChunkCount {
				t.Error("expected", c.expectedChunkCount, "but got", chunkCount)
			}
		})
	}
}

func TestBasename(t *testing.T) {
	name, err := basename(8)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if size := len(name); size != 16 {
		// 8 bytes => 16 chars
		t.Error("expected 16 but got", size)
	}
}
