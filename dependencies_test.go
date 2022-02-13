package main

import (
	"sort"
	"testing"
)

func TestGetDependencies(t *testing.T) {
	result := GetDependencies(".")

	got := result["github.com/egon12/dep-drawer/."]
	want := []string{
		"flag",
		"bytes",
		"os",
		"strings",
		"time",
		"io/ioutil",
		"github.com/egon12/dep-drawer/gomod",
		"html/template",
		"errors",
		"github.com/egon12/dep-drawer/browser",
		"compress/gzip",
		"fmt",
		"path/filepath",
		"go/ast",
		"io",
		"go/parser",
		"go/token",
		"log",
	}

	sort.Strings(got)
	sort.Strings(want)

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("want: %v\n got: %v\n", want[i], got[i])
		}
	}
}

func TestGetRecursiveDependencies(t *testing.T) {
	a := GetRecursiveDependencies("diffpackage")
	for k, v := range a {
		t.Errorf("%v: %v", k, v)
	}
}
