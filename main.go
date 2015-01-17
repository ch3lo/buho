package main

import (
	//	"errors"
	"flag"
	"fmt"
	//	"github.com/mkindahl/gograph/directed"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

func services(process *Process) {

	// Sin proceso ni servicio
	if !process.hasProcess() && !process.hasServices() {
		return
	}

	// Sin proceso y con servicio
	/*	if !process.hasProcess() && process.hasServices() {
		for index, service := range process.Mount {
			runService(index, &service)
		}
		return
	}*/

	// Con proceso y sin servicios
	if process.hasProcess() && !process.hasServices() {
		for _, proc := range process.Import {
			services(&proc)
		}
		return
	}

	// Con proceso y servicios
	// Primero iteramos por los procesos y luego se corre el servicio
	for _, proc := range process.Import {
		services(&proc)
	}
	/*
		for index, service := range process.Mount {
			runService(index, &service)
		}*/
}

func runService(index string, service *Service) {
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

	//project := flag.String("pp", "prefix_", "Project Prefix")

	configFile := flag.String("config", "./file.yml", "Configuration file")
	process := flag.String("process", "lan", "Process")

	flag.Parse()

	config := readConfiguration(*configFile)
	proc := config.getProcess(*process)

	for id, srv := range proc.servicex() {
		runService(id, &srv)
	}
	//services(proc)
}
