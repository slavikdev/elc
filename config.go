package main

import "flag"
import "fmt"

// Config includes client configuration.
type Config struct {
	Host string
	Port int
}

// NewConfig creates new instance of the client configuration.
func NewConfig() *Config {
	config := &Config{}
	readHost(&config.Host)
	readPort(&config.Port)
	flag.Parse()
	return config
}

func readHost(host *string) {
	helpMessage := fmt.Sprintf(
		"ElasticSearch host. %s is used by default.", defHost)
	flag.StringVar(host, "h", defHost, helpMessage)
}

func readPort(port *int) {
	helpMessage := fmt.Sprintf(
		"ElasticSearch port. %d is used by default.", defPort)
	flag.IntVar(port, "p", defPort, helpMessage)
}
