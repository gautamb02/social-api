package main

import (
	"flag"
	"log"

	"github.com/gautamb02/social-api/configreader"
)

func main() {
	configFilePath := flag.String("config", "", "Path to the configuration file")
	// Parse the command-line flags
	flag.Parse()

	if *configFilePath == "" {
		log.Fatalf("No config file path provided. Use the -config flag to specify the config file.")
	}
	config, err := configreader.LoadConfig(*configFilePath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	log.Printf("App Name: %v", config.AppName)
}
