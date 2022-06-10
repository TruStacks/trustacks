package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

// toolchain cli command flags.
var (
	toolchainName   string
	toolchainSource string
	toolchainTag    string
)

// rootCmd is the cobra start command.
var rootCmd = &cobra.Command{
	Use:   "tsctl",
	Short: "Trustacks is the workflow driven value steam delivery platform",
}

// toolchainCmd contains subcommands for managing factories.
var toolchainCmd = &cobra.Command{
	Use:   "toolchain",
	Short: "manage toolchains",
}

// toolchainInstallCmd
var toolchainInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "install a toolchain",
	Run: func(cmd *cobra.Command, args []string) {
		if err := installToolchain(toolchainName, toolchainSource, toolchainTag, git.PlainClone); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(toolchainCmd)
	toolchainCmd.AddCommand(toolchainInstallCmd)

	toolchainInstallCmd.Flags().StringVar(&toolchainName, "name", "", "name of the toolchain")
	toolchainInstallCmd.MarkFlagRequired("name")

	toolchainInstallCmd.Flags().StringVar(&toolchainSource, "source", "", "software toolchain git repository")
	toolchainInstallCmd.MarkFlagRequired("source")

	toolchainInstallCmd.Flags().StringVar(&toolchainTag, "version", "", "software toolchain version")
	toolchainInstallCmd.MarkFlagRequired("ref")
}
