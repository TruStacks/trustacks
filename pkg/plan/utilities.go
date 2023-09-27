package plan

import (
	"dagger.io/dagger"
)

type ActionUtilities struct {
	*ArtifactStore
	client *dagger.Client
	config *Config
}

func (util *ActionUtilities) SetSecret(name, plaintext string) *dagger.Secret {
	return util.client.SetSecret(name, plaintext)
}

func (util *ActionUtilities) GetConfig() *Config {
	return util.config
}

func newActionUtilities(client *dagger.Client, artifacts *ArtifactStore, config *Config) *ActionUtilities {
	return &ActionUtilities{artifacts, client, config}
}
