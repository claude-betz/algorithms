package main

import (
	"fmt"
)

var (
	program = "(a|b)*"
)

func main() {
	fmt.Printf("regex: %s\n", program)

	ast, err := union(&program)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("printing ast\n")
	ast.PrintNFA()
}
