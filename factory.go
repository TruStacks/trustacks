package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
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

// factory represents a factory helm chart.
type factory struct {
	name string
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

// addChart downloads the charts and adds it to the factory sub
// charts.
func (c *factory) addSubchart(repo, name, version string) error {
	pull := action.NewPullWithOpts(action.WithConfig(&action.Configuration{}))
	pull.Settings = cli.New()
	pull.UntarDir = fmt.Sprintf("%s/chart/charts", c.path())
	pull.Untar = true
	url := fmt.Sprintf("%s/%s-%s.tgz", repo, name, version)
	_, err := pull.Run(url)
	if err != nil {
		return err
	}
	return os.Remove(fmt.Sprintf("%s/chart/charts/%s-%s.tgz", c.path(), name, version))
}

// path returns the filesystem path of the factory metadata.
func (c *factory) path() string {
	return filepath.Join(factoryRoot, c.name)
}

// newFactory creates a new factory chart instance.
func newFactory(name, source, tag string, cloneFunc func(string, bool, *git.CloneOptions) (*git.Repository, error)) (*factory, error) {
	chart := &factory{name}
	if _, err := cloneFunc(chart.path(), false, &git.CloneOptions{
		URL:           source,
		Depth:         1,
		SingleBranch:  true,
		ReferenceName: plumbing.NewTagReferenceName(tag),
	}); err != nil {
		return nil, err
	}
	if err := os.Mkdir(fmt.Sprintf("%s/chart/charts", chart.path()), 0755); err != nil {
		return nil, err
	}
	return chart, nil
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
			if _, err := newFactory(
				factoryInstallCmdName,
				factoryInstallCmdSource,
				factoryInstallCmdTag,
				git.PlainClone,
			); err != nil {
				log.Fatal(err)
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
