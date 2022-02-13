package gomod

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFind_Success(t *testing.T) {
	_, err := Find()
	if err != nil {
		t.Error("It should not be error, because we have go.mod in parent ancestor")
	}
}

func TestFind_InAnotherCWD(t *testing.T) {
	oriPath, _ := filepath.Abs(".")

	os.Chdir(os.TempDir())
	defer os.Chdir(oriPath)

	_, err := Find()
	if err == nil {
		t.Errorf("It should be error, because we dont have go.mod in %s", os.TempDir())
	}
}

func Test_findRecursiveGoMod_InRoot(t *testing.T) {
	val, err := findRecursiveGoMod("/")
	if err == nil {
		t.Errorf("want error because we search in root, got nil and val %s", val)
	}
}

func BenchmarkFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Find()
	}
}
