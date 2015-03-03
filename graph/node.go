package graph

import (
	"fmt"
)

type Node struct {
	Id    string
	Nodes map[string]*Node
}

func NewNode(id string) *Node {
	n := new(Node)
	n.Id = id
	n.Nodes = map[string]*Node{}
	return n
}

func (n *Node) addNeighbor(nb *Node) {
	fmt.Printf("Adding neighbor: %#s\n", nb.Id)
	n.Nodes[nb.Id] = nb
}

func (n *Node) isLeaf() bool {
	if len(n.Nodes) > 0 {
		return false
	}
	return true
}
