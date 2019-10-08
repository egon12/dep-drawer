package main

import (
	"github.com/egon12/dep-drawer/browser"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func ShowInDagBrowser(dag string) {
	t, _ := template.ParseFiles("dag.html")

	f, err := ioutil.TempFile(os.TempDir(), "dep_drawer_*.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = t.Execute(f, struct{ DAG string }{dag})
	if err != nil {
		log.Fatal(err)
	}

	browser.Open(f.Name())
}
