package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/trustacks/trustacks/pkg"
	"github.com/trustacks/trustacks/pkg/secrets"
)

var (
	// sopsURL is the url of the sops binary.
	sopsURL = "https://github.com/mozilla/sops/releases/download/v3.7.3/sops-v3.7.3.linux.amd64"
)

// secrets cli command flags.
var (
	secretName   string
	secretSource string
)

// secretsCmd contains subcommands for managing factories.
var secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "manage secrets",
}

// secretsCreateCmd
var secretsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a secret",
	Run: func(cmd *cobra.Command, args []string) {
		if err := secrets.DownloadSops(sopsURL, pkg.BinDir); err != nil {
			fmt.Printf("error downloading sops: %s\n", err)
			return
		}
		if err := secrets.New(secretName, secretSource); err != nil {
			fmt.Printf("error creating secret: %s\n", err)
			return
		}
	},
}

func init() {
	secretsCmd.AddCommand(secretsCreateCmd)

	secretsCreateCmd.Flags().StringVar(&secretName, "name", "", "name of the secret")
	secretsCreateCmd.MarkFlagRequired("name")

	secretsCreateCmd.Flags().StringVar(&secretSource, "source", "", "secret file path")
	secretsCreateCmd.MarkFlagRequired("source")

	rootCmd.AddCommand(secretsCmd)
}
