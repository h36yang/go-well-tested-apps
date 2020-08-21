package messages

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Gopher")
	expected := "Hello, Gopher!\n"

	if got != expected {
		t.Errorf("Did not get expected result. Got: %q, expected: %q\n", got, expected)
	}
}

// A table driven test
func TestGreetTableDriven(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "Gopher", expect: "Hello, Gopher!\n"},
		{input: "", expect: "Hello, !\n"},
	}
	for _, s := range scenarios {
		got := Greet(s.input)
		if got != s.expect {
			t.Errorf("Did not get expected result for input %q. Got: %q, expected: %q\n", s.input, got, s.expect)
		}
	}
}

func TestDepart(t *testing.T) {
	got := depart("Gopher")
	expected := "Goodbye, Gopher\n"

	if got != expected {
		t.Errorf("Did not get expected result. Got: %q, expected: %q\n", got, expected)
	}
}
