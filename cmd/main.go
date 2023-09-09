package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/trustacks/trustacks/internal"
	_ "trustacks.io/trustacks/actions"
	_ "trustacks.io/trustacks/engine/rules"
)

var (
	version                    string
	planCmdName                string
	planCmdSource              string
	explainCmdPlan             string
	explainCmdDocsURL          string
	runCmdPlan                 string
	runCmdInputsFile           string
	runCmdSource               string
	runCmdStages               []string
	runCmdForce                bool
	runCmdContext              string
	stackInitializeCmdFromPlan string
	stackInitializeCmdOutput   string
	loginCmdUsername           string
	loginCmdPassword           string
	rootCmdServer              string
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
		if err := internal.RunPlan(planCmdSource, planCmdName); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explain an action plan",
	Run: func(cmd *cobra.Command, args []string) {
		if err := internal.ExplainCmd(explainCmdPlan, explainCmdDocsURL); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an action plan",
	Run: func(cmd *cobra.Command, args []string) {
		planFile := ""
		if len(args) > 0 {
			planFile = args[0]
		}
		if err := internal.RunCmd(runCmdSource, runCmdPlan, planFile, runCmdInputsFile, runCmdContext, rootCmdServer, runCmdStages, runCmdForce); err != nil {
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
		if err := internal.StackInitializeCmd(stackInitializeCmdFromPlan, stackInitializeCmdOutput); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to trustacks",
	Run: func(cmd *cobra.Command, _ []string) {
		if err := internal.LoginCmd(rootCmdServer, loginCmdUsername, loginCmdPassword); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func main() {
	rootCmd.Flags().StringVar(&rootCmdServer, "server", "", "rpc server host")

	rootCmd.AddCommand(versionCmd)

	planCmd.Flags().StringVar(&planCmdName, "name", "", "name of the application")
	if err := planCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
		return
	}
	planCmd.Flags().StringVar(&planCmdSource, "source", "./", "path to the application source")
	rootCmd.AddCommand(planCmd)

	explainCmd.Flags().StringVar(&explainCmdPlan, "plan", "", "path to the plan file")
	explainCmd.Flags().StringVar(&explainCmdDocsURL, "docs", "https://docs.trustacks.io", "documentation url")
	rootCmd.AddCommand(explainCmd)

	runCmd.Flags().StringVar(&runCmdPlan, "plan", "", "name of a hosted action plan")
	runCmd.Flags().StringVar(&runCmdSource, "source", "./", "application source path")
	runCmd.Flags().StringVar(&runCmdInputsFile, "inputs", "inputs.yaml", "stack inputs file")
	runCmd.Flags().StringVar(&runCmdContext, "context", "default", "input context from a hosted action plan")
	runCmd.Flags().BoolVar(&runCmdForce, "force", false, "plan and execute in one command")
	runCmd.Flags().StringSliceVar(&runCmdStages, "stages", []string{"feedback", "package", "stage", "qa"}, "activity phases to run during the action plan")
	runCmd.Flags().StringVar(&rootCmdServer, "server", "", "rpc server host")

	rootCmd.AddCommand(runCmd)

	stackInitializeCmd.Flags().StringVar(&stackInitializeCmdFromPlan, "from-plan", "", "path to the plan file")
	if err := stackInitializeCmd.MarkFlagRequired("from-plan"); err != nil {
		fmt.Println(err)
		return
	}
	stackInitializeCmd.Flags().StringVar(&stackInitializeCmdOutput, "output", "inputs.yaml", "inputs output path")
	stackCmd.AddCommand(stackInitializeCmd)
	rootCmd.AddCommand(stackCmd)

	loginCmd.Flags().StringVar(&loginCmdUsername, "username", "", "login username")
	if err := loginCmd.MarkFlagRequired("username"); err != nil {
		fmt.Println(err)
		return
	}
	loginCmd.Flags().StringVar(&loginCmdPassword, "password", "", "login password")
	if err := loginCmd.MarkFlagRequired("password"); err != nil {
		fmt.Println(err)
		return
	}
	loginCmd.Flags().StringVar(&rootCmdServer, "server", "", "rpc server host")
	rootCmd.AddCommand(loginCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("error executing the command: %s\n", err)
		os.Exit(1)
	}
}
