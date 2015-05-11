package service

import (
	"github.com/ch3lo/wakeup/util"
	"github.com/op/go-logging"
)

type Service interface {
	Id() string
	Healthy() Healthy
	Type() ServiceType
	Run()
}

type Healthy struct {
	Mode string
	Ping string
	Pong string
}

type ServiceType int

const (
	EXTERNAL ServiceType = iota
	DOCKER
)

var config = util.GetConfiguration()
var log = logging.MustGetLogger("main")
