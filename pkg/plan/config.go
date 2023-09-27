package plan

import (
	"os"
	"path/filepath"

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

type Config struct {
	Common ConfigCommon `toml:"common"`
	Python ConfigPython `toml:"python"`
}

func NewConfig() (*Config, error) {
	var config Config
	if _, err := os.Stat(filepath.Join(configPath)); os.IsNotExist(err) {
		return &Config{}, nil
	}
	data, err := os.ReadFile(filepath.Join(configPath))
	if err != nil {
		return nil, err
	}
	if err := toml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
