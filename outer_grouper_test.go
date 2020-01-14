package main

import (
	"reflect"
	"testing"
)

func TestOuterGrouper(t *testing.T) {

	primitive := map[string][]string{
		"github.com/egon12/dep-drawer":         []string{"github.com/egon12/dep-drawer/browers"},
		"github.com/egon12/dep-drawer/browser": []string{"github.com/egon12/outer/pkg", "github.com/egon12/outer2/pkg"},
	}

	pkg := "github.com/egon12/dep-drawer"

	got := OuterPackageGrouper(primitive, pkg)

	want := map[string][]string{
		"github.com/egon12/dep-drawer":         []string{"github.com/egon12/dep-drawer/browers"},
		"github.com/egon12/dep-drawer/browser": []string{"vendor"},
		"vendor": []string{},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
