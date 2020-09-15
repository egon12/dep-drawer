package main

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"

	"strings"
	"testing"
	"text/template"

	"github.com/egon12/dep-drawer/browser"
)

func TestTemplate(t *testing.T) {
	t.Skip()
	temp, _ := template.ParseFiles("a.txt")
	b := &strings.Builder{}

	f, err := ioutil.TempFile(os.TempDir(), "dep_drawer_*.html")
	if err != nil {
		t.Error(err)
	}

	err = temp.Execute(f, struct{ Dag string }{"something"})
	if err != nil {
		t.Error(err)
	}

	browser.Open(f.Name())

	t.Error(b.String())

}

func TestAst(t *testing.T) {
	fs := token.NewFileSet()
	f, err := parser.ParseDir(fs, ".", nil, parser.ImportsOnly)
	if err != nil {
		panic(err)
	}
	t.Errorf("%T\n %+v", f, f)

}

func TestGetDependencies(t *testing.T) {
	a := GetRecursiveDependencies("diffpackage")
	for k, v := range a {
		t.Errorf("%v: %v", k, v)
	}
}
