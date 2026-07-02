package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Taj")
	want := "Hello, Taj!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
