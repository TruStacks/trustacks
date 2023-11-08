package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/trustacks/trustacks/pkg/engine"
	"github.com/trustacks/trustacks/pkg/plan"
)

func PlanCmd(source, name string) error {
	spec, err := engine.New().CreateActionPlan(source, true)
	if err != nil {
		return err
	}
	if err := os.WriteFile(name, []byte(spec), 0644); err != nil {
		return err
	}
	fmt.Printf("plan filed saved at: %s\n", name)
	return nil
}

func ConfigInitCmd(planFile string) error {
	planContents := map[string]interface{}{}
	contents, err := os.ReadFile(planFile)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(contents, &planContents); err != nil {
		return err
	}
	output := fmt.Sprintf("./%s.cfgu.json", strings.Replace(planFile, ".plan", "", 1))
	schema := map[string]interface{}{}
	if _, ok := planContents["inputs"]; ok {
		for _, field := range planContents["inputs"].([]interface{}) {
			input := plan.GetInput(field.(string))
			schema[field.(string)] = input.Schema()
		}
	}
	data, err := json.Marshal(schema)
	if err != nil {
		return err
	}
	var prettySchema bytes.Buffer
	if err := json.Indent(&prettySchema, data, "", "  "); err != nil {
		return err
	}
	if err := os.WriteFile(output, prettySchema.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}

func LoginCmd(host, username, password string) error {
	var err error
	var client struct {
		NewSessionToken func(string, string) (string, error)
	}
	closer, err := jsonrpc.NewClient(context.Background(), fmt.Sprintf("%s/rpc", host), "v1", &client, nil)
	if err != nil {
		return err
	}
	defer closer()
	token, err := client.NewSessionToken(username, password)
	if err != nil {
		return err
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := filepath.Join(homeDir, ".trustacks")
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(path, "auth-token"), []byte(token), 0600); err != nil {
		return err
	}
	fmt.Println("login successful")
	return nil
}

func ExplainCmd(path string, fromSource bool) error {
	var data []byte
	var actionPlan plan.ActionPlan
	if fromSource {
		spec, err := engine.New().CreateActionPlan("./", false)
		if err != nil {
			return err
		}
		data = []byte(spec)
	} else {
		var err error
		data, err = os.ReadFile(path)
		if err != nil {
			return err
		}
	}
	if err := json.Unmarshal(data, &actionPlan); err != nil {
		return err
	}
	if len(actionPlan.Actions) > 0 {
		fmt.Printf("\nActions\n-------\n\n")
		for _, action := range actionPlan.Actions {
			spec := engine.GetActionSpec(action)
			if spec != nil {
				fmt.Printf(
					"▸ %s - %s\n\n",
					lipgloss.NewStyle().Foreground(lipgloss.Color("#897dbb")).Render(spec.DisplayName),
					spec.Description,
				)
			}
		}
		if len(actionPlan.Inputs) > 0 {
			fmt.Printf("\nInputs\n------\n\n")
			for _, input := range actionPlan.Inputs {
				schema := plan.GetInput(input).Schema()
				fmt.Printf(
					"▸ %s - %s\n\n",
					lipgloss.NewStyle().Foreground(lipgloss.Color("#897dbb")).Render(input),
					schema.Description,
				)
			}
			fmt.Println("Run the following command to generate a configu input schema")
			fmt.Println(lipgloss.NewStyle().Render("\n  ⤷ tsctl config init\n"))
		}
	} else {
		fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#FFBF00")).Render("* No actions could be generated from the provided source"))
	}
	return nil
}
