package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Servers map[string]Server `yaml:"servers"`
}

type Server struct {
	IP       string `yaml:"ip"`
	User     string `yaml:"user"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

func YamlLoad() (*Config, error) {
	data, err := os.ReadFile("configs/config.yml")
	if err != nil {
		return nil, fmt.Errorf("Error in load file : %v", err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("Error in parsing YAML: %v", err)
	}

	return &config, nil
}
