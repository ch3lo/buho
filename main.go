package main

import (
	"flag"
	"fmt"
	"github.com/ch3lo/wakeup/graph"
	"github.com/kr/pretty"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

func createGraph(config *Configuration) *graph.Graph {
	g := graph.NewGraph()

	for id, _ := range config.DockerServices {
		g.AddNode(graph.NewNode(&config.DockerServices[id]))
	}

	for id, _ := range config.ExternalServices {
		g.AddNode(graph.NewNode(&config.ExternalServices[id]))
	}

	var from *graph.Node
	var to *graph.Node

	for _, srv_a := range config.DockerServices {
		from = g.GetNode(srv_a.Id())

		for _, srv_b := range srv_a.Uses {
			to = g.GetNode(srv_b.Id())
			g.AddEdge(from, to)
		}

		for _, srv_b := range srv_a.Checks {
			to = g.GetNode(srv_b.Id())
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
	fmt.Printf("DIC ID %#v %#v\n", g.Nodes["acc"].Neighbors["pcc"].Neighbors["dic"], pretty.Formatter(g.Nodes["acc"].Neighbors["pcc"].Neighbors["dic"].Change))

	g.Nodes["acc"].Neighbors["pcc"].Neighbors["dic"].Change = g.Nodes["acc"].Neighbors["pcc"].Neighbors["dic"].Change + "?"

	fmt.Printf("DIC ID %#v %#v\n", g.Nodes["acc"].Neighbors["dic"], pretty.Formatter(g.Nodes["acc"].Neighbors["dic"].Change))

	g.Nodes["acc"].Neighbors["dic"].Change = g.Nodes["acc"].Neighbors["dic"].Change + "?"

	fmt.Printf("DIC ID %#v %#v\n", g.Nodes["dic"], pretty.Formatter(g.Nodes["dic"].Change))
}

func main() {
	configFile := flag.String("config", "./config.yml", "Process configuration file")

	flag.Parse()

	config := readConfiguration(*configFile)
	fmt.Printf("GRAPH % #v\n", config)
	var g *graph.Graph = createGraph(&config)
	//var g2 *graph.Graph = createGraph(&config)

	//fmt.Printf("GRAPH %#v\n", pretty.Formatter(g))
	//fmt.Printf("GRAPH2 %#v\n", pretty.Formatter(g2))
	//testPointer(g)
	//fmt.Printf("GRAPH DIC ID %#v %#v\n", g.Nodes["dic"], pretty.Formatter(g.Nodes["dic"].Change))
	//fmt.Printf("GRAPH2 DIC ID %#v %#v\n", g2.Nodes["dic"], pretty.Formatter(g2.Nodes["dic"].Change))
	//fmt.Printf("GRAPH %#v\n", pretty.Formatter(g))
	//fmt.Printf("GRAPH2 %#v\n", pretty.Formatter(g2))

	var nodes *[]*graph.Node = g.ReverseChildrens("acc")

	for key, value := range *nodes {
		fmt.Printf("Node %#s retrieved % #v\n", key, value)
	}

	fmt.Printf("GRAPH %#v\n", pretty.Formatter(nodes))

	//runner(nodes)
	serviceRunner(g.Nodes["acc"])
}
