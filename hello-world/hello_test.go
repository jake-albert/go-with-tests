package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("says hello to a person", func(t *testing.T) {
		got := Hello("Jake", "")
		want := "Hello, Jake"
		assertCorrectMessage(t, got, want)
	})
	t.Run("says 'Hello, World' when an given an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("says hello in Korean", func(t *testing.T) {
		got := Hello("Jake", "Korean")
		want := "안녕, Jake"
		assertCorrectMessage(t, got, want)
	})
	t.Run("says hello in Georgian", func(t *testing.T) {
		got := Hello("Jake", "Georgian")
		want := "გამარჯობა, Jake"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
