package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("teslazonda")
	want := "Hello, teslazonda"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
