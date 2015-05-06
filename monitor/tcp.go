package monitor

import (
	"net"
	"time"
)

type TcpMonitor struct {
	Ip       string
	Port     string
	Endpoint string
	Expect   string
}

func (tcp *TcpMonitor) Check(retries int) bool {
	try := 1

	for retries == -1 || try <= retries {
		log.Info("Checking TCP connection %s", tcp.Endpoint)
		conn, err := net.Dial("tcp", tcp.Endpoint)

		if err == nil {
			conn.Close()
			return true
		} else {
			log.Warning("%s", err)
		}

		log.Warning("%s", err)
		try++
		time.Sleep(2 * 1e9)
	}

	return false
}

func (tcp *TcpMonitor) SetIp(ip string) {
	tcp.Ip = ip
}

func (tcp *TcpMonitor) GetIp() string {
	return tcp.Endpoint
}

func (tcp *TcpMonitor) SetPort(port string) {
	tcp.Port = port
}

func (tcp *TcpMonitor) GetPort() string {
	return tcp.Port
}

func (tcp *TcpMonitor) SetEndpoint(ep string) {
	tcp.Endpoint = ep
}

func (tcp *TcpMonitor) GetEndpoint() string {
	return tcp.Endpoint
}

func (tcp *TcpMonitor) SetExpect(ex string) {
	tcp.Expect = ex
}

func (tcp *TcpMonitor) GetExpect() string {
	return tcp.Expect
}
