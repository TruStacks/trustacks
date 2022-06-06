package main

import (
	"github.com/spf13/cobra"
)

func main() {
	rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:   "tsctl",
	Short: "Trustacks is the workflow driven value steam delivery platform",
}
