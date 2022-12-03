/*
	nfa_to_dfa.go

	Conversion from an NFA to a DFA via subset construction
*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"algorithms/automata/graph"
)

const (
	eps = 'ε'
)

var (
	alphabet = []rune{'a', 'b'}
)

type NFAState struct {
	id int
}

func (n NFAState) Id() int {
	return n.id
}

type DFAState struct {
	id int

	// idea index is to concatenate ordered nfaStates for an efficient lookup
	index string
	nfaStates []graph.Node
}

func (n DFAState) Id() int {
	return n.id
}

func computeTIndex(T []graph.Node) string {
	var tIndex bytes.Buffer
	for _, n := range T {
		tIndex.WriteString(fmt.Sprintf("%d", n.Id()))
	}
	return tIndex.String()
}

func SubsetConstruction(nfa *graph.Graph) *graph.Graph {
	// DStates
	var dStates []DFAState
		
	// map
	var markMap = make(map[string]bool, 0)

	count := 0
	dfa := graph.NewGraph()
	
	// first DState = epsilonClosure(s0)
	nfaStartNode := nfa.Nodes[0]
	epsClosure := nfa.EpsilonClosure([]graph.Node{nfaStartNode})
	index := computeTIndex(epsClosure)

	// add to DStates
	state := DFAState{
		id: count,
		index: index,
		nfaStates: epsClosure,
	}
	
	dStates = append(dStates, state)
	
	// add unmarked
	markMap[index] = false
	
	// while there is an unmarked state T in DStates
	for {	 
		if len(dStates) == 0 {
			break
		}
		
		// check states 
		T := dStates[0]	
	
		// mark T
		markMap[T.index] = true

		// loop possible input characters 
		for _, c := range alphabet {
			M := nfa.Move(T.nfaStates, c) 
			U := nfa.EpsilonClosure(M)

			// U not in DStates add as unmarked
			index = computeTIndex(U)
			
			_, ok := markMap[index]

			if !ok && len(U) != 0 {
				// increment id count
				count = count +1
				
				// create dfaState
				dfaState := DFAState{
					id: count,
					index: index,
					nfaStates: U,
				}

				// add to dStates
				dStates = append(dStates, dfaState)
				
				// add E(T,U,c) to dfaGraph
				dfa.AddEdge(T, dfaState, c)
			}
		}

		// pop off T
		dStates = dStates[1:]	
	}

	return dfa
}

func NfaSimulation(g *graph.Graph, input string, acceptingStates []int) bool {
	buf := bytes.NewBufferString(input)
	alreadyOn := make(map[int]bool, len(g.Nodes))

	s0 := g.Nodes[0]
	S := g.EpsilonClosure([]graph.Node{s0})
	c, _, err := buf.ReadRune()
			
	for {
		if err == io.EOF {
			break
		}	

		S = g.EpsilonClosure(g.Move(S, c))
		c, _, err = buf.ReadRune()
	}

	for _, s := range S {
		alreadyOn[s.Id()] = true	
	} 

	for _, accepting := range acceptingStates {
		if alreadyOn[accepting] {
			return true
		}
	}
	return false
}

//func (g *Graph) AddState(s Node, oldStates *[]Node, alreadyOn map[Node]bool) {
//	*oldStates = append(*oldStates, s.Id())
//	alreadyOn[s.Id()] = true
//
//	for _, state := range g.Move([]Node{s}, eps) {
//		if !alreadyOn[state] {
//			g.AddState(state, oldStates)
//		}		
//	}
//}



