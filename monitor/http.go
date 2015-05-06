package monitor

import (
	"net/http"
	"time"
)

type HttpMonitor struct {
	Ip       string
	Port     string
	Endpoint string
	Expect   string
}

func (h *HttpMonitor) Check(retries int) bool {

	var url string = h.Endpoint

	log.Info("Checking HTTP connection %s", url)

	try := 1
	for retries == -1 || try <= retries {
		resp, err := http.Get(url)

		if err == nil {
			log.Info("Response from %s received with status %s", url, resp.StatusCode)

			if resp.StatusCode == 200 {
				return true
			}
		} else {
			log.Warning("%s", err)
		}

		try++
		time.Sleep(2 * 1e9)
	}

	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)

	return false
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
