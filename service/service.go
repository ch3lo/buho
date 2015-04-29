package service

type Service interface {
	Id() string
	Run()
	Healthy() Healthy
}
