package go_hello_world

import "testing"

func TestGreeter(t *testing.T) {
	expected := "Hello World!"
	actual := HelloWorld()
	if actual != expected {
		t.Error("expected %q but got %q", expected, actual)
	}
}