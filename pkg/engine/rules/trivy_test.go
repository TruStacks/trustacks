package rules

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSonarProjectPropertiesExistsRule(t *testing.T) {
	t.Run("SonarProjectPropertiesExists is true", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, "sonar-project.properties"), []byte("sonar.organization=trustacks"), 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := sonarProjectPropertiesExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, SonarProjectPropertiesExists)
	})

	t.Run("SonarProjectPropertiesExists is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		fact, err := sonarProjectPropertiesExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, SonarProjectPropertiesExists)
	})
}
