package service

import (
	"github.com/ch3lo/wakeup/monitor"
)

type Status int

const (
	CREATED Status = iota
	INIT
	READY
)

var statusList = []string{
	"CREATED",
	"INIT",
	"READY",
}

func (p Status) String() string {
	return statusList[p]
}

type ServiceManager struct {
	Status       Status
	Service      Service
	Channel      chan string
	Suscribers   []chan string
	Monitor      monitor.Monitor
	dependencies []*ServiceManager
}

func NewServiceManager(s Service) *ServiceManager {
	sm := new(ServiceManager)
	sm.Status = CREATED
	sm.Service = s
	sm.Channel = make(chan string)
	sm.Monitor = sm.initMonitor(s.Healthy())

	return sm
}

func (s *ServiceManager) initMonitor(checker Healthy) monitor.Monitor {
	var mon monitor.Monitor

	log.Debug("%s has checker %s", s.Id(), checker)
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
	if s.Status == INIT {
		log.Info("Allowed only one %s instance", s.Id())
		return
	}

	log.Info("Queuing %s", s.Id())
	s.Status = INIT

	go s.gooo()
}

func (s *ServiceManager) gooo() {
	waitDependencies := len(s.dependencies) != 0

	for waitDependencies {
		log.Info("%s waiting for signal", s.Id())
		signal := <-s.Channel

		log.Info("%s has signal received from %s", s.Id(), signal)

		waitDependencies = false
		for id, _ := range s.dependencies {
			if s.dependencies[id].Status != READY {
				log.Info("%s waiting for dependency %s", s.Id(), s.dependencies[id].Id())
				waitDependencies = true
			}
		}
	}

	log.Info("%s dependencies ready", s.Id())

	if s.check(3) == false {
		s.Service.Run()
	}

	s.check(-1)
}

func (s *ServiceManager) check(retries int) bool {
	if s.Monitor != nil {
		status := s.Monitor.Check(retries)

		if status {
			//AGREGAR LOGICA DE VALIDACION
			s.Status = READY

			for _, sus := range s.Suscribers {
				log.Info("%s service sending signal to %s", s.Id(), sus)
				sus <- "READY " + s.Id()
			}
			return true
		}

		log.Info("%s check failed", s.Id())
		return false
	}

	log.Info("has not checker defined", s.Id())

	return false
	//		s.Status = "ready"

	//		for _, sus := range s.Suscribers {
	//			fmt.Println("Service", s.Id(), "sending signal to", sus)
	//			sus <- "READY WO CHECK"
	//		}
}
