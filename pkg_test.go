package main

import (
	"testing"
)

func TestGetPkg(t *testing.T) {
	if GetPkg() != "github.com/egon12/dep-drawer" {
		t.Errorf("got wrong package name, %v", GetPkg())
	}
}
