package service

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
)

type DockerService struct {
	Name       string
	Uses       []DockerService
	Check      Check
	Config     docker.Config
	HostConfig docker.HostConfig
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

	opts := docker.CreateContainerOptions{Name: s.Id(), Config: &s.Config}
	container, err := client.CreateContainer(opts)

	if err != nil {
		fmt.Println("FATAL: ", err)
	}

	client.StartContainer(container.ID, &s.HostConfig)
	bla, err := client.InspectContainer(container.ID)

	if err != nil {
		fmt.Println("FATAL: ", err)
	}

	fmt.Println("CONTAINER", bla)
}
