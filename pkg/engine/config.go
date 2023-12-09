package engine

import (
	"os"

	"github.com/pelletier/go-toml"
)

const (
	configPath = "./trustacks.toml"
)

type ConfigCommon struct {
	Version string `toml:"version"`
}

type ConfigPython struct {
	Version         string   `toml:"version"`
	Libraries       []string `toml:"libs"`
	DevRequirements string   `toml:"dev_reqs"`
}

type ConfigGolang struct {
	Version string `toml:"version"`
	LDFlags string `toml:"ldflags"`
}

type ConfigArgoCD struct {
	GRPCWeb  bool `toml:"grpcWeb"`
	Insecure bool `toml:"insecure"`
}

type Config struct {
	Common ConfigCommon `toml:"common"`
	Python ConfigPython `toml:"python"`
	Golang ConfigGolang `toml:"golang"`
	ArgoCD ConfigArgoCD `toml:"argocd"`
}

func NewConfig() (*Config, error) {
	var config Config
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &Config{}, nil
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	if err := toml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
