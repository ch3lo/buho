package main

import (
	"github.com/ch3lo/wakeup/graph"
	"github.com/ch3lo/wakeup/service"
	"time"
)

func runNode(node *graph.Node) {
	runChildrens(node)

	log.Info("runNode waiting for Node %s", node.Id())

	for {
		if node.ServiceManager.Status == service.READY || node.ServiceManager.Status == service.FAILED {
			return
		}

		time.Sleep(1 * time.Second)
	}
}

func runChildrens(node *graph.Node) {
	for id, _ := range node.Neighbors {
		log.Info("%s needs %s", node.Id(), node.Neighbors[id].Id())

		node.Neighbors[id].ServiceManager.Suscribe(node.ServiceManager.Channel)
		runChildrens(node.Neighbors[id])
	}

	node.ServiceManager.EnqueueService()
}
