package main

import(
	"fmt"
	"testing"
)

var tests = []struct {
	pattern string
	output []int
}{
	{"ababd", []int{0,0,1,2,0}},
	{"ababaa", []int{0,0,1,2,3,1}},
}

func TestFailureFunction(t *testing.T) {
	
	for _, tt := range tests {
		fmt.Printf("\npattern: %s\n", tt.pattern)

		ff := computeFailureFunction(tt.pattern)
		eq := testEquality(tt.output, ff)
		if eq == false {
			t.Errorf("expected: %v, got: %v\n", tt.output, ff)

		}
	}
}

func testEquality(a, b []int) bool {
	fmt.Printf("a: %v, b: %v\n", a, b)
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
