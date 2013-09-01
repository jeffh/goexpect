package goexpect

import (
	. "github.com/jeffh/goexpect/utils"
	"reflect"
)

// The interface that Expect requires for its first argument. It is used to report
// the results of the expectation.
type Reporter interface {
	FailNow()
	Logf(format string, args ...interface{})
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
	var argValues []reflect.Value

	argValues = AppendValueFor(argValues, obj)
	for _, v := range args {
		argValues = AppendValueFor(argValues, v)
	}

	testfn := reflect.ValueOf(test)
	if testfn.Kind() != reflect.Func {
		stacktrace := Tabulate("Stacktrace: ", GetStackTrace(3), "\n")
		r.Logf("Expect() requires 3rd argument to be a matcher func\n       got: %#v\n\n%s", test, stacktrace)
		r.FailNow()
		return
	}

	returnValues := testfn.Call(argValues)
	str, ok := returnValues[0].String(), returnValues[1].Bool()
	if !ok {
		stacktrace := Tabulate(" stacktrace: ", GetStackTrace(3), "\n")
		r.Logf("expected %s %s\n%s", ValueAsString(obj), str, stacktrace)
		r.FailNow()
	}
}

// Fails the test immediately with the given message.
//
// Example:
//
//    Fail(t, "every time")
//
func Fail(r Reporter, message string) {
	stacktrace := Tabulate(" stacktrace: ", GetStackTrace(3), "\n")
	r.Logf("Fail %s:\n%s", message, stacktrace)
	r.FailNow()
}

// Shorthand for Expect(t, value, ToBeNil)
// Used to assert that errors are nil.
//
// Example:
//
//    var foo interface{} = nil
//    Must(t, foo)
//
func Must(r Reporter, value error) {
	if value != nil {
		stacktrace := Tabulate(" stacktrace: ", GetStackTrace(3), "\n")
		r.Logf("Must failed: got %s\n%s", value, stacktrace)
		r.FailNow()
	}
}
