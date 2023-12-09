package python

import (
	"context"

	"dagger.io/dagger"
)

func InstallPythonDependencies(container *dagger.Container) (*dagger.Container, error) {
	entries, err := container.Directory("/src").Entries(context.Background())
	if err != nil {
		return container, err
	}
	installed := false
	container = container.WithExec([]string{"pip", "install", "--upgrade", "pip"})
EntriesLoop:
	for _, entry := range entries {
		if installed {
			break
		}
		switch entry {
		case "poetry.lock":
			container = container.
				WithExec([]string{"pip", "install", "poetry"}).
				WithExec([]string{"poetry", "install"})
			break EntriesLoop
		case "pdm.lock":
			container = container.
				WithExec([]string{"/bin/sh", "-c", "curl -sSL https://pdm-project.org/install-pdm.py | python -"}).
				WithExec([]string{"/root/.local/bin/pdm", "install", "-d"})
		case "requirements.txt":
			container = container.
				WithEnvVariable("XDG_CACHE_HOME", "/src/.cache/pip").
				WithExec([]string{"pip", "install", "-r", "requirements.txt"})
			break EntriesLoop
		default:
			continue
		}
		installed = true
	}
	return container, nil
}
