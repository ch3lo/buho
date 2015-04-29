package service

import (
	"bufio"
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"os"
	"path"
	"strings"
)

type DockerService struct {
	Name         string            `json:"name" yaml:"name"`
	Uses         []DockerService   `json:"uses" yaml:"uses"`
	Checks       []ExternalService `json:"checks" yaml:"checks"`
	HealthyCheck Healthy           `json:"healthy" yaml:"healthy"`
	Config       docker.Config     `json:"config" yaml:"config"`
	HostConfig   docker.HostConfig `json:"hostconfig" yaml:"hostconfig"`
	EnvFiles     []string          `json:"envfiles" yaml:"envfiles"`
}

func (s *DockerService) Id() string {
	return s.Name
}

func (s *DockerService) Healthy() Healthy {
	return s.HealthyCheck
}

func (s *DockerService) Run() {
	//SACAR DE ACA
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)

	s.Config.Env = s.parseEnvFiles()

	fmt.Println("Starting container", s.Id(), "with config", s.Config)

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

	containerName := strings.Split(path.Base(bla.Config.Image), ":")[0]
	containerName = containerName
	fmt.Println("Container", s.Id(), "is running")
}

func (s *DockerService) parseEnvFiles() []string {
	var envs []string

	for _, path := range s.EnvFiles {
		parsedEnvs, _ := parseEnvFiles(path)
		envs = append(envs, parsedEnvs...)
	}

	return envs
}

func parseEnvFiles(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
