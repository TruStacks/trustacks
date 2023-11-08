package run

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/trustacks/trustacks/pkg/actions"
	"github.com/trustacks/trustacks/pkg/plan"
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
		plan.GetStage(plan.FeedbackStage),
		plan.GetStage(plan.PreleaseStage),
		plan.GetStage(plan.ReleaseStage),
	}
	assert.Contains(t, stages, plan.GetStage(plan.ReleaseStage))
	stages = removeReleaseStage(stages)
	assert.NotContains(t, stages, plan.GetStage(plan.ReleaseStage))
}

func TestRunCmdFromPlanIntegration(t *testing.T) {
	d, close := makeTestdata(t)
	defer close()
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
		Source: filepath.Join(d),
		Stages: []string{
			plan.GetStage(plan.FeedbackStage),
		},
	}); err != nil {
		t.Fatal(err)
	}
}

func TestRunCmdFromSourceIntegration(t *testing.T) {
	d, close := makeTestdata(t)
	defer close()
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	if err := RunCmd(&RunCmdOptions{
		Source: filepath.Join(d),
		Stages: []string{
			plan.GetStage(plan.FeedbackStage),
		},
	}); err != nil {
		t.Fatal(err)
	}
}
