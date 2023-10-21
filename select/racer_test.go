package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns URL of the faster server", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)

		if err != nil {
			t.Fatalf("did not expect an error but go one: %v", err)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns an error if a server doesn't respons within 10s", func(t *testing.T) {
		server := makeDelayedServer(2 * time.Second)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return slowServer
}
