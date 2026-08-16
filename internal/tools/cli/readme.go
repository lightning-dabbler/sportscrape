package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/lightning-dabbler/sportscrape/internal/catalogdocs"
	"github.com/spf13/cobra"
)

const (
	dataProvidersTableStart = "<!-- DATA_PROVIDERS_TABLE_START -->"
	dataProvidersTableEnd   = "<!-- DATA_PROVIDERS_TABLE_END -->"
)

var dataProvidersTableRe = regexp.MustCompile(
	regexp.QuoteMeta(dataProvidersTableStart) + `(?s).*?` + regexp.QuoteMeta(dataProvidersTableEnd),
)

// createReadmeCmd creates the readme-table subcommand
// Returns the readme-table Command object
func createReadmeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "readme-table",
		Short: "Generate or check the Data providers table",
		Long:  "Regenerates the Data providers table (see --readme) from internal/catalogdocs, or checks (--check) that it's already up to date without writing.",
		RunE: func(cmd *cobra.Command, args []string) error {
			readmePath, err := cmd.Flags().GetString("readme")
			if err != nil {
				return err
			}
			check, err := cmd.Flags().GetBool("check")
			if err != nil {
				return err
			}
			return runReadmeTable(readmePath, check)
		},
		SilenceUsage: true,
	}
	cmd.Flags().String("readme", "README.md", "Path to the markdown file containing the DATA_PROVIDERS_TABLE_START/END markers")
	cmd.Flags().Bool("check", false, "Check the table is up to date instead of writing changes; exits non-zero on drift")
	return cmd
}

func runReadmeTable(readmePath string, check bool) error {
	original, err := os.ReadFile(readmePath)
	if err != nil {
		return err
	}

	// ModelPath (in internal/catalogdocs/registry.go) is always repo-root-relative,
	// but GitHub-flavored markdown relative links resolve against the file
	// they're rendered into. Compute how far readmePath's directory sits from
	// the repo root so links still resolve when the target isn't at the root
	// (e.g. docs/DATA_PROVIDERS.md needs a "../" prefix; README.md needs none).
	linkPrefix, err := filepath.Rel(filepath.Dir(readmePath), ".")
	if err != nil {
		return err
	}
	// filepath.Rel(dir, ".") returns the unclean "../." rather than "..", so
	// Clean before comparing/using it, or the "." leaks into the prefix.
	linkPrefix = filepath.Clean(linkPrefix)
	if linkPrefix == "." {
		linkPrefix = ""
	} else {
		linkPrefix += "/"
	}

	table := catalogdocs.RenderTable(catalogdocs.FeedDocs, linkPrefix)
	replacement := dataProvidersTableStart + "\n" + table + "\n" + dataProvidersTableEnd

	if !dataProvidersTableRe.Match(original) {
		return fmt.Errorf("could not find %s / %s markers in %s", dataProvidersTableStart, dataProvidersTableEnd, readmePath)
	}
	updated := dataProvidersTableRe.ReplaceAll(original, []byte(replacement))

	if string(updated) == string(original) {
		fmt.Println("README Data providers table is up to date.")
		return nil
	}

	if check {
		return fmt.Errorf("README Data providers table is out of date; run `make generate-readme`")
	}

	if err := os.WriteFile(readmePath, updated, 0644); err != nil {
		return err
	}
	fmt.Printf("Updated %s\n", readmePath)
	return nil
}
