package toolchain

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/trustacks/trustacks/pkg"
	"github.com/trustacks/trustacks/pkg/secrets"
	"gopkg.in/yaml.v3"
)

// toolchainRoot is where software toolchain metadata is stored.
var toolchainRoot = filepath.Join(pkg.RootDir, "toolchains")

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
	Name         string                 `json:"name"`
	Source       string                 `json:"source"`
	Version      string                 `json:"version"`
	Parameters   map[string]interface{} `json:"parameters"`
	LocalSecrets bool                   `json:"localSecrets"`
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
func Install(configPath string, cloneFunc func(string, bool, *git.CloneOptions) (*git.Repository, error)) error {
	config, err := loadToolchainConfig(configPath)
	if err != nil {
		return fmt.Errorf("error loading the toolchain config: %s", err)
	}
	if config.LocalSecrets {
		if err := secrets.Initialize(); err != nil {
			return err
		}
	}
	toolchain, err := newToolchain(config.Name, config.Source, config.Version, cloneFunc)
	if err != nil {
		return fmt.Errorf("error creating the toolchian: %s", err)
	}
	for _, dep := range toolchain.Dependencies {
		catalog, err := pkg.GetCatalog(dep.Catalog)
		if err != nil {
			return fmt.Errorf("error fetching catalog: %s", err)
		}
		parameters := pkg.Join(config.Parameters, catalog.Config.Parameters)
		if err := pkg.AddSubcharts(toolchain.path(), dep.Components, catalog); err != nil {
			return fmt.Errorf("error adding subcharts: %s", err)
		}
		if err := pkg.AddHooks(toolchain.path(), dep.Components, catalog, parameters); err != nil {
			return fmt.Errorf("error adding hook templates: %s", err)
		}
		if err := pkg.AddSubChartValues(toolchain.path(), dep.Components, catalog, parameters); err != nil {
			return fmt.Errorf("error adding subchart values: %s", err)
		}
		if err := pkg.InstallResource("toolchain", config.Name, toolchain.path()); err != nil {
			return fmt.Errorf("error installing the toolchain chart: %s", err)
		}
		if err := pkg.InstallResourceComponents("toolchain", config.Name, toolchain.path()); err != nil {
			return fmt.Errorf("error installing the toolchain components: %s", err)
		}
	}
	return nil
}

// WorkflowPath returns the filesystem path to the toolchain
// workflows.
func WorkflowPath(toolchain string) string {
	return filepath.Join(toolchainRoot, toolchain, "workflows")
}
