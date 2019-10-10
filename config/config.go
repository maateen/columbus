package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Consul is an exportable type
type Consul struct {
	Enabled  bool   `mapstructure: "enabled"`
	Hostname string `mapstructure: "hostname"`
	Port     string `mapstructure: "port"`
}

// Traefik is an exportable type
type Traefik struct {
	Enabled bool `mapstructure: "enabled"`
}

// Node is an exportable type
type Node struct {
	Hostname string `mapstructure: "hostname"`
	Port     string `mapstructure: "port"`
	Scheme   string `mapstructure: "scheme"`
	Weight   int    `mapstructure: "weight"`
}

// Config is an exportable type
type Config struct {
	Consul  Consul  `mapstructure: "consul"`
	Traefik Traefik `mapstructure: "traefik"`
	Node    Node    `mapstructure: "node"`
}

var config *Config

// GetConfig exports config settings
func GetConfig() *Config {
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.AddConfigPath("/etc/columbus/") // path to look for the config file in
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(err)
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return config
}
