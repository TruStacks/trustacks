package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"dagger.io/dagger"
	"github.com/charmbracelet/lipgloss"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/trustacks/trustacks/pkg/engine"
	"github.com/trustacks/trustacks/pkg/plan"
)

func RunPlan(source, name string) error {
	spec, err := engine.New().CreateActionPlan(source, true)
	if err != nil {
		return err
	}
	planPath := fmt.Sprintf("%s.plan", name)
	if err := os.WriteFile(planPath, []byte(spec), 0600); err != nil {
		return err
	}
	fmt.Printf("plan filed saved at: %s\n", planPath)
	return nil
}

func getLocalSpec(source, planFile string, force bool) (map[string]interface{}, error) {
	var err error
	var planJson []byte
	planData := map[string]interface{}{}
	if force {
		spec, err := engine.New().CreateActionPlan(source, true)
		if err != nil {
			return nil, err
		}
		planJson = []byte(spec)
	} else {
		planJson, err = os.ReadFile(planFile)
	}
	if err != nil {
		return nil, fmt.Errorf("failed opening plan file: %s", err)
	}
	if err := json.Unmarshal(planJson, &planData); err != nil {
		return nil, fmt.Errorf("failed parsing plan file: %s", err)
	}
	return planData, nil
}

func getHostedSpec(planName, server string) (map[string]interface{}, string, error) {
	var client struct {
		GetActionPlan func(string, string) (map[string]interface{}, error)
	}
	closer, err := jsonrpc.NewClient(context.Background(), fmt.Sprintf("%s/rpc", server), "v1", &client, nil)
	if err != nil {
		return nil, "", err
	}
	defer closer()
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, "", err
	}
	path := filepath.Join(homeDir, ".trustacks")
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, "", err
	}
	token, err := os.ReadFile(filepath.Join(path, "auth-token"))
	if err != nil {
		return nil, "", errors.New("error: auth token not found. login to generate one")
	}
	actionPlanData, err := client.GetActionPlan(planName, string(token))
	if err != nil {
		return nil, "", err
	}
	var sourceSubpath string
	if _, ok := actionPlanData["path"].(string); ok {
		sourceSubpath = actionPlanData["path"].(string)
	}
	planData := map[string]interface{}{
		"actions": []map[string]interface{}{},
	}
	planSpec, ok := actionPlanData["plan"].(map[string]interface{})
	if ok {
		for _, action := range planSpec["actions"].([]interface{}) {
			exclude := false
			actionName := action.(map[string]interface{})["name"].(string)
			if exclusions, ok := actionPlanData["exclusions"].([]interface{}); ok {
				for _, exclusion := range exclusions {
					if actionName == exclusion {
						exclude = true
					}
				}
			}
			if exclude {
				continue
			}
			planData["actions"] = append(planData["actions"].([]map[string]interface{}), action.(map[string]interface{}))
		}
	}
	return planData, sourceSubpath, nil
}

func RunCmd(source, planName, planFile, server string, stages []string, force bool) error {
	var err error
	var planData map[string]interface{}
	if planFile != "" {
		planData, err = getLocalSpec(source, planFile, force)
	} else {
		var sourceSubpath string
		planData, sourceSubpath, err = getHostedSpec(planName, server)
		if err != nil {
			return err
		}
		source = filepath.Join(source, sourceSubpath)
	}
	if err != nil {
		return err
	}
	spec, err := json.Marshal(planData)
	if err != nil {
		return fmt.Errorf("failed converting plan file to spec: %s", err)
	}
	client, err := dagger.Connect(context.Background(), dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	if err := plan.Run(source, string(spec), client, stages); err != nil {
		return err
	}
	return nil
}

func StackInitializeCmd(planFile, output string) error {
	plan := map[string]interface{}{}
	contents, err := os.ReadFile(planFile)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(contents, &plan); err != nil {
		return err
	}
	if _, err := os.Stat(output); !os.IsNotExist(err) {
		return fmt.Errorf("error: %s already exists", output)
	}
	if _, ok := plan["inputs"]; ok {
		data := bytes.NewBufferString("")
		for _, field := range plan["inputs"].([]interface{}) {
			if _, err := data.WriteString(fmt.Sprintf("export %s=\n", field.(string))); err != nil {
				return err
			}
		}
		if err := os.WriteFile(output, data.Bytes(), 0644); err != nil {
			return err
		}
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

func ExplainCmd(path, docsURL string) error {
	var actionPlan plan.ActionPlan
	data, err := os.ReadFile(path)
	if err != nil {
		return err
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
				spec := plan.GetInputSpec(input)
				if spec == nil {
					continue
				}
				fmt.Printf(
					"▸ %s - %s\n%s\n",
					lipgloss.NewStyle().Foreground(lipgloss.Color("#897dbb")).Render(input),
					fmt.Sprintf("%s/inputs/%s", docsURL, spec.Link()),
					spec.Description(),
				)
			}
			fmt.Println("Run the following command to generate a keyed input file for this plan")
			fmt.Println(lipgloss.NewStyle().Render(fmt.Sprintf("\n  ⤷ tsctl stack init --from-plan %s\n", path)))
		}
	} else {
		fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#FFBF00")).Render("* No actions could be generated from the provided source"))
	}
	return nil
}
