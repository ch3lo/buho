package main

import (
	"flag"
	"github.com/ch3lo/wakeup/graph"
	"github.com/kr/pretty"
	"github.com/op/go-logging"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var log = logging.MustGetLogger("main")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli sec
var format = logging.MustStringFormatter(
	"%{color}%{time:15:04:05.000} %{level:.5s} %{id:03x} ▶%{color:reset} %{message}",
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

		for _, srv_b := range srv_a.Externals {
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
	log.Debug("DIC ID %#v %# v", g.Nodes["acc"].Neighbors["pcc"].Neighbors["dic"], pretty.Formatter(g.Nodes["acc"].Neighbors["pcc"].Neighbors["dic"].Change))

	g.Nodes["acc"].Neighbors["pcc"].Neighbors["dic"].Change = g.Nodes["acc"].Neighbors["pcc"].Neighbors["dic"].Change + "?"

	log.Debug("DIC ID %#v %#v\n", g.Nodes["acc"].Neighbors["dic"], pretty.Formatter(g.Nodes["acc"].Neighbors["dic"].Change))

	g.Nodes["acc"].Neighbors["dic"].Change = g.Nodes["acc"].Neighbors["dic"].Change + "?"

	log.Debug("DIC ID %#v %#v\n", g.Nodes["dic"], pretty.Formatter(g.Nodes["dic"].Change))
}

func main() {
	bkd := logging.NewLogBackend(os.Stderr, "", 0)
	bkdFormatter := logging.NewBackendFormatter(bkd, format)
	logging.SetBackend(bkdFormatter)

	// AFTER LOGGING SETUP
	configFile := flag.String("config", "./config.yml", "Process configuration file")
	var logLevel string
	flag.StringVar(&logLevel, "logging", "INFO", "Logging level")

	flag.Parse()

	level, _ := logging.LogLevel(logLevel)

	logging.SetLevel(level, "main")

	config := readConfiguration(*configFile)
	log.Debug("GRAPH %#v\n\n\n", config)
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

	log.Debug("Reverse nodes:")
	for key, value := range *nodes {
		log.Debug("Node %# i: %s", key, value.Id())
	}

	log.Debug("GRAPH %#v", pretty.Formatter(nodes))

	//runner(nodes)
	serviceRunner(g.Nodes["acc"])
}
