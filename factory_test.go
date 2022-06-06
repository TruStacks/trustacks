package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-git/go-git/v5"
)

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

func TestFactoryGetCatalog(t *testing.T) {
	defer patchFactoryRoot(t)()
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
	mockCloneFunc := func(path string, bare bool, opts *git.CloneOptions) (*git.Repository, error) {
		if err := os.MkdirAll(path+"/chart", 0755); err != nil {
			return nil, err
		}
		return nil, nil
	}
	factory, err := newFactory("test", "https://test.com/test.git", "master", mockCloneFunc)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(factory.path())
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

func TestFactoryAddSubChart(t *testing.T) {
	defer patchFactoryRoot(t)()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/charts/helloworld-1.0.0.tgz" {
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
	mockCloneFunc := func(path string, bare bool, opts *git.CloneOptions) (*git.Repository, error) {
		if err := os.MkdirAll(path+"/chart", 0755); err != nil {
			return nil, err
		}
		return nil, nil
	}
	factory, err := newFactory("test", "https://test.com/test.git", "master", mockCloneFunc)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(factory.path())
	repo := fmt.Sprintf("%s/charts", ts.URL)
	if err := factory.addSubchart(repo, "helloworld", "1.0.0"); err != nil {
		t.Fatal(err)
	}
}
