package goexpect

import (
	"errors"
	"strings"
	"testing"
)

func TestItExpects(t *testing.T) {
	reporter := newFakeReporter()
	it := NewIt(reporter)
	it.Expects(1, ToBe, 1)

	if reporter.DidFailNow {
		t.Fatalf("Failed to support equality it.Expects(1, ToBe, 1)")
	}

	it.Expects(1, ToBe, 2)
	if !reporter.DidFailNow {
		t.Fatalf("Failed to support equality it.Expects(1, ToBe, 2)")
	}
	assertStartsWith(t, reporter.Messages[0], "expected 1 (int) to be 2 (int)", "it.Expects(1, ToBe, 2)")

	if strings.Contains(reporter.Messages[0], "expectations.go") {
		t.Fatalf("it.Expects() should not give a stacktrace to expectations.go; got: %v", reporter.Messages[0])
	}
}

func TestItFails(t *testing.T) {
	reporter := newFakeReporter()
	it := NewIt(reporter)
	it.Fails("to see blue")

	if !reporter.DidFailNow {
		t.Fatalf("Failed to support it.Fails(\"because stuff is required\")")
	}
	assertStartsWith(t, reporter.Messages[0], "Failed to see blue:", "Failure message")
	if strings.Contains(reporter.Messages[0], "expectations.go") {
		t.Fatalf("it.Fails() should not give a stacktrace to expectations.go; got: %v", reporter.Messages[0])
	}
}

func TestToBeShouldPerformSimpleEquality(t *testing.T) {
	msg, ok := ToBe(1, 1)
	if !ok {
		t.Fatalf("ToBe should be ok for 1 == 1")
	}
	if msg != "to be 1 (int)" {
		t.Fatalf("ToBe should return message of 'to be 1', got '%s'", msg)
	}

	msg, ok = ToBe(1, 2)
	if ok {
		t.Fatalf("ToBe should not be ok for 1 == 2")
	}
	if msg != "to be 2 (int)" {
		t.Fatalf("ToBe should return message of 'to be 2', got '%s'", msg)
	}
}

func TestExpectHasThirdArgAsFunc(t *testing.T) {
	reporter := newFakeReporter()
	Expect(reporter, 1, 1)

	if !reporter.DidFailNow {
		t.Fatalf("Expect requires third argument to be a function")
	}

	assertStartsWith(t, reporter.Messages[0], "Expect() requires 3rd argument to be a matcher func", "Expect")
}

func TestExpectWithToBe(t *testing.T) {
	reporter := newFakeReporter()
	Expect(reporter, 1, ToBe, 1)

	if reporter.DidFailNow {
		t.Fatalf("Failed to support equality Expect(t, 1, ToBe, 1)")
	}

	Expect(reporter, 1, ToBe, 2)
	if !reporter.DidFailNow {
		t.Fatalf("Failed to support equality Expect(t, 1, ToBe, 2)")
	}

	assertStartsWith(t, reporter.Messages[0], "expected 1 (int) to be 2 (int)", "Expect(t, 1, ToBe, 2)")
	if strings.Contains(reporter.Messages[0], "expectations.go") {
		t.Fatalf("Expect() should not give a stacktrace to expectations.go; got: %v", reporter.Messages[0])
	}
}

func TestMustShouldFailWhenNotNil(t *testing.T) {
	reporter := newFakeReporter()
	Must(reporter, nil)
	if reporter.DidFailNow {
		t.Fatalf("Reporter failed when it got nil")
	}

	value := errors.New("An Error")
	Must(reporter, value)
	if !reporter.DidFailNow {
		t.Fatalf("Failed to report non-nil error: %#v", value)
	}

	assertStartsWith(t, reporter.Messages[0], "Must failed: got", "Failure message")
	if strings.Contains(reporter.Messages[0], "expectations.go") {
		t.Fatalf("Fail() should not give a stacktrace to expectations.go; got: %v", reporter.Messages[0])
	}
}

func TestFailShouldAlwaysFail(t *testing.T) {
	reporter := newFakeReporter()
	Fail(reporter, "as intended")

	if !reporter.DidFailNow {
		t.Fatalf("Failed to fail!")
	}

	assertStartsWith(t, reporter.Messages[0], "Failed as intended:", "Failure message")
}
