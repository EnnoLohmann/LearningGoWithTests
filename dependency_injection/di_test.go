package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	actual := buffer.String()
	expected := "Hello, Chris"

	if actual != expected {
		t.Errorf("got %q want %q", actual, expected)
	}
}
