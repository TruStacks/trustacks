package secrets

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"

	"github.com/trustacks/trustacks/pkg"
	"gopkg.in/yaml.v3"
)

// patchRoot changes the secrets root to a temporary
// directory.
//
// call the return function in a deferral to return the secrets root
// to the original state.
func patchRootDir(t *testing.T) func() {
	previousRootDir := pkg.RootDir
	d, err := ioutil.TempDir("", "root")
	if err != nil {
		t.Fatal(err)
	}
	pkg.RootDir = d
	return func() {
		if err := os.RemoveAll(d); err != nil {
			t.Fatal(err)
		}
		pkg.RootDir = previousRootDir
	}
}

// patchSecretsRoot changes the secrets root to a temporary
// directory.
//
// call the return function in a deferral to return the secrets root
// to the original state.
func patchSecretsRoot(t *testing.T) func() {
	previousSecretsRoot := secretsRoot
	d, err := ioutil.TempDir("", "secrets-root")
	if err != nil {
		t.Fatal(err)
	}
	secretsRoot = d
	return func() {
		if err := os.RemoveAll(d); err != nil {
			t.Fatal(err)
		}
		secretsRoot = previousSecretsRoot
	}
}

func TestGenerateKey(t *testing.T) {
	defer patchRootDir(t)()
	if err := generateKey(); err != nil {
		t.Fatal(err)
	}
	private, err := ioutil.ReadFile(path.Join(pkg.RootDir, "age.key"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(private), "AGE-SECRET-KEY") {
		t.Fatal("expected age secret key")
	}
	public, err := ioutil.ReadFile(path.Join(pkg.RootDir, "age.pub"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(public), "age") {
		t.Fatal("expected age public key")
	}

	// test noop.
	if err := generateKey(); err != nil {
		t.Fatal(err)
	}
	noop, err := ioutil.ReadFile(path.Join(pkg.RootDir, "age.key"))
	if err != nil {
		t.Fatal(err)
	}
	if string(noop) != string(private) {
		t.Fatal("expected noop")
	}
}

type testSecret struct {
	Data map[string]string `yaml:"data"`
}

func TestEncryptSecret(t *testing.T) {
	defer patchRootDir(t)()
	defer patchSecretsRoot(t)()
	if err := generateKey(); err != nil {
		t.Fatal(err)
	}
	if err := encryptSecret("test", path.Join("testdata", "secret.yaml")); err != nil {
		t.Fatal(err)
	}
	data, err := ioutil.ReadFile(path.Join(secretsRoot, "test.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	var secret testSecret
	if err := yaml.Unmarshal(data, &secret); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(secret.Data["username"], "ENC[AES256_GCM") {
		t.Fatal("expected the secret to be encrypted")
	}
	if err := encryptSecret("test", path.Join("testdata", "secret.yaml")); err == nil {
		t.Fatal("expected a duplicate error")
	}
}

func TestDownloadSops(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`#!/bin/sh
echo "test"
`))
	}))
	d, err := ioutil.TempDir("", "bin")
	if err != nil {
		t.Fatal(err)
	}
	if err := DownloadSops(ts.URL, d); err != nil {
		t.Fatal(err)
	}
	var outBuf, errBuf bytes.Buffer
	cmd := exec.Command(path.Join(d, "sops"))
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	if err := cmd.Run(); err != nil {
		t.Fatal(errBuf.String())
	}
	if outBuf.String() != "test\n" {
		t.Fatal("got an unexpected command output")
	}
}
