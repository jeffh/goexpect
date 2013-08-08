package goexpect

import (
	"errors"
	"testing"
)

func TestToBeShouldPerformSimpleEquality(t *testing.T) {
	msg, ok := ToBe(1, 1)
	if !ok {
		t.Fatalf("ToBe should be ok for 1 == 1")
	}
	if msg != "to be 1" {
		t.Fatalf("ToBe should return message of 'to be 1'")
	}

	msg, ok = ToBe(1, 2)
	if ok {
		t.Fatalf("ToBe should not be ok for 1 == 2")
	}
	if msg != "to be 2" {
		t.Fatalf("ToBe should return message of 'to be 2'")
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

	assertStartsWith(t, reporter.Messages[0], "expected 1 to be 2", "Expect(t, 1, ToBe, 2)")
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
}

func TestFailShouldAlwaysFail(t *testing.T) {
	reporter := newFakeReporter()
	Fail(reporter, "as intended")

	if !reporter.DidFailNow {
		t.Fatalf("Failed to fail!")
	}

	assertStartsWith(t, reporter.Messages[0], "Fail as intended:", "Failure message")
}
