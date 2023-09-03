package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/thegenem0/dot-mango/pkg/fileops"
	"github.com/thegenem0/dot-mango/pkg/state"
)

func InitializeGitRepo(path string) {
	os.MkdirAll(path, 0755)

	if err := os.Chdir(path); err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	if _, err := os.Stat(".git"); err == nil {
		fmt.Println("Git repository already exists!")
		return
	}

	cmd := exec.Command("git", "init")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error initializing git repository:", err)
		return
	}
	fileops.InitDefaultMangoConfig(path)

	fmt.Println("Initialization successful!")
}

func GenerateFoldersFromConfig() {
	appState := state.GetAppState()

	appState.LoadConfig()

	fileops.GenerateRepoFileStructure(appState)

	fmt.Println("Folders generated successfully!")
}

func PrintHelp() {
	fmt.Println("Usage: mango [command]")
	fmt.Println("Commands:")
	fmt.Println("  init      Initialize dot-mango in the current directory")
	fmt.Println("  generate  Generate folders from the config.toml")
	fmt.Println("  help      Print this help message")
}

func PrintLoadError() {
	redPrint := color.New(color.FgRed).PrintlnFunc()

	redPrint("Could not find mangoConfig.yaml in current directory \nEnsure you run Mango from your mango git Repository")
	fmt.Println("If you are not in your mango repository, you can specify the path to your dotfiles repository with the --path flag")
	fmt.Println("If you do not have a mango repository, you can create one with the --init flag")
	fmt.Println("\n")

	os.Exit(1)
}
