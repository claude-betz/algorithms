package main

import (
	"testing"
)

var tests = []struct {
	regex  []rune
	tokens []*Token
}{
	{
		[]rune("a|b"),
		[]*Token{
			&Token{Tag: TagId, Value: "a"},
			&Token{Tag: TagPunct, Value: "|"},
			&Token{Tag: TagId, Value: "b"},
		},
	},
}

func testLexing(t *testing.T) {
	for _, test := range tests {
		lex := &Lexer{input: test.regex}

		var lexedTokens []*Token
		for tk, _ := lex.ReadToken(); tk.Value != EOF; tk, _ = lex.ReadToken() {
			lexedTokens = append(lexedTokens, tk)
		}

		// validate
		for i, expectedToken := range test.tokens {
			if !areEqual(expectedToken, lexedTokens[i]) {
				t.Errorf("expected: %v but got: %v", expectedToken, lexedTokens[i])
			}
		}
	}
}

func areEqual(a, b *Token) bool {
	if a.Tag == b.Tag && a.Value == b.Value {
		return true
	}
	return false
}
