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

	"golang.org/x/sync/errgroup"
)

func CreateChunkFiles(r io.Reader, chunkSize int64, tmpDir string) ([]string, error) {
	base, err := basename(8)
	if err != nil {
		return nil, err
	}

	var idx int
	var files []string
	var eg errgroup.Group
	s := bufio.NewScanner(r)

	for {
		bytes, err := createChunkBytes(s, chunkSize)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}

		file := filepath.Join(tmpDir, fmt.Sprintf("%s-%d", base, idx))
		files = append(files, file)
		idx++

		eg.Go(func() error {
			if err := createChunkFile(file, bytes); err != nil {
				return err
			}
			return nil
		})
	}

	return files, eg.Wait()
}

func createChunkFile(filename string, bytes Bytes) error {
	w, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer w.Close()

	for _, b := range bytes {
		if _, err = w.Write(append(b, byte('\n'))); err != nil {
			return err
		}
	}

	return nil
}

func createChunkBytes(s *bufio.Scanner, chunkSize int64) (Bytes, error) {
	var curBuf [][]byte
	var curSize int64

	for s.Scan() {
		buf := allocBytes(s.Bytes())
		curBuf = append(curBuf, buf)
		curSize += int64(len(buf))

		if curSize >= chunkSize {
			break
		}
	}

	if curSize == 0 {
		return nil, io.EOF
	}

	sort.Sort(Bytes(curBuf))

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
