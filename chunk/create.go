package chunk

import (
	"bufio"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"github.com/yowcow/xsort/types"
)

func CreateChunkFiles(r io.Reader, chunkSize int, tmpDir string) ([]string, error) {
	base, err := basename(8)
	if err != nil {
		return nil, err
	}

	var idx int
	var files []string
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
		wg.Add(1)

		go func(file string, g *sync.WaitGroup) {
			defer g.Done()
			if err := createChunkFile(file, bytes); err != nil {
				panic(err)
			}
		}(filename, &wg)

		files = append(files, filename)
		idx++
	}

	wg.Wait()

	return files, nil
}

func createChunkFile(filename string, bytes types.Bytes) error {
	w, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer w.Close()

	for _, b := range bytes {
		_, err = w.Write(b)
		if err != nil {
			return err
		}

		_, err = w.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}

	return nil
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
