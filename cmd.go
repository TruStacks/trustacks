package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

// factory cli command flags.
var (
	factoryName   string
	factorySource string
	factoryTag    string
)

// rootCmd is the cobra start command.
var rootCmd = &cobra.Command{
	Use:   "tsctl",
	Short: "Trustacks is the workflow driven value steam delivery platform",
}

// factoryCmd contains subcommands for managing factories.
var factoryCmd = &cobra.Command{
	Use:   "factory",
	Short: "manage software factories",
}

// factoryInstallCmd
var factoryInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "install a software factory",
	Run: func(cmd *cobra.Command, args []string) {
		if err := installFactory(factoryName, factorySource, factoryTag, git.PlainClone); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(factoryCmd)
	factoryCmd.AddCommand(factoryInstallCmd)

	factoryInstallCmd.Flags().StringVar(&factoryName, "name", "", "name of the factory")
	factoryInstallCmd.MarkFlagRequired("name")

	factoryInstallCmd.Flags().StringVar(&factorySource, "source", "", "software factory git repository")
	factoryInstallCmd.MarkFlagRequired("source")

	factoryInstallCmd.Flags().StringVar(&factoryTag, "version", "", "software factory version")
	factoryInstallCmd.MarkFlagRequired("ref")
}
