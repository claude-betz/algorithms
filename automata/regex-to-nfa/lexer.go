package main

import (
	"errors"
	"fmt"
	"unicode"
)

var (
	tagError = errors.New("can not tag")
	errEOF   = errors.New("end of file")
)

type tagval struct {
	Tag   Tag
	Value string
}

type Lexer struct {
	input []rune
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: []rune(input)}
}

func (l *Lexer) ReadToken() (*Token, error) {
	n, err := skipWhitespace(l.input[l.pos:])
	if err != nil {
		if err == errEOF {
			return eof, nil
		}
		return nil, err
	}
	l.pos += n

	taggers := []func(input []rune) (*tagval, error){
		readPunctuation,
		readId,
	}

	for _, tagger := range taggers {
		tv, err := tagger(l.input[l.pos:])
		if err == nil {
			tk := &Token{l.pos, tv.Tag, tv.Value}
			l.pos++
			return tk, nil
		}
		if err != tagError {
			panic(fmt.Sprintf("unknown error: %s", err))
		}
	}
	return nil, fmt.Errorf("cannot recognise rune: %c", l.input[l.pos])
}

func (l *Lexer) UnreadToken(t *Token) {
	l.pos -= len(t.Value)
}

func skipWhitespace(input []rune) (int, error) {
	for n := 0; n < len(input); n++ {
		if !unicode.IsSpace(input[n]) {
			return n, nil
		}
	}
	return 0, errEOF
}

func readPunctuation(input []rune) (*tagval, error) {
	r := input[0]
	if isPunctionSymbol[r] {
		return &tagval{TagPunct, fmt.Sprintf("%c", r)}, nil
	}
	return nil, tagError
}

func readId(input []rune) (*tagval, error) {
	r := input[0]
	if !unicode.IsLetter(input[0]) {
		return nil, tagError
	}
	return &tagval{TagId, string(r)}, nil
}
