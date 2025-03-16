package main

import "testing"

func TestGreet(t *testing.T) {
	want := "Hello, Earth!"

	got := greet()

	if got != want {
		// Test failed
		t.Errorf("greet() = %q, want %q", got, want)
	} else {
		// Test passed
		t.Logf("greet() = %q, want %q", got, want)
	}
}
