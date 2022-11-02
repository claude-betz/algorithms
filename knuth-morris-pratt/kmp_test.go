package main

import(
	"testing"
)

var ffTests = []struct {
	pattern string
	output []int
}{
	{"ababd", []int{0,0,1,2,0}},
	{"ababaa", []int{0,0,1,2,3,1}},
}

var kmpTests = []struct {
	str 	string
	pattern string
	output []int
}{
	{"ababd", "b", []int{1,3}},
	{"ababaa", "aba", []int{0,2}},
	{"abababs", "abs", []int{4}},
}

func TestFailureFunction(t *testing.T) {
	
	for _, tt := range ffTests {
		ff := computeFailureFunction(tt.pattern)
		eq := testEquality(tt.output, ff)
		if eq == false {
			t.Errorf("expected: %v, got: %v\n", tt.output, ff)
		}
	}
}

func TestKMP(t *testing.T) {
	
	for _, tt := range kmpTests {	
		kmp := KMP(tt.str, tt.pattern) 
		eq := testEquality(tt.output, kmp)
		if eq == false {
			t.Errorf("expected: %v, got: %v\n", tt.output, kmp)
		}
	}
}

func testEquality(a, b []int) bool {
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
