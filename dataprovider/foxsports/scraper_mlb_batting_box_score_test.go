//go:build integration

package foxsports

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/writer"
)

func TestMLBBattingBoxScoreScraper(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	// Get matchups
	matchupScraper := NewMatchupScraper(
		MatchupScraperLeague(MLB),
		MatchupScraperSegmenter(&GeneralSegmenter{Date: "2024-10-30"}),
	)

	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)

	// Get boxscore data
	boxscoreScraper := NewMLBBattingBoxScoreScraper()
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConcurrency(1),
		runner.EventDataRunnerScraper(
			boxscoreScraper,
		),
	)
	boxScoreStats, err := boxscorerunner.Run(matchups...)
	assert.NoError(t, err)
	n_stats := len(boxScoreStats)
	n_expected := 19
	assert.Equal(t, n_expected, n_stats, "13 statlines")
	GavinLuxTested := false
	fs := afero.NewOsFs()
	tmpDir, err := afero.TempDir(fs, "./", "test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v\n", err)
	}
	defer fs.RemoveAll(tmpDir)
	filePath := filepath.Join(tmpDir, "foo.parquet")

	// Write parquet
	fw, err := local.NewLocalFileWriter(filePath)
	if err != nil {
		t.Fatalf("Error creating file: %v\n", err)
	}
	defer fw.Close()

	pw, err := writer.NewParquetWriter(fw, new(model.MLBBattingBoxScoreStats), 1)
	if err != nil {
		t.Fatalf("Can't create parquet writer %v\n", err)
	}
	pw.CompressionType = parquet.CompressionCodec_SNAPPY
	for _, statline := range boxScoreStats {
		s := statline.(model.MLBBattingBoxScoreStats)
		if s.Player == "Gavin Lux" {
			GavinLuxTested = true
			assert.Equal(t, int64(91226), s.EventID, "EventID")
			assert.Equal(t, int64(24), s.TeamID, "TeamID")
			assert.Equal(t, int64(1730333280000), s.EventTimeParquet, "EventTimeParquet")
			assert.Equal(t, "Los Angeles Dodgers", s.Team, "Team")
			assert.Equal(t, int64(6), s.OpponentID, "OpponentID")
			assert.Equal(t, "New York Yankees", s.Opponent, "Opponent")
			assert.Equal(t, int64(8677), s.PlayerID, "PlayerID")
			assert.Equal(t, "Gavin Lux", s.Player, "Player")
			assert.Equal(t, "2B", s.Position, "Position")
			assert.Equal(t, int32(2), s.AtBat, "AtBat")
			assert.Equal(t, int32(0), s.Runs, "Runs")
			assert.Equal(t, int32(0), s.Hits, "Hits")
			assert.Equal(t, int32(1), s.RunsBattedIn, "RunsBattedIn")
			assert.Equal(t, int32(1), s.Walks, "Walks")
			assert.Equal(t, int32(1), s.Strikeouts, "Strikeouts")
			assert.Equal(t, int32(3), s.LeftOnBase, "LeftOnBase")
			assert.Equal(t, float32(0.176), s.BattingAverage, "BattingAverage")
		}
		if err = pw.Write(statline); err != nil {
			t.Fatalf("Write error %v\n", err)
		}
	}
	if err = pw.WriteStop(); err != nil {
		t.Fatalf("WriteStop error %v\n", err)
	}
	log.Printf("Parquet file '%s' written\n", filePath)
	// Read newly written parquet
	fr, err := local.NewLocalFileReader(filePath)
	if err != nil {
		t.Fatalf("Can't open file\n")
		return
	}
	defer fr.Close()

	pr, err := reader.NewParquetReader(fr, new(model.MLBBattingBoxScoreStats), 1)
	if err != nil {
		t.Fatalf("Can't create parquet reader %v\n", err)
	}
	defer pr.ReadStop()

	num := int(pr.GetNumRows())
	assert.Equal(t, n_expected, num)
	assert.True(t, GavinLuxTested, "Gavin Lux statline tested")
}
