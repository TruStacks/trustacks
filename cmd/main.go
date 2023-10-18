package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/trustacks/trustacks/internal"
	_ "github.com/trustacks/trustacks/pkg/actions"
	_ "github.com/trustacks/trustacks/pkg/engine/rules"
)

var (
	version            string
	planCmdName        string
	planCmdSource      string
	explainCmdDocsURL  string
	runCmdPlan         string
	runCmdSource       string
	runCmdStages       []string
	runCmdForce        bool
	configInitFromPlan string
	loginCmdUsername   string
	loginCmdPassword   string
	rootCmdServer      string
)

var rootCmd = &cobra.Command{
	Use:   "tsctl",
	Short: "TruStacks software delivery engine",
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
		if len(args) == 0 {
			fmt.Println("path to plan is required")
			os.Exit(1)
		}
		path := args[0]
		if err := internal.ExplainCmd(path, explainCmdDocsURL); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an action plan",
	Run: func(cmd *cobra.Command, args []string) {
		planFile := "trustacks.plan"
		if len(args) > 0 {
			planFile = args[0]
		}
		if err := internal.RunCmd(runCmdSource, runCmdPlan, planFile, rootCmdServer, runCmdStages, runCmdForce); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "manage action plan inputs",
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "create a Configu schema from an action plan",
	Run: func(cmd *cobra.Command, _ []string) {
		if err := internal.ConfigInitCmd(configInitFromPlan); err != nil {
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

	planCmd.Flags().StringVar(&planCmdName, "name", "trustacks", "name of the plan")
	planCmd.Flags().StringVar(&planCmdSource, "source", "./", "path to the application source")
	rootCmd.AddCommand(planCmd)

	explainCmd.Flags().StringVar(&explainCmdDocsURL, "docs", "https://docs.trustacks.io", "documentation url")
	rootCmd.AddCommand(explainCmd)

	runCmd.Flags().StringVar(&runCmdPlan, "plan", "", "name of a hosted action plan")
	runCmd.Flags().StringVar(&runCmdSource, "source", "./", "application source path")
	runCmd.Flags().BoolVar(&runCmdForce, "force", false, "plan and execute in one command")
	runCmd.Flags().StringSliceVar(&runCmdStages, "stages", []string{"feedback", "package", "prerelease"}, "activity phases to run during the action plan")
	// runCmd.Flags().StringVar(&rootCmdServer, "server", "", "rpc server host")
	rootCmd.AddCommand(runCmd)

	rootCmd.AddCommand(configCmd)

	configInitCmd.Flags().StringVar(&configInitFromPlan, "from-plan", "", "path to the plan file")
	if err := configInitCmd.MarkFlagRequired("from-plan"); err != nil {
		fmt.Println(err)
		return
	}
	configCmd.AddCommand(configInitCmd)

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
