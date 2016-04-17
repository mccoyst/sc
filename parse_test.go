package main

import (
	"testing"
)

func TestEmptyParse(t *testing.T) {
	parts := parseLine(nil)
	if len(parts) != 0 {
		t.Fatalf("expected nothing, got something: %v", parts)
	}
}

func TestSimpleParse(t *testing.T) {
	parts := parseLine([]byte("hello there"))
	if len(parts) != 2 {
		t.Fatalf("expected two parts, got something else: %v", parts)
	}
	for i, ex := range []string{ "hello", "there" } {
		if parts[i] != ex {
			t.Fatalf("expected %q, got %q", ex, parts[i])
		}
	}
}
