package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jake")
	want := "Hello, Jake"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
