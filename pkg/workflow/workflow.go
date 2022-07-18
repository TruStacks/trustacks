package workflow

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/trustacks/trustacks/pkg"
	"github.com/trustacks/trustacks/pkg/toolchain"
	"gopkg.in/yaml.v3"
)

// workflowDependencies contains the catalog and required components.
type workflowDependencies struct {
	Catalog    string   `json:"catalog"`
	Components []string `json:"components"`
}

type workflow struct {
	Name         string                 `json:"name"`
	Dependencies []workflowDependencies `json:"dependencies"`
}

type workflowsSource struct {
	Workflows map[string]workflow `json:"workflows"`
}

func getWorkflowsSource(source, version string, cloneFunc func(string, bool, *git.CloneOptions) (*git.Repository, error)) (*workflowsSource, error) {
	catalog := &workflowsSource{}
	d, err := ioutil.TempDir("", "workflows-source")
	if err != nil {
		return nil, err
	}
	if _, err := cloneFunc(d, false, &git.CloneOptions{
		URL:           source,
		Depth:         1,
		SingleBranch:  true,
		ReferenceName: plumbing.NewTagReferenceName(version),
	}); err != nil {
		return nil, err
	}
	config, err := ioutil.ReadFile(path.Join(d, "config.yaml"))
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(config, &catalog); err != nil {
		return nil, err
	}
	return catalog, nil
}

type workflowConfig struct {
	Name       string                 `json:"name"`
	Workflow   string                 `json:"workflow"`
	Source     string                 `json:"source"`
	Version    string                 `json:"version"`
	Vars       map[string]interface{} `json:"vars"`
	Secrets    map[string]interface{} `json:"secrets"`
	Parameters map[string]interface{} `json:"parameters"`
}

// loadWorkflowConfig loads the config file at the provided path.
func loadWorkflowConfig(path string) (*workflowConfig, error) {
	rawConfig, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config *workflowConfig
	if err := yaml.Unmarshal(rawConfig, &config); err != nil {
		return nil, err
	}
	return config, nil
}

// InstallWorkflow installs the workflow in the toolchain.
func InstallWorkflow(toolchainName string, configPath string, cloneFunc func(string, bool, *git.CloneOptions) (*git.Repository, error)) error {
	config, err := loadWorkflowConfig(configPath)
	if err != nil {
		return fmt.Errorf("error loading the toolchain config: %s", err)
	}
	workflows, err := getWorkflowsSource(config.Source, config.Version, git.PlainClone)
	if err != nil {
		return fmt.Errorf("error fetching workflow source: %s", err)
	}
	workflow, ok := workflows.Workflows[config.Name]
	if !ok {
		return fmt.Errorf("error: the workflow '%s' was not found in the workflow source", config.Name)
	}
	toolchainPath := toolchain.WorkflowPath(toolchainName)
	if err := os.MkdirAll(toolchainPath, 0755); err != nil {
		return fmt.Errorf("error creating the workflow path: %s", err)
	}
	for _, dep := range workflow.Dependencies {
		catalog, err := pkg.GetCatalog(config.Source)
		if err != nil {
			return fmt.Errorf("error fetching catalog: %s", err)
		}
		parameters := pkg.Join(config.Parameters, catalog.Config.Parameters)
		if err := pkg.AddSubcharts(toolchainPath, dep.Components, catalog); err != nil {
			return fmt.Errorf("error adding subcharts: %s", err)
		}
		if err := pkg.AddHooks(toolchainPath, dep.Components, catalog, parameters); err != nil {
			return fmt.Errorf("error adding hook templates: %s", err)
		}
		if err := pkg.AddSubChartValues(toolchainPath, dep.Components, catalog, parameters); err != nil {
			return fmt.Errorf("error adding subchart values: %s", err)
		}
		if err := pkg.InstallResource("workflow", config.Name, toolchainPath); err != nil {
			return fmt.Errorf("error installing the toolchain chart: %s", err)
		}
	}
	return nil
}
