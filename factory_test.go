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

// patchFactoryRoot changes the factory root to a temporary
// directory.
//
// call the return function in a deferral to return the factory root
// to the original state.
func patchFactoryRoot(t *testing.T) func() {
	previousFactoryRoot := factoryRoot
	d, err := ioutil.TempDir("", "factory-root")
	if err != nil {
		t.Fatal(err)
	}
	factoryRoot = d
	return func() {
		if err := os.RemoveAll(d); err != nil {
			t.Fatal(err)
		}
		factoryRoot = previousFactoryRoot
	}
}

// newTestFactory creates a factory instance with a mocked git clone.
func newTestFactory(t *testing.T) *factory {
	mockCloneFunc := func(path string, bare bool, opts *git.CloneOptions) (*git.Repository, error) {
		manifest := []byte(`name: test
version: 0.0.1
dependencies:
- catalog: https://test-catalog.trustacks.io
  components: []`)
		if err := os.MkdirAll(filepath.Join(path, "chart"), 0755); err != nil {
			return nil, err
		}
		if err := ioutil.WriteFile(filepath.Join(path, "factory.yaml"), manifest, 0666); err != nil {
			return nil, err
		}
		return nil, nil
	}
	factory, err := newFactory("test", "https://test.com/test.git", "1.0.0", mockCloneFunc)
	if err != nil {
		t.Fatal(err)
	}
	return factory
}

func TestFactoryGetCatalog(t *testing.T) {
	patchFactoryRoot(t)
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
	factory := newTestFactory(t)
	catalog, err := factory.getCatalog(ts.URL)
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

func TestFactoryAddSubcharts(t *testing.T) {
	defer patchFactoryRoot(t)()
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
	factory := newTestFactory(t)
	if err := factory.addSubcharts(catalog); err != nil {
		t.Fatal(err)
	}
	_, err := os.Stat(fmt.Sprintf("%s/chart/charts/helloworld", factory.path()))
	if os.IsNotExist(err) {
		t.Fatal("expected chart to exist")
	}
}

func TestAddHooks(t *testing.T) {
	defer patchFactoryRoot(t)()
	catalog := &componentCatalog{
		HookSource: "quay.io/trustacks/test-catalog:latest",
		Components: map[string]component{
			"helloworld": {
				Hooks: []string{"post-install"},
			},
		},
	}
	factory := newTestFactory(t)
	cmd := exec.Command("cp", "-R", "testdata/helloworld", fmt.Sprintf("%s/chart/charts/helloworld", factory.path()))
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
	if err := factory.addHooks(catalog); err != nil {
		t.Fatal(err)
	}
	_, err := os.Stat(fmt.Sprintf("%s/chart/charts/helloworld/templates/post-install-trustacks.io.yaml", factory.path()))
	if os.IsNotExist(err) {
		t.Fatal("expected post install hook to exist")
	}
}

func TestFactoryAddSubchartValues(t *testing.T) {
	defer patchFactoryRoot(t)()
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
	factory := newTestFactory(t)
	if err := factory.addSubChartValues(catalog, parameters); err != nil {
		t.Fatal(err)
	}
	values, err := ioutil.ReadFile(filepath.Join(factory.path(), "chart", "values.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	expectedValues := `helloworld: {"password":"password","username":"username"}
` // don't delete his newline or the test will break. ;-)

	if string(values) != expectedValues {
		t.Fatal("got an unexpected values output")
	}
}
