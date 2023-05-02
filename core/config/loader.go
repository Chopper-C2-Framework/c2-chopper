package config

import (
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	PluginsDir string `yaml:"plugins_path"`
	ClientPort int    `yaml:"client_port"`
	ServerPort int    `yaml:"server_port"`
	Host       string `yaml:"host"`
}

const (
	DEFAULT_CONFIG_DIR = "~/.c2chopper/"
)

const (
	CONFIG_NOT_FOUND       string = "The configuration file was not found to its specified path. If it's the first time using the framework add --init-config to generate a config file"
	CONFIG_PARSING_FAILURE string = "Error: Parsing configuration file has failed"
)

func ParseConfigFromPath() *Config {
	configFilePath := filepath.Join(DEFAULT_CONFIG_DIR, "config.yaml")

	config, err := ParseConfigFromFile(configFilePath)
	if err != nil {
		log.Panicln(err)
	}

	return config
}

func ParseConfigFromFile(filePath string) (*Config, error) {
	f, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, errors.New(CONFIG_NOT_FOUND)
	}

	var config Config

	if err := yaml.Unmarshal(f, &config); err != nil {
		return nil, errors.New(CONFIG_PARSING_FAILURE)
	}

	return &config, nil

}


