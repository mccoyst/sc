package main

import (
	"bytes"
	"io"
	"unicode"
)

func parseLine(line []byte) []string {
	var parts []string
	l := newLexer(bytes.NewReader(line))

	err := consumeSpaces(l)
	if err == io.EOF {
		return parts
	}
	if err != nil {
		panic("Oops, I need to return an error here!")
	}

	for {
		s, err := consumeString(l)
		if err == io.EOF {
			parts = append(parts, s)
			return parts
		}
		if err != nil {
			panic("Oops, I need to return an error here!")
		}
		parts = append(parts, s)

		err = consumeSpaces(l)
		if err == io.EOF {
			return parts
		}
		if err != nil {
			panic("Oops, I need to return an error here!")
		}
	}

	return parts
}

func consumeSpaces(l lexer) error {
	for {
		r, err := l.Next()
		if err != nil {
			return err
		}

		if !unicode.IsSpace(r) {
			l.Undo()
			break
		}
	}
	return nil
}

		
func consumeString(l lexer) (string, error) {
	var b bytes.Buffer
	for {
		r, err := l.Next()
		if err != nil {
			return b.String(), err
		}

		if unicode.IsSpace(r) {
			l.Undo()
			break
		}

		b.WriteRune(r)
	}

	return b.String(), nil
}