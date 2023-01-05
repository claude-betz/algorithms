package main

import (
	"fmt"
)

type Tag string

const (
	TagPunct Tag = "punctuation"
	TagId    Tag = "identifier"
)

const EOF string = "!"

var eof *Token = &Token{Tag: TagPunct, Value: EOF}

type Token struct {
	position int
	Tag
	Value string
}

func (tk Token) String() string {
	return fmt.Sprintf("{%q}", tk.Value)
}

var punctuationSymbol = []rune{
	'(', // left parenthesis
	')', // right parenthesis
	'|', // union
	'.', // concatenation
	'*', // closure
}

var isPunctionSymbol = map[rune]bool{}

func init() {
	for _, r := range punctuationSymbol {
		isPunctionSymbol[r] = true
	}
}
