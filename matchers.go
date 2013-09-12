package goexpect

import (
	"fmt"
	. "github.com/jeffh/goexpect/utils"
	"reflect"
)

// A matcher generator that negates the given matcher.
// Unlike other matchers, this function directly accepts the matcher in question:
//
// Example:
//
//    Expect(t, "red", Not(ToEqual), "blue")
//
func Not(test interface{}) func(actual interface{}, args ...interface{}) (string, bool) {
	return func(actual interface{}, args ...interface{}) (string, bool) {
		var argValues []reflect.Value

		argValues = AppendValueFor(argValues, actual)
		for _, v := range args {
			argValues = AppendValueFor(argValues, v)
		}

		returnValues := reflect.ValueOf(test).Call(argValues)
		str, ok := returnValues[0].String(), returnValues[1].Bool()

		return fmt.Sprintf("not %s", str), !ok
	}
}

// A matcher that expects the value to be true
//
// Example:
//
//    Expect(t, value, ToBeTrue)
//
func ToBeTrue(actual interface{}) (string, bool) {
	return "to be true", actual == true
}

// A matcher that expects the value to be false
//
// Example:
//
//    Expect(t, value, ToBeTrue)
//
func ToBeFalse(actual interface{}) (string, bool) {
	return "to be false", actual == false
}

// A matcher that expects the value to be nil
//
// Example:
//
//    Expect(t, value, ToBeNil)
//
func ToBeNil(actual interface{}) (string, bool) {
	value := reflect.ValueOf(actual)
	return "to be nil", value.Kind() == reflect.Ptr && value.IsNil()
}

// Expects the given value to have a length of the provided value
//
// Example:
//
//    Expect(t, value, ToBeLengthOf, 5)
//
func ToBeLengthOf(actual interface{}, size int) (msg string, ok bool) {
	defer func() {
		err := recover()
		if err != nil {
			msg = fmt.Sprintf("to be length of %d, but has no length", size)
			ok = false
		}
	}()
	value := reflect.ValueOf(actual)
	ok = (value.Len() == size)
	msg = fmt.Sprintf("to be length of %d, got %d", size, value.Len())
	return
}

// Expects the given value to be have a length of zero
//
// Example:
//
//    Expect(t, value, ToBeEmpty)
//
func ToBeEmpty(actual interface{}) (msg string, ok bool) {
	defer func() {
		err := recover()
		if err != nil {
			msg = fmt.Sprintf("to be empty, but has no length")
			ok = false
		}
	}()
	value := reflect.ValueOf(actual)
	ok = (value.Len() == 0)
	msg = fmt.Sprintf("to be empty (%s; size: %d)", ValueAsString(actual), value.Len())
	return
}

// Performs a simple equality comparison. Does not perform a deep equality.
//
// Example:
//
//    Expect(t, value, ToBe, "foo")
//
func ToBe(actual, expected interface{}) (string, bool) {
	return fmt.Sprintf("to be %s", ValueAsString(expected)), actual == expected
}

// Performs a deep equal - comparing struct, arrays, and slices items too.
//
// Example:
//
//    Expect(t, value, ToEqual, []string{"Foo", "Bar"})
//
func ToEqual(actual, expected interface{}) (string, bool) {
	return fmt.Sprintf("to equal %s", ValueAsString(expected)), reflect.DeepEqual(actual, expected)
}
