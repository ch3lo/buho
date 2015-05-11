package main

import (
	"github.com/ch3lo/wakeup/graph"
	"github.com/ch3lo/wakeup/service"
	"github.com/ch3lo/wakeup/util"
	"github.com/kr/pretty"
	"github.com/op/go-logging"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

var log = logging.MustGetLogger("main")

var config = util.GetConfiguration()

type ServicesConfiguration struct {
	DockerServices   []service.DockerService   `json:"docker_services" yaml:"docker_services"`
	ExternalServices []service.ExternalService `json:"external_services" yaml:"external_services"`
}

func main() {

	services := servicesConfigurationReader(config.File)
	var g *graph.Graph = createGraph(&services)

	//var g2 *graph.Graph = createGraph(&config)

	//fmt.Printf("GRAPH %#v\n", pretty.Formatter(g))
	//fmt.Printf("GRAPH2 %#v\n", pretty.Formatter(g2))
	//testPointer(g)
	//fmt.Printf("GRAPH DIC ID %#v %#v\n", g.Nodes["dic"], pretty.Formatter(g.Nodes["dic"].Change))
	//fmt.Printf("GRAPH2 DIC ID %#v %#v\n", g2.Nodes["dic"], pretty.Formatter(g2.Nodes["dic"].Change))
	//fmt.Printf("GRAPH %#v\n", pretty.Formatter(g))
	//fmt.Printf("GRAPH2 %#v\n", pretty.Formatter(g2))

	/*var nodes *[]*graph.Node = g.ReverseChildrens("acc")

	log.Debug("Reverse nodes:")
	for key, value := range *nodes {
		log.Debug("Node %# i: %s", key, value.Id())
	}

	log.Debug("GRAPH %#v", pretty.Formatter(nodes))
	*/

	runNode(g.Nodes["acc"])
}

func servicesConfigurationReader(configFile string) ServicesConfiguration {
	filename, _ := filepath.Abs(configFile)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var servicesConfig ServicesConfiguration

	err = yaml.Unmarshal(yamlFile, &servicesConfig)
	if err != nil {
		panic(err)
	}

	log.Debug("Configuration %#v\n\n\n", servicesConfig)

	return servicesConfig
}

func createGraph(services *ServicesConfiguration) *graph.Graph {
	g := graph.NewGraph()

	for id, _ := range services.DockerServices {
		g.AddNode(graph.NewNode(&services.DockerServices[id]))
	}

	for id, _ := range services.ExternalServices {
		g.AddNode(graph.NewNode(&services.ExternalServices[id]))
	}

	var from *graph.Node
	var to *graph.Node

	for _, srv_a := range services.DockerServices {
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

func testPointer(g *graph.Graph) {
	log.Debug("DIC ID %#v %# v", g.Nodes["acc"].Neighbors["pcc"].Neighbors["dic"], pretty.Formatter(g.Nodes["acc"].Neighbors["pcc"].Neighbors["dic"].Change))

	g.Nodes["acc"].Neighbors["pcc"].Neighbors["dic"].Change = g.Nodes["acc"].Neighbors["pcc"].Neighbors["dic"].Change + "?"

	log.Debug("DIC ID %#v %#v\n", g.Nodes["acc"].Neighbors["dic"], pretty.Formatter(g.Nodes["acc"].Neighbors["dic"].Change))

	g.Nodes["acc"].Neighbors["dic"].Change = g.Nodes["acc"].Neighbors["dic"].Change + "?"

	log.Debug("DIC ID %#v %#v\n", g.Nodes["dic"], pretty.Formatter(g.Nodes["dic"].Change))
}
