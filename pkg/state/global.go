package state

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/thegenem0/dot-mango/pkg/types"
	"gopkg.in/yaml.v2"
)

var (
	instance *AppState
	once     sync.Once
)

const (
	DotMangoConfigFile = "mangoConfig.yaml"
)

type AppState struct {
	sync.RWMutex

	MangoConfigPath  string
	SystemConfigPath string
	MangoConfigs     []types.MangoConfig
	Overrides        []types.Override
}

func GetAppState() *AppState {
	once.Do(func() {
		instance = &AppState{}
	})
	return instance
}

func (self *AppState) SetMangoConfigPath(path string) {
	self.Lock()
	defer self.Unlock()
	self.MangoConfigPath = path
}

func (self *AppState) SetSystemConfigPath(path string) {
	self.Lock()
	defer self.Unlock()
	self.SystemConfigPath = path
}

func (self *AppState) GetMangoConfigPath() string {
	self.RLock()
	defer self.RUnlock()
	return self.MangoConfigPath
}

func (self *AppState) GetSystemConfigPath() string {
	self.RLock()
	defer self.RUnlock()
	return self.SystemConfigPath
}

func (self *AppState) SetMangoConfigs(configs []types.MangoConfig) {
	self.Lock()
	defer self.Unlock()
	self.MangoConfigs = configs
}

func (self *AppState) GetMangoConfigs() []types.MangoConfig {
	self.RLock()
	defer self.RUnlock()
	return self.MangoConfigs
}

func (self *AppState) SetOverrides(overrides []types.Override) {
	self.Lock()
	defer self.Unlock()
	self.Overrides = overrides
}

func (self *AppState) GetOverrides() []types.Override {
	self.RLock()
	defer self.RUnlock()
	return self.Overrides
}

func (self *AppState) LoadConfig() error {
	self.Lock()
	defer self.Unlock()

	var config types.AppConfig

	bytes, err := os.ReadFile(filepath.Clean(filepath.Join(self.MangoConfigPath, DotMangoConfigFile)))
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return fmt.Errorf("failed to unmarshal yaml: %w", err)
	}

	for i, mangoConfig := range config.MangoConfigs {
		var overrides []types.ConfigBoundOverride
		if config.Overrides != nil {
			for _, override := range config.Overrides {
				if override.Name == mangoConfig.Name || override.Name == "all" {
					overrides = append(mangoConfig.Overrides, types.ConfigBoundOverride{
						Path:   override.Dotfile,
						Target: override.Target,
					})
				}
			}
		}

		fullPath := filepath.Join(config.MangoConfigPath, mangoConfig.Path)

		config.MangoConfigs[i].Path = fullPath
		config.MangoConfigs[i].Overrides = overrides
	}

	self.MangoConfigPath = config.MangoConfigPath
	self.SystemConfigPath = config.SystemConfigPath
	self.MangoConfigs = config.MangoConfigs
	self.Overrides = config.Overrides

	return nil
}
