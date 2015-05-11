package monitor

import (
	"github.com/ch3lo/wakeup/util"
	"github.com/op/go-logging"
)

type Monitor interface {
	Check(retries int) bool
	SetEndpoint(ep string)
	SetExpect(ex string)
}

var config = util.GetConfiguration()
var log = logging.MustGetLogger("main")
