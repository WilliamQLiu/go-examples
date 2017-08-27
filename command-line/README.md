# cmd-example

Trying out Command Line functionality with Go.

## Build


    go build cmd-example.go

## Run

Run with:

* Arguments
* Flags
* Environment Variables

### Run with Arguments

Example

    ./cmd-example <args>
    ./cmd-example a b c d e

### Run with Flags

When run with arguments and flags, then flags must appear BEFORE positional arguments. Use `-h` or `--help` flags to get automatically generated help text

Example

    ./cmd-example <flags> <args>
    ./cmd-example -word=opt -numb=7 -fork -svar=flag arg1 arg2 arg3
    ./cmd-example -h
        Usage of ./cmd-example:
            -fork
                  a bool
            -numb int
                  an int (default 42)
            -svar string
                  a string var (default "bar")
            -word string
                  a string (default "foo")

### Environment Variables