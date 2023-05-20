package plugins

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	Cfg "github.com/chopper-c2-framework/c2-chopper/core/config"
)

type IPluginManager interface {
	ListAllPlugins() ([]string, error)
	ListLoadedPlugins() []string
	LoadAllPlugins() ([]IPlugin, error)
	LoadPlugin(filePath string) (IPlugin, error)
	GetPlugin(filePath string) (IPlugin, error)
}

type PluginManager struct {
	config        *Cfg.Config
	loadedPlugins map[string]IPlugin
}

func CreatePluginManager(cfg *Cfg.Config) PluginManager {
	return PluginManager{
		config:        cfg,
		loadedPlugins: make(map[string]IPlugin)}
}

func lookupError(currErr error, errorMsg string) error {
	return errors.Join(currErr, errors.New(fmt.Sprintln("[-] Error: function lookup error in plugin", errorMsg)))
}

func reflectionError(currErr error, errorMsg string) error {
	return errors.Join(currErr, errors.New(fmt.Sprintln("[-] Error: type reflection error in plugin", errorMsg)))
}

func (manager PluginManager) GetPlugin(filePath string) (IPlugin, error) {
	loadedPlugin, ok := manager.loadedPlugins[filePath]
	if !ok {
		return nil, errors.New("plugin not loaded")
	}
	return loadedPlugin, nil
}

func (manager PluginManager) ListLoadedPlugins() []string {
	keys := make([]string, 0, len(manager.loadedPlugins))
	for k := range manager.loadedPlugins {
		keys = append(keys, k)
	}
	return keys
}

func (manager PluginManager) ListAllPlugins() ([]string, error) {
	var (
		plugins []string
	)

	files, err := os.ReadDir(manager.config.PluginsDir)
	if err != nil {
		// log.Panicln("Error: Cannot load plugins, No directory found.", err)
		return nil, err
	}

	for idx := range files {
		file := files[idx]

		if file.IsDir() || !strings.Contains(file.Name(), ".so") {
			continue
		}

		plugins = append(plugins, file.Name())
	}
	return plugins, nil
}

func (manager PluginManager) LoadPlugin(filePath string) (IPlugin, error) {
	loadedPlugin, ok := manager.loadedPlugins[filePath]
	if ok {
		fmt.Println("[+] Plugin already loaded:", filePath)
		return loadedPlugin, nil
	}

	fullPath := filepath.Join(manager.config.PluginsDir, filePath)

	fmt.Println("Loading plugin:", filePath)

	info, err := os.Lstat(fullPath)
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		return nil, errors.New("bad file path, path is a directory")
	}

	p, err := plugin.Open(fullPath)
	if err != nil {
		// log.Panicln(err)
		return nil, err
	}

	NewFnSymb, err := p.Lookup("New")
	if err != nil {
		err = lookupError(err, filePath)
		// log.Panicln(err)
		return nil, err
	}

	NewFn, ok := NewFnSymb.(func() IPlugin)
	if !ok {
		return nil, errors.New("new function is not defined")
	}

	pluginInstance := NewFn()

	log.Println("[+] Loaded plugin ", pluginInstance.Info().Name)
	manager.loadedPlugins[filePath] = pluginInstance
	return pluginInstance, nil
}

func (manager PluginManager) LoadAllPlugins() ([]IPlugin, error) {
	var (
		plugins []IPlugin
	)

	files, err := manager.ListAllPlugins()
	if err != nil {
		return nil, err
	}

	for idx := range files {
		file := files[idx]

		loadedPlugin, err := manager.LoadPlugin(file)
		if err != nil {
			return plugins, err
		}

		plugins = append(plugins, loadedPlugin)
	}
	return plugins, nil
}
