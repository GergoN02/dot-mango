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

func FormatErrors(errors []error) string {
	var errorString string
	for _, err := range errors {
		if err != nil {
			errorString += fmt.Sprintf("%s\n", err)
		}
	}
	return errorString
}

func FormatActions(actions []string) string {
	var actionString string
	for _, action := range actions {
		actionString += fmt.Sprintf("(%s)    ", action)
	}
	return actionString
}
