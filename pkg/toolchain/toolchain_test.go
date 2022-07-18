package toolchain

import (
	"path/filepath"
	"testing"

	"github.com/trustacks/trustacks/pkg"
)

func TestLoadToolchainConfig(t *testing.T) {
	config, err := loadToolchainConfig(filepath.Join("testdata", "config.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	if config.Parameters["test"].(string) != "value" {
		t.Fatal("got an unexpected config parameter value")
	}
}

func TestConfigJoinParameters(t *testing.T) {
	catalogConfig := pkg.ComponentCatalogConfig{
		Parameters: []pkg.ComponentCatalogConfigParameters{
			{Name: "test", Default: ""},
			{Name: "port", Default: "8080"},
		},
	}
	config := &toolchainConfig{
		Parameters: map[string]interface{}{
			"test": "value",
		},
	}
	joined := pkg.Join(config.Parameters, catalogConfig.Parameters)
	if joined["test"].(string) != "value" {
		t.Fatal("expected test value to be set")
	}
	if joined["port"].(string) != "8080" {
		t.Fatal("expected default port value to be set")
	}
}
