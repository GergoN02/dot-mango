package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/thegenem0/dot-mango/pkg/app"
	"github.com/thegenem0/dot-mango/pkg/cmd"
)

func main() {
	pathPtr := flag.Bool("path", false, "path to dotfiles")
	initPtr := flag.Bool("init", false, "initialize a mango repository")
	generatePtr := flag.Bool("generate", false, "generate folders from config")
	helpPtr := flag.Bool("help", false, "print help")

	flag.Parse()

	if *generatePtr {
		cmd.GenerateFoldersFromConfig()
		return
	}

	if *helpPtr {
		flag.PrintDefaults()
		return
	}

	if *initPtr {
		if len(flag.Args()) > 0 {
			initPath := flag.Arg(0)
			cmd.InitializeGitRepo(initPath)
		} else {
			cmd.InitializeGitRepo(".")
		}
		return
	}

	if *pathPtr {
		if len(flag.Args()) > 0 {
			path := flag.Arg(0)
			homeDir, err := os.UserHomeDir()
			if err != nil {
				panic(err)
			}
			println("Changing directory to " + filepath.Clean(filepath.Join(homeDir, path)))
			if err := os.Chdir(filepath.Clean(filepath.Join(homeDir, path))); err != nil {
				panic(err)
			}
		}
	}
	app.LaunchUI()
}
