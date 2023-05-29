package plugins

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"plugin"
	"strings"

	Cfg "github.com/chopper-c2-framework/c2-chopper/core/config"
	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
	"github.com/chopper-c2-framework/c2-chopper/core/services"
)

type LoadedPluginInfo struct {
	Plugin  IPlugin
	Channel chan *entity.TaskResultModel
}

type IPluginManager interface {
	ListAllPlugins() ([]string, error)
	ListLoadedPlugins() []string
	LoadAllPlugins() ([]*LoadedPluginInfo, error)
	LoadPlugin(filePath string) (*LoadedPluginInfo, error)
	GetPlugin(filePath string) (*LoadedPluginInfo, error)
}

type PluginManager struct {
	config        *Cfg.Config
	loadedPlugins map[string]*LoadedPluginInfo
	TaskService   services.ITaskService
}

func CreatePluginManager(cfg *Cfg.Config, taskService services.ITaskService) PluginManager {
	return PluginManager{
		config:        cfg,
		loadedPlugins: make(map[string]*LoadedPluginInfo),
		TaskService:   taskService,
	}
}

func lookupError(currErr error, errorMsg string) error {
	return errors.Join(currErr, errors.New(fmt.Sprintln("[-] Error: function lookup error in plugin", errorMsg)))
}

func reflectionError(currErr error, errorMsg string) error {
	return errors.Join(currErr, errors.New(fmt.Sprintln("[-] Error: type reflection error in plugin", errorMsg)))
}

func (manager PluginManager) GetPlugin(filePath string) (*LoadedPluginInfo, error) {
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

	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	fmt.Println(path.Join(homedir, manager.config.PluginsDir))
	files, err := os.ReadDir(path.Join(homedir, manager.config.PluginsDir))
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

func (manager PluginManager) LoadPlugin(filePath string) (*LoadedPluginInfo, error) {

	loadedPlugin, ok := manager.loadedPlugins[filePath]

	if ok {
		fmt.Println("[+] Plugin already loaded:", filePath)
		return loadedPlugin, nil
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	fullPath := filepath.Join(homedir, manager.config.PluginsDir, filePath)

	fmt.Println("Loading plugin:", filePath)

	info, err := os.Lstat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("plugin not found")
		}
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

	NewFn, ok := NewFnSymb.(func(service services.ITaskService) IPlugin)
	if !ok {
		return nil, errors.New("new function is not defined")
	}

	pluginInstance := NewFn(manager.TaskService)

	log.Println("[+] Loaded plugin ", pluginInstance.Info().Name)
	manager.loadedPlugins[filePath] = &LoadedPluginInfo{
		Plugin: pluginInstance,
	}
	return manager.loadedPlugins[filePath], nil
}

func (manager PluginManager) LoadAllPlugins() ([]*LoadedPluginInfo, error) {
	var (
		plugins []*LoadedPluginInfo
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
