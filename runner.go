package main

import (
	"fmt"
	"github.com/ch3lo/buho/graph"
	"github.com/fsouza/go-dockerclient"
)

func runner(g *graph.Graph) {
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)
	imgs, _ := client.ListContainers(docker.ListContainersOptions{})

	for _, img := range imgs {
		fmt.Printf("Image: %#v\n", img)
	}
}
