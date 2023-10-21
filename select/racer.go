package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondtimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondtimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
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
