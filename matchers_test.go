package goexpect

import (
	"testing"
)

func TestNotDecoratorShouldInvertMatcher(t *testing.T) {
	msg, ok := Not(ToBeTrue)(true)
	if ok {
		t.Fatalf("Not(ToBeTrue) should be ok for true")
	}
	assertStringEquals(t, msg, "not to be true", "Not(ToBeTrue)")

	_, ok = Not(ToBeTrue)(false)
	if !ok {
		t.Fatalf("Not(ToBeTrue) should be ok for false")
	}
}

func TestToBeTrueShouldCheckForFalse(t *testing.T) {
	ptr := false
	msg, ok := ToBeTrue(ptr)
	if ok {
		t.Fatalf("ToBeTrue should not be ok for %#v", ptr)
	}
	assertStringEquals(t, msg, "to be true", "ToBeTrue")

	ptr = true
	msg, ok = ToBeTrue(ptr)
	if !ok {
		t.Fatalf("ToBeTrue should be ok for %#v", ptr)
	}
}

func TestToBeFalseShouldCheckForFalse(t *testing.T) {
	ptr := true
	msg, ok := ToBeFalse(ptr)
	if ok {
		t.Fatalf("ToBeFalse should not be ok for %#v", ptr)
	}
	assertStringEquals(t, msg, "to be false", "ToBeFalse")

	ptr = false
	msg, ok = ToBeFalse(ptr)
	if !ok {
		t.Fatalf("ToBeFalse should be ok for %#v", ptr)
	}
}

func TestToBeNilShouldCheckForNil(t *testing.T) {
	ptr := new(struct{})
	msg, ok := ToBeNil(ptr)
	if ok {
		t.Fatalf("ToBeNil should not be ok for %#v", ptr)
	}
	assertStringEquals(t, msg, "to be nil", "ToBeNil")

	ptr = nil
	msg, ok = ToBeNil(ptr)
	if !ok {
		t.Fatalf("ToBeNil should be ok for %#v", ptr)
	}
}

func TestToBeLengthOfShouldCheckLen(t *testing.T) {
	filledArray := []string{"Foo"}
	msg, ok := ToBeLengthOf(filledArray, 0)
	if ok {
		t.Fatalf("ToBeLengthOf should not be ok for (%#v, 0)", filledArray)
	}
	assertStringEquals(t, msg, "to be length of 0, got (size: 1; value: []string{\"Foo\"})", "ToBeLengthOf")
}

func TestToBeEmptyShouldCheckEmptinessInCollections(t *testing.T) {
	filledArray := []string{"Foo"}
	emptyArray := []string{}
	msg, ok := ToBeEmpty(filledArray)
	if ok {
		t.Fatalf("ToBeEmpty should not be ok for %#v", filledArray)
	}
	assertStringEquals(t, msg, "to be empty (size: 1; value: []string{\"Foo\"})", "ToBeEmpty")

	msg, ok = ToBeEmpty(emptyArray)
	if !ok {
		t.Fatalf("ToBeEmpty should be ok for %#v", filledArray)
	}

	hash := make(map[string]string)
	msg, ok = ToBeEmpty(hash)
	if !ok {
		t.Fatalf("ToBeEmpty should be ok for %#v (size %d)", hash, len(hash))
	}
}
