package firstmod

import (
	"testing"
)

// test funcs
func TestHello(t *testing.T) {
	want := "Hello"
	if got := sayHello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestHi(t *testing.T) {
	want := "hi :)"
	if got := sayHi(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
