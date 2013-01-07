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

    func(actual interface{})

Where `actual` is the actual value that `Expect` received. Any value parameters are 