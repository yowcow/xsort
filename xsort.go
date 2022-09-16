package xsort

import (
	"errors"
	"io"
	"os"

	"github.com/yowcow/xsort/chunk"
)

type SortOptions struct {
	Input     io.Reader
	Output    io.Writer
	ChunkSize int64
	TmpDir    string
}

func Sort(opt *SortOptions) error {
	files, err := chunk.CreateChunkFiles(opt.Input, opt.ChunkSize, opt.TmpDir)
	if err != nil {
		return err
	}

	var reducer chunk.Reducer
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

		reducer.AddChunk(chunk)
	}

	for {
		b, err := reducer.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}

		opt.Output.Write(append(b, byte('\n')))
	}

	return nil
}
