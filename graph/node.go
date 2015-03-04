package graph

import (
	"github.com/ch3lo/buho/service"
)

type Node struct {
	ServiceManager *service.ServiceManager
	Nodes          map[string]*Node
	Change         string
}

func NewNode(s service.Service) *Node {
	n := new(Node)
	n.ServiceManager = service.NewServiceManager(s)
	n.Change = s.Id()
	n.Nodes = map[string]*Node{}
	return n
}

func (n *Node) Id() string {
	return n.ServiceManager.Id()
}

func (n *Node) addNeighbor(nb *Node) {
	//fmt.Printf("Adding neighbor: %#s\n", nb.Id())
	n.Nodes[nb.Id()] = nb
	n.ServiceManager.AddDependency(nb.ServiceManager)
}

func (n *Node) isLeaf() bool {
	if len(n.Nodes) > 0 {
		return false
	}
	return true
}
