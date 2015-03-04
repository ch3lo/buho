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
	nodes := []*graph.Node{}

	runChildrens(&nodes, node)
}

func runChildrens(nodes *[]*graph.Node, node *graph.Node) {
	//go checkChannel(node)

	for id, _ := range node.Nodes {
		node.Nodes[id].ServiceManager.Suscribe(node.ServiceManager.Channel)
		runChildrens(nodes, node.Nodes[id])
	}

	//for _, val := range *nodes {
	//	if val == node {
	//		return
	//	}
	//}

	*nodes = append(*nodes, node)
	node.ServiceManager.Run()
}

//func checkChannel(node *graph.Node) {
//	for {
//		s := <-node.ServiceManager.Channel
//		fmt.Printf("Check %#v : %#v\n", node.Id(), s)
//	}
//}

//func runner(nodes *[]*graph.Node) {
//	endpoint := "unix:///var/run/docker.sock"
//	client, _ := docker.NewClient(endpoint)

//	cs := make(chan string)
//	go check(cs)

//	for _, value := range *nodes {
//		cmd := []string{}
//		//cmd = append(cmd, value.Service.Command)
//		cmd = append(cmd)

//		config := docker.Config{Image: value.ServiceManager.Service.Image}
//		opts := docker.CreateContainerOptions{Name: value.Id(), Config: &config}
//		container, err := client.CreateContainer(opts)

//		if err != nil {
//			fmt.Println("FATAL: ", err)
//		}

//		client.StartContainer(container.ID, &docker.HostConfig{PublishAllPorts: true})

//		serviceOk := "Service " + value.Change + " - " + container.ID
//		cs <- serviceOk
//		//time.Sleep(1 * 1e9)
//	}
//}

//func check(cs chan string) {
//	for {
//		s := <-cs
//		fmt.Println("Check HTTP: ", s)
//	}
//}
