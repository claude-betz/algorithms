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
	n.printNFA(0, 0)
}

func (n nfa) printNFA(state, next int) {
	var queue []*nfa
	
	for k, v := range n.edges {	
		next++
		fmt.Printf("[%d]-%s->[%d]\n", state, string(k), next)		
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

		curr.printNFA(state, next)
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
	c.AddState('d', true)

	n.Print()
}
