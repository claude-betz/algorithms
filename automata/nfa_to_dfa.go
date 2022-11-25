/*
	nfa_to_dfa.go

	Conversion from an NFA to a DFA via subset construction
*/

package main

import (
	"bytes"
	"fmt"
)

const (
	eps = -1
)

var (
	alphabet = []rune{'a', 'b'}
)

// graph
type Graph struct {
	nodes map[int]Node
	adjList map[Node][]*Edge
}

type Node interface{
	Id() int
} 

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
	nfaStates []Node
}

func (n DFAState) Id() int {
	return n.id
}

type Edge struct {
	src, dst Node
	accepts rune
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int]Node),
		adjList: make(map[Node][]*Edge),
	}
}

func (g *Graph) AddEdge(src, dst Node, accepts rune) {	
	// add nodes
	g.nodes[src.Id()] = src 
	g.nodes[dst.Id()] = dst

	// add edge
	e := &Edge{
		src: src,
		dst: dst,
		accepts: accepts,
	}

	_, ok := g.adjList[e.src]  

	if !ok {
		g.adjList[e.src] = []*Edge{e}
	} else {  
		g.adjList[e.src] = append(g.adjList[e.src], e)
	}
}

func (g *Graph) RecursiveBFS(res *[]Node, n []Node) {
	// print everything in array
	for _, node := range n {
		*res = append(*res, node)
	}
	fmt.Println()

	// recurse 
	for _, node := range n {
		// add all children to array
		var nodeArr []Node 

		for _, edge := range g.adjList[node] {
			nodeArr = append(nodeArr, edge.dst)
		}
		
		if len(nodeArr) != 0 {
			g.RecursiveBFS(res, nodeArr)
		}
	}
}

func (g *Graph) IterativeBFS(n Node) []Node {
	var res []Node

	// queue for BFS
	var queue []Node
	
	// visited array: default false
	var visited = make(map[Node]bool, len(g.nodes))

	// queue start
	queue = append(queue, n)
	visited[n] = true

	for {
		if len(queue) == 0 {
			break
		}

		// deque front node
		v := queue[0]
		// pop
		queue = queue[1:] 

		// append to res
		res = append(res, v)

		for _, edge := range g.adjList[v] {

			// if not visited add to queue
			dst := edge.dst
			if !visited[dst] {
				visited[dst] = true
				queue = append(queue, dst)
			}
		}
	}
	return res
}

func (g *Graph) EpsilonClosure(T []Node) []Node {
	// initialise epsilon closure
	var epsClosure []Node

	// use stack
	stack := make([]Node, 0)

	// push all states of T onto stack
	for _, state := range T {
		epsClosure = append(epsClosure, state)
		stack = append(stack, state)
	}

	// while stack is not empty
	for {
		if len(stack) == 0 {
			break
		}

		// deque t (last item)
		t := stack[len(stack)-1]
		// pop t
		stack = stack[:len(stack)-1]

		// iterate all states u with edge from t to u
		for _, edge := range g.adjList[t] {
			// only epsilon edges
			if edge.accepts == eps{
				u := edge.dst		

				// add to eps closure
				epsClosure = append(epsClosure, u)

				// push to stack
				stack = append(stack, u)
			}
		}
	}
	return epsClosure
}

func computeTIndex(T []Node) string {
	var tIndex bytes.Buffer
	for _, n := range T {
		tIndex.WriteString(fmt.Sprintf("%d", n.Id()))
	}
	return tIndex.String()
}


func (g *Graph) Move(T []Node, c rune) []Node {
	var res []Node
	// for each of the nodes in T
	for _, node := range T {
		// edges leaving this node
		edges := g.adjList[node]

		for _, edge := range edges {
			// check for edge that accepts c
			if edge.accepts == c {
				res = append(res, edge.dst)	
			}	
		}	
	}
	return res
}

func SubsetConstruction(nfa *Graph) *Graph {
	// DStates
	var dStates []DFAState
		
	// map
	var markMap map[string]bool

	count := 0
	dfa := NewGraph()
	
	// first DState = epsilonClosure(s0)
	nfaStartNode := nfa.nodes[0]
	epsClosure := nfa.EpsilonClosure([]Node{nfaStartNode})
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
			if !markMap[index] {
				// increment id count
				count++
				
				// create dfaState
				dfaState := DFAState{
					id: count,
					index: index,
					nfaStates: U,
				}

				// add to dStates
				dStates = append(dStates, state)
				
				// add to markMap
				markMap[index] = false

				// add E(T,U,c) to dfaGraph
				dfa.AddEdge(T, dfaState, c)
			}
		}

		// pop off T
		dStates = dStates[1:]	
	}

	return dfa
}

func (g *Graph) Print() {
	// for each node
	for n := range g.adjList {
		fmt.Printf("[%d]:", n)

		for _, e := range g.adjList[n] {
			fmt.Printf(" Move[%d, %s] = %d", n, string(e.accepts), e.dst)
		}

		fmt.Println()
	} 
}

func main() {

}


