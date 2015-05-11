package monitor

import (
	"github.com/op/go-logging"
)

type Monitor interface {
	Check(retries int) bool
	SetEndpoint(ep string)
	SetExpect(ex string)
}

var log = logging.MustGetLogger("main")
