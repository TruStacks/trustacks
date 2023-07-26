package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"dagger.io/dagger"
	"github.com/trustacks/pkg/engine"
	"github.com/trustacks/pkg/plan"
	"go.mozilla.org/sops/v3/decrypt"
	"gopkg.in/yaml.v2"
)

func RunPlan(source, name string) error {
	spec, err := engine.New().CreateActionPlan(source)
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

func RunCmd(source, planFile, inputsFile string) error {
	planData := map[string]interface{}{}
	planJson, err := os.ReadFile(planFile)
	if err != nil {
		return fmt.Errorf("failed opening plan file: %s", err)
	}
	if err := json.Unmarshal(planJson, &planData); err != nil {
		return fmt.Errorf("failed parsing plan file: %s", err)
	}
	inputs := map[string]interface{}{}
	inputsYAMLEncrypted, err := os.ReadFile(inputsFile)
	if err != nil {
		return fmt.Errorf("failed opening inputs file: %s", err)
	}
	inputsYAML, err := decrypt.Data(inputsYAMLEncrypted, "yaml")
	if err != nil {
		return fmt.Errorf("failed decryptiong inputs file: %s", err)
	}
	if err := yaml.Unmarshal(inputsYAML, &inputs); err != nil {
		return fmt.Errorf("failed parsing inputs file: %s", err)
	}
	for k, v := range inputs {
		planData["inputs"].(map[string]interface{})[k] = v
	}
	spec, err := json.Marshal(planData)
	if err != nil {
		return fmt.Errorf("failed converting plan file to spec: %s", err)
	}
	client, err := dagger.Connect(context.Background(), dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	if err := plan.Run(source, string(spec), client, []plan.State{plan.FeebackState, plan.PackageState, plan.StageState, plan.QAState}); err != nil {
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
		data, err := yaml.Marshal(plan["inputs"])
		if err != nil {
			return err
		}
		if err := os.WriteFile(output, data, 0644); err != nil {
			return err
		}
	}
	return nil
}
