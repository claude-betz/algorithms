package main

import (
	"testing"
)

var r2nTestCases= []struct {
	regex string
	values []string
	expected []bool
}{
	{
		"a",
		[]string{"a", "aa", "aab", "c"},
		[]bool{true, false, false, false},
	},
	{
		"a|b",
		[]string{"a", "b", "c", "aa"},
		[]bool{true, true, false, false},
	},
	{
		"c*",
		[]string{"c", "ccc", "ac", "cccb"},
		[]bool{true, true, false, false},
	},
	{
		"(a|b)*",
		[]string{"aa", "bb", "ab", "c"},
		[]bool{true, true, true, false},
	},
}

func TestRegexToNFA(t *testing.T) {
	for _, tc := range r2nTestCases {
		nfa, _ := union(&tc.regex)
		results := make([]bool, 0)
		for _, val := range tc.values {
			res := nfa.Simulate(val)
			results = append(results, res) 
		}

		eq := checkBoolEquality(tc.expected, results)
		
		if !eq {
			t.Fail()
		}
	}
}

