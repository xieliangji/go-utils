package hello

import "testing"

func TestSayHello(t *testing.T) {
	helloWorld := SayHello("world")
	if helloWorld != "Hello world" {
		t.Errorf("Hello got %q, want Hello world", helloWorld)
	}
}
