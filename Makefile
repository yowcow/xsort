ARTIFACTS := _examples/unsorted.csv

all: $(ARTIFACTS)

_examples/unsorted.csv: _examples
	seq -f '%04g' 1 1000 | shuf > $@

_examples:
	mkdir -p $@

clean:
	rm -rf $(ARTIFACTS)

.PHONY: all clean
