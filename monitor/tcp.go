package monitor

import (
	"net"
	"time"
)

type TcpMonitor struct {
	Endpoint string
	Expect   string
}

func (tcp *TcpMonitor) Check(retries int) bool {
	try := 1

	for retries == -1 || try <= retries {
		log.Debug("Checking TCP connection %s", tcp.Endpoint)
		conn, err := net.Dial("tcp", tcp.Endpoint)

		if err == nil {
			log.Info("Response from %s ... OK", tcp.Endpoint)
			conn.Close()
			return true
		} else {
			log.Debug("%s", err)
		}

		try++
		time.Sleep(2 * 1e9)
	}

	return false
}

func (tcp *TcpMonitor) SetEndpoint(ep string) {
	tcp.Endpoint = ep
}

func (tcp *TcpMonitor) SetExpect(ex string) {
	tcp.Expect = ex
}
