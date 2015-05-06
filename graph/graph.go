package graph

import (
	"github.com/kr/pretty"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	g := new(Graph)
	g.Nodes = map[string]*Node{}
	return g
}

func (g *Graph) AddNode(node *Node) {
	log.Debug("Adding node %s - %#v", node.Id(), pretty.Formatter(node))
	g.Nodes[node.Id()] = node
}

func (g *Graph) GetNode(id string) *Node {
	return g.Nodes[id]
}

func (g *Graph) AddEdge(from *Node, to *Node) {
	from.addNeighbor(to)
}

func (g *Graph) ReverseChildrens(id string) *[]*Node {
	nodes := []*Node{}

	childrens(&nodes, g.Nodes[id])

	return &nodes
}

func childrens(nodes *[]*Node, node *Node) {
	// recorre cada uno de los nodos vecinos
	for id, _ := range node.Neighbors {
		childrens(nodes, node.Neighbors[id])
	}

	//log.Debug("MSR %#v\n", pretty.Formatter(node))
	//node.Change = node.Change + "?"

	// si el nodo ya fue agregado se omite
	for _, val := range *nodes {
		if val == node {
			return
		}
	}

	*nodes = append(*nodes, node)
	//log.Debug("MSR NODES %#v\n", nodes)
}
