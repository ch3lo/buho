package util

import (
	"flag"
	"github.com/op/go-logging"
	"os"
)

var config *configuration = nil

var log = logging.MustGetLogger("main")

const logFormat = "%{color}%{time:15:04:05.000} %{level:.4s} %{id:03x} ▶%{color:reset} %{message}"

type configuration struct {
	File            string
	DockerApi       string
	PreCheckRetries int
	Checks          int
	CheckInterval   int
}

func GetConfiguration() *configuration {
	if config != nil {
		return config
	}

	log.Info("Setting up configuration")

	config = new(configuration)

	bkd := logging.NewLogBackend(os.Stderr, "", 0)
	format := logging.MustStringFormatter(logFormat)
	bkdFormatter := logging.NewBackendFormatter(bkd, format)
	logging.SetBackend(bkdFormatter)

	// AFTER LOGGING SETUP
	flag.StringVar(&config.File, "config", "./config.yml", "Process configuration file.")
	flag.StringVar(&config.DockerApi, "da", "unix:///var/run/docker.sock", "Api de docker.")
	flag.IntVar(&config.PreCheckRetries, "pcr", 3, "Cantidad de chequeos antes de correr el servicio.")
	flag.IntVar(&config.Checks, "c", -1, "Cantidad de chequeos despues de correr el servicio. -1 es infinito")
	flag.IntVar(&config.CheckInterval, "ci", 2, "Cantidad de segundos entre chequeos.")
	logLevel := flag.String("log", "INFO", "Logging level")

	flag.Parse()

	var level logging.Level
	var err error

	level, err = logging.LogLevel(*logLevel)

	if err != nil {
		log.Error("Invalid log level")
		level, _ = logging.LogLevel("INFO")
	}

	logging.SetLevel(level, "main")

	return config
}
