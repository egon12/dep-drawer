package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	IgnoreInit()
	ShortenInit()
	ColorerInit()

	flag.Parse()

	dep := GetRecursiveDependencies(GetPath())
	//dep = RemoveMissingImport(dep)
	//dep = OuterPackageGrouper(dep, GetPkg())
	dep = GroupStdlibDependency(dep)
	dep = OuterPackageAdder(dep, GetPkg())
	dep = AddColor(dep)
	result := PrintForDAG(dep)
	result = Shorten(result)
	ShowInDagBrowser(result)
}

func RemoveMissingImport(dependencies map[string][]string) map[string][]string {

	for k, v := range dependencies {
		lenv := len(v)
		newImport := make([]string, lenv, lenv)
		index := 0
		for _, i := range v {
			_, ok := dependencies[i]
			if ok {
				newImport[index] = i
				index += 1
			}
		}
		dependencies[k] = newImport
	}
	return dependencies

}

func PrintForDAG(dependencies map[string][]string) string {
	result := ""
	for k, v := range dependencies {
		importList := strings.Join(v, " ")
		importList = strings.Replace(importList, "\"", "", -1)
		result = result + fmt.Sprintf("%s %s\n", k, importList)
	}
	return result
}

func GetRecursiveDependencies(rootPath string) map[string][]string {

	result := map[string][]string{}

	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			return nil
		}

		if IsIgnored(path) {
			return nil
		}

		for k, v := range GetDependencies(path) {
			result[k] = v
		}
		return nil
	})

	return result
}
