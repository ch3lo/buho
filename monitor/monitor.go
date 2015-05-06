package monitor

import (
	"github.com/op/go-logging"
)

type Monitor interface {
	Check(retries int) bool
	SetIp(ep string)
	GetIp() string
	SetPort(port string)
	GetPort() string
	SetEndpoint(ep string)
	GetEndpoint() string
	SetExpect(ex string)
	GetExpect() string
}

var log = logging.MustGetLogger("main")
