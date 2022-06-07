package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

// factoryRoot is where software factory metadata is stored.
var factoryRoot = fmt.Sprintf("%s/%s/%s", os.Getenv("HOME"), ".trustacks", "factories")

// component represents a factory component.
type component struct {
	Repo    string   `json:"repository"`
	Chart   string   `json:"chart"`
	Version string   `json:"version"`
	Values  string   `json:"values"`
	Hooks   []string `json:"hooks"`
}

// componentCatalog contains the component manifests.
type componentCatalog struct {
	HookSource string               `json:"hookSource"`
	Components map[string]component `json:"components"`
}

type factoryDependencies struct {
	Catalog    string   `yaml:"catalog"`
	Components []string `yaml:"components"`
}

// factory represents a factory helm chart.
type factory struct {
	name         string
	Dependencies []factoryDependencies `yaml:"dependencies`
}

// getCatalog gets the component catalog.
func (f *factory) getCatalog(url string) (*componentCatalog, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var catalog *componentCatalog
	if err := json.Unmarshal(data, &catalog); err != nil {
		return nil, err
	}
	return catalog, nil
}

// addChart downloads the component charts and adds them to the
// factory subcharts.
func (f *factory) addSubcharts(catalog *componentCatalog) error {
	for _, component := range catalog.Components {
		pull := action.NewPullWithOpts(action.WithConfig(&action.Configuration{}))
		pull.Settings = cli.New()
		pull.UntarDir = fmt.Sprintf("%s/chart/charts", f.path())
		pull.Untar = true
		url := fmt.Sprintf("%s/%s-%s.tgz", component.Repo, component.Chart, component.Version)
		_, err := pull.Run(url)
		if err != nil {
			return err
		}
		if err := os.Remove(fmt.Sprintf("%s/chart/charts/%s-%s.tgz", f.path(), component.Chart, component.Version)); err != nil {
			return err
		}
	}

	return nil
}

func (f *factory) addHooks(catalog *componentCatalog) error {
	hookTemplate := `apiVersion: batch/v1
kind: Job
metadata:
  name: trustacks-{{.Kind}}
  annotations:
    "helm.sh/hook": {{.Kind}}
spec:
  template:
    spec:
      restartPolicy: Never
      containers:
      - name: trustacks-{{.Kind}}
      image: {{.Image}}
      env:
      - name: CATALOG_MODE
        value: hook
      - name: HOOK_COMPONENT
        value: {{.Component}}
      - name: HOOK_KIND
        value: {{.Kind}}`

	t := template.Must(template.New("hook").Parse(hookTemplate))
	type job struct {
		Kind, Image, Component string
	}
	for name, component := range catalog.Components {
		for _, kind := range component.Hooks {
			var buf bytes.Buffer
			if err := t.Execute(&buf, job{kind, catalog.HookSource, name}); err != nil {
				return err
			}
			path := fmt.Sprintf("%s/chart/charts/%s/templates/%s-trustacks.io.yaml", f.path(), name, kind)
			if err := ioutil.WriteFile(path, buf.Bytes(), 0666); err != nil {
				return err
			}
		}
	}
	return nil
}

// path returns the filesystem path of the factory metadata.
func (f *factory) path() string {
	return filepath.Join(factoryRoot, f.name)
}

// newFactory creates a new factory chart instance.
func newFactory(name, source, tag string, cloneFunc func(string, bool, *git.CloneOptions) (*git.Repository, error)) (*factory, error) {
	factory := &factory{name: name}
	if _, err := cloneFunc(factory.path(), false, &git.CloneOptions{
		URL:           source,
		Depth:         1,
		SingleBranch:  true,
		ReferenceName: plumbing.NewTagReferenceName(tag),
	}); err != nil {
		return nil, err
	}
	if err := os.Mkdir(filepath.Join(factory.path(), "chart", "charts"), 0755); err != nil {
		return nil, err
	}
	manifest, err := ioutil.ReadFile(filepath.Join(factory.path(), "factory.yaml"))
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(manifest, factory); err != nil {
		return nil, err
	}
	return factory, nil
}

// factory cli command flags.
var (
	factoryInstallCmdName   string
	factoryInstallCmdSource string
	factoryInstallCmdTag    string
)

// factory cli commands.
var (
	factoryCmd = &cobra.Command{
		Use:   "factory",
		Short: "manage software factories",
	}
	factoryInstallCmd = &cobra.Command{
		Use:   "install",
		Short: "install a new software factory",
		Run: func(cmd *cobra.Command, args []string) {
			factory, err := newFactory(
				factoryInstallCmdName,
				factoryInstallCmdSource,
				factoryInstallCmdTag,
				git.PlainClone,
			)
			if err != nil {
				log.Fatal(err)
			}
			for _, dep := range factory.Dependencies {
				catalog, err := factory.getCatalog(dep.Catalog)
				if err != nil {
					log.Fatal("error fetching catalog: ", err)
				}
				if err := factory.addSubcharts(catalog); err != nil {
					log.Fatal("error adding subcharts: ", err)
				}
				if err := factory.addHooks(catalog); err != nil {
					log.Fatal("error adding hook templates: ", err)
				}
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(factoryCmd)
	factoryCmd.AddCommand(factoryInstallCmd)
	factoryInstallCmd.Flags().StringVar(&factoryInstallCmdName, "name", "", "name of the factory")
	factoryInstallCmd.MarkFlagRequired("name")
	factoryInstallCmd.Flags().StringVar(&factoryInstallCmdSource, "source", "", "software factory git repository")
	factoryInstallCmd.MarkFlagRequired("source")
	factoryInstallCmd.Flags().StringVar(&factoryInstallCmdTag, "version", "", "software factory version")
	factoryInstallCmd.MarkFlagRequired("ref")
}
