package monitor

type Monitor interface {
	Check() bool
	SetEndpoint(ep string)
	GetEndpoint() string
	SetExpect(ex string)
	GetExpect() string
}
