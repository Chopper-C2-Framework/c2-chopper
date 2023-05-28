package config_test

import (
	"log"
	"os"
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/chopper-c2-framework/c2-chopper/core/config"
)

func TestParseConfigFile(t *testing.T) {
	configFilePath := "./config.yaml"
	defer deleteConfigFile(configFilePath)
	generateConfigFile(configFilePath)

	file, err := config.ParseConfigFromFile(configFilePath)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(file, config.CreateDefaultConfig()) {
		t.Fatalf("Error the configuration wasn't as intended")
	}

}

func generateConfigFile(filePath string) {

	frameworkConfig := config.CreateDefaultConfig()

	b, err := yaml.Marshal(frameworkConfig)

	if err != nil {
		log.Fatalln("Error marshelling frameworkConfig struct")
	}

	f, err := os.Create(filePath)

	if err != nil {
		log.Panicln("Error creating file")
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Panicln("Error loading config in test")
		}
	}(f)

	if err != nil {
		log.Panicln("Error creating file")
	}

	_, err = f.Write(b)

	if err != nil {
		log.Panicln(err)
	}

}

func deleteConfigFile(filepath string) {
	err := os.Remove(filepath)

	if err != nil {
		log.Panicln(err)
	}

}
