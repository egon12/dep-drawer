package main

import (
	"flag"
	"log"
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
	log.Println(allValues)

	result := allValues

	for i, _ := range pkgToBeColoredVar {
		p := pkgToBeColoredVar[i]
		c := pkgColorListVar[i]

		result[p] = append(result[p], "| "+c)
	}

	log.Println("result")
	log.Println(result)
	return result
}

func ColorerInit() {
	flag.Var(&pkgToBeColoredVar, "cp", "package name to be colored")
	flag.Var(&pkgColorListVar, "cn", "color name to the package")
}
