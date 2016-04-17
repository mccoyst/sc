package main

import (
	"bytes"
)

var meta = map[rune]rune{
	'n': '\n',
	't': '\t',
}

type lexer struct {
	r *bytes.Reader
}

func newLexer(r *bytes.Reader) lexer {
	return lexer{r}
}

func (l *lexer) Next() (rune, error) {
	r, _, err := l.r.ReadRune()
	if err != nil {
		return 0, err
	}

	if r == '\\' {
		r, _, err = l.r.ReadRune()
		if err != nil {
			return 0, err
		}
		m := meta[r]
		if m != 0 {
			r = m
		}
	}

	return r, nil
}

func (l *lexer) Undo() {
	l.r.UnreadRune()
}
