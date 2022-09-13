package main

import (
	"flag"
	"io"
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
	var r io.Reader
	var w io.Writer
	var c int
	var d string

	if inputFile == "" {
		r = os.Stdin
	} else {
		f, err := os.Open(inputFile)
		if err != nil {
			panic(err)
		}
		r = f
	}

	if outputFile == "" {
		w = os.Stdout
	} else {
		f, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		w = f
	}

	if strings.HasSuffix(chunkSize, "M") {
		chunkSize = strings.TrimSuffix(chunkSize, "M")
		size, err := strconv.Atoi(chunkSize)
		if err != nil {
			panic(err)
		}
		c = size * 1000000
	} else {
		size, err := strconv.Atoi(chunkSize)
		if err != nil {
			panic(err)
		}
		c = size
	}

	if tmpDir == "" {
		dir, err := os.MkdirTemp("", "xsort")
		if err != nil {
			panic(err)
		}
		defer os.RemoveAll(dir)
		d = dir
	} else {
		d = tmpDir
	}

	xsort.Sort(r, w, c, d)
}
