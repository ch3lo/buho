// Services
package main

type Service struct {
	Name       string
	Image      string
	Command    string
	Running    bool
	Parameters []string
}

func (s *Service) String() string {
	return s.Image
}

func (s *Service) isRunning() bool {
	return s.Running
}
