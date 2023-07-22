package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"dagger.io/dagger"
	"github.com/spf13/cobra"
	_ "github.com/trustacks/pkg/actions"
	"github.com/trustacks/pkg/engine"
	"github.com/trustacks/pkg/plan"
	"go.mozilla.org/sops/v3/decrypt"
)

var version string

var rootCmd = &cobra.Command{
	Use:   "tsctl",
	Short: "Trustacks software delivery engine",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the cli version",
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Println(version)
	},
}

var (
	generateCmdName   string
	generateCmdSource string
	runCmdPlanFile    string
	runCmdStackFile   string
)

var generateCmd = &cobra.Command{
	Use:   "plan",
	Short: "Generate an action plan",
	Run: func(cmd *cobra.Command, _ []string) {
		spec, err := engine.New().CreateActionPlan(generateCmdSource)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		planPath := fmt.Sprintf("%s.plan", generateCmdName)
		if err := os.WriteFile(planPath, []byte(spec), 0600); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("plan filed saved at: %s\n", planPath)
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an action plan",
	Run: func(cmd *cobra.Command, _ []string) {
		planData := map[string]interface{}{}
		planJson, err := os.ReadFile(runCmdPlanFile)
		if err != nil {
			fmt.Println("Failed opening plan file:", err)
			os.Exit(1)
		}
		if err := json.Unmarshal(planJson, &planData); err != nil {
			fmt.Println("failed parsing plan file:", err)
			os.Exit(1)
		}
		stack := map[string]interface{}{}
		stackJsonEncrypted, err := os.ReadFile(runCmdStackFile)
		if err != nil {
			fmt.Println("Failed opening stack file:", err)
			os.Exit(1)
		}
		stackJson, err := decrypt.Data(stackJsonEncrypted, "json")
		if err != nil {
			fmt.Println("failed decryptiong stack file:", err)
			os.Exit(1)
		}
		if err := json.Unmarshal(stackJson, &stack); err != nil {
			fmt.Println("failed parsing stack file:", err)
			os.Exit(1)
		}
		for k, v := range stack {
			planData["inputs"].(map[string]interface{})[k] = v
		}
		spec, err := json.Marshal(planData)
		if err != nil {
			fmt.Println("failed converting plan file to spec:", err)
			os.Exit(1)
		}
		client, err := dagger.Connect(context.Background(), dagger.WithLogOutput(os.Stdout))
		if err != nil {
			log.Fatal(err)
		}
		if err := plan.Run("./", string(spec), client, []plan.State{plan.FeebackState, plan.StageState, plan.QAState}); err != nil {
			log.Fatal(err)
		}
	},
}

func main() {
	generateCmd.Flags().StringVar(&generateCmdName, "name", "", "the name of the application")
	if err := generateCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
		return
	}
	generateCmd.Flags().StringVar(&generateCmdSource, "src", "./", "the path to the application source")
	runCmd.Flags().StringVar(&runCmdPlanFile, "plan", "", "the path to the plan file")
	if err := runCmd.MarkFlagRequired("plan"); err != nil {
		fmt.Println(err)
		return
	}
	runCmd.Flags().StringVar(&runCmdStackFile, "stack", "stack.json", "the path to the stack file")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(runCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("error executing the command: %s\n", err)
		os.Exit(1)
	}
}
