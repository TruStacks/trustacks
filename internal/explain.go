package internal

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/trustacks/trustacks/pkg/engine"
)

func ExplainCmd(path string) error {
	var data []byte
	var actionPlan *engine.ActionPlan
	if path == "" {
		var err error
		actionPlan, err = engine.New().CreateActionPlan("./")
		if err != nil {
			return err
		}
	} else {
		var err error
		data, err = os.ReadFile(path)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(data, &actionPlan); err != nil {
			return err
		}
	}
	if len(actionPlan.Actions) > 0 {
		actions := []*engine.Action{}
		for _, name := range actionPlan.Actions {
			actions = append(actions, engine.GetAction(name))
		}
		inputs := []engine.InputField{}
		for _, action := range actions {
			inputs = append(inputs, action.Inputs...)
		}
		fmt.Printf("\nActions:\n\n")
		for _, action := range actions {
			if action != nil {
				fmt.Printf("▸ %s - %s\n\n", lipgloss.NewStyle().Foreground(lipgloss.Color("#897DBB")).Render(action.DisplayName), action.Description)
			}
		}
		if len(inputs) > 0 {
			fmt.Printf("Inputs:\n\n")
			for _, input := range inputs {
				schema := engine.GetInput(string(input)).Schema()
				fmt.Printf("▸ %s - %s\n\n", lipgloss.NewStyle().Foreground(lipgloss.Color("#897DBB")).Render(string(input)), schema.Description)
			}
			fmt.Println("Run the following command to generate a configu input schema")
			fmt.Println(lipgloss.NewStyle().Render("\n  ⤷ tsctl config init\n"))
		}
	} else {
		fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#5555FF")).Bold(true).Render("[WRN]") + " No actions could be generated from the provided source")
	}
	return nil
}
