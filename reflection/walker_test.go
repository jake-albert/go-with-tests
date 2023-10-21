package main

import "testing"

func TestWalk(t *testing.T) {
	expected := "Joey"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function clls, got %d want %d", len(got), 1)
	}
}
