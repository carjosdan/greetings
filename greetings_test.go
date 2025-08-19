package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Alice"
	want := regexp.MustCompile(`\b` + name + `\b`)
	got, err := Hello("Alice")
	if err != nil {
		t.Fatalf("Hello(%q) returned an error: %v", name, err)
	}
	if !want.MatchString(got) {
		t.Fatalf("Hello(%q) = %q; want %#q", name, got, want)
	}
}

func TestHelloEmptyName(t *testing.T) {
	got, err := Hello("")
	if err == nil {
		t.Fatalf("Hello(%q) = %q; want an error", got, err)
	}
}
