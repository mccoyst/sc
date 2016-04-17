package main

import (
	"bytes"
	"io"
	"testing"
)

func TestEmptyLex(t *testing.T) {
	l := newLexer(bytes.NewReader(nil))
	r, err := l.Next()
	if err != io.EOF {
		t.Fatalf("expected EOF, got %c, %v", r, err)
	}
}

func TestHello(t *testing.T) {
	l := newLexer(bytes.NewReader([]byte("hello")))

	for _, ex := range []rune("hello") {
		r, err := l.Next()
		if err != nil {
			t.Fatalf("expected no error, got one: %v", err)
		}
		if r != ex {
			t.Fatalf("expected %c, got %c", ex, r)
		}
	}

	r, err := l.Next()
	if err != io.EOF {
		t.Fatalf("expected EOF, got %c, %v", r, err)
	}
}

func TestLiteralEscapes(t *testing.T) {
	l := newLexer(bytes.NewReader([]byte(`\h\e\l\l\o`)))

	for _, ex := range []rune("hello") {
		r, err := l.Next()
		if err != nil {
			t.Fatalf("expected no error, got one: %v", err)
		}
		if r != ex {
			t.Fatalf("expected %c, got %c", ex, r)
		}
	}

	r, err := l.Next()
	if err != io.EOF {
		t.Fatalf("expected EOF, got %c, %v", r, err)
	}
}

func TestMetaEscapes(t *testing.T) {
	l := newLexer(bytes.NewReader([]byte(`\h\e\n\t`)))

	for _, ex := range []rune("he\n\t") {
		r, err := l.Next()
		if err != nil {
			t.Fatalf("expected no error, got one: %v", err)
		}
		if r != ex {
			t.Fatalf("expected %c, got %c", ex, r)
		}
	}

	r, err := l.Next()
	if err != io.EOF {
		t.Fatalf("expected EOF, got %c, %v", r, err)
	}
}
