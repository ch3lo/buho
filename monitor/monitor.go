package monitor

type Monitor interface {
	Check(cs chan string)
}
