package internal

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/trustacks/trustacks/pkg/actions"
	"github.com/trustacks/trustacks/pkg/engine"
)

// makeTestdata creates the golang source fixutres programatically
// because the golang action pattern matche excludes the testdata
// directory.
func makeTestdata(t *testing.T) (string, func()) {
	d, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	goMod := `
	module test
	go 1.20`
	if err := os.WriteFile(filepath.Join(d, "go.mod"), []byte(goMod), 0744); err != nil {
		t.Fatal(err)
	}
	equalTest := `
	package test
	import "testing"	
	func Test(t *testing.T) { }`
	if err := os.WriteFile(filepath.Join(d, "run_test.go"), []byte(equalTest), 0744); err != nil {
		t.Fatal(err)
	}
	return d, func() { os.RemoveAll(d) }
}

func TestRemoveReleaseStages(t *testing.T) {
	stages := []string{
		engine.GetStage(engine.CommitStage),
		engine.GetStage(engine.DeployStage),
		engine.GetStage(engine.ReleaseStage),
	}
	assert.Contains(t, stages, engine.GetStage(engine.ReleaseStage))
	stages = removeReleaseStage(stages)
	assert.NotContains(t, stages, engine.GetStage(engine.ReleaseStage))
}

func TestRunCmdFromPlanIntegration(t *testing.T) {
	d, clean := makeTestdata(t)
	defer clean()
	f, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	if _, err := f.Write([]byte(`{"actions": ["golangTest"]}`)); err != nil {
		t.Fatal(err)
	}
	f.Close()
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	if err := RunCmd(&RunCmdOptions{
		Plan:   f.Name(),
		Source: d,
		Stages: []string{
			engine.GetStage(engine.CommitStage),
		},
	}); err != nil {
		t.Fatal(err)
	}
}

func TestRunCmdFromSourceIntegration(t *testing.T) {
	d, clean := makeTestdata(t)
	defer clean()
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	if err := RunCmd(&RunCmdOptions{
		Source: d,
		Stages: []string{
			engine.GetStage(engine.CommitStage),
		},
	}); err != nil {
		t.Fatal(err)
	}
}
