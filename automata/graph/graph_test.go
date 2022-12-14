package graph

import (
	"testing"
)

var testCases = []struct {
	edges []Edge
	outputBFS []int
	outputEps []int
	outputMove []int
}{
	{
		[]Edge{
			Edge{State{0}, State{1}, eps},
			Edge{State{1}, State{2}, 'a'},
			Edge{State{0}, State{3}, eps},
			Edge{State{3}, State{4}, 'b'},
		},
		[]int{0, 1, 3, 2, 4},
		[]int{0, 1, 3},
		[]int{1, 3},
	},
}

func TestBFS(t *testing.T) {
	for _, tc := range testCases {
		// build graph
		g := buildGraph(tc.edges)

		// perform iterative BFS from start=0
		itrBFS := g.IterativeBFS(State{0})
		
		// perform recursive BFS from start=0
		var recBFS []Node
		g.RecursiveBFS(&recBFS, []Node{State{0}})
		
		// check equality of recursive BFS to expected
		eq := checkEquality(recBFS, tc.outputBFS) 
		if !eq {
			t.Errorf("recBFS: %v expected: %v", recBFS, tc.outputBFS)
		}

		// check equality of iterative BFS to expected
		eq = checkEquality(itrBFS, tc.outputBFS)	
		if !eq {
			t.Errorf("itrBFS: %v expected: %v", itrBFS, tc.outputBFS)
		}
	}
}

func TestEpsClosure(t *testing.T) {
	for _, tc := range testCases {
		// build graph
		g := buildGraph(tc.edges)

		// get eps closure
		epsClosure := g.EpsilonClosure([]Node{State{0}})

		eq := checkEquality(epsClosure, tc.outputEps)
		if !eq {
			t.Errorf("epsClosure: %v expected: %v", epsClosure, tc.outputEps)	
		} 
	}
}

func TestMove(t *testing.T) {
	for _, tc := range testCases {
		// build graph
		g := buildGraph(tc.edges)

		// move from {0} with eps
		startNode := g.nodes[0]
		validMoves := g.Move([]Node{startNode}, eps)

		eq := checkEquality(validMoves, tc.outputMove)
		if !eq {
			t.Errorf("validMoves: %v expected: %v", validMoves, tc.outputMove)
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

func checkEquality(res []Node, expected []int) bool {
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
