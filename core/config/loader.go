package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	NotFound       string = "the configuration file was not found to its specified path. If it's the first time using the framework add `gen-config` or `genc` to generate a config file"
	parsingFailure string = "error: Parsing configuration file has failed"
)

func GetConfig() (*Config, error) {

	err := GenerateConfigIfNotExists()

	if err != nil {
		log.Fatalln("Error getting configuration")
	}

	return ParseConfigFromPath(), nil
}

func ParseConfigFromPath() *Config {
	userHomeDir, err := os.UserHomeDir()

	if err != nil {
		log.Panicln("Error retrieving user home directory")
	}
	configFilePath := filepath.Join(userHomeDir, defaultPath)

	config, err := ParseConfigFromFile(configFilePath)
	if err != nil {
		log.Panicln(err)
	}

	return config
}

func ParseConfigFromFile(filePath string) (*Config, error) {
	f, err := os.ReadFile(filePath)

	if err != nil {
		return nil, errors.New(NotFound)
	}

	var config Config

	if err := yaml.Unmarshal(f, &config); err != nil {
		return nil, errors.New(parsingFailure)
	}

	return &config, nil
}
