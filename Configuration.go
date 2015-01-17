// Configuration
package main

import "fmt"

type Configuration struct {
	Process  map[string]Process
	Services []Service
}

func (c *Configuration) getProcess(name string) *Process {
	fmt.Printf("Using configuration %#v\n\n", c)

	for index, element := range c.Process {
		if index == name {
			fmt.Printf("Using process %s\n%#v\n\n", index, element)
			return &element
		}
	}

	return nil
}
