package main

import(
	"testing"
)

var tests = []struct {
	patterns []string
	output [][]int
}{
	{
		[]string{
			"car",
			"dog",
		},
		[][]int{
			{0,1,2,3},
			{0,4,5,6},
		},
	},
	{
		[]string{
			"car",
			"cars",
			"cards",
		},
		[][]int{
			{0,1,2,3},
			{0,1,2,3,4},
			{0,1,2,3,5,6},
		},
	},
	{
		[]string{
			"he",
			"she",
			"his",
			"hers",
		},
		[][]int{
			{0,1,2},
			{0,3,4,5},
			{0,1,6,7},
			{0,1,2,8,9},
		},
	},
}

func TestInsertAndRetrieveWord(t *testing.T) {
	
	for _, test := range tests {

		// Create root node
		root := NewState() 
		
		for i:=0; i<len(test.patterns); i++ {
			// insert
			root.InsertKeyword([]rune(test.patterns[i]))

			// get transitions
			validStates := root.GetKeywordStates([]rune(test.patterns[i]))

			// validate
			eq := checkEquality(validStates, test.output[i])
			if eq == false {
				t.Errorf("expected: %v, got: %v\n", []int(test.output[i]), validStates)
			}
		}
	} 
} 

func checkEquality(a, b []int) bool {
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

