package configreader

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AppName string `yaml:"appname"`
}

func LoadConfig(filePath string) (*Config, error) {
	// Open the YAML configuration file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %w", err)
	}
	defer file.Close()

	// Parse the YAML file into a Config struct
	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("unable to parse config file: %w", err)
	}

	// Return the populated Config struct
	return &config, nil
}
