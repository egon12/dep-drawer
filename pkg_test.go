package main

import (
	"fmt"
	"testing"
)

func TestGetPkg(t *testing.T) {
	if GetPkg() != "github.com/egon12/dep-drawer" {
		t.Errorf("got wrong package name, %v", GetPkg())
	}
}

func TestGetModPkg(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  ".",
			output: "github.com/egon12/dep-drawer",
		},
		{
			input:  "./diffpackage",
			output: "github.com/egon12/dep-drawer/diffpackage",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("input %s", tt.input), func(t *testing.T) {
			if GetModPkg(tt.input) != tt.output {
				t.Errorf("want: %s\n got: %s\n", tt.output, GetModPkg(tt.input))
			}
		})
	}
}
