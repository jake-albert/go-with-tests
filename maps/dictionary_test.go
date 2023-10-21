package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"

		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("known word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word"

		assertError(t, err, ErrNotFound)
		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	definition := "this is just a test"
	word := "test"

	dictionary.Add(word, "this is just a test")

	assertDefinition(t, dictionary, word, definition)
}

func assertDefinition(t *testing.T, dictionary Dictionary, word string, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got != want {
		t.Errorf("got error %q, wanted error %q", got.Error(), want)
	}
}