package main

import (
	"fmt"
)

var (
	program = []rune("(a)")
)

func main() {
	fmt.Println(string(program))
	lex := &Lexer{input: program}

//var lexedTokens []*Token
//for tk, _ := lex.ReadToken(); tk.Value != EOF; tk, _ = lex.ReadToken() {
//	lexedTokens = append(lexedTokens, tk)
//}
//
//fmt.Printf("%v\n", lexedTokens)

	closure(lex)	
}


