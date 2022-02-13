package main

import (
	"fmt"
	"io/ioutil"
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

func IsPkg(path string) bool {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		// if cannot be read then it's not pkg
		return false
	}

	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".go" {
			return true
		}
	}

	return false
}

func GetModPkg(path string) string {
	rootPkg, err := gomod.ModuleName()
	if err != nil {
		panic(err)
	}

	return rootPkg + removeRootPackageDir(path)
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
	rootPkgDir, _ = filepath.Abs(rootPkgDir)
	p = strings.Replace(p, rootPkgDir, "", 1)
	return p
}
