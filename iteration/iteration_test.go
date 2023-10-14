package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("returns empty string when length is 0", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""
		assertCorrect(t, repeated, expected)
	})
	t.Run("repeats given string five times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrect(t, repeated, expected)
	})
}

func assertCorrect(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q but got %q", got, want)
	}
}

// go test -bench="."
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5000)
	}
}
