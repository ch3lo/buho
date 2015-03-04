package service

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
)

type DockerService struct {
	Name       string
	Image      string
	Command    string
	Parameters []string
	Uses       []DockerService
	Check      Check
}

func (s *DockerService) Id() string {
	return s.Name
}

func (s *DockerService) Checker() Check {
	return s.Check
}

func (s *DockerService) Run() {
	//SACAR DE ACA
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)

	config := docker.Config{Image: s.Image}
	opts := docker.CreateContainerOptions{Name: s.Id(), Config: &config}
	container, err := client.CreateContainer(opts)

	if err != nil {
		fmt.Println("FATAL: ", err)
	}

	client.StartContainer(container.ID, &docker.HostConfig{PublishAllPorts: true})
}
