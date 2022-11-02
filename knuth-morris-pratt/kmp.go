package main

import(
	"fmt"
)

const (
	s = "aaaba"
	pattern = "ab"
)

func main() {
	fmt.Printf("indices: %v", KMP(s, pattern))
}

func KMP(s, pattern string) []int {
	// compute failure function
	ff := computeFailureFunction(pattern)

	// to hold indices of found patterns
	indices := make([]int, 0, 0)

	// loop through string
	i := 0
	j := 0
	for {
		fmt.Printf("i:%d, j:%d\n", i, j)

		if j == len(pattern) {
			indices = append(indices, i-len(pattern)) 
			j = ff[j-1]
		}

		// termination condition
		if i == len(s) {
			break
		}
	
		// iterate
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

func computeFailureFunction(pattern string) []int {

	// allocate space for failure function
	n := len(pattern)
	ff := make([]int, n, n)	

	t := 0
	ff[0] = 0

	for s:=1; s<n; s++ {
		fmt.Printf("s:%d, t:%d\n", s , t)
		for { 
			fmt.Printf("val t:%d\n", t)
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

	fmt.Printf("ff computed:\t%v\n", ff)
	return ff
}

func runeToString(b rune) string {
	return string(b)
}
