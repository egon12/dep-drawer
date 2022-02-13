package gomod

import "testing"

func TestModuleName(t *testing.T) {
	val, err := ModuleName()
	if err != nil {
		t.Errorf("want no error got %v", err)
	}

	if val != "github.com/egon12/dep-drawer" {
		t.Errorf("got wrong module name: %s", val)
	}
}
