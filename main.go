package main

import (
	"flag"
	"log"

	"github.com/gautamb02/social-api/configreader"
	"github.com/gautamb02/social-api/server"
)

func main() {
	configFilePath := flag.String("config", "", "Path to the configuration file")
	// Parse the command-line flags
	flag.Parse()
	configreader.SetConfigPath(*configFilePath)
	config, err := configreader.GetConfig()
	if err != nil {
		log.Fatal("%w", err)
		return
	}
	// log.Printf("%s", config.DB.MySQl.SocialAPIDB.Host)

	server := server.Server{Config: config}
	server.Setup()
	server.Start()
}
