GoExpect
----------

A simple assertion library that provides BDD-styled expectations.

It's pretty simple:

    go get github.com/jeffh/goexpect

Then use it in your tests:

    import (
        "testing"
        . "github.com/jeffh/goexpect"
    )

    func TestFoo(t *testing.T) {
        Expect(t, 1, ToEqual, 1)
        Expect(t, true, ToBeTrue)
    }

Use the `Expect` function to perform assertions on the codebase, followed by
the value in question and a matcher. Additional arguments are passed to `Expect` after the matcher.

Two other functions which may be interesting:

 - `Must(*testing.T, error)` checks if error is nil. Shorthand to Expect(t, err, ToBeNil).
 - `Fail(*testing.T, string)` always fails a test.

Matchers
------------

The following matches are available:

 - `ToBeTrue` checks if the given value is true. Takes no extra arguments.
 - `ToBeFalse` checks if the given value is false. Takes no extra arguments
 - `ToBeNil` checks if the given value is nil. Takes no extra arguments.
 - `ToBeEmpty` checks if the given value has the length of zero. Takes no extra arguments.
 - `ToBeLengthOf` checks if the given value has the length of a given value. Takes an integer as the expected length.
 - `ToBe` checks if the given value is equal. Performs simple equality. Takes the expected value as an additional argument.
 - `ToEqual` checks if the given value is equal. Performs deep equality to match structs, maps and arrays. Takes the expected value as the expected length.
 - `Not` is a matcher generator that takes an existing match and returns the negated version.

Matchers that take additional arguments (such as `ToEqual`) are passed to `Expect` like so:

    Expect(t, 1, ToEqual, 1)

`Not` is a function that returns a *new matcher*. So it takes the matcher as direct invocation:

    Expect(t, 1, Not(ToEqual), 1)

Writing Custom Matchers
-----------------------

A matcher is a function that has one argument plus any additional argument it chooses to have:

    func(actual interface{}) (string, bool)

Where `actual` is the actual value that `Expect` received. Any value additional parameters
can be specified. The return values are the error message string and the bool indicating if
the assertion has passed. The error message string should always be returned for decorators,
such as `Not` to modify the message string.


Also, typed values can be provided if you expect certain types:

    func(haystack, needle string) (string, bool)

Including the `utils` subpackage provides a unified way to output the value arguments:

    import (
        "fmt"
        "github.com/jefh/goexpect/utils"
        "reflect"
    )

    func ToEqual(actual, expected interface{}) (string, bool) {
        return fmt.Sprintf("to equal %s", utils.ValueAsString(expected)), reflect.DeepEqual(actual, expected)
    }

`utils.ValuesAsString` provides the actual type, which is useful for trouble shooting two
value that are the same underlying type.

The `utils` package also includes a simple way to append arbitrary values to a reflect.Value slice:

    func Not(test interface{}) func(actual interface{}, args ...interface{}) (string, bool) {
        return func(actual interface{}, args ...interface{}) (string, bool) {
            var argValues []reflect.Value

            argValues = utils.AppendValueFor(argValues, actual)
            for _, v := range args {
                argValues = utils.AppendValueFor(argValues, v)
            }

            returnValues := reflect.ValueOf(test).Call(argValues)
            str, ok := returnValues[0].String(), returnValues[1].Bool()

            return fmt.Sprintf("not %s", str), !ok
        }
    }

Unlike a naive append() with reflect.ValueOf, `AppendValueFor` correctly marks nil values for
use in `reflect.Value.Call`.


Using Another Reporter
----------------------

`Expect` `Must` and `Fail` don't directly depend on `*testing.T`, but use an interface that requires the
following methods:

    Logf(format string, args...interface{})
    FailNow()

Where `Logf` is a 'Sprintf' styled format string and `FailNow` tells the test framework to abort the currently
running code.
