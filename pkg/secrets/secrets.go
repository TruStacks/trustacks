package secrets

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"

	"filippo.io/age"
	"github.com/trustacks/trustacks/pkg"
)

// secretsRoot is where secrets metadata is stored.
var secretsRoot = path.Join(pkg.RootDir, "secrets")

// generateKey creates the age private and public keys.
func generateKey() error {
	if _, err := os.Stat(path.Join(pkg.RootDir, "age.key")); !os.IsNotExist(err) {
		return nil
	}
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(path.Join(pkg.RootDir, "age.key"), []byte(identity.String()), 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(path.Join(pkg.RootDir, "age.pub"), []byte(identity.Recipient().String()), 0600); err != nil {
		return err
	}
	return nil
}

// encryptSecret encrypts the kubernetse secret.
func encryptSecret(name, secret string) error {
	secretPath := path.Join(secretsRoot, fmt.Sprintf("%s.yaml", name))
	if _, err := os.Stat(secretPath); !os.IsNotExist(err) {
		return fmt.Errorf("secret '%s' alreay exists", name)
	}
	public, err := ioutil.ReadFile(path.Join(pkg.RootDir, "age.pub"))
	if err != nil {
		return err
	}
	var errBuf bytes.Buffer
	cmd := exec.Command(
		"sops",
		"-e",
		"--age",
		string(public),
		"--input-type",
		"yaml",
		"--encrypted-regex",
		"^(data|stringData)$",
		"--output",
		secretPath,
		secret,
	)
	cmd.Stderr = &errBuf
	if err := cmd.Run(); err != nil {
		return errors.New(errBuf.String())
	}
	return nil
}

// DownloadSops downloads the sops binary.
func DownloadSops(url, bin string) error {
	if _, err := os.Stat(path.Join(bin, "sops")); !os.IsNotExist(err) {
		return nil
	}
	if err := os.MkdirAll(path.Join(bin), 0755); err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	f, err := os.OpenFile(path.Join(bin, "sops"), os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	if _, err := io.Copy(f, resp.Body); err != nil {
		return err
	}
	f.Close()
	return nil
}

// New creates a new encrypted secret.
func New(name, secret string) error {
	if err := Initialize(); err != nil {
		return err
	}
	return encryptSecret(name, secret)
}

// Initialize the secrets assets.
func Initialize() error {
	if err := os.MkdirAll(secretsRoot, 0755); err != nil {
		return err
	}
	if err := generateKey(); err != nil {
		return err
	}
	return nil
}
