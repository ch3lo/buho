package monitor

type Monitor interface {
	Check() bool
	SetIp(ep string)
	GetIp() string
	SetPort(port string)
	GetPort() string
	SetEndpoint(ep string)
	GetEndpoint() string
	SetExpect(ex string)
	GetExpect() string
}
