package config

import (
	"fmt"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

func CheckIfConfigExistsInPath() (bool, error) {

	_, err := os.ReadFile(CONFIG_DEFAULT_PATH)

	if err != nil {
		// if err == os.ErrInvalid {
		return false, nil
		// }
		//
		// return false, err
	}

	return true, nil
}

func GenerateConfigIfNotExists() error {

	exists, err := CheckIfConfigExistsInPath()
	if err != nil {
		return err
	}

	if !exists {
		createDefaultConfigs()
	}

	return nil
}

func createDefaultConfigs() {

	currentUserHomeDirectory, err := os.UserHomeDir()

	if err != nil {
		log.Fatalln("Error getting user home directory")
	}

	configDirPath := path.Join(currentUserHomeDirectory, CONFIG_DEFAULT_DIR)

	_, err = os.ReadDir(configDirPath)


	if err != nil {
		err := os.Mkdir(configDirPath, 0777)
		if err != nil {
			log.Fatalln("Error creating .c2-chopper folder at home directory", err)
		}
	}

	configFilePath := path.Join(currentUserHomeDirectory, CONFIG_DEFAULT_PATH)
	config := CreateDefaultConfig()
	err = writeConfigToPath(config, configFilePath)

	if err != nil {
		log.Fatalln("Error marshalling configuration sturct")
	}

}

func writeConfigToPath(config *Config, path string) error {
	configBytes, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, configBytes, 0644)
	if err != nil {
		log.Fatalln("Error creating configuration file", err)
	}

	return nil
}
