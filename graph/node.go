package graph

type Node struct {
	id    string
	nodes map[string]Node
}
