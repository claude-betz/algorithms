package main

import (
	"testing"
)

var readNextTestCases = []struct{
	program string
	expected string
}{
	{
		"abc",
		"a",
	},
}

func TestReadNext(t *testing.T) {
	for _, tc := range readNextTestCases {
		val, _ := readNext(&tc.program)
		if val != tc.expected {
			t.Fail()
		}
	}
}
