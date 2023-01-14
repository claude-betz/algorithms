package main

import (
	"fmt"
)

type nfa struct {
	accepting bool
	edges map[rune]*nfa
}

func (n *nfa) AddState(char rune, accepting bool) *nfa {
	nfa := &nfa{
		accepting: accepting,
		edges: make(map[rune]*nfa),
	}

	// add state transition
	n.edges[char] = nfa

	// return new state
	return nfa
}

func (n *nfa) Print() {
	// need to track assigned state numbers
	var seen = make(map[*nfa]int)
	
	// queue for bfs
	var queue []*nfa

	// nextState
	var stateId = 0

	// populate level 0  
	for key, nextState := range n.edges {	
		// increment stateId
		stateId++

		// print state
		fmt.Printf("[%d]-%s->[%d]\n", 0, string(key), stateId)

		// add to seen map
		seen[nextState] = stateId

		// add to queue
		queue = append(queue, nextState)
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

		for char, nextState := range curr.edges {			
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

func main() {
	n := nfa{
		accepting: false,
		edges: make(map[rune]*nfa),
	}

	a := n.AddState('a', false)
	n.AddState('b', false)

	c := a.AddState('c', false)
	d := c.AddState('d', false)
	
	d.AddState('e', true)


	n.Print()
}
