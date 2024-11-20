package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("teslazonda")
		want := "Hello, teslazonda"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty sting defaults to 'world'", func(t *testing.T) {

		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
