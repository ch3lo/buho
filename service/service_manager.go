package service

import (
	"fmt"
	"github.com/ch3lo/wakeup/monitor"
)

type ServiceManager struct {
	Status       string
	Service      Service
	Channel      chan string
	Suscribers   []chan string
	Monitor      monitor.Monitor
	dependencies []*ServiceManager
}

func NewServiceManager(s Service) *ServiceManager {
	sm := new(ServiceManager)
	sm.Status = "created"
	sm.Service = s
	sm.Channel = make(chan string)
	sm.Monitor = initMonitor(s.Healthy())

	return sm
}

func initMonitor(checker Healthy) monitor.Monitor {
	var mon monitor.Monitor

	fmt.Println("Creating check", checker)
	if checker.Mode == "" {
		return nil
	} else if checker.Mode == "tcp" {
		mon = new(monitor.TcpMonitor)
	} else {
		mon = new(monitor.HttpMonitor)
	}

	mon.SetEndpoint(checker.Ping)
	mon.SetExpect(checker.Pong)

	return mon
}

func (s *ServiceManager) Id() string {
	return s.Service.Id()
}

func (s *ServiceManager) Suscribe(cs chan string) {
	s.Suscribers = append(s.Suscribers, cs)
}

func (s *ServiceManager) AddDependency(sm *ServiceManager) {
	s.dependencies = append(s.dependencies, sm)
}

func (s *ServiceManager) Run() {
	if s.Status == "init" {
		fmt.Println("Allowed only one", s.Id(), "instance")
		return
	}

	fmt.Println("Queuing", s.Id())
	s.Status = "init"

	go s.gooo()
}

func (s *ServiceManager) gooo() {
	waitDependencies := len(s.dependencies) != 0

	//fmt.Println("waitDependencies ", waitDependencies)

	for waitDependencies {
		fmt.Println(s.Id(), "waiting for signal")
		signal := <-s.Channel

		fmt.Println("Signal received from", signal)

		waitDependencies = false
		for id, _ := range s.dependencies {
			if s.dependencies[id].Status != "ready" {
				fmt.Println(s.Id(), "waiting for dependency", s.dependencies[id].Id())
				waitDependencies = true
			}
		}
	}

	s.Service.Run()

	if s.Monitor != nil {
		status := s.Monitor.Check()

		if status {
			//AGREGAR LOGICA DE VALIDACION
			s.Status = "ready"

			for _, sus := range s.Suscribers {
				fmt.Println("Service", s.Id(), "sending signal to", sus)
				sus <- "READY " + s.Monitor.GetEndpoint()
			}
		} else {
			fmt.Println("CHECK FAILED")
		}
	} else {
		s.Status = "ready"
		for _, sus := range s.Suscribers {
			fmt.Println("Service", s.Id(), "sending signal to", sus)
			sus <- "READY WO CHECK"
		}
		fmt.Println("No checker defined")
	}
}
