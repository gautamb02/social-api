package configreader

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AppName string   `yaml:"appname"`
	DB      DBConfig `yaml:"db"`
}

type DBConfig struct {
	MySQl SQLDBsConfig `yaml:"mysql"`
}

type SQLDBsConfig struct {
	SocialAPIDB MySQLConnectionConfig `yaml:"socialapi"`
}
type MySQLConnectionConfig struct {
	Host                  string `yaml:"host"`
	Port                  int    `yaml:"port"`
	User                  string `yaml:"user"`
	Password              string `yaml:"password"`
	Database              string `yaml:"database"`
	MaxOpenConnections    int    `yaml:"max_open_connections"`
	MaxIdleConnection     int    `yaml:"max_idle_connections"`
	MaxConnectionLifetime int    `yaml:"conn_max_lifetime"`
}

var config *Config
var cFilePath string

func SetConfigPath(filePath string) {
	cFilePath = filePath
}
func loadConfig() (*Config, error) {
	// Open the YAML configuration file
	file, err := os.Open(cFilePath)
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

func GetConfig() (*Config, error) {
	if config == nil {
		var err error = nil
		config, err = loadConfig()
		return config, err
	}
	return config, nil
}
