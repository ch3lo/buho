// Configuration
package main

import (
	"github.com/ch3lo/wakeup/service"
)

type Configuration struct {
	DockerServices   []service.DockerService   `json:"docker_services" yaml:"docker_services"`
	ExternalServices []service.ExternalService `json:"external_services" yaml:"external_services"`
}
