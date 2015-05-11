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
	File      string
	DockerApi string
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
	flag.StringVar(&config.File, "config", "./config.yml", "Process configuration file. Default: ./config.yml")
	flag.StringVar(&config.DockerApi, "docker_api", "unix:///var/run/docker.sock", "Api de docker. Default: unix:///var/run/docker.sock")
	logLevel := flag.String("logging", "INFO", "Logging level")

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
