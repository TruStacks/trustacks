package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/trustacks/trustacks/pkg/engine"
)

func ConfigInitCmd(planFile string) error {
	var actionPlan engine.ActionPlan
	contents, err := os.ReadFile(planFile)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(contents, &actionPlan); err != nil {
		return err
	}
	output := fmt.Sprintf("./%s.cfgu.json", strings.Replace(planFile, ".plan", "", 1))
	schema := map[string]interface{}{}
	for _, name := range actionPlan.Actions {
		action := engine.GetAction(name)
		for _, field := range action.Inputs {
			input := engine.GetInput(string(field))
			schema[string(field)] = input.Schema()
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
	//nolint:gosec
	return os.WriteFile(output, []byte(prettySchema.String()+"\n"), 0644) //nolint:gomnd
}
