package pkg

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var (
	// RootDir is the asset root directory.
	RootDir = filepath.Join(os.Getenv("HOME"), ".trustacks")

	// BinDir is the binary dependencies directory.
	BinDir = filepath.Join(RootDir, "bin")
)

// patchRoot changes the secrets root to a temporary
// directory.
//
// call the return function in a deferral to return the secrets root
// to the original state.
func PatchRootDir(t *testing.T) func() {
	previousRootDir := RootDir
	d, err := ioutil.TempDir("", "root")
	if err != nil {
		t.Fatal(err)
	}
	RootDir = d
	return func() {
		if err := os.RemoveAll(d); err != nil {
			t.Fatal(err)
		}
		RootDir = previousRootDir
	}
}
