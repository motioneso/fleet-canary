package main

import "testing"

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2, 3) = %d, want 5", got)
	}
}

func TestGreet(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"with name", "Ben", "Hello, Ben!"},
		{"no name", "", "Hello, world!"},
		{"spaces only", "   ", "Hello, world!"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Greet(c.in); got != c.want {
				t.Fatalf("Greet(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}

func TestFarewell(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"with name", "Ben", "Goodbye, Ben!"},
		{"no name", "", "Goodbye, world!"},
		{"spaces only", "   ", "Goodbye, world!"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Farewell(c.in); got != c.want {
				t.Fatalf("Farewell(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
