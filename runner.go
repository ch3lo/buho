package main

import (
	"fmt"
	"github.com/ch3lo/buho/graph"
	"time"
)

func serviceRunner(node *graph.Node) {
	runNode(node)
	fmt.Println("waiting Node ", node.Id())
	for {
		if node.ServiceManager.Status == "ready" {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func runNode(node *graph.Node) {
	for id, _ := range node.Nodes {
		node.Nodes[id].ServiceManager.Suscribe(node.ServiceManager.Channel)
		runNode(node.Nodes[id])
	}
	node.ServiceManager.Run()
}
