package main

import (
	"flag"
	"fmt"
	"github.com/ch3lo/buho/graph"
	"github.com/kr/pretty"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

func createGraph(config *Configuration) *graph.Graph {
	g := graph.NewGraph()

	for id, _ := range config.Services {
		g.AddNode(graph.NewNode(&config.Services[id]))
	}

	var from *graph.Node
	var to *graph.Node

	for _, srv_a := range config.Services {
		from = g.GetNode(srv_a.Name)

		for _, srv_b := range srv_a.Uses {
			to = g.GetNode(srv_b.Name)
			g.AddEdge(from, to)
		}
	}

	return g
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

func testPointer(g *graph.Graph) {
	fmt.Printf("DIC ID %#v %#v\n", g.Nodes["acc"].Nodes["pcc"].Nodes["dic"], pretty.Formatter(g.Nodes["acc"].Nodes["pcc"].Nodes["dic"].Change))
	g.Nodes["acc"].Nodes["pcc"].Nodes["dic"].Change = g.Nodes["acc"].Nodes["pcc"].Nodes["dic"].Change + "?"
	fmt.Printf("DIC ID %#v %#v\n", g.Nodes["acc"].Nodes["dic"], pretty.Formatter(g.Nodes["acc"].Nodes["dic"].Change))
	g.Nodes["acc"].Nodes["dic"].Change = g.Nodes["acc"].Nodes["dic"].Change + "?"
	fmt.Printf("DIC ID %#v %#v\n", g.Nodes["dic"], pretty.Formatter(g.Nodes["dic"].Change))
}

func main() {
	configFile := flag.String("config", "./config.yml", "Process configuration file")

	flag.Parse()

	config := readConfiguration(*configFile)

	var g *graph.Graph
	g = createGraph(&config)

	fmt.Printf("GRAPH %#v\n", pretty.Formatter(g))
	testPointer(g)
	fmt.Printf("GRAPH %#v\n", pretty.Formatter(g))

	var nodes *[]*graph.Node
	nodes = g.ReverseChildrens("acc")

	for key, value := range *nodes {
		fmt.Printf("Node %#s retrieved %#v\n", key, value)
	}

	fmt.Printf("GRAPH %#v\n", pretty.Formatter(nodes))

	runner(nodes)
}
