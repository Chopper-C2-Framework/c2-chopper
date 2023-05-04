package config_test

import (
	"log"
	"os"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/chopper-c2-framework/c2-chopper/core/config"
)

const (
	PluginsDir    = "./plugins"
	ClientPort    = 3000
	ServerPort    = 3000
	Host          = "localhost"
	ServerCert    = "./cert/server-cert.pem"
	ServerCertKey = "./cert/server-key.pem"
)

func TestParseConfigFile(t *testing.T) {
	configFilePath := "./config.yaml"
	defer deleteConfigFile(configFilePath)
	generateConfigFile(configFilePath)

	sconfig, err := config.ParseConfigFromFile(configFilePath)
	if err != nil {
		t.Fatal(err)
	}

	if sconfig.ClientPort != ClientPort || sconfig.ServerPort != ServerPort || sconfig.Host != Host || sconfig.PluginsDir != PluginsDir {
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
