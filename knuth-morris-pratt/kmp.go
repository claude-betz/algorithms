package main

import(
	"fmt"
)

const (
	s = "ababd"
)

func main() {
	computeFailureFunction(s)
}

func computeFailureFunction(pattern string) []int {

	// allocate space for failure function
	n := len(pattern)
	ff := make([]int, n, n)	
	fmt.Printf("ff init:\t%v\n", ff)

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
