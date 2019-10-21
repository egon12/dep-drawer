package main

import (
	"flag"
	"strings"
)

type namesToBeIgnore []string

func (n *namesToBeIgnore) String() string {
	return "Dir names to be ignore"
}

func (n *namesToBeIgnore) Set(value string) error {
	*n = append(*n, value)
	return nil
}

var namesToBeIgnoreVar namesToBeIgnore

func IsIgnored(path string) bool {
	for _, s := range namesToBeIgnoreVar {
		if strings.Contains(path, s) {
			return true
		}
	}
	return false
}

func IgnoreInit() {
	flag.Var(&namesToBeIgnoreVar, "i", "names to be ignore")
}
