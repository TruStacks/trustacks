package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/trustacks/internal"
	_ "github.com/trustacks/pkg/actions"
)

var (
	version                    string
	planCmdName                string
	planCmdSource              string
	runCmdPlanFile             string
	runCmdInputsFile           string
	runCmdSource               string
	stackInitializeCmdPlanFile string
	stackInitializeCmdOutput   string
)

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

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Generate an action plan",
	Run: func(cmd *cobra.Command, _ []string) {
		if err := internal.RunPlan(planCmdName, planCmdName); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an action plan",
	Run: func(cmd *cobra.Command, _ []string) {
		if err := internal.RunCmd(runCmdSource, runCmdPlanFile, runCmdInputsFile); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var stackCmd = &cobra.Command{
	Use:   "stack",
	Short: "manage input stacks",
}

var stackInitializeCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize an input file from a plan",
	Run: func(cmd *cobra.Command, _ []string) {
		if err := internal.StackInitializeCmd(stackInitializeCmdPlanFile, stackInitializeCmdOutput); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func main() {
	rootCmd.AddCommand(versionCmd)

	planCmd.Flags().StringVar(&planCmdName, "name", "", "the name of the application")
	if err := planCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
		return
	}
	planCmd.Flags().StringVar(&planCmdSource, "source", "./", "the path to the application source")
	rootCmd.AddCommand(planCmd)

	runCmd.Flags().StringVar(&runCmdPlanFile, "plan", "", "the path to the plan file")
	if err := runCmd.MarkFlagRequired("plan"); err != nil {
		fmt.Println(err)
		return
	}
	runCmd.Flags().StringVar(&runCmdInputsFile, "inputs", "inputs.yaml", "the path to the stack inputs file")
	runCmd.Flags().StringVar(&runCmdSource, "source", "./", "the path to the application source")
	rootCmd.AddCommand(runCmd)

	stackInitializeCmd.Flags().StringVar(&stackInitializeCmdPlanFile, "plan", "", "the path to the plan file")
	if err := stackInitializeCmd.MarkFlagRequired("plan"); err != nil {
		fmt.Println(err)
		return
	}
	stackInitializeCmd.Flags().StringVar(&stackInitializeCmdOutput, "output", "inputs.yaml", "the inputs output path")
	stackCmd.AddCommand(stackInitializeCmd)
	rootCmd.AddCommand(stackCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("error executing the command: %s\n", err)
		os.Exit(1)
	}
}
