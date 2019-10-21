package main

import (
	"github.com/egon12/dep-drawer/browser"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func ShowInDagBrowser(input string) {
	t := template.New("dag.html")
	t, err := t.Parse(dagTemplate)
	if err != nil {
		log.Fatal(err)
	}

	f, err := ioutil.TempFile(os.TempDir(), "dep_drawer_*.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = t.Execute(f, struct{ DAG string }{input})
	if err != nil {
		log.Fatal(err)
	}

	browser.Open(f.Name())
}
