package main // Dependency Injection
import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	//fmt.Printf("Buffer : %s \n", &buffer)
	received := buffer.String()
	expected := "Hello, Chris"

	if received != expected {
		t.Errorf("received %q expected %q", received, expected)
	}
}
