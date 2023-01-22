package main

import (
	"io"
	"fmt"
	"bytes"
)

const (
	eps = 'ε'
)

type nfa struct {
	accepting bool
	edges map[rune][]*nfa
}

func GetEndState(n *nfa) *nfa {
	for _, nextList := range n.edges {
		for _, elem := range nextList {
			if elem.accepting {
				return elem
			}
			return GetEndState(elem)
		}
		
	}	
	return nil
}

func epsilonClosure(T []*nfa) []*nfa {
	// initialise
	var epsClosure []*nfa

	// stack
	stack := make([]*nfa, 0)

	// push all initial states to epsClosure and stack 
	for _, nfa := range T {
		epsClosure = append(epsClosure, nfa)
		stack = append(stack, nfa) 
	}

	// while stack not empty
	for {
		if len(stack) == 0 {
			break
		}

		// deque last item
		t := stack[len(stack)-1]
		// pop
		stack = stack[:len(stack)-1]

		// iterate all states reachable via eps
		for _, nfa := range t.edges[eps] {
			epsClosure = append(epsClosure, nfa)
			stack = append(stack, nfa)
		}
	}

	return epsClosure 
}

func Move(T []*nfa, c rune) []*nfa {
	var res []*nfa

	for _, nfa := range T {
		val, ok := nfa.edges[c]
		if ok {
			res = append(res, val...)
		}
	}

	return res
}

func (n *nfa) Simulate(input string) bool {
	buf := bytes.NewBufferString(input)

	S := epsilonClosure([]*nfa{n})
	c, _, err := buf.ReadRune()

	for {
		if err == io.EOF {
			break
		}

		S = epsilonClosure(Move(S, c))
		c, _, err = buf.ReadRune()
	}

	for _, s := range S {
		if s.accepting {
			return true
		}
	}
	return false
}

func (n *nfa) PrintNFA() {
	// need to track assigned state numbers
	var seen = make(map[*nfa]int)
	
	// queue for bfs
	var queue []*nfa

	// nextState
	var stateId = 0

	// populate level 0  
	for key, nextStates := range n.edges {	
		for _, nextState := range nextStates {
			// increment stateId
			stateId++

			// print state
			fmt.Printf("[%d]-%s->[%d]\n", 0, string(key), stateId)

			// add to seen map
			seen[nextState] = stateId

			// add to queue
			queue = append(queue, nextState)
		}
	}

	for {
		if len(queue) == 0 {
			break
		}

		// deque
		curr := queue[0]
		queue = queue[1:]

		// currStateId
		currStateId := seen[curr]

		for char, nextStates := range curr.edges {			
			for _, nextState := range nextStates {
				val, ok := seen[nextState]

				// we have not seen this state before
				if !ok {
					// increment state
					stateId++

					// add to val
					val = stateId

					// add to seen
					seen[nextState] = stateId	

					// add to queue
					queue = append(queue, nextState)
				}

				// print
				fmt.Printf("[%d]-%s->[%d]\n", currStateId, string(char), val)
			}
		}		 
	}
}

