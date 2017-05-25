package hello

import (
	"testing"
)

func TestHelloFunction(t *testing.T) {
	logic := Logic{}
	hello := logic.helloFunction("John")

	if hello.Name != "John" {
		t.Errorf("got %s, want John", hello.Name)
	}
}
