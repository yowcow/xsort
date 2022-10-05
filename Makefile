SIZE := 4
COUNT := 10000000

all:

test: gotest diff

gotest:
	go test ./...

diff: _testdata/sorted-sort.csv _testdata/sorted-xsort.csv
	diff -u $^

_testdata/unsorted.csv:
	mkdir -p $(dir $@)
	go run ./cmd/xgen/main.go -size $(SIZE) -count $(COUNT) > $@

_testdata/sorted-sort.csv: _testdata/unsorted.csv
	cat $< | sort > $@

_testdata/sorted-xsort.csv: _testdata/unsorted.csv FORCE
	cat $< | go run ./cmd/xsort/main.go > $@

clean:
	rm -rf _testdata/sorted-*

realclean: clean
	rm -rf _testdata/unsorted.csv

.PHONY: all test gotest diff clean realclean

FORCE:
