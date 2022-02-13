package main

import (
	"sort"
	"testing"
)

func TestGetRecursiveDependencies(t *testing.T) {
	a := GetRecursiveDependencies("diffpackage")
	for k, v := range a {
		t.Errorf("%v: %v", k, v)
	}
}

func TestGetImports(t *testing.T) {
	got := GetImports(".")

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

func TestGetImports_diffpackage(t *testing.T) {
	got := GetImports("./diffpackage")

	want := []string{
		"github.com/egon12/dep-drawer/diffpackage/diffpackage2",
		"github.com/egon12/dep-drawer/diffpackage/diffpackage3",
	}

	sort.Strings(got)

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("want: %v\n got: %v\n", want[i], got[i])
		}
	}
}
