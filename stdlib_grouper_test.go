package main

import (
	"reflect"
	"testing"
)

func TestStdlibGrouper(t *testing.T) {

	primitive := map[string][]string{
		"github.com/egon12/dep-drawer":         []string{"log", "fmt", "net/http", "go/ast", "github.com/egon12/dep-drawer/browers"},
		"github.com/egon12/dep-drawer/browser": []string{"fmt", "exec/command"},
	}

	got := GroupStdlibDependency(primitive)

	want := map[string][]string{
		"github.com/egon12/dep-drawer":         []string{"stdlib", "github.com/egon12/dep-drawer/browers"},
		"github.com/egon12/dep-drawer/browser": []string{"stdlib"},
		"stdlib":                               []string{},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
