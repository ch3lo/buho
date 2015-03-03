package graph

import (
	"fmt"
	"github.com/kr/pretty"
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
	fmt.Printf("Adding node %#s - %#v - %#v\n", node.Id(), &node, pretty.Formatter(node))
	g.Nodes[node.Id()] = node
}

func (g *Graph) GetNode(id string) *Node {
	var node *Node
	node = g.Nodes[id]
	fmt.Printf("Getting node %#s - %#v - %#v\n", id, *node, node)
	return g.Nodes[id]
}

func (g *Graph) AddEdge(from *Node, to *Node) {
	fmt.Printf("Adding edge from %#v to %#v\n", from.Id(), to.Id())
	from.addNeighbor(to)
}

func (g *Graph) ReverseChildrens(id string) *[]*Node {
	nodes := []*Node{}
	fmt.Printf("MSR NODES ONE %v\n", &nodes)

	childrens(&nodes, g.Nodes[id])

	return &nodes
}

func childrens(nodes *[]*Node, node *Node) {
	for id, _ := range node.Nodes {
		childrens(nodes, node.Nodes[id])
	}

	fmt.Printf("MSR %#v", pretty.Formatter(node))
	node.Change = node.Change + "?"

	*nodes = append(*nodes, node)
	fmt.Printf("MSR NODES %#v\n", nodes)
}
