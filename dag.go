package main

//go:generate go-bindata resources
var dagTemplate string

func init() {
	b, _ := Asset("resources/index.html")
	dagTemplate = string(b)
}
