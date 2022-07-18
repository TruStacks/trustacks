package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/trustacks/trustacks/pkg/workflow"
)

// application cli command flags.
var (
	applicationConfig    string
	applicationToolchain string
)

// applicationCmd contains subcommands for managing factories.
var applicationCmd = &cobra.Command{
	Use:   "applications",
	Short: "manage applications",
}

// applicationInstallCmd
var applicationInstallCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new application",
	Run: func(cmd *cobra.Command, args []string) {
		if err := workflow.InstallWorkflow(applicationToolchain, applicationConfig, git.PlainClone); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	applicationCmd.AddCommand(applicationInstallCmd)

	applicationInstallCmd.Flags().StringVar(&applicationConfig, "toolchain", "", "toolchain instance to use")
	applicationInstallCmd.MarkFlagRequired("toolchain")

	applicationInstallCmd.Flags().StringVar(&applicationConfig, "config", "", "configuration file")
	applicationInstallCmd.MarkFlagRequired("config")

	rootCmd.AddCommand(applicationCmd)
}
