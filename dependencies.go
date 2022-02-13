package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func filterTest(f os.FileInfo) bool {
	return !strings.HasSuffix(f.Name(), "_test.go")
}

func GetDependencies(dirName string) map[string][]string {
	fs := token.NewFileSet()
	f, err := parser.ParseDir(fs, dirName, filterTest, parser.ImportsOnly)
	if err != nil {
		panic(err)
	}

	packagesImportSet := map[string]map[string]bool{}
	packageName := filepath.Base(dirName)

	for _, i := range f {
		//packageName = i.Name

		importSet, ok := packagesImportSet[packageName]
		if !ok {
			importSet = map[string]bool{}
		}

		ast.Inspect(i, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.ImportSpec:
				importName := strings.Replace(x.Path.Value, "\"", "", -1)
				importSet[importName] = true
			}
			return true
		})

		packagesImportSet[packageName] = importSet
	}

	packagesImportList := map[string][]string{}

	for k, v := range packagesImportSet {
		fullPacakgeName := RemoveGoDir(filepath.Dir(dirName)) + "/" + k
		importList := MapKeyIntoSlice(v)
		packagesImportList[fullPacakgeName] = importList
	}

	return packagesImportList

}

func MapKeyIntoSlice(theMap map[string]bool) []string {
	theMapLen := len(theMap)
	slice := make([]string, theMapLen, theMapLen)
	index := 0
	for k, _ := range theMap {
		slice[index] = k
		index += 1
	}
	return slice
}
