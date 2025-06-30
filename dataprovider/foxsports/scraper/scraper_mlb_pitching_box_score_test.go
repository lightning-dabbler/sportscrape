//go:build integration

package scraper

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/writer"
)

func TestMLBPitchingBoxScoreScraper(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	// Get matchups
	matchupScraper := NewMatchupScraper(
		MatchupScraperLeague(foxsports.MLB),
		MatchupScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2024-10-30"}),
	)

	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.RunMatchupsScraper()
	if err != nil {
		t.Error(err)
	}

	// Get boxscore data
	boxscoreScraper := MLBPitchingBoxScoreScraper{}
	boxscoreScraper.League = foxsports.MLB
	runner := sportscrape.NewEventDataRunner(
		sportscrape.EventDataRunnerConcurrency(1),
		sportscrape.EventDataRunnerScraper(
			&boxscoreScraper,
		),
	)
	boxScoreStats, err := runner.RunEventsDataScraper(matchups...)
	if err != nil {
		t.Error(err)
	}
	n_stats := len(boxScoreStats)
	n_expected := 13
	assert.Equal(t, n_expected, n_stats, "13 statlines")
	LukeWeaverTested := false
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

	pw, err := writer.NewParquetWriter(fw, new(model.MLBPitchingBoxScoreStats), 1)
	if err != nil {
		t.Fatalf("Can't create parquet writer %v\n", err)
	}
	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	for _, statline := range boxScoreStats {
		s := statline.(model.MLBPitchingBoxScoreStats)
		if s.Player == "Luke Weaver" {
			LukeWeaverTested = true
			assert.Equal(t, int64(91226), s.EventID, "EventID")
			assert.Equal(t, int64(6), s.TeamID, "TeamID")
			assert.Equal(t, "New York Yankees", s.Team, "Team")
			assert.Equal(t, int64(24), s.OpponentID, "OpponentID")
			assert.Equal(t, "Los Angeles Dodgers", s.Opponent, "Opponent")
			assert.Equal(t, int64(8034), s.PlayerID, "PlayerID")
			assert.Equal(t, "Luke Weaver", s.Player, "Player")
			assert.Equal(t, "BS (3)", *s.Record, "Record")
			assert.Equal(t, int32(4), s.PitchingOrder, "PitchingOrder")
			assert.Equal(t, float32(1.1), s.InningsPitched, "InningsPitched")
			assert.Equal(t, int32(1), s.HitsAllowed, "HitsAllowed")
			assert.Equal(t, int32(0), s.RunsAllowed, "RunsAllowed")
			assert.Equal(t, int32(0), s.EarnedRunsAllowed, "EarnedRunsAllowed")
			assert.Equal(t, int32(1), s.Walks, "Walks")
			assert.Equal(t, int32(1), s.Strikeouts, "Strikeouts")
			assert.Equal(t, int32(0), s.HomeRunsAllowed, "HomeRunsAllowed")
			assert.Equal(t, float32(1.76), s.EarnedRunAverage, "EarnedRunAverage")
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

	pr, err := reader.NewParquetReader(fr, new(model.MLBPitchingBoxScoreStats), 1)
	if err != nil {
		t.Fatalf("Can't create parquet reader %v\n", err)
	}
	defer pr.ReadStop()

	num := int(pr.GetNumRows())
	assert.Equal(t, n_expected, num)
	assert.True(t, LukeWeaverTested, "Luke Weaver statline tested")
}
