package monitor

import (
	"fmt"
	"net/http"
	"time"
)

type HttpMonitor struct {
	Ip       string
	Port     string
	Endpoint string
	Expect   string
}

func (h *HttpMonitor) Check() bool {

	var url string = "http://" + h.Ip + ":" + h.Port + h.Endpoint

	for {
		fmt.Println("Check HTTP: ", url)

		resp, err := http.Get(url)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("RESP", resp)
			break
		}

		time.Sleep(1 * 1e9)
	}

	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)

	return true
}

func (http *HttpMonitor) SetIp(ip string) {
	http.Ip = ip
}

func (http *HttpMonitor) GetIp() string {
	return http.Endpoint
}

func (http *HttpMonitor) SetPort(port string) {
	http.Port = port
}

func (http *HttpMonitor) GetPort() string {
	return http.Port
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
