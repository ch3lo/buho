package graph

import (
	"github.com/ch3lo/buho/service"
)

type Node struct {
	Service *service.DockerService // Change it to an interface to support another services
	Nodes   map[string]*Node
	Change  string
}

func NewNode(service *service.DockerService) *Node {
	n := new(Node)
	n.Service = service
	n.Change = service.Id()
	n.Nodes = map[string]*Node{}
	return n
}

func (n *Node) Id() string {
	return n.Service.Id()
}

func (n *Node) addNeighbor(nb *Node) {
	//fmt.Printf("Adding neighbor: %#s\n", nb.Id())
	n.Nodes[nb.Id()] = nb
}

func (n *Node) isLeaf() bool {
	if len(n.Nodes) > 0 {
		return false
	}
	return true
}
