package cli

import (
	"fmt"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/feed"
	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/shared"

	"github.com/spf13/cobra"
)

func CreateBaseballSavantCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "baseballsavant",
		Aliases: []string{"bs", "bsmlb"},
		Short:   "Extract baseball savant MLB data",
		Long:    "Extract baseball savant MLB data",
		RunE: func(cmd *cobra.Command, args []string) error {
			return shared.Run(cmd, "baseballsavant", "")
		},
	}
	cmd.Flags().IntP("concurrency", "c", 1, fmt.Sprintf("Max number of concurrent goroutines. Dependent on data feed (%s)", feed.BaseballSavantConcurrencyOptions))
	cmd.Flags().String("feed", "", fmt.Sprintf("The data feed to extract. Options: %s", feed.BaseballSavantOptions))
	shared.EmbedDateFlag(cmd)
	shared.EmbedDestinationFlag(cmd)
	shared.EmbedFileFormatFlag(cmd)
	shared.EmbedParquetFlags(cmd)
	shared.EmbedS3Flags(cmd)
	return cmd
}
