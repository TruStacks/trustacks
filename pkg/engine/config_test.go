package engine

import (
	"os"
	"testing"

	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	config := &Config{
		Common: ConfigCommon{
			Version: "1.1.1",
		},
	}
	data, err := toml.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile("trustacks.toml", data, 0744); err != nil {
		t.Fatal(err)
	}
	defer os.Remove("trustacks.toml")
	conf, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "1.1.1", conf.Common.Version)
}
