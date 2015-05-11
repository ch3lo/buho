package monitor

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

type HttpMonitor struct {
	Endpoint string
	Expect   string
}

func (h *HttpMonitor) Check(retries int) bool {
	log.Info("Checking HTTP connection %s", h.Endpoint)

	r, _ := regexp.Compile(h.Expect)

	try := 1
	for retries == -1 || try <= retries {
		resp, err := http.Get(h.Endpoint)

		if err == nil {
			log.Debug("Response from %s received with status %s", h.Endpoint, resp.StatusCode)

			if resp.StatusCode == 200 {
				log.Info("Checking response from %s ...", h.Endpoint)
				result := false

				body, _ := ioutil.ReadAll(resp.Body)

				if r.MatchString(string(body)) {
					log.Info("Response from %s ... OK", h.Endpoint)
					result = true
				} else {
					log.Warning("Response from %s ... FAILED with content %s", h.Endpoint, string(body))
				}

				resp.Body.Close()
				return result
			}
		} else {
			log.Debug("%s", err)
		}

		try++
		time.Sleep(2 * 1e9)
	}

	return false
}

func (http *HttpMonitor) SetEndpoint(ep string) {
	http.Endpoint = ep
}

func (http *HttpMonitor) SetExpect(ex string) {
	http.Expect = ex
}
