package goexpect

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

var DefaultFailer = runtime.Goexit

type WriterReporter struct {
	Writer io.Writer
	Failer func()
}

func NewWriterReporter(w io.Writer) *WriterReporter {
	return &WriterReporter{w, DefaultFailer}
}

func (r *WriterReporter) WithFailer(f func()) *WriterReporter {
	r.Failer = f
	return r
}

func (r *WriterReporter) Logf(format string, values ...interface{}) {
	fmt.Fprintf(r.Writer, format, values...)
}

func (r *WriterReporter) FailNow() {
	r.Failer()
}

var StdoutReporter = NewWriterReporter(os.Stdout)
var StderrReporter = NewWriterReporter(os.Stderr)
