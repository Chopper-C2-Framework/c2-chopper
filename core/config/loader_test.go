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

	sconfig, err := config.ParseConfigFromFile(configFilePath)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(sconfig, config.DefaultConfig) {
		t.Fatalf("Error the configuration wasn't as intended")
	}

}

func generateConfigFile(filePath string) {

	config := config.CreateDefaultConfig()

	b, err := yaml.Marshal(config)

	if err != nil {
		log.Fatalln("Error marshelling config struct")
	}

	f, err := os.Create(filePath)

	if err != nil {
		log.Panicln("Error creating file")
	}

	defer f.Close()

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
