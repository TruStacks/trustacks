package main

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/trustacks/trustacks/pkg/toolchain"
)

// toolchain cli command flags.
var (
	toolchainConfig string
	toolchainForce  bool
)

// toolchainCmd contains subcommands for managing factories.
var toolchainCmd = &cobra.Command{
	Use:   "toolchain",
	Short: "manage toolchains",
}

// toolchainInstallCmd install the toolchain.
var toolchainInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "install a toolchain",
	Run: func(cmd *cobra.Command, args []string) {
		if err := toolchain.Install(toolchainConfig, toolchainForce, git.PlainClone); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	toolchainCmd.AddCommand(toolchainInstallCmd)
	toolchainInstallCmd.Flags().StringVar(&toolchainConfig, "config", "", "configuration file")
	if err := toolchainInstallCmd.MarkFlagRequired("config"); err != nil {
		log.Fatal(err)
	}
	toolchainInstallCmd.Flags().BoolVar(&toolchainForce, "force", false, "force update (experimental: use at your own risk)")
	rootCmd.AddCommand(toolchainCmd)
}
