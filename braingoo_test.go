package braingoo

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	source := "+[-->-[>>+>-----<<]<--<---]>-.>>>+.>>..+++[.>]<<<<.+++.------.<<-.>>>>+."
	testProgram := NewInterpreter(source, []int8{})
	testProgram.parse()
	output := testProgram.output

	if output != expected {
		t.Errorf("Expected output '%s' got '%s'", expected, output)
	}
}

func TestInput(t *testing.T) {
	expected := "abc"
	input := []int8{'a', 'b', 'c'}
	source := ",.,.,."
	testProgram := NewInterpreter(source, input)
	testProgram.parse()
	output := testProgram.output

	if output != expected {
		t.Errorf("Expected output '%s' got '%s'", expected, output)
	}
}
