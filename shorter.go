package main

import (
	"flag"
	"strings"
)

type namesToBeShorten []string

func (n *namesToBeShorten) String() string {
	return "Dir names to be shorten"
}

func (n *namesToBeShorten) Set(value string) error {
	*n = append(*n, value)
	return nil
}

var namesToBeShortenVar namesToBeShorten

func Shorten(allValues string) string {
	result := allValues

	for _, s := range namesToBeShortenVar {
		firstLetter := string(s[0])
		result = strings.Replace(result, s, firstLetter, -1)
	}

	return result
}

func ShortenInit() {
	flag.Var(&namesToBeShortenVar, "s", "names to be shorten")
}
