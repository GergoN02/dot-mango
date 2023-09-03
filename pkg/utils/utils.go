package utils

import (
	"fmt"

	"github.com/thegenem0/dot-mango/pkg/state"
)

const DotMangoConfigFile = "mangoConfig.yaml"

func GetConfigNames(appState *state.AppState) []string {
	var configNames []string
	for _, mangoConfig := range appState.MangoConfigs {
		configNames = append(configNames, mangoConfig.Name)
	}
	return configNames
}

func GetConfigPathByName(appState *state.AppState, name string) (string, error) {
	for _, mangoConfig := range appState.MangoConfigs {
		if mangoConfig.Name == name {
			return mangoConfig.Path, nil
		}
	}
	return "", fmt.Errorf("config with name '%s' not found", name)
}
