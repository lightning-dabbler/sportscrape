package main

import (
	"fmt"
	"log"

	"github.com/lightning-dabbler/sportscrape/tools/common"
	"github.com/lightning-dabbler/sportscrape/version"
	"github.com/spf13/cobra"
)

// createVersionCmd creates the version subcommand
// Returns the version Command object
func createVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Ouput project version.",
		Long:    "Display the project's version at compilation.",
		RunE: func(cmd *cobra.Command, args []string) error {
			v, err := common.LoadSemVer(version.Version)
			if err != nil {
				log.Printf("Issue parsing project semver %s", version.Version)
				return err
			}
			fmt.Printf("%s\n", v.Original())
			return nil
		},
	}
}
