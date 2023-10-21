package main

import (
	"bytes"
	"reflect"
	"testing"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const (
	write = "write"
	sleep = "sleep"
)

func TestCounter(t *testing.T) {
	t.Run("writes correct output", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spyCountdownOperations := &SpyCountdownOperations{}

		Countdown(buffer, spyCountdownOperations)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("sleeps after all but final write", func(t *testing.T) {
		spyCountdownOperations := &SpyCountdownOperations{}

		Countdown(spyCountdownOperations, spyCountdownOperations)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spyCountdownOperations.Calls) {
			t.Errorf("got calls %v, want %v", spyCountdownOperations.Calls, want)
		}
	})
}
