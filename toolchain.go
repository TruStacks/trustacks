package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Masterminds/sprig/v3"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/manifoldco/promptui"
	helmclient "github.com/mittwald/go-helm-client"
	"gopkg.in/yaml.v3"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

// toolchainRoot is where software toolchain metadata is stored.
var toolchainRoot = fmt.Sprintf("%s/%s/%s", os.Getenv("HOME"), ".trustacks", "toolchains")

// component represents a toolchain component.
type component struct {
	Repo       string                   `json:"repository"`
	Chart      string                   `json:"chart"`
	Version    string                   `json:"version"`
	Hooks      []string                 `json:"hooks"`
	Values     map[string]interface{}   `json:"values"`
	Parameters []map[string]interface{} `json:"parameters"`
}

// componentCatalog contains the component manifests.
type componentCatalog struct {
	HookSource string               `json:"hookSource"`
	Components map[string]component `json:"components"`
}

// toolchainDependencies contains the catalog and required components.
type toolchainDependencies struct {
	Catalog    string   `yaml:"catalog"`
	Components []string `yaml:"components"`
}

// toolchain represents a toolchain helm chart.
type toolchain struct {
	name         string
	Dependencies []toolchainDependencies `yaml:"dependencies"`
}

// getCatalog gets the component catalog.
func (f *toolchain) getCatalog(url string) (*componentCatalog, error) {
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

// addSubcharts downloads the component charts and adds them to the
// toolchain subcharts.
func (f *toolchain) addSubcharts(catalog *componentCatalog) error {
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

// addHooks creates the hook template file in the toolchain chart.
func (f *toolchain) addHooks(catalog *componentCatalog) error {
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

// parseParameters prompts for the component parameters.
func (f *toolchain) parseParameters(catalog *componentCatalog) (map[string]interface{}, error) {
	parameters := map[string]interface{}{}
	for _, component := range catalog.Components {
		for _, parameter := range component.Parameters {
			templates := &promptui.PromptTemplates{
				Prompt:  "{{ . }}: ",
				Valid:   "{{ . }}: ",
				Invalid: "{{ . }}: ",
				Success: "{{ . }}: ",
			}
			var mask rune
			if _, ok := parameter["mask"]; ok && parameter["mask"].(bool) {
				mask = '*'
			}
			prompt := promptui.Prompt{
				Templates: templates,
				Label:     parameter["name"],
				Mask:      mask,
			}
			result, err := prompt.Run()
			if err != nil {
				return nil, err
			}
			parameters[parameter["name"].(string)] = result
		}
	}
	return parameters, nil
}

// addSubchartValues adds the subchart values to the helm values
// file.
func (f *toolchain) addSubChartValues(catalog *componentCatalog, parameters map[string]interface{}) error {
	values, err := os.Create(filepath.Join(f.path(), "chart", "values.yaml"))
	if err != nil {
		return err
	}
	for name, component := range catalog.Components {
		v, err := json.Marshal(component.Values)
		if err != nil {
			return err
		}
		t := template.Must(template.New("values").Funcs(sprig.FuncMap()).Parse(string(v)))
		var buf bytes.Buffer
		if err := t.Execute(&buf, parameters); err != nil {
			return err
		}
		if _, err := values.Write([]byte(name + ": ")); err != nil {
			return err
		}
		if _, err := values.Write([]byte(buf.String() + "\n")); err != nil {
			return err
		}
	}
	return nil
}

// install runs helm install and install the toolchain.
func (f *toolchain) install() error {
	name := fmt.Sprintf("toolchain-%s", f.name)
	chartSpec := helmclient.ChartSpec{
		ReleaseName:     name,
		ChartName:       filepath.Join(f.path(), "chart"),
		Namespace:       name,
		UpgradeCRDs:     true,
		CreateNamespace: true,
		CleanupOnFail:   true,
	}
	kubeconfig, err := ioutil.ReadFile(filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	if err != nil {
		return err
	}
	helmClient, err := helmclient.NewClientFromKubeConf(&helmclient.KubeConfClientOptions{
		Options:    &helmclient.Options{Namespace: name},
		KubeConfig: kubeconfig,
	})
	if err != nil {
		return err
	}
	_, err = helmClient.InstallOrUpgradeChart(context.Background(), &chartSpec, nil)
	return err
}

// path returns the filesystem path of the toolchain metadata.
func (f *toolchain) path() string {
	return filepath.Join(toolchainRoot, f.name)
}

// newToolchain creates a new toolchain chart instance.
func newToolchain(name, source, version string, cloneFunc func(string, bool, *git.CloneOptions) (*git.Repository, error)) (*toolchain, error) {
	toolchain := &toolchain{name: name}
	if _, err := os.Stat(toolchain.path()); !os.IsNotExist(err) {
		return nil, fmt.Errorf("error: toolchain '%s' already exists", name)
	}
	if _, err := cloneFunc(toolchain.path(), false, &git.CloneOptions{
		URL:           source,
		Depth:         1,
		SingleBranch:  true,
		ReferenceName: plumbing.NewTagReferenceName(version),
	}); err != nil {
		return nil, err
	}
	if err := os.Mkdir(filepath.Join(toolchain.path(), "chart", "charts"), 0755); err != nil {
		return nil, err
	}
	manifest, err := ioutil.ReadFile(filepath.Join(toolchain.path(), "config.yaml"))
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(manifest, toolchain); err != nil {
		return nil, err
	}
	return toolchain, nil
}

// installToolchain installs the toolchain.
func installToolchain(name, source, version string, cloneFunc func(string, bool, *git.CloneOptions) (*git.Repository, error)) error {
	toolchain, err := newToolchain(name, source, version, cloneFunc)
	if err != nil {
		return err
	}
	for _, dep := range toolchain.Dependencies {
		catalog, err := toolchain.getCatalog(dep.Catalog)
		if err != nil {
			return fmt.Errorf("error fetching catalog: %s", err)
		}
		parameters, err := toolchain.parseParameters(catalog)
		if err != nil {
			return fmt.Errorf("error parsing parameters: %s", err)
		}
		if err := toolchain.addSubChartValues(catalog, parameters); err != nil {
			return fmt.Errorf("error adding subchart values: %s", err)
		}
		if err := toolchain.addSubcharts(catalog); err != nil {
			return fmt.Errorf("error adding subcharts: %s", err)
		}
		if err := toolchain.addHooks(catalog); err != nil {
			return fmt.Errorf("error adding hook templates: %s", err)
		}
		if err := toolchain.install(); err != nil {
			return fmt.Errorf("error installing the toolchain chart: %s", err)
		}
	}
	return nil
}
