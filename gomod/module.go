package gomod

import (
	"fmt"
	"os"
)

func ModuleName() (string, error) {
	gomodPath, err := Find()
	if err != nil {
		return "", err
	}

	file, err := os.Open(gomodPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var mod string
	_, err = fmt.Fscanf(file, "module %s", &mod)
	if err != nil {
		return "", err
	}

	return mod, nil
}
