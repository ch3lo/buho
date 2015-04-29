package main

import (
	"fmt"
	"github.com/ch3lo/wakeup/graph"
	"time"
)

func serviceRunner(node *graph.Node) {
	runNode(node)
	fmt.Println("serviceRunner waiting for Node", node.Id())
	for {
		if node.ServiceManager.Status == "ready" {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func runNode(node *graph.Node) {
	for id, _ := range node.Neighbors {
		fmt.Println(node.Id(), "needs", node.Neighbors[id].Id())
		node.Neighbors[id].ServiceManager.Suscribe(node.ServiceManager.Channel)
		runNode(node.Neighbors[id])
	}
	node.ServiceManager.Run()
}
