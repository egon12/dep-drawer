package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

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

		if !IsPkg(path) {
			return nil
		}

		k := GetModPkg(path)
		v := GetImports(path)
		result[k] = v
		return nil
	})

	return result
}

func GetImports(path string) []string {
	fs := token.NewFileSet()
	f, err := parser.ParseDir(fs, path, filterTest, parser.ImportsOnly)
	if err != nil {
		panic(err)
	}

	result := make(map[string]struct{})

	for _, i := range f {
		ast.Inspect(i, func(n ast.Node) bool {
			x, ok := n.(*ast.ImportSpec)
			if ok {
				importName := strings.Replace(x.Path.Value, "\"", "", -1)
				result[importName] = struct{}{}
			}
			return true
		})
	}

	return mapKeyToSlice(result)
}

func mapKeyToSlice(input map[string]struct{}) []string {
	result := make([]string, len(input))
	index := 0
	for val, _ := range input {
		result[index] = val
		index += 1
	}
	return result
}

func filterTest(f os.FileInfo) bool {
	return !strings.HasSuffix(f.Name(), "_test.go")
}
