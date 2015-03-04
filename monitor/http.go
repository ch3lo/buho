package monitor

import (
	"fmt"
	"time"
)

type HttpMonitor struct {
	Endpoint string
	Expect   string
}

func (http *HttpMonitor) Check() bool {
	try := 1
	for try <= 5 {
		fmt.Println("Check HTTP: ", http.Endpoint)
		time.Sleep(1 * 1e9)
		try++
	}

	return true
}

func (http *HttpMonitor) SetEndpoint(ep string) {
	http.Endpoint = ep
}

func (http *HttpMonitor) GetEndpoint() string {
	return http.Endpoint
}

func (http *HttpMonitor) SetExpect(ex string) {
	http.Expect = ex
}

func (http *HttpMonitor) GetExpect() string {
	return http.Expect
}
