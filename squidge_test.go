package main

import "testing"

func TestSquidge(t *testing.T) {
	got := Squidge()
	want := "Squidge!"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
