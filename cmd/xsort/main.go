package main

import (
	"flag"
	"os"
	"strconv"
	"strings"

	"github.com/yowcow/xsort"
)

var inputFile string
var outputFile string
var chunkSize string
var tmpDir string

func init() {
	flag.StringVar(&inputFile, "i", "", "input file (default: STDIN)")
	flag.StringVar(&outputFile, "o", "", "output file (default: STDOUT)")
	flag.StringVar(&chunkSize, "c", "100M", "chunk size in mega bytes (default: 100M)")
	flag.StringVar(&tmpDir, "d", "", "tmp dir")
	flag.Parse()
}

func main() {
	var opt xsort.SortOptions

	if inputFile == "" {
		opt.Input = os.Stdin
	} else {
		f, err := os.Open(inputFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		opt.Input = f
	}

	if outputFile == "" {
		opt.Output = os.Stdout
	} else {
		f, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		opt.Output = f
	}

	if strings.HasSuffix(chunkSize, "M") {
		chunkSize = strings.TrimSuffix(chunkSize, "M")
		size, err := strconv.Atoi(chunkSize)
		if err != nil {
			panic(err)
		}
		opt.ChunkSize = int64(size * xsort.MiB)
	} else {
		size, err := strconv.Atoi(chunkSize)
		if err != nil {
			panic(err)
		}
		opt.ChunkSize = int64(size)
	}

	if tmpDir == "" {
		dir, err := os.MkdirTemp("", "xsort")
		if err != nil {
			panic(err)
		}
		defer os.RemoveAll(dir)
		opt.TmpDir = dir
	} else {
		opt.TmpDir = tmpDir
	}

	xsort.Sort(&opt)
}
