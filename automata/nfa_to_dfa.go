/*
	nfa_to_dfa.go

	Conversion from an NFA to a DFA via subset construction
*/

package main

import (
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

type Node int 

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

func (g *Graph) RecursiveBFS(n []Node) {
	// print everything in array
	for _, node := range n {
		fmt.Printf("%d ", node)
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
			g.RecursiveBFS(nodeArr)
		}
	}
}

func (g *Graph) IterativeBFS(n Node) {
	// queue for BFS
	var queue []Node
	
	// visited array: default false
	var visited []bool = make([]bool, len(g.nodes))

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

		// print
		fmt.Printf("%d ", v)

		for _, edge := range g.adjList[v] {

			// if not visited add to queue
			dst := edge.dst
			if !visited[dst] {
				visited[dst] = true
				queue = append(queue, dst)
			}
		}
	}
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
	g := NewGraph()	
	
	// connect n0 and n1 accepting 'a'
	g.AddEdge(0, 1, eps)
	g.AddEdge(1, 2, 'a')
	g.AddEdge(0, 3, eps)
	g.AddEdge(3, 4, 'b')

	// print
	g.Print()

	// BFS
	nodeArr := []Node{0}
	fmt.Println("Recursive BFS")

	g.RecursiveBFS(nodeArr)

	fmt.Println("Iterative BFS")
	g.IterativeBFS(0)
}


