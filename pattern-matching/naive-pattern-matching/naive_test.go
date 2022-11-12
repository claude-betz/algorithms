package main

import (
	"testing"
)

var testCases = []struct{
	str string
	pattern string
	output []int
}{
	{"ababd", "b", []int{1,3}},
	{"ababaa", "aba", []int{0,2}},
}

func TestNaivePatternMatching(t *testing.T) {
	for _, tc := range testCases {
		indices := NaivePatternMatching(tc.str, tc.pattern)
		eq := arraysEqual(tc.output, indices)
		if eq == false {
			t.Errorf("expected: %v, got: %v", tc.output, indices)
		}
	}
} 

func arraysEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	} 

	for i:=0; i<len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

