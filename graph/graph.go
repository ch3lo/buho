package graph

import (
	"fmt"
)

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	g := new(Graph)
	g.Nodes = map[string]*Node{}
	return g
}

func (g *Graph) AddNode(node *Node) {
	fmt.Printf("Adding node %#s: %#v\n", node.Id, node)
	g.Nodes[node.Id] = node
}

func (g *Graph) GetNode(id string) *Node {
	fmt.Printf("Getting node %#s: %#v\n", id, g.Nodes[id])
	return g.Nodes[id]
}

func (g *Graph) AddEdge(from *Node, to *Node) {
	fmt.Printf("Adding edge from %#s to %#s\n", from.Id, to.Id)
	from.addNeighbor(to)
}

func (g *Graph) Print(id string) {
	g.fPrint(g.Nodes[id])
}

func (g *Graph) fPrint(node *Node) {
	for _, srv := range node.Nodes {
		g.fPrint(srv)
	}
	fmt.Printf("Node %#v\n", node)
}
