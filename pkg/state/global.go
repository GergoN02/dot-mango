package state

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/thegenem0/dot-mango/pkg/models"
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
	MangoConfigs     []models.MangoConfig
	Overrides        []models.Override
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

func (self *AppState) SetMangoConfigs(configs []models.MangoConfig) {
	self.Lock()
	defer self.Unlock()
	self.MangoConfigs = configs
}

func (self *AppState) GetMangoConfigs() []models.MangoConfig {
	self.RLock()
	defer self.RUnlock()
	return self.MangoConfigs
}

func (self *AppState) SetOverrides(overrides []models.Override) {
	self.Lock()
	defer self.Unlock()
	self.Overrides = overrides
}

func (self *AppState) GetOverrides() []models.Override {
	self.RLock()
	defer self.RUnlock()
	return self.Overrides
}

func (self *AppState) LoadConfig() error {
	self.Lock()
	defer self.Unlock()

	var config models.AppConfig

	bytes, err := os.ReadFile(filepath.Clean(filepath.Join(self.MangoConfigPath, DotMangoConfigFile)))
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return fmt.Errorf("failed to unmarshal yaml: %w", err)
	}

	if len(config.MangoConfigs) == 0 {
		log.Panicln("no mango configs found, please add some to the config file")
	}

	for i, mangoConfig := range config.MangoConfigs {
		var overrides []models.ConfigBoundOverride
		if config.Overrides != nil {
			for _, override := range config.Overrides {
				if override.Name == mangoConfig.Name || override.Name == "all" {
					overrides = append(mangoConfig.Overrides, models.ConfigBoundOverride{
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

	sysConfigPath, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get user config dir: %w", err)
	}

	if config.SystemConfigPath != "" {
		self.SystemConfigPath = config.SystemConfigPath
	} else {
		self.SystemConfigPath = sysConfigPath
	}
	if config.MangoConfigPath == "" {
		log.Panicln("mango config path is not set, please set it in the config file")
	} else {
		self.MangoConfigPath = config.MangoConfigPath
	}
	self.MangoConfigs = config.MangoConfigs
	self.Overrides = config.Overrides

	return nil
}
