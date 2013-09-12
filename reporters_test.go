package goexpect

import (
	"bytes"
	"testing"
)

type FakeFailer struct {
	DidFail bool
}

func (f *FakeFailer) Fail() {
	f.DidFail = true
}

func TestWriterReporterCanReport(t *testing.T) {
	f := &FakeFailer{}
	b := bytes.NewBufferString("")
	r := NewWriterReporter(b).WithFailer(f.Fail)
	r.Logf("Something: %v", 1)

	if b.String() != "Something: 1" {
		t.Fatalf("Failed to log error")
	}

	r.FailNow()
	if !f.DidFail {
		t.Fatalf("Log should fail")
	}
}
