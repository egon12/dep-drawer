package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/egon12/dep-drawer/gomod"
)

func GetPkg() string {
	_, err := gomod.Find()

	// if got no go.mod
	if err != nil {
		path := GetPath()
		return RemoveGoDir(path)
	}

	rootPkg, err := gomod.ModuleName()
	if err != nil {
		panic(err)
	}

	return rootPkg + removeRootPackageDir(GetPath())
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

func removeRootPackageDir(path string) string {
	p, _ := filepath.Abs(path)
	rootPkgDir, _ := gomod.FindDir()
	p = strings.Replace(p, rootPkgDir, "", 1)
	return p
}
