package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "tools",
		Short: "Utility suite",
		Long:  "Utility suite of subcommands for automation",
	}
	// Store subcommands (git, version, data-providers-table)
	rootCmd.AddCommand(createGitCmd(), createVersionCmd(), createDataProvidersTableCmd())
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
