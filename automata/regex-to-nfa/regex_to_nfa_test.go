package main

import (
	"testing"
)

var tests = []struct {
	alphabet []rune
	regex    string
	input    []string
	output   []bool
}{
	{
		[]rune{'a', 'b'},
		"a|b",
		[]string{"a", "b", "c"},
		[]bool{true, true, false},
	},
}

func TestRegexToNFA(t *testing.T) {

}
