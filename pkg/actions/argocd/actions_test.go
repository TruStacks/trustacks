package argocd

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"dagger.io/dagger"
	"github.com/stretchr/testify/assert"
	"github.com/trustacks/trustacks/pkg/plan"
)

func TestExtraGlobalOptions(t *testing.T) {
	config := &plan.Config{
		ArgoCD: plan.ConfigArgoCD{
			Insecure: true,
			GRPCWeb:  true,
		},
	}
	assert.Contains(t, extraGlobalOptions(config), "--insecure", "--grpc-web")
}

func TestGetArgoApplicationInfoIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	src, err := os.MkdirTemp("", "get-argo-application-info")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(src)
	applicationYaml := `apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: test-argo-app
  namespace: argo-cd
spec:
  source:
    repoURL: https://github.com/trustacks/test-argo-app.git
    path: helm
  destination:
    namespace: test-argo-app
    name: test
`
	if err := ioutil.WriteFile(filepath.Join(src, "application.yaml"), []byte(applicationYaml), 0644); err != nil {
		t.Fatal(err)
	}
	client, err := dagger.Connect(context.Background(), dagger.WithLogOutput(os.Stdout))
	if err != nil {
		log.Fatal(err)
	}
	container := client.Container().
		From("argoproj/argocd").
		WithMountedDirectory("/src", client.Host().Directory(src)).
		WithWorkdir("/src")
	path, name, err := getArgoApplicationInfo(container)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "application.yaml", path)
	assert.Equal(t, "test-argo-app", name)
}
