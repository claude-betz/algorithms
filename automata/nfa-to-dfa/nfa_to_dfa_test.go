/*
	nfa_to_dfa_test.go

	testing conversion of nfa to dfa
*/

package main

import (
	"testing"
	"fmt"
	"algorithms/automata/graph"
)

var testCases = []struct {
	edges []graph.Edge
	outputBFS []int
	outputEps []int
	outputMove []int
}{
	{
		[]graph.Edge{
			graph.Edge{NFAState{0}, NFAState{1}, eps},
			graph.Edge{NFAState{1}, NFAState{2}, 'a'},
			graph.Edge{NFAState{0}, NFAState{3}, eps},
			graph.Edge{NFAState{3}, NFAState{4}, 'b'},
		},
		[]int{0, 1, 3, 2, 4},
		[]int{0, 1, 3},
		[]int{1, 3},
	},
}

func TestNFAToDFA(t *testing.T) {
	for _, tc := range testCases {
		// build nfa
		nfa := buildGraph(tc.edges)

		// convert nfa to dfa
		dfa := SubsetConstruction(nfa)

		// print for now
		fmt.Println("nfa")
		nfa.Print()

		fmt.Println("dfa")
		dfa.Print()
	}
} 

func TestNFASimulation(t *testing.T) {
	for _, tc := range testCases {
		// build NFA
		nfa := buildGraph(tc.edges)

		// simulate with 
		validA := NfaSimulation(nfa, "a", []int{2})	
		validB := NfaSimulation(nfa, "b", []int{4})
		invalidA := NfaSimulation(nfa, "a", []int{4})
		invalidB := NfaSimulation(nfa, "b", []int{2})	
		fmt.Printf("validA: %v, invalidA: %v\n", validA, invalidA)
		fmt.Printf("validB: %v, invalidB: %v\n", validB, invalidB)
	}
}

func buildGraph(edges []graph.Edge) *graph.Graph {
	g := graph.NewGraph()
	for _, e := range edges {
		g.AddEdge(e.Src, e.Dst, e.Accepts)
	}
	return g
}

func checkEquality(res []graph.Node, expected []int) bool {
	if len(res) != len(expected) {
		return false
	}

	for i, _ := range res {
		if res[i].Id() != expected[i] {
			return false
		}
	}
	return true
}
