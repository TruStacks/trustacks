package internal

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/trustacks/trustacks/pkg/engine"
)

func PlanCmd(source, name string, force bool) error {
	if _, err := os.Stat(name); !force && !os.IsNotExist(err) {
		return errors.New("plan file already exists")
	}
	actionPlan, err := engine.New().CreateActionPlan(source)
	if err != nil {
		return fmt.Errorf("failed creating the action plan: %s", err)
	}
	if len(actionPlan.Actions) == 0 {
		fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#FFBF00")).Bold(true).Render("[WRN]") + " No actions could be matched in the provided source")
		return nil
	}
	spec, err := actionPlan.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to serialize the action plan to json: %s", err)
	}
	if err := os.WriteFile(name, []byte(spec), 0644); err != nil { //nolint:gosec,gomnd
		return fmt.Errorf("failed writing to the action plan file: %s", err)
	}
	fmt.Printf("%s plan filed saved at: %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("#897DBB")).Bold(true).Render("[INF]"), name)
	return nil
}
