package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

var (
	in  = flag.String("i", "-", "path to the pipeline")
	out = flag.String("output", "-", "path to dot destination")
)

func main() {
	flag.Parse()

	r, err := reader(*in)
	if err != nil {
		panic(err)
	}

	w, err := writer(*out)
	if err != nil {
		panic(err)
	}

	p, err := Parse(r)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprintf(w, ToDot(ToDigraph(p)))
	if err != nil {
		panic(err)
	}
}

func writer(value string) (w io.Writer, err error) {
	var file *os.File

	if value == "-" {
		w = os.Stdout
		return
	}

	file, err = os.Create(value)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to create file %s", value)
		return
	}

	w = file
	return
}

func reader(value string) (r io.Reader, err error) {
	var file *os.File

	if value == "-" {
		r = os.Stdin
		return
	}

	file, err = os.Open(value)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to open dpkg status file at %s", value)
		return
	}

	r = file
	return
}
