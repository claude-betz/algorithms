package graph

import (
	"fmt"
)

const (
	eps = 'ε'
)

type Node interface{
	Id() int
}

type State struct {
	id int
}

func (n State) Id() int {
	return n.id
}

type Edge struct {
	Src, Dst Node
	Accepts rune
}

type Graph struct {
	Nodes map[int]Node
	AdjList map[int][]*Edge
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]Node),
		AdjList: make(map[int][]*Edge),
	}
}

func (g *Graph) AddEdge(Src, Dst Node, Accepts rune) {	
	// add Nodes
	g.Nodes[Src.Id()] = Src 
	g.Nodes[Dst.Id()] = Dst

	// add edge
	e := &Edge{
		Src: Src,
		Dst: Dst,
		Accepts: Accepts,
	}

	_, ok := g.AdjList[e.Src.Id()]  

	if !ok {
		g.AdjList[e.Src.Id()] = []*Edge{e}
	} else {  
		g.AdjList[e.Src.Id()] = append(g.AdjList[e.Src.Id()], e)
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

		for _, edge := range g.AdjList[node.Id()] {
			nodeArr = append(nodeArr, edge.Dst)
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
	var visited = make(map[Node]bool, len(g.Nodes))

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

		for _, edge := range g.AdjList[v.Id()] {

			// if not visited add to queue
			Dst := edge.Dst
			if !visited[Dst] {
				visited[Dst] = true
				queue = append(queue, Dst)
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
		for _, edge := range g.AdjList[t.Id()] {
			// only epsilon edges
			if edge.Accepts == eps{
				u := edge.Dst		

				// add to eps closure
				epsClosure = append(epsClosure, u)

				// push to stack
				stack = append(stack, u)
			}
		}
	}
	return epsClosure
}

func (g *Graph) Move(T []Node, c rune) []Node {
	var res []Node
	// for each of the Nodes in T
	for _, node := range T {
		// edges leaving this node
		edges := g.AdjList[node.Id()]

		for _, edge := range edges {
			// check for edge that Accepts c
			if edge.Accepts == c {
				res = append(res, edge.Dst)	
			}	
		}	
	}
	return res
}

func (g *Graph) Print() {
	// for each node
	for n := range g.AdjList {
		fmt.Printf("[%d]:", n)

		for _, e := range g.AdjList[n] {
			fmt.Printf(" Move[%d, %s] = %d", n, string(e.Accepts), e.Dst.Id())
		}

		fmt.Println()
	} 
}
