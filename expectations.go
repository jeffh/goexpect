package goexpect

import (
	. "github.com/jeffh/goexpect/utils"
	"reflect"
)

const (
	itStackOffset         = 1 // to exclude the method that calls GetStackTrace
	functionalStackOffset = 2 // +1 more since functional wraps it methods
)

// The interface that Expect requires for its first argument. It is used to report
// the results of the expectation.
type Reporter interface {
	FailNow()
	Logf(format string, args ...interface{})
}

// Wrapper around Expectations that encloses the reporter
//
// This is a more object-oriented approach to expectations:
//
// - Instead of Expect(t, 1, ToBe, 1) => it.Expects(1, ToBe, 1)
// - Instead of Fail(t, "to parse") => it.Fails("to parse")
// - Instead of Must(t, Parse(...)) => it.Must(Parse(...))
//
// This also is the first step to custom test harnesses, which
// has not been implemented yet.
type It struct {
	r                   Reporter
	stackOffset         int
	missingMatcherError string
}

// Creates a new wrapper around expectations to a given reporter
func NewIt(r Reporter) *It {
	return &It{
		r,
		itStackOffset,
		"it.Expects() requires 2nd argument to be a matcher func",
	}
}

func (i *It) withStackOffset(offset int) *It {
	return &It{i.r, offset, i.missingMatcherError}
}

// Idential to Expect(t, obj, test, args...), where t is the reporter
// given to NewIt(t).
func (i *It) Expects(obj interface{}, test interface{}, args ...interface{}) {
	var argValues []reflect.Value

	argValues = AppendValueFor(argValues, obj)
	for _, v := range args {
		argValues = AppendValueFor(argValues, v)
	}

	testfn := reflect.ValueOf(test)
	if testfn.Kind() != reflect.Func {
		stacktrace := Tabulate("Stacktrace: ", GetStackTrace(i.stackOffset), "\n")
		i.r.Logf("%s\n       got: %#v\n\n%s", i.missingMatcherError, test, stacktrace)
		i.r.FailNow()
		return
	}

	returnValues := testfn.Call(argValues)
	str, ok := returnValues[0].String(), returnValues[1].Bool()
	if !ok {
		stacktrace := Tabulate(" stacktrace: ", GetStackTrace(i.stackOffset), "\n")
		i.r.Logf("expected %s %s\n%s", ValueAsString(obj), str, stacktrace)
		i.r.FailNow()
	}
}

// Idential to Fail(t, msg), where t is the reporter given to NewIt(t).
func (i *It) Fails(message string) {
	stacktrace := Tabulate(" stacktrace: ", GetStackTrace(i.stackOffset), "\n")
	i.r.Logf("Failed %s:\n%s", message, stacktrace)
	i.r.FailNow()
}

// Idential to Must(t, err), where t is the reporter given to NewIt(t).
func (i *It) Must(err error) {
	if err != nil {
		stacktrace := Tabulate(" stacktrace: ", GetStackTrace(i.stackOffset), "\n")
		i.r.Logf("Must failed: got %s\n%s", err, stacktrace)
		i.r.FailNow()
	}
}

// Performs an expectation. It takes a Reporter (which testing.T satisfies), followed
// by the value under test, then a matcher. Any additional arguments, after that
// are passed directly to the matcher. Certain matches may require more arguments.
//
// Example:
//
//    func TestBoolean(t *testing.T) {
//      Expect(t, true, ToBeTrue)
//      Expect(t, 1, ToEqual, 1)
//    }
//
func Expect(r Reporter, obj interface{}, test interface{}, args ...interface{}) {
	it := NewIt(r).withStackOffset(functionalStackOffset)
	it.missingMatcherError = "Expect() requires 3rd argument to be a matcher func"
	it.Expects(obj, test, args...)
}

// Fails the test immediately with the given message.
//
// Example:
//
//    Fail(t, "every time")
//
func Fail(r Reporter, message string) {
	NewIt(r).withStackOffset(functionalStackOffset).Fails(message)
}

// Shorthand for Expect(t, value, ToBeNil)
// Used to assert that errors are nil.
//
// Example:
//
//    Must(t, MakeWriteToSocket(...))
//
func Must(r Reporter, value error) {
	NewIt(r).withStackOffset(functionalStackOffset).Must(value)
}
