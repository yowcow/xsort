xsort
=====

Sort extremely large files with constant memory.

```
❯ xsort -h
Usage of xsort:
  -c string
        chunk size in mega bytes (default: 100M) (default "100M")
  -d string
        tmp dir
  -i string
        input file (default: STDIN)
  -o string
        output file (default: STDOUT)
```


HOW TO INSTALL
--------------

```
go install github.com/yowcow/xsort/cmd/...
```


HOW TO USE
----------

Sort a file:

```
xsort -i large-input.csv -o large-output.csv
```

Sort a file with a specified chunk size:

```
xsort -i large-input.csv -o large-output.csv -c 10M
```

Sort input from STDIN, and output to STDOUT:

```
cat large-input.csv | xsort > large-output.csv
```
