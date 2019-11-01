package main

import (
	"flag"
)

type pkgToBeColored []string

func (c *pkgToBeColored) String() string {
	return "Package name to be colored"
}

func (c *pkgToBeColored) Set(value string) error {
	*c = append(*c, value)
	return nil
}

type pkgColorList []string

func (c *pkgColorList) String() string {
	return "Package name to be colored"
}

func (c *pkgColorList) Set(value string) error {
	*c = append(*c, value)
	return nil
}

var pkgToBeColoredVar pkgToBeColored
var pkgColorListVar pkgColorList

func AddColor(allValues map[string][]string) map[string][]string {
	if len(pkgToBeColoredVar) != len(pkgColorListVar) {
		panic("Need same size of pkg color and the color? need -cp and -cn")
	}

	result := allValues

	for i, _ := range pkgToBeColoredVar {
		p := pkgToBeColoredVar[i]
		c := pkgColorListVar[i]

		deps, ok := result[p]
		if !ok {
			panic("Cannot find package " + p)
		}

		result[p] = append(deps, "| "+c)
	}

	return result
}

func ColorerInit() {
	flag.Var(&pkgToBeColoredVar, "cp", "package name to be colored")
	flag.Var(&pkgColorListVar, "cn", "color name to the package")
}
