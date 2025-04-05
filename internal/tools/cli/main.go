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
	// Store subcommands (git, version)
	rootCmd.AddCommand(createGitCmd(), createVersionCmd())
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
