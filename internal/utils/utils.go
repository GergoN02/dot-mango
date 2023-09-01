package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const DotMangoConfigFile = ".config/dot-mango/config.toml"

type UserConfig struct {
	Name string `toml:"name"`
	Path string `toml:"path"`
}

type UserConfigDir struct {
	FolderName    string
	SymlinkTarget string
	Selected      bool
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

		config.UserConfig[i].Path = fullPath
	}

	return &config, nil
}

func GetConfigNames(c *AppConfig) []string {
	var configNames []string
	for _, userConfig := range c.UserConfig {
		configNames = append(configNames, userConfig.Name)
	}
	return configNames
}

func GetUserConfigDirs(path string) ([]UserConfigDir, error) {
	var configs []UserConfigDir

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %v", err)
	}

	for _, file := range files {
		folderName := file.Name()
		symlinkTarget := os.ExpandEnv(fmt.Sprintf("$HOME/%s", folderName))
		config := UserConfigDir{
			FolderName:    folderName,
			SymlinkTarget: symlinkTarget,
			Selected:      true,
		}
		configs = append(configs, config)
	}

	return configs, nil
}
