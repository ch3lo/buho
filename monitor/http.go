package monitor

import (
	"fmt"
)

type HttpCheck struct {
}

func (http HttpCheck) Check(cs chan string) {
	for {
		s := <-cs
		fmt.Println("Check HTTP: ", s)
	}
}
