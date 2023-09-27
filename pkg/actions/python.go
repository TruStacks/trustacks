package actions

import (
	"context"

	"dagger.io/dagger"
)

func installPythonDependencies(container *dagger.Container) (*dagger.Container, error) {
	entries, err := container.Directory("/src").Entries(context.Background())
	if err != nil {
		return container, err
	}
	installed := false
	for _, entry := range entries {
		if installed {
			break
		}
		container = container.WithExec([]string{"pip", "install", "--upgrade", "pip"})
		switch entry {
		case "poetry.lock":
			container = container.
				WithExec([]string{"pip", "install", "poetry"}).
				WithExec([]string{"poetry", "install"})
		case "requirements.txt":
			container = container.WithEnvVariable("XDG_CACHE_HOME", "/src/.cache/pip")
			container = container.WithExec([]string{"pip", "install", "-r", "requirements.txt"})
		default:
			continue
		}
		installed = true
	}
	return container, nil
}
