//go:build integration

package nba

import (
	"log"
	"path/filepath"
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference/nba/model"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/writer"
)

func TestGetMatchups(t *testing.T) {

	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	tests := []struct {
		name               string
		date               string
		expectedNumMatches int
	}{
		{
			name:               "2025-02-13 NBA matches",
			date:               "2025-02-13",
			expectedNumMatches: 5,
		},
	}
	fs := afero.NewOsFs()
	tmpDir, err := afero.TempDir(fs, "./", "test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v\n", err)
	}
	defer fs.RemoveAll(tmpDir)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := NewMatchupRunner(
				WithMatchupTimeout(4 * time.Minute),
			)
			matchups := runner.GetMatchups(tt.date)
			assert.Equal(t, tt.expectedNumMatches, len(matchups))
			filePath := filepath.Join(tmpDir, "foo.parquet")
			// file, err := fs.Create("/tmp/test-dir/foo.parquet")

			// Write parquet
			fw, err := local.NewLocalFileWriter(filePath)

			if err != nil {
				t.Fatalf("Error creating file: %v\n", err)
			}
			defer fw.Close()

			pw, err := writer.NewParquetWriter(fw, new(model.NBAMatchup), 2)
			if err != nil {
				t.Fatalf("Can't create parquet writer %v\n", err)
			}
			pw.CompressionType = parquet.CompressionCodec_SNAPPY

			for _, matchup := range matchups {
				if err = pw.Write(matchup); err != nil {
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

			pr, err := reader.NewParquetReader(fr, new(model.NBAMatchup), 2)
			if err != nil {
				t.Fatalf("Can't create parquet reader %v\n", err)
			}
			defer pr.ReadStop()

			num := int(pr.GetNumRows())
			assert.Equal(t, tt.expectedNumMatches, num)
		})
	}
}
