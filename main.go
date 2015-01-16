package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	Process  map[string]Process
	Services map[string]Service
}

type Process struct {
	Import []Process
	Mount  []Service
}

type Service struct {
	Image      string
	Command    string
	Parameters []string
}

func main() {
	filename, _ := filepath.Abs("./file.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", config)
}
