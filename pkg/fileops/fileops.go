package fileops

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/thegenem0/dot-mango/pkg/models"
	"github.com/thegenem0/dot-mango/pkg/state"
	"gopkg.in/yaml.v2"
)

func InitDefaultMangoConfig(path string) {
	mangoPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}

	config := models.AppConfig{
		MangoConfigPath:  mangoPath,
		SystemConfigPath: "",
		MangoConfigs:     []models.MangoConfig{},
		Overrides:        []models.Override{},
	}

	destPath := filepath.Join(path, "mangoConfig.yaml")
	modifiedConfig, err := yaml.Marshal(config)
	if err != nil {
		fmt.Println("Error marshaling YAML:", err)
		return
	}

	err = os.WriteFile(destPath, modifiedConfig, 0644)
	if err != nil {
		fmt.Println("Error copying config:", err)
		return
	}
}

func GenerateRepoFileStructure(appState *state.AppState) {
	for _, userConfig := range appState.MangoConfigs {
		fmt.Printf("Generating folder: %s\n", userConfig.Path)
		err := os.MkdirAll(userConfig.Path, 0755)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
	}
}

func GetUserConfigDirs(path string, overrides []models.ConfigBoundOverride, sysConfigPath string) []models.DotfileDirectory {
	var configs []models.DotfileDirectory

	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		folderName := file.Name()
		var resolvedPath string

		if overrides == nil {
			resolvedPath = filepath.Clean(filepath.Join(GetOsConfigPath(sysConfigPath), folderName))
		} else {
			for _, override := range overrides {
				if folderName == override.Path {
					resolvedPath = filepath.Clean(filepath.Join(GetOsConfigPath(sysConfigPath), override.Target))
				} else {
					resolvedPath = filepath.Clean(filepath.Join(GetOsConfigPath(sysConfigPath), folderName))
				}
			}
		}

		config := models.DotfileDirectory{
			Name:          folderName,
			SymlinkTarget: resolvedPath,
			Selected:      true,
			IsFolder:      file.IsDir(),
		}
		configs = append(configs, config)
	}

	return configs
}

func GetOsConfigPath(sysConfigPath string) string {
	var configPath string

	defaultConfigDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return ""
	}

	if sysConfigPath != "" {
		configPath = sysConfigPath
	} else {
		configPath = defaultConfigDir
	}

	return configPath
}

// func modifyAppConfig(configFlePath string, update models.AppConfig) (models.AppConfig, error) {
// 	var c models.AppConfig
//
// 	srcFile, err := os.ReadFile("/home/gergon02/Dev/Personal/dot-mango/mangoConfig.example.yaml")
// 	if err != nil {
// 		fmt.Println("Error reading default config:", err)
// 		return c, err
// 	}
//
// 	err = yaml.Unmarshal(srcFile, &c)
// 	if err != nil {
// 		fmt.Println("Error unmarshaling YAML:", err)
// 		return c, err
// 	}
//
// 	if update.SystemConfigPath != "" {
// 		c.SystemConfigPath = update.SystemConfigPath
// 	}
//
// 	if update.MangoConfigPath != "" {
// 		c.MangoConfigPath = update.MangoConfigPath
// 	}
//
// 	if update.MangoConfigs != nil {
// 		c.MangoConfigs = update.MangoConfigs
// 	}
//
// 	return c, nil
// }
