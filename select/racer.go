package main

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

// [NOTE] struct{} is smaller even than bool in Go and seems to require no memory allocation
// See: https://dave.cheney.net/2014/03/25/the-empty-struct
func ping(url string) chan struct{} {
	// [NOTE] Creating a channel with `var ch chan struct{}` leads to the "zero" value, nil.
	// See: https://go.dev/play/p/IIbeAox5jKA
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
