package main_test

import (
	"testing"
)

// A passing test
func TestAddition(t *testing.T) {
	got := 2 + 2
	expected := 4
	if got != expected {
		t.Errorf("Did not get expected result. Got: '%v', expected: '%v'\n", got, expected)
	}
}

// A failing test
func TestSubtraction(t *testing.T) {
	got := 10 - 5
	expected := 4
	if got != expected {
		t.Errorf("Did not get expected result. Got: '%v', expected: '%v'\n", got, expected)
	}
}

func TestFailureTypes(t *testing.T) {
	t.Error("Error signals a failed test, but does not stop the rest of the tests from executing")
	t.Fatal("Fatal will fail the test and stop its execution")
	t.Error("This will never print since it is preceeded by an immediate failure")
}
