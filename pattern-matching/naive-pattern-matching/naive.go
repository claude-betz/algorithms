package main

func NaivePatternMatching(text, pattern string) []int {
	res := make([]int, 0)
	// iternate over string
	for i:=0; i<len(text)-len(pattern); i++ {
		index := i
		match := true
		// iterate over pattern
		for j:=0; j<len(pattern); j++ {
			if text[index] != pattern[j] {
				match = false
				break
			} else {
				index++		
			} 
		} 	

		// if we didn't break early we matched
		if match {
			// pattern found at index i
			res = append(res, i)
		}
	} 

	return res
}
