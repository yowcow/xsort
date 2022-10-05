package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/sync/errgroup"
)

var size int
var count int

func init() {
	flag.IntVar(&size, "size", 4, "generating row size in bytes")
	flag.IntVar(&count, "count", 100000, "generating row count")
	flag.Parse()
}

const (
	maxChunkSize = 10000
	workers      = 8
)

func main() {
	chunks := makeChunks(count)

	var inEg errgroup.Group
	var outEg errgroup.Group
	inCh := make(chan int)
	outCh := make(chan []byte)

	outEg.Go(func() error {
		for out := range outCh {
			if _, err := fmt.Fprintf(os.Stdout, "%x\n", string(out)); err != nil {
				return err
			}
		}
		return nil
	})

	for i := 0; i < workers; i++ {
		inEg.Go(generateBytes(inCh, outCh, size))
	}

	for _, chunk := range chunks {
		inCh <- chunk
	}

	close(inCh)
	if err := inEg.Wait(); err != nil {
		panic(err)
	}

	close(outCh)
	if err := outEg.Wait(); err != nil {
		panic(err)
	}
}

func generateBytes(inCh <-chan int, outCh chan<- []byte, s int) func() error {
	return func() error {
		fi, err := os.Open("/dev/urandom")
		if err != nil {
			panic(err)
		}
		defer fi.Close()

		tmp := make([]byte, s)
		for count := range inCh {
			for i := 0; i < count; i++ {
				if _, err := fi.Read(tmp); err != nil {
					return err
				}
				b := make([]byte, len(tmp))
				copy(b, tmp)
				outCh <- b
			}
		}

		return nil
	}
}

func makeChunks(count int) []int {
	var chunks []int

	for {
		chunk := maxChunkSize
		if count < chunk {
			chunk = count
		}
		if chunk == 0 {
			break
		}

		chunks = append(chunks, chunk)
		count -= chunk
	}

	return chunks
}
