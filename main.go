package main

import (
	"flag"
	"fmt"
	//"github.com/ch3lo/buho/graph"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

func runService(index int, service *Service) {
	fmt.Printf("Running Service %#s:\n%#v\n\n", index, service)
	service.Running = true
}

func readConfiguration(configFile string) Configuration {
	filename, _ := filepath.Abs(configFile)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config Configuration

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func main() {
	configFile := flag.String("config", "./config.yml", "Process configuration file")

	flag.Parse()

	config := readConfiguration(*configFile)

	for id, srv := range config.Services {
		runService(id, &srv)
	}
}
