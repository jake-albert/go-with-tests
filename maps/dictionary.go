package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// [NOTE] Maps values are pointers, so when passing by value the pointer is copied.
// https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it
func (d Dictionary) Add(word string, definition string) {
	d[word] = definition
}
