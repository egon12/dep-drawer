// Package gomod package that contain function to find absolute path for go.mod
package gomod

import (
	"fmt"
	"os"
	"path/filepath"
)

// Find find absolutepath that have go.mod in parent folder.
// only test works in UNIX. Maybe wouldn't work in windows.
func Find() (string, error) {
	p, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("Cannot get working dir: %v", err)
	}

	return findRecursiveGoMod(p)
}

func FindDir() (string, error) {
	p, err := Find()
	if err != nil {
		return "", err
	}
	return filepath.Dir(p), nil
}

func findRecursiveGoMod(currentPath string) (string, error) {
	gomodPath := filepath.Join(currentPath, "go.mod")

	fileInfo, err := os.Stat(gomodPath)
	if err == nil {
		if !fileInfo.IsDir() {
			return gomodPath, nil
		}
	}

	newPath := filepath.Dir(currentPath)

	// maybe got to root directory
	if currentPath == newPath {
		return "", fmt.Errorf("can't find go.mod in parent ancestor: stuck in %s", newPath)
	}

	return findRecursiveGoMod(newPath)
}
