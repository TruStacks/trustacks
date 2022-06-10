package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/go-git/go-git/v5"
)

// patchToolchainRoot changes the toolchain root to a temporary
// directory.
//
// call the return function in a deferral to return the toolchain root
// to the original state.
func patchToolchainRoot(t *testing.T) func() {
	previousToolchainRoot := toolchainRoot
	d, err := ioutil.TempDir("", "toolchain-root")
	if err != nil {
		t.Fatal(err)
	}
	toolchainRoot = d
	return func() {
		if err := os.RemoveAll(d); err != nil {
			t.Fatal(err)
		}
		toolchainRoot = previousToolchainRoot
	}
}

// newTestToolchain creates a toolchain instance with a mocked git clone.
func newTestToolchain(t *testing.T) *toolchain {
	mockCloneFunc := func(path string, bare bool, opts *git.CloneOptions) (*git.Repository, error) {
		manifest := []byte(`name: test
version: 0.0.1
dependencies:
- catalog: https://test-catalog.trustacks.io
  components: []`)
		if err := os.MkdirAll(filepath.Join(path, "chart"), 0755); err != nil {
			return nil, err
		}
		if err := ioutil.WriteFile(filepath.Join(path, "config.yaml"), manifest, 0666); err != nil {
			return nil, err
		}
		return nil, nil
	}
	toolchain, err := newToolchain("test", "https://test.com/test.git", "1.0.0", mockCloneFunc)
	if err != nil {
		t.Fatal(err)
	}
	return toolchain
}

func TestToolchainGetCatalog(t *testing.T) {
	patchToolchainRoot(t)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
  "hookSource":"quay.io/trustacks/test:latest",
  "components":{
    "test":{
      "repository":"https://test-charts.trustacks.io",
      "chart":"test/test",
      "version":"1.1.1"
    }
  }
}`))
	}))
	toolchain := newTestToolchain(t)
	catalog, err := toolchain.getCatalog(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	if catalog.HookSource != "quay.io/trustacks/test:latest" {
		t.Fatal("got an unexpected hook source")
	}
	if catalog.Components["test"].Repo != "https://test-charts.trustacks.io" {
		t.Fatal("got an unexpected repo")
	}
	if catalog.Components["test"].Chart != "test/test" {
		t.Fatal("got an unexpected chart")
	}
	if catalog.Components["test"].Version != "1.1.1" {
		t.Fatal("got an unexpected chart")
	}
}

func TestToolchainAddSubcharts(t *testing.T) {
	defer patchToolchainRoot(t)()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/charts/helloworld-1.0.0.tgz":
			data, err := ioutil.ReadFile("testdata/helloworld-1.0.0.tgz")
			if err != nil {
				t.Fatal(err)
			}
			_, err = w.Write(data)
			if err != nil {
				t.Fatal(err)
			}
		}
	}))
	catalog := &componentCatalog{
		Components: map[string]component{
			"helloworld": {
				Repo:    fmt.Sprintf("%s/charts", ts.URL),
				Chart:   "helloworld",
				Version: "1.0.0",
				Hooks:   []string{"post-install"},
			},
		},
	}
	toolchain := newTestToolchain(t)
	if err := toolchain.addSubcharts(catalog); err != nil {
		t.Fatal(err)
	}
	_, err := os.Stat(fmt.Sprintf("%s/chart/charts/helloworld", toolchain.path()))
	if os.IsNotExist(err) {
		t.Fatal("expected chart to exist")
	}
}

func TestAddHooks(t *testing.T) {
	defer patchToolchainRoot(t)()
	catalog := &componentCatalog{
		HookSource: "quay.io/trustacks/test-catalog:latest",
		Components: map[string]component{
			"helloworld": {
				Hooks: []string{"post-install"},
			},
		},
	}
	toolchain := newTestToolchain(t)
	cmd := exec.Command("cp", "-R", "testdata/helloworld", fmt.Sprintf("%s/chart/charts/helloworld", toolchain.path()))
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
	if err := toolchain.addHooks(catalog); err != nil {
		t.Fatal(err)
	}
	_, err := os.Stat(fmt.Sprintf("%s/chart/charts/helloworld/templates/post-install-trustacks.io.yaml", toolchain.path()))
	if os.IsNotExist(err) {
		t.Fatal("expected post install hook to exist")
	}
}

func TestToolchainAddSubchartValues(t *testing.T) {
	defer patchToolchainRoot(t)()
	catalog := &componentCatalog{
		HookSource: "quay.io/trustacks/test-catalog:latest",
		Components: map[string]component{
			"helloworld": {
				Values: map[string]interface{}{
					"username": "{{ .username }}",
					"password": "{{ .password }}",
				},
			},
		},
	}
	parameters := map[string]interface{}{
		"username": "username",
		"password": "password",
	}
	toolchain := newTestToolchain(t)
	if err := toolchain.addSubChartValues(catalog, parameters); err != nil {
		t.Fatal(err)
	}
	values, err := ioutil.ReadFile(filepath.Join(toolchain.path(), "chart", "values.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	expectedValues := `helloworld: {"password":"password","username":"username"}
` // don't delete his newline or the test will break. ;-)

	if string(values) != expectedValues {
		t.Fatal("got an unexpected values output")
	}
}
