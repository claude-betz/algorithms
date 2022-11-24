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
	any = -2
)

// graph
type Graph struct {
	nodes map[Node]bool
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
		nodes: make(map[Node]bool),
		adjList: make(map[Node][]*Edge),
	}
}

func (g *Graph) AddEdge(src, dst Node, accepts rune) {	
	// add nodes
	g.nodes[src] = true
	g.nodes[dst] = true

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

func (g *Graph) RecursiveBFS(res *[]int, n []Node) {
	// print everything in array
	for _, node := range n {
		fmt.Printf("%d ", node.Id())
		*res = append(*res, node.Id())
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

func (g *Graph) IterativeBFS(n Node) []int {
	var res []int

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
		res = append(res, v.Id())

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

func (g *Graph) EpsilonClosure(T []Node) []int {
	// initialise epsilon closure
	var epsClosure []int

	// use stack
	stack := make([]Node, 0)

	// push all states of T onto stack
	for _, state := range T {
		epsClosure = append(epsClosure, state.Id())
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
				epsClosure = append(epsClosure, u.Id())

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
		tIndex.WriteString(fmt.Sprintf("%d",n.Id()))
	}
	return tIndex.String()
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


