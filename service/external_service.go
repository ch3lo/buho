package service

import (
	"fmt"
)

type ExternalService struct {
	Name         string  `json:"name" yaml:"name"`
	HealthyCheck Healthy `json:"check" yaml:"check"`
}

func (s *ExternalService) Id() string {
	return s.Name
}

func (s *ExternalService) Healthy() Healthy {
	return s.HealthyCheck
}

func (s *ExternalService) Run() {
	fmt.Println("Nada que correr en", s.Id())
}
