/*
	nfa_to_dfa.go

	Conversion from an NFA to a DFA via subset construction
*/

package main

import (
	"fmt"
)

// graph
type Graph struct {
	nodeCount, edgeCount int
	nodes []*Node
	edges []*Edge	
	mapping map[Node][]*Edge
}

type Node struct {
	id int
} 

type Edge struct {
	src, dst *Node
	accepts rune
}

func NewGraph() *Graph {
	return &Graph{
		nodeCount: 0,
		edgeCount: 0,
		nodes: make([]*Node, 0),
		edges: make([]*Edge, 0),
		mapping: make(map[Node][]*Edge),
	}
}

func (g *Graph) AddNode(n *Node) {
	// increment node count
	g.nodeCount++

	g.nodes = append(g.nodes, n)
}

func (g *Graph) AddEdge(src, dst *Node, accepts rune) {
	// increment edge count
	g.edgeCount++

	// create edge
	e := &Edge{
		src: src,
		dst: dst,
		accepts: accepts,
	}

	g.edges = append(g.edges, e)

	g.mapping[*src] = append(g.mapping[*src], e) 
}

func (g *Graph) Print() {
	fmt.Printf("nodeCount: %d, edgeCount: %d\n", g.nodeCount, g.edgeCount)

	// for each node
	for _, n := range g.nodes {
		fmt.Printf("[%d]", n.id)

		for _, e := range g.mapping[*n] {
			fmt.Printf(" --%s--> [%d]", string(e.accepts), e.dst.id)
		}

		fmt.Println()
	} 
}


func main() {
	// create graph
	g := NewGraph()	

	// create nodes
	n0 := &Node{id: 0}
	n1 := &Node{id: 1}
	n2 := &Node{id: 2}
	n3 := &Node{id: 3}
	n4 := &Node{id: 4}

	// add node
	g.AddNode(n0)
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)

	// connect n0 and n1 accepting 'a'
	g.AddEdge(n0, n1, -1)
	g.AddEdge(n1, n2, 'a')
	g.AddEdge(n2, n2, 'a')
	g.AddEdge(n0, n3, -1)
	g.AddEdge(n3, n4, 'b')
	g.AddEdge(n4, n4, 'b')

	// print
	g.Print()
}

