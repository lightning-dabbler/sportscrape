package cli

import (
	"fmt"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/feed"
	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/shared"

	"github.com/spf13/cobra"
)

func createPropFinderMLBCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mlb",
		Short: "Extract MLB data",
		Long:  "Extract MLB data",
		RunE: func(cmd *cobra.Command, args []string) error {
			return shared.Run(cmd, "propfinder", "mlb")
		},
	}
	cmd.Flags().IntP("concurrency", "c", 1, fmt.Sprintf("Max number of concurrent goroutines. Dependent on data feed (%s)", feed.PropFinderMLBOptions))
	cmd.Flags().String("feed", "", fmt.Sprintf("The data feed to extract. Options: %s", feed.PropFinderMLBOptions))
	shared.EmbedDateFlag(cmd)
	shared.EmbedDestinationFlag(cmd)
	shared.EmbedFileFormatFlag(cmd)
	shared.EmbedParquetFlags(cmd)
	shared.EmbedS3Flags(cmd)
	return cmd
}

func CreatePropFinderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "propfinder",
		Aliases: []string{"pf"},
		Short:   "Extract prop finder data",
		Long:    "Extract prop finder data",
	}
	// Store subcommands (mlb)
	cmd.AddCommand(createPropFinderMLBCmd())
	return cmd
}
