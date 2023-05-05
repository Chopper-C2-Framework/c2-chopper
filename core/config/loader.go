package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	CONFIG_NOT_FOUND       string = "The configuration file was not found to its specified path. If it's the first time using the framework add `gen-config` or `genc` to generate a config file"
	CONFIG_PARSING_FAILURE string = "Error: Parsing configuration file has failed"
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
	configFilePath := filepath.Join(userHomeDir, CONFIG_DEFAULT_PATH)

	config, err := ParseConfigFromFile(configFilePath)
	if err != nil {
		log.Panicln(err)
	}

	return config
}

func ParseConfigFromFile(filePath string) (*Config, error) {
	f, err := os.ReadFile(filePath)

	if err != nil {
		return nil, errors.New(CONFIG_NOT_FOUND)
	}

	var config Config

	if err := yaml.Unmarshal(f, &config); err != nil {
		return nil, errors.New(CONFIG_PARSING_FAILURE)
	}

	return &config, nil
}
