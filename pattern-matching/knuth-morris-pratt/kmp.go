package main

import (
	"fmt"
)

var(
	patterns = []string{
		"abababaab",
		"aaaaaa",
		"abbaabb",
	}
	
)

func main() {
	for _, p := range patterns {
		fmt.Printf("string: %s, ff: %v\n", p, FailureFunction(p))	
	}
}

func KMP(s, pattern string) []int {
	// compute failure function
	ff := FailureFunction(pattern)

	// to hold indices of found patterns
	indices := make([]int, 0, 0)

	// loop through string
	i := 0
	j := 0
	for {
		// recognise pattern
		if j == len(pattern) {
			indices = append(indices, i-len(pattern)) 
			j = ff[j-1]
		}

		// termination condition
		if i == len(s) {
			break
		}
	
		// evaluate and move pointers
		if s[i] == pattern[j] {
			i++
			j++
		} else {
			for {
				if j > 0 {
					j = ff[j]
					continue
				}
				break
			}
			i++
		}
	}

	return indices
}

func FailureFunction(pattern string) []int {

	// allocate space for failure function
	n := len(pattern)
	ff := make([]int, n, n)	

	t := 0
	ff[0] = 0

	for s:=1; s<n; s++ {
		for { 
			if t>0 && pattern[s] != pattern[t] {
				t = ff[t-1]
				continue
			}
			break
		}	
		
		if pattern[s] == pattern[t] {
			t++
			ff[s] = t 
		} else {
			ff[s] = 0
		}
	}

	return ff
}

func runeToString(b rune) string {
	return string(b)
}
