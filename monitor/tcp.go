package monitor

import (
	"fmt"
	"time"
)

type TcpMonitor struct {
	Ip       string
	Port     string
	Endpoint string
	Expect   string
}

func (tcp *TcpMonitor) Check() bool {
	try := 1
	for try <= 3 {
		fmt.Println("Checking TCP connection", tcp.Endpoint)
		time.Sleep(1 * 1e9)
		try++
	}

	return true
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
