package fileops

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/thegenem0/dot-mango/pkg/models"
)

func CheckSymlinkPaths(dirs []models.DotfileDirectory) []error {
	errorMap := make([]error, len(dirs))

	for _, dir := range dirs {
		if dir.Selected {

			// Check if config directory exists
			targetDir := dir.SymlinkTarget
			if _, err := os.Stat(targetDir); err == nil {

				errMsg := fmt.Sprintf("Path %s already exists", targetDir)
				errorMap = append(errorMap, errors.New(errMsg))
			} else if !os.IsNotExist(err) {
				// Some other error occurred while trying to get the stat
				errorMap = append(errorMap, err)
			}
		}
	}
	if len(errorMap) > 0 {
		return errorMap
	}
	return nil
}

func SymlinkSelectedDotfiles(selectedConfig string, dirs []models.DotfileDirectory) error {
	for _, dir := range dirs {
		if dir.Selected {
			// Create the target directory if it doesn't exist
			targetDir := filepath.Dir(dir.SymlinkTarget)
			if err := os.MkdirAll(targetDir, 0755); err != nil {
				return err
			}

			// Delete the target if it already exists
			if _, err := os.Stat(dir.SymlinkTarget); !os.IsNotExist(err) {
				err := os.RemoveAll(dir.SymlinkTarget)
				if err != nil {
					panic(err)
					return err
				}
			}

			// Create the symlink
			err := os.Symlink(filepath.Join(selectedConfig, dir.Name), dir.SymlinkTarget)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
