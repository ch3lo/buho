package main

import (
	"github.com/ch3lo/wakeup/graph"
	"github.com/ch3lo/wakeup/service"
	"time"
)

func serviceRunner(node *graph.Node) {
	runChildrensFirst(node)
	log.Info("serviceRunner waiting for Node %s", node.Id())
	for {
		if node.ServiceManager.Status == service.READY {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func runChildrensFirst(node *graph.Node) {
	for id, _ := range node.Neighbors {
		log.Info("%s needs %s", node.Id(), node.Neighbors[id].Id())
		node.Neighbors[id].ServiceManager.Suscribe(node.ServiceManager.Channel)
		runChildrensFirst(node.Neighbors[id])
	}
	node.ServiceManager.Run()
}
