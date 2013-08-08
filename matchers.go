package goexpect

import (
	"fmt"
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

		argValues = appendValueFor(argValues, actual)
		for _, v := range args {
			argValues = appendValueFor(argValues, v)
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
	return ToBe(actual, true)
}

// A matcher that expects the value to be false
//
// Example:
//
//    Expect(t, value, ToBeTrue)
//
func ToBeFalse(actual interface{}) (string, bool) {
	return ToBe(actual, false)
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
func ToBeLengthOf(actual interface{}, size int) (string, bool) {
	value := reflect.ValueOf(actual)
	msg := fmt.Sprintf("to be length of %d, got (size: %d; value: %#v)", size, value.Len(), actual)
	return msg, value.Len() == size
}

// Expects the given value to be have a length of zero
//
// Example:
//
//    Expect(t, value, ToBeEmpty)
//
func ToBeEmpty(actual interface{}) (string, bool) {
	value := reflect.ValueOf(actual)
	msg := fmt.Sprintf("to be empty (size: %d; value: %#v)", value.Len(), actual)
	return msg, value.Len() == 0
}

// Performs a simple equality comparison. Does not perform a deep equality.
//
// Example:
//
//    Expect(t, value, ToBe, "foo")
//
func ToBe(actual, expected interface{}) (string, bool) {
	return fmt.Sprintf("to be %#v", expected), actual == expected
}

// Performs a deep equal - comparing struct, arrays, and slices items too.
//
// Example:
//
//    Expect(t, value, ToEqual, []string{"Foo", "Bar"})
//
func ToEqual(actual, expected interface{}) (string, bool) {
	return fmt.Sprintf("to deeply equal %#v", expected), reflect.DeepEqual(actual, expected)
}
