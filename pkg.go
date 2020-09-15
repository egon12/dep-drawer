package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPkg() string {
	path := GetPath()

	return RemoveGoDir(path)
}

func GetPath() string {
	path := "."

	path, err := filepath.Abs(path)
	if err != nil {
		fmt.Errorf("Cannot get absolute path from %s", path)
		return ""
	}

	return path
}

func RemoveGoDir(path string) string {
	p, _ := filepath.Abs(path)
	p = strings.Replace(p, os.Getenv("GOPATH")+"/src/", "", 1)
	return p
}
