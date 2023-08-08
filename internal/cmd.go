package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"dagger.io/dagger"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/trustacks/pkg/engine"
	"github.com/trustacks/pkg/plan"
	"go.mozilla.org/sops/v3/decrypt"
	"gopkg.in/yaml.v2"
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

func getLocalSpec(source, planFile, inputsFile string, force bool) (map[string]interface{}, error) {
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
	inputs := map[string]interface{}{}
	inputsYAMLEncrypted, err := os.ReadFile(inputsFile)
	if err != nil {
		return nil, fmt.Errorf("failed opening inputs file: %s", err)
	}
	inputsYAML, err := decrypt.Data(inputsYAMLEncrypted, "yaml")
	if err != nil {
		return nil, fmt.Errorf("failed decryptiong inputs file: %s", err)
	}
	if err := yaml.Unmarshal(inputsYAML, &inputs); err != nil {
		return nil, fmt.Errorf("failed parsing inputs file: %s", err)
	}
	if _, ok := planData["inputs"]; ok {
		for k, v := range inputs {
			planData["inputs"].(map[string]interface{})[k] = v
		}
	}
	return planData, nil
}

func getHostedSpec(planName, contextName, server string) (map[string]interface{}, string, error) {
	var client struct {
		GetActionPlan   func(string, string) (map[string]interface{}, error)
		GetStackContext func(string, string, string) (map[string]interface{}, error)
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
	stackContext, err := client.GetStackContext(contextName, os.Getenv("SOPS_AGE_KEY"), string(token))
	if err != nil {
		return nil, "", err
	}
	for k, v := range stackContext {
		if k != "_overrides" {
			stackContext[k] = v
		}
	}
	if overrides, ok := stackContext["_overrides"]; ok {
		if actionPlanOverrides, ok := overrides.(map[string]interface{})[planName]; ok {
			for k, v := range actionPlanOverrides.(map[string]interface{}) {
				stackContext[k] = v
			}
		}
	}
	delete(stackContext, "_overrides")
	planData["inputs"] = stackContext
	return planData, sourceSubpath, nil
}

func RunCmd(source, planName, planFile, inputsFile, contextName, server string, phases []string, force bool) error {
	var err error
	var planData map[string]interface{}
	if planFile != "" {
		planData, err = getLocalSpec(source, planFile, inputsFile, force)
	} else {
		var sourceSubpath string
		planData, sourceSubpath, err = getHostedSpec(planName, contextName, server)
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
	if err := plan.Run(source, string(spec), client, phases); err != nil {
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
