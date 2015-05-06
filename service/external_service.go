package service

type ExternalService struct {
	Name         string  `json:"name" yaml:"name"`
	HealthyCheck Healthy `json:"healthy" yaml:"healthy"`
}

func (s *ExternalService) Id() string {
	return s.Name
}

func (s *ExternalService) Healthy() Healthy {
	return s.HealthyCheck
}

func (s *ExternalService) Type() ServiceType {
	return EXTERNAL
}

func (s *ExternalService) Run() {
	log.Debug("Nada que correr en %s", s.Id())
}
