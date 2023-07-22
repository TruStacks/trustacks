package plan

import "dagger.io/dagger"

type ActionUtilities struct {
	*ArtifactStore
	client *dagger.Client
}

func (util *ActionUtilities) SetSecret(name, plaintext string) *dagger.Secret {
	return util.client.SetSecret(name, plaintext)
}

func newActionUtilities(client *dagger.Client, artifacts *ArtifactStore) *ActionUtilities {
	return &ActionUtilities{artifacts, client}
}
