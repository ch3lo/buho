package service

import (
	"github.com/op/go-logging"
)

type Service interface {
	Id() string
	Healthy() Healthy
	Type() ServiceType
	Run()
}

type ServiceType int

const (
	EXTERNAL ServiceType = iota
	DOCKER
)

var log = logging.MustGetLogger("main")
