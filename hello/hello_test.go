package hello

import "testing"

func TestSayHello(t *testing.T) {
	helloGo := SayHello("world")
	if helloGo != "Hello world" {
		t.Errorf("Hello got %q, want Hello world", helloGo)
	}
}
