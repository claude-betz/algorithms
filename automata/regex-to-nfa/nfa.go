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

	_, ok := nfa.edges[char]
	if !ok {
		// error
	}

	// add state transition
	n.edges[char] = nfa

	// return new state
	return nfa
}

func (n nfa) Print() {
	state := 0
	next := state
	var queue []*nfa
	
	for k, v := range n.edges {
		fmt.Printf("[%d]-%s->[%d]\n", state, string(k), next)		
		next++
		queue = append(queue, v)
	}

	for {
		state++
		if len(queue) == 0 {
			break
		}

		// deque
		curr := queue[0]
		queue = queue[1:]

		for k, v := range curr.edges {
			fmt.Printf("[%d]-%s->[%d]\n", state, string(k), next)		
			next++
			queue = append(queue, v)
		}
	}
}

func main() {
	n := nfa{
		accepting: false,
		edges: make(map[rune]*nfa),
	}

	next :=	n.AddState('a', false)
	n.AddState('b', false)

	next.AddState('c', true)
	n.Print()
}
