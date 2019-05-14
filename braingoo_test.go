package braingoo

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	source := "+[-->-[>>+>-----<<]<--<---]>-.>>>+.>>..+++[.>]<<<<.+++.------.<<-.>>>>+."
	testProgram := NewInterpreter(source)
	testProgram.parse()
	output := testProgram.output

	if output != expected {
		t.Errorf("Expected output '%s' got '%s'", expected, output)
	}
}
