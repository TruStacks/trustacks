package toolchain

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"gopkg.in/yaml.v3"
)

// workflowDependencies contains the catalog and required components.
type workflowDependencies struct {
	Catalog    string   `json:"catalog"`
	Components []string `json:"components"`
}

// workflow contains workflow metadata and dependencies.
type workflow struct {
	Name         string                 `json:"name"`
	Dependencies []workflowDependencies `json:"dependencies"`
}

// workflowCatalog contains the list of application workflows.
type workflowCatalog struct {
	Workflows []*workflow `json:"workflows"`
}

// getWorkflowCatalog gets the catalog containing the workflows.
func getWorkflowCatalog(source, version string, cloneFunc func(string, bool, *git.CloneOptions) (*git.Repository, error)) (*workflowCatalog, error) {
	catalog := &workflowCatalog{}
	d, err := ioutil.TempDir("", "workflows-catalog")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(d)
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

// applicationConfig contains the parameters for the application.
type applicationConfig struct {
	Name     string            `json:"name"`
	Workflow string            `json:"workflow"`
	Source   string            `json:"source"`
	Version  string            `json:"version"`
	Vars     map[string]string `json:"vars"`
	Secrets  map[string]string `json:"secrets"`
}

// application is an instance of a toolchain application.
type application struct {
	toolchain *toolchain
	name      string
}

// createChart creates the application helm chart for installation.
func (app *application) createChart() error {
	if err := os.MkdirAll(path.Join(app.path(), "templates"), 0755); err != nil {
		return err
	}
	chart := fmt.Sprintf(`apiVersion: v1
version: 0.0.0
name: %s
`, app.name)
	if err := ioutil.WriteFile(path.Join(app.path(), "Chart.yaml"), []byte(chart), 0644); err != nil {
		return err
	}
	return nil
}

// addVars adds the application variables config map to
// the toolchain chart.
func (app *application) addVars(vars map[string]string) error {
	configMap := map[string]interface{}{
		"apiVersion": "v1",
		"kind":       "ConfigMap",
		"metadata": map[string]interface{}{
			"name": fmt.Sprintf("application-%s", app.name),
		},
		"data": vars,
	}
	data, err := yaml.Marshal(configMap)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(app.path(), "templates", "vars-configmap.yaml"), data, 0644)
}

// addSecrets adds the application varaibles secrets to
// the toolchain chart.
func (app *application) addSecrets(secrets map[string]string) error {
	secret := map[string]interface{}{
		"apiVersion": "v1",
		"kind":       "Secret",
		"metadata": map[string]interface{}{
			"name": fmt.Sprintf("application-%s", app.name),
		},
		"data": map[string]string{},
	}
	for k, v := range secrets {
		secret["data"].(map[string]string)[k] = base64.StdEncoding.EncodeToString([]byte(v))
	}
	data, err := yaml.Marshal(secret)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(app.path(), "templates", "secrets-secret.yaml"), data, 0644)
}

// path returns the application filesystem root.
func (app *application) path() string {
	return path.Join(app.toolchain.applicationsPath(), app.name)
}

// newApplication creates the application chart and input assets.
func newApplication(name string, config *applicationConfig, tc *toolchain, force bool) (*application, error) {
	app := &application{name: name, toolchain: tc}
	if _, err := os.Stat(app.path()); !os.IsNotExist(err) && !force {
		return nil, fmt.Errorf("error: application '%s' already exists", name)
	}
	if err := app.createChart(); err != nil {
		return nil, err
	}
	if err := app.addVars(config.Vars); err != nil {
		return nil, err
	}
	if err := app.addSecrets(config.Secrets); err != nil {
		return nil, err
	}
	return app, nil
}

// CreateApplication creates a new application instance and installs
// the application workflow dependencies.
func CreateApplication(name string, force bool, configPath string, cloneFunc func(string, bool, *git.CloneOptions) (*git.Repository, error)) error {
	config, err := loadToolchainConfig(configPath)
	if err != nil {
		return fmt.Errorf("error loading the toolchain config: %s", err)
	}
	var appConfig *applicationConfig
	for _, a := range config.Applications {
		if a.Name == name {
			appConfig = &a
			break
		}
	}
	if appConfig == nil {
		return fmt.Errorf("error: config for '%s' was not found in '%s'", name, configPath)
	}
	tc := &toolchain{name: config.Name}
	_, err = newApplication(name, appConfig, tc, force)
	if err != nil {
		return fmt.Errorf("error creating the application %s", err)
	}
	catalog, err := getWorkflowCatalog(appConfig.Source, appConfig.Version, cloneFunc)
	if err != nil {
		return fmt.Errorf("error getting the workflow catalog: %s", err)
	}
	var wf *workflow
	for _, workflow := range catalog.Workflows {
		if workflow.Name == appConfig.Workflow {
			wf = workflow
		}
	}
	if wf == nil {
		return fmt.Errorf("error: workflow '%s' was not found in the catalog", appConfig.Workflow)
	}
	for _, dep := range wf.Dependencies {
		catalog, err := getToolchainCatalog(dep.Catalog)
		if err != nil {
			return fmt.Errorf("error fetching catalog: %s", err)
		}
		parameters := tc.join(config.Parameters, catalog.Config.Parameters)
		if err := tc.addComponents(dep.Components, catalog); err != nil {
			return fmt.Errorf("error adding subcharts: %s", err)
		}
		if err := tc.addHooks(dep.Components, catalog, parameters); err != nil {
			return fmt.Errorf("error adding hook templates: %s", err)
		}
		if err := tc.addSubChartValues(dep.Components, catalog, parameters); err != nil {
			return fmt.Errorf("error adding subchart values: %s", err)
		}
		if err := tc.install(); err != nil {
			return fmt.Errorf("error installing the toolchain chart: %s", err)
		}
		if err := tc.installComponents(); err != nil {
			return fmt.Errorf("error installing the toolchain components: %s", err)
		}
	}
	return nil
}
