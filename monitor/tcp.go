package monitor

import (
	"fmt"
	"time"
)

type TcpMonitor struct {
	Endpoint string
	Expect   string
}

func (tcp *TcpMonitor) Check() bool {
	try := 1
	for try <= 3 {
		fmt.Println("Check TCP: ", tcp.Endpoint)
		time.Sleep(1 * 1e9)
		try++
	}

	return true
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
