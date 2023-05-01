package plugins

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"plugin"
)

const PluginsDir = "./plugins"

func lookupError(currErr error, errorMsg string) error {
	return errors.Join(currErr, errors.New(fmt.Sprintln("[-] Error: function lookup error in plugin", errorMsg)))
}

func reflectionError(currErr error, errorMsg string) error {
	return errors.Join(currErr, errors.New(fmt.Sprintln("[-] Error: type reflection error in plugin", errorMsg)))
}

func LoadPlugins() ([]Plugin, error) {

	var (
		files   []os.FileInfo
		plugins []Plugin
		err     error
		p       *plugin.Plugin
		n       plugin.Symbol
	)

	if files, err = ioutil.ReadDir(PluginsDir); err != nil {
		log.Panicln("Error: Cannot load plugins, error occured")
	}

	for idx := range files {
		file := files[idx]
		fmt.Println("Loading plugin: ", files[idx].Name())

		if p, err = plugin.Open(PluginsDir + files[idx].Name()); err != nil {
			// TODO change this to append to the error object
			log.Panicln(err)
		}

		n, err = p.Lookup("New")

		if err != nil {
			err = lookupError(err, file.Name())
			continue
		}

		newPlugin, ok := n.(func() Plugin)

		if !ok {
			err = reflectionError(err, fmt.Sprintf("New function for plugin %s", file.Name()))
			continue
		}

		plugin := newPlugin()
		log.Println("[+] Loaded plugin ", plugin.Name)
		plugins = append(plugins, plugin)
	}

	return plugins, err

}
