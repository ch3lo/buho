package service

import (
	"fmt"
)

type DockerService struct {
	Name       string
	Image      string
	Command    string
	Running    bool
	Parameters []string
	Uses       []DockerService
	Check      Check
}

func (s *DockerService) String() string {
	return s.Image
}

func (s *DockerService) isRunning() bool {
	return s.Running
}

func (s *DockerService) Run() {
	fmt.Printf("Running Service %#s:\n%#v\n", s.Name, s)
}
