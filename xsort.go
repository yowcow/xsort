package xsort

import (
	"errors"
	"io"
	"os"

	"github.com/yowcow/xsort/chunk"
)

func Sort(r io.Reader, w io.Writer, chunkSize int, tmpDir string) error {
	files, err := chunk.CreateChunkFiles(r, chunkSize, tmpDir)
	if err != nil {
		return err
	}

	var chunks []*chunk.Chunk
	for _, f := range files {
		r, err := os.Open(f)
		if err != nil {
			return err
		}
		defer r.Close()

		chunk, err := chunk.NewChunk(r)
		if err != nil {
			return err
		}

		chunks = append(chunks, chunk)
	}

	reducer := chunk.NewReducer(chunks)

	for {
		b, err := reducer.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}

		w.Write(append(b, byte('\n')))
	}

	return nil
}
