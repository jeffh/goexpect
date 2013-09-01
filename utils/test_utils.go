package utils

import (
	"fmt"
	"strings"
	"testing"
)

type fakeReporter struct {
	DidFailNow bool
	Messages   []string
}

func newFakeReporter() *fakeReporter {
	return &fakeReporter{Messages: make([]string, 0)}
}

func (r *fakeReporter) FailNow() {
	r.DidFailNow = true
}

func (r *fakeReporter) Logf(format string, args ...interface{}) {
	r.Messages = append(r.Messages, fmt.Sprintf(format, args...))
}

func assertStartsWith(t *testing.T, input string, prefix string, noun string) {
	if !strings.HasPrefix(input, prefix) {
		t.Fatalf("%s should have %#v message, got %#v", noun, input, prefix)
	}
}

func assertStringEquals(t *testing.T, input string, output string, noun string) {
	if input != output {
		t.Fatalf("%s should have %#v message, got %#v", noun, input, output)
	}
}
