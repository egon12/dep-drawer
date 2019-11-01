package main

import (
	"testing"
)

func TestAddColor(t *testing.T) {
	pkgToBeColoredVar = []string{"g/e/A"}
	pkgColorListVar = []string{"red"}

	input := map[string][]string{
		"g/e/B": []string{},
		"g/e/C": []string{},
		"g/e/A": []string{"g/e/B", "g/e/C"},
	}

	got := AddColor(input)

	if len(got["g/e/A"]) != 3 {
		t.Error("Color is not aded")
	}

	if got["g/e/A"][2] != "| red" {
		t.Error("Color value is wrong")
	}

}
