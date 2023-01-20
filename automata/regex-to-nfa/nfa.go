package main

import (
	"fmt"
)

const (
	eps = 'ε'
)

type nfa struct {
	accepting bool
	edges map[rune][]*nfa
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


func main() {
	end := []*nfa{
		&nfa{
			false,
			map[rune][]*nfa{},
		},
	}
	
	nfa1 := &nfa{
		false,
		map[rune][]*nfa{
			'a': end,
			'b': end,
		},
	}
	
	nfa2 := &nfa{
		false,
		map[rune][]*nfa{
			'c': end,
		},
	}

	start := &nfa{
		false,
		map[rune][]*nfa{
			eps: []*nfa{
				nfa1,
				nfa2,
			},
		},
	}	

	start.PrintNFA()
}

