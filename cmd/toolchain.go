package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/trustacks/trustacks/pkg/toolchain"
)

// toolchain cli command flags.
var (
	toolchainConfig string
)

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
		if err := toolchain.Install(toolchainConfig, git.PlainClone); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	toolchainCmd.AddCommand(toolchainInstallCmd)
	toolchainInstallCmd.Flags().StringVar(&toolchainConfig, "config", "", "configuration file")
	toolchainInstallCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(toolchainCmd)
}
