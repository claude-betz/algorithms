/*
	nfa_to_dfa_test.go

	testing conversion of nfa to dfa
*/

package main

import (
	"testing"
)

var testCases = []struct {
	edges []Edge
	output []int
}{
	{
		[]Edge{
			Edge{NFAState{0}, NFAState{1}, eps},
			Edge{NFAState{1}, NFAState{2}, 'a'},
			Edge{NFAState{0}, NFAState{3}, eps},
			Edge{NFAState{3}, NFAState{4}, 'b'},
		},
		[]int{0, 1, 3, 2, 4},
	},
}

func TestBFS(t *testing.T) {
	for _, tc := range testCases {
		// build Graph
		g := buildGraph(tc.edges)

		// perform iterative BFS from start=0
		itrBFS := g.IterativeBFS(NFAState{0})
		
		// perform recursive BFS from start=0
		var recBFS []int
		g.RecursiveBFS(&recBFS, []Node{NFAState{0}})
		
		// check equality of iterative and recursive BFS
		eq := checkEquality(itrBFS, recBFS) 
		if eq == false {
			t.Errorf("itrBFS: %v not equal recBFS: %v", itrBFS, recBFS)
		}

		// check equality of one of the arrays to expected
		correct := checkEquality(itrBFS, tc.output)	
		if correct == false {
			t.Errorf("itrBFS: %v not equal to expected output: %v", itrBFS, tc.output)
		}
	}
}

func buildGraph(edges []Edge) *Graph {
	g := NewGraph()
	for _, e := range edges {
		g.AddEdge(e.src, e.dst, e.accepts)
	}
	return g
}

func checkEquality(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i, _ := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true

}

