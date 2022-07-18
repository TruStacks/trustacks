package pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/Masterminds/sprig/v3"
	helmclient "github.com/mittwald/go-helm-client"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

var (
	// RootDir is the asset root directory.
	RootDir = path.Join(os.Getenv("HOME"), ".trustacks")

	// BinDir is the binary dependencies directory.
	BinDir = path.Join(RootDir, "bin")
)

// component represents a toolchain component.
type Component struct {
	Repo       string                   `json:"repository"`
	Chart      string                   `json:"chart"`
	Version    string                   `json:"version"`
	Values     string                   `json:"values"`
	Hooks      string                   `json:"hooks"`
	Parameters []map[string]interface{} `json:"parameters"`
}

// ComponentCatalogConfigParameters .
type ComponentCatalogConfigParameters struct {
	Name    string `json:"name"`
	Default string `json:"default"`
}

// ComponentCatalogConfig .
type ComponentCatalogConfig struct {
	Parameters []ComponentCatalogConfigParameters `json:"parameters"`
}

// componentCatalog contains the component manifests.
type ComponentCatalog struct {
	HookSource string                  `json:"hookSource"`
	Components map[string]Component    `json:"components"`
	Config     *ComponentCatalogConfig `json:"config"`
}

// GetCatalog gets the component catalog.
func GetCatalog(url string) (*ComponentCatalog, error) {
	resp, err := http.Get(fmt.Sprintf("%s/.well-known/catalog-manifest", url))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var catalog *ComponentCatalog
	if err := json.Unmarshal(data, &catalog); err != nil {
		return nil, err
	}
	return catalog, nil
}

// AddSubcharts downloads the component charts and adds them to the
// resource subcharts.
func AddSubcharts(basePath string, components []string, catalog *ComponentCatalog) error {
	for _, name := range components {
		component := catalog.Components[name]
		pull := action.NewPullWithOpts(action.WithConfig(&action.Configuration{}))
		pull.Settings = cli.New()
		pull.UntarDir = path.Join(basePath, "components")
		pull.Untar = true
		url := fmt.Sprintf("%s/%s-%s.tgz", component.Repo, component.Chart, component.Version)
		_, err := pull.Run(url)
		if err != nil {
			return err
		}
		if err := os.Remove(path.Join(basePath, "components", fmt.Sprintf("%s-%s.tgz", component.Chart, component.Version))); err != nil {
			return err
		}
	}
	return nil
}

// AddHooks creates the hook template file in the chart.
func AddHooks(basePath string, components []string, catalog *ComponentCatalog, params map[string]interface{}) error {
	for _, name := range components {
		params["image"] = catalog.HookSource
		component := catalog.Components[name]

		var buf bytes.Buffer
		t := template.Must(template.New("hook").Parse(component.Hooks))
		if err := t.Execute(&buf, params); err != nil {
			return err
		}
		path := filepath.Join(basePath, "components", name, "templates", "trustacks-hooks.yaml")
		if err := ioutil.WriteFile(path, buf.Bytes(), 0666); err != nil {
			return err
		}
	}
	return nil
}

// AddSubchartValues adds the subchart values to the helm values
// file.
func AddSubChartValues(basePath string, components []string, catalog *ComponentCatalog, parameters map[string]interface{}) error {
	for _, name := range components {
		values, err := os.Create(path.Join(basePath, "components", name, "override-values.yaml"))
		if err != nil {
			return err
		}
		component := catalog.Components[name]
		t := template.Must(template.New("values").Funcs(sprig.FuncMap()).Parse(component.Values))
		var buf bytes.Buffer
		if err := t.Execute(&buf, parameters); err != nil {
			return err
		}
		if _, err := values.Write(buf.Bytes()); err != nil {
			return err
		}
	}
	return nil
}

func InstallResource(resource, slug, basePath string) error {
	slug = fmt.Sprintf("%s-%s", resource, slug)
	chartSpec := helmclient.ChartSpec{
		ReleaseName:     slug,
		ChartName:       filepath.Join(basePath, "chart"),
		Namespace:       slug,
		UpgradeCRDs:     true,
		CreateNamespace: true,
		CleanupOnFail:   true,
	}
	kubeconfig, err := ioutil.ReadFile(filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	if err != nil {
		return err
	}
	helmClient, err := helmclient.NewClientFromKubeConf(&helmclient.KubeConfClientOptions{
		Options:    &helmclient.Options{Namespace: slug},
		KubeConfig: kubeconfig,
	})
	if err != nil {
		return err
	}
	_, err = helmClient.InstallOrUpgradeChart(context.Background(), &chartSpec, nil)
	return err
}

// InstallResourceComponents installs the component helm charts.
func InstallResourceComponents(resource, slug, basePath string) error {
	slug = fmt.Sprintf("%s-%s", resource, slug)
	components, err := ioutil.ReadDir(path.Join(basePath, "components"))
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	for _, component := range components {
		wg.Add(1)
		go func(name string, wg *sync.WaitGroup) {
			values, err := ioutil.ReadFile(filepath.Join(basePath, "components", name, "override-values.yaml"))
			if err != nil {
				log.Fatalf("error reading override values: %s", err)
			}
			chartSpec := helmclient.ChartSpec{
				ReleaseName:     name,
				ChartName:       filepath.Join(basePath, "components", name),
				Namespace:       slug,
				UpgradeCRDs:     true,
				CreateNamespace: true,
				CleanupOnFail:   true,
				ValuesYaml:      string(values),
			}
			kubeconfig, err := ioutil.ReadFile(filepath.Join(os.Getenv("HOME"), ".kube", "config"))
			if err != nil {
				log.Fatalf("error reading kubeconfig: %s", err)
			}
			helmClient, err := helmclient.NewClientFromKubeConf(&helmclient.KubeConfClientOptions{
				Options:    &helmclient.Options{Namespace: slug},
				KubeConfig: kubeconfig,
			})
			if err != nil {
				log.Fatalf("error creating helm client: %s", err)
			}
			_, err = helmClient.InstallOrUpgradeChart(context.Background(), &chartSpec, nil)
			if err != nil {
				log.Fatalf("error deploying '%s': %s", name, err)
			}
			wg.Done()
		}(component.Name(), &wg)
	}
	wg.Wait()
	return nil
}

// Join combines the toolchain configuration parameters with the
// component parameters
//
// Parameter defaults are set if required.
func Join(parameters map[string]interface{}, catalogParameters []ComponentCatalogConfigParameters) map[string]interface{} {
	joined := make(map[string]interface{})
	for _, param := range catalogParameters {
		if _, ok := parameters[param.Name]; !ok {
			if param.Default != "" {
				joined[param.Name] = param.Default
			}
		} else {
			joined[param.Name] = parameters[param.Name]
		}
	}
	return joined
}
