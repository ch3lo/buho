package main

import (
	"fmt"
	"github.com/ch3lo/buho/graph"
	"github.com/fsouza/go-dockerclient"
)

func runner(nodes *[]*graph.Node) {
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)

	cs := make(chan string)
	go check(cs)

	for _, value := range *nodes {
		cmd := []string{}
		//cmd = append(cmd, value.Service.Command)
		cmd = append(cmd)

		config := docker.Config{Image: value.Service.Image, AttachStdout: false, AttachStderr: false}
		opts := docker.CreateContainerOptions{Name: value.Id(), Config: &config}
		container, err := client.CreateContainer(opts)

		if err != nil {
			fmt.Println("FATAL: ", err)
		}

		serviceOk := "Service " + value.Change + " - " + container.ID
		cs <- serviceOk
		//time.Sleep(1 * 1e9)
	}
}

func check(cs chan string) {
	for {
		s := <-cs
		fmt.Println("Check HTTP: ", s)
	}
}
