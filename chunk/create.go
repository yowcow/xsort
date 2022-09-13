package chunk

import (
	"bufio"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"sort"
	"sync"

	"github.com/yowcow/xsort/types"
)

func CreateChunks(r io.Reader, chunkSize int, tmpDir string) ([]*Chunk, error) {
	base, err := basename(8)
	if err != nil {
		return nil, err
	}

	var idx int
	var chunks []*Chunk
	var wg sync.WaitGroup
	s := bufio.NewScanner(r)

	for {
		bytes, err := createChunkBytes(s, chunkSize)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}

		filename := filepath.Join(tmpDir, fmt.Sprintf("%s-%d", base, idx))
		chunk := &Chunk{file: filename}

		wg.Add(1)
		go func(g *sync.WaitGroup) {
			defer g.Done()
			if err := chunk.createFile(bytes); err != nil {
				panic(err)
			}
		}(&wg)

		chunks = append(chunks, chunk)
		idx++
	}

	wg.Wait()

	return chunks, nil
}

func createChunkBytes(s *bufio.Scanner, chunkSize int) (types.Bytes, error) {
	var curBuf [][]byte
	var curSize int

	for s.Scan() {
		buf := s.Bytes()
		curBuf = append(curBuf, buf)
		curSize += len(buf)

		if curSize >= chunkSize {
			break
		}
	}

	if curSize == 0 {
		return nil, io.EOF
	}

	sort.Sort(types.Bytes(curBuf))

	return curBuf, nil
}

func basename(size int) (string, error) {
	var s string
	b := make([]byte, size)

	_, err := rand.Read(b)
	if err != nil {
		return s, err
	}

	return fmt.Sprintf("%x", b), err
}
