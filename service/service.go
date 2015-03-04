package service

type Service interface {
	Id() string
	Run()
	Checker() Check
}
