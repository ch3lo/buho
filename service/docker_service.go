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

func (s *DockerService) isRunning() bool {
	return s.Running
}

func (s *DockerService) Run() {
	if s.isRunning() {
		fmt.Printf("Service %#s is running with check %#s\n", s.Name, s.Check.Mode)
		return
	}

	fmt.Printf("Service command: docker run %#v %#s %#s\n", s.Parameters, s.Image, s.Command)
	s.Running = true
}

func (s *DockerService) Id() string {
	return s.Name
}
