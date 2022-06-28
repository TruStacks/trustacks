package toolchain

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
	"strings"

	"github.com/Masterminds/sprig/v3"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	helmclient "github.com/mittwald/go-helm-client"
	"github.com/trustacks/trustacks/pkg"
	"gopkg.in/yaml.v3"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

// toolchainRoot is where software toolchain metadata is stored.
var toolchainRoot = filepath.Join(pkg.RootDir, "toolchains")

// component represents a toolchain component.
type component struct {
	Repo       string                   `json:"repository"`
	Chart      string                   `json:"chart"`
	Version    string                   `json:"version"`
	Values     string                   `json:"values"`
	Hooks      []string                 `json:"hooks"`
	Parameters []map[string]interface{} `json:"parameters"`
}

type componentCatalogConfigParameters struct {
	Name    string `json:"name"`
	Default string `json:"default"`
}

type componentCatalogConfig struct {
	Parameters []componentCatalogConfigParameters `json:"parameters"`
}

// componentCatalog contains the component manifests.
type componentCatalog struct {
	HookSource string                  `json:"hookSource"`
	Components map[string]component    `json:"components"`
	Config     *componentCatalogConfig `json:"config"`
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
	resp, err := http.Get(fmt.Sprintf("%s/.well-known/catalog-manifest", url))
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
func (f *toolchain) addSubcharts(components []string, catalog *componentCatalog) error {
	for _, name := range components {
		component := catalog.Components[name]
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
func (f *toolchain) addHooks(components []string, catalog *componentCatalog) error {
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
	for _, name := range components {
		component := catalog.Components[name]
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

// addSubchartValues adds the subchart values to the helm values
// file.
func (f *toolchain) addSubChartValues(components []string, catalog *componentCatalog, parameters map[string]interface{}) error {
	values, err := os.Create(filepath.Join(f.path(), "chart", "values.yaml"))
	if err != nil {
		return err
	}
	for _, name := range components {
		component := catalog.Components[name]
		t := template.Must(template.New("values").Funcs(sprig.FuncMap()).Parse(component.Values))
		var buf bytes.Buffer
		if err := t.Execute(&buf, parameters); err != nil {
			return err
		}
		if _, err := values.Write([]byte(name + ":\n")); err != nil {
			return err
		}
		for _, line := range strings.Split(buf.String(), "\n") {
			if _, err := values.WriteString(fmt.Sprintf("  %s\n", line)); err != nil {
				return err
			}
		}
	}
	return nil
}

// install runs helm install and installs the toolchain.
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

type toolchainConfig struct {
	Parameters map[string]interface{} `json:"parameters"`
}

func (config *toolchainConfig) join(parameters []componentCatalogConfigParameters) map[string]interface{} {
	joined := make(map[string]interface{})
	for _, param := range parameters {
		if _, ok := config.Parameters[param.Name]; !ok {
			if param.Default != "" {
				joined[param.Name] = param.Default
			}
		} else {
			joined[param.Name] = config.Parameters[param.Name]
		}
	}
	return joined
}

// loadToolchainConfig loads the config file at the provided path.
func loadToolchainConfig(path string) (*toolchainConfig, error) {
	rawConfig, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config *toolchainConfig
	if err := yaml.Unmarshal(rawConfig, &config); err != nil {
		return nil, err
	}
	return config, nil
}

// Install installs the toolchain.
func Install(name, source, version, configPath string, cloneFunc func(string, bool, *git.CloneOptions) (*git.Repository, error)) error {
	config, err := loadToolchainConfig(configPath)
	if err != nil {
		return fmt.Errorf("error loading the toolchain config: %s", err)
	}
	toolchain, err := newToolchain(name, source, version, cloneFunc)
	if err != nil {
		return fmt.Errorf("error creating the toolchian: %s", err)
	}
	for _, dep := range toolchain.Dependencies {
		catalog, err := toolchain.getCatalog(dep.Catalog)
		if err != nil {
			return fmt.Errorf("error fetching catalog: %s", err)
		}
		if err := toolchain.addSubChartValues(dep.Components, catalog, config.Parameters); err != nil {
			return fmt.Errorf("error adding subchart values: %s", err)
		}
		if err := toolchain.addSubcharts(dep.Components, catalog); err != nil {
			return fmt.Errorf("error adding subcharts: %s", err)
		}
		if err := toolchain.addHooks(dep.Components, catalog); err != nil {
			return fmt.Errorf("error adding hook templates: %s", err)
		}
		if err := toolchain.install(); err != nil {
			return fmt.Errorf("error installing the toolchain chart: %s", err)
		}
	}
	return nil
}
