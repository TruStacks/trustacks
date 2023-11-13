package engine

import (
	"context"
	"os"
	"testing"

	"dagger.io/dagger"
	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	utils := &ActionUtilities{config: &Config{
		Common: ConfigCommon{
			Version: "1.1.1",
		},
	}}
	assert.Equal(t, "1.1.1", utils.GetConfig().Common.Version)
}

func TestSetSecretIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	client, err := dagger.Connect(context.Background(), dagger.WithLogOutput(os.Stdout))
	if err != nil {
		t.Fatal(err)
	}
	utils := &ActionUtilities{client: client}
	secret := utils.SetSecret("test", "password1!")
	plainText, err := secret.Plaintext(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "password1!", plainText)
}
