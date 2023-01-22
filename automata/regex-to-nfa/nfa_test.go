package main

import (
	"testing"
)

/*
	[0]-ε->[1]
	[0]-ε->[2]
	[1]-b->[3]
	[1]-a->[3]
	[2]-c->[3]
*/
var (
	end = []*nfa{
		&nfa{
			false,
			map[rune][]*nfa{},
		},
	}

	nfa1 = &nfa{
		false,
		map[rune][]*nfa{
			'a': end,
			'b': end,
		},
	}
	
	nfa2 = &nfa{
		false,
		map[rune][]*nfa{
			'c': end,
		},
	}

	start = &nfa{
		false,
		map[rune][]*nfa{
			eps: []*nfa{
				nfa1,
				nfa2,
			},
		},
	}
)

var testCases = []struct {
	nfas []*nfa
	expectedEpsClosure []*nfa
	expectedMoves []*nfa
}{
	{
		[]*nfa{
			start,
		},
		[]*nfa{
			start,
			nfa1,
			nfa2,
		},
		[]*nfa{
			nfa1,
			nfa2,
		},
	},
}

func TestEpsClosure(t *testing.T) {
	for _, tc := range testCases {
		epsClosure := epsilonClosure(tc.nfas)		
		expectedEpsClosure := tc.expectedEpsClosure

		equal := checkEquality(epsClosure, expectedEpsClosure)
		if !equal {
			t.Errorf("epsClosure: %v\n, expectedEpsClosure: %v\n", epsClosure, expectedEpsClosure)
		}
	}
}

func TestMove(t *testing.T) {
	for _, tc := range testCases {
		moves := Move(tc.nfas, eps)
		expectedMoves := tc.expectedMoves

		equal := checkEquality(moves, expectedMoves)
		if !equal {
			t.Errorf("moves: %v\n, expectedMoves: %v\n", moves, expectedMoves)
		}
	}
}

func checkEquality(res []*nfa, expected []*nfa) bool {
	if len(res) != len(expected) {
		return false
	}

	for i, _ := range res {
		if res[i] != expected[i] {
			return false
		}
	}
	return true
}
	
