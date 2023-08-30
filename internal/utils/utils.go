package utils

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const DotMangoConfigFile = ".config/dot-mango/config.toml"

type UserConfig struct {
	Name string `toml:"name"`
	Path string `toml:"path"`
}

type AppConfig struct {
	UserConfig []UserConfig `toml:"config_version"`
}

func LoadConfig() (*AppConfig, error) {
	var config AppConfig

	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	if _, err := toml.DecodeFile(filepath.Join(homedir, DotMangoConfigFile), &config); err != nil {
		return nil, err
	}

	for i, userConfig := range config.UserConfig {
		fullPath := filepath.Join(homedir, userConfig.Path)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			return nil, err
		}

		relativePath, err := filepath.Rel(homedir, fullPath)
		if err != nil {
			return nil, err
		}
		config.UserConfig[i].Path = relativePath
	}

	return &config, nil
}
