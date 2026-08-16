# sportscrape
A Go package for collecting and transforming sports statistics from various sources into standardized formats.

[![Deploy][sportscrape-ci-status]][sportscrape-ci]
[![Go Reference][goref-sportscrape-status]][goref-sportscrape]
[![Releases][release-status]][releases]

## CLI

The `sportscrape` CLI extracts sports data from all supported providers and exports to JSONL or Parquet, locally or to S3-compatible storage.

### Installation
```console
go install github.com/lightning-dabbler/sportscrape/cmd/sportscrape@latest
```

### Commands

| Command | Subcommand | Provider |
|---------|------------|----------|
| `sportscrape baseballsavant` | | baseballsavant.mlb.com |
| `sportscrape propfinder` | `mlb` | api.propfinder.app |
| `sportscrape foxsports` | `mlb`, `nba`, `wnba` | foxsports.com |
| `sportscrape espn` | `ufc` | espn.com/mma |
| `sportscrape nba` | | nba.com |
| `sportscrape wnba` | | wnba.com |

Run `sportscrape <command> --help` for feeds, flags, and defaults per provider.

### Example
Extract `2025-06-05` NBA traditional box score data from https://nba.com and store in `./tmp/nba-traditional-box-score-2025-06-05.jsonl` in JSONL format
```console
sportscrape nba \
  --feed traditional-box-score \
  --date 2025-06-05 \
  --destination ./tmp/nba-traditional-box-score-2025-06-05.jsonl
```

## Go Package

### Installation
```console
go get github.com/lightning-dabbler/sportscrape
```

### Quick start
Retrieve and output `2025-06-05` NBA traditional box score data from https://nba.com
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

func main() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-05"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchupScraper.NetworkHeaders = nba.NetworkHeaders
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper:   matchupScraper,
			KeepAlive: true,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		matchupScraper.Close()
		panic(err)
	}

	boxscorescraper := nba.NewBoxScoreTraditionalScraper(
		nba.WithBoxScoreTraditionalTimeout(2*time.Minute),
		nba.WithBoxScoreTraditionalPeriod(nba.Full),
	)
	boxscorescraper.DocumentRetriever = matchupScraper.DocumentRetriever

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreTraditional]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, record := range records {
		jsonBytes, err := json.MarshalIndent(record, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

```

### Usage
- [basketball-reference.com NBA scrape examples](dataprovider/basketballreferencenba/example_test.go) (Deprecated)
- [baseball-reference.com MLB scrape examples](dataprovider/baseballreferencemlb/example_test.go) (Deprecated)
- [foxsports.com scraping examples](dataprovider/foxsports/example_test.go)
- [baseballsavant.mlb.com scraping examples](dataprovider/baseballsavantmlb/example_test.go)
- [propfinder scraping examples](dataprovider/propfinder/mlb/example_test.go)
- [ESPN MMA scraping examples](dataprovider/espn/mma/example_test.go)
- [nba.com NBA scraping examples](dataprovider/nba/example_test.go)
- [wnba.com WNBA scraping examples](dataprovider/wnba/example_test.go)

## Data providers

See [docs/DATA_PROVIDERS.md](docs/DATA_PROVIDERS.md) for the full list of supported providers, feeds, and their data models.

## Supported Formats
File formats the constructed data models support on export and import.
|Format|Export|Import|Go Package|
|:------|:----:|:-----:|:-----|
|Parquet|✅|✅|[xitongsys/parquet-go](https://pkg.go.dev/github.com/xitongsys/parquet-go)|
|JSON|✅|✅|[encoding/json](https://pkg.go.dev/encoding/json)|

## Development
### Prerequisites
Go 1.24 or higher

### Building the CLI
```console
make build-sportscrape
```

### Testing
This project is using [mockery](https://github.com/vektra/mockery) v3.5.0 to mock interfaces.

To run unit tests:
```console
make unit-tests
```

To run unit and integration tests:
```console
make all-tests
```

Tests are also being ran as CI workflows on Github Actions.

## License
MIT

[sportscrape-ci]: https://github.com/lightning-dabbler/sportscrape/actions/workflows/deploy.yml (Deploy CI)
[sportscrape-ci-status]: https://github.com/lightning-dabbler/sportscrape/actions/workflows/deploy.yml/badge.svg (Deploy CI)
[goref-sportscrape]: https://pkg.go.dev/github.com/lightning-dabbler/sportscrape
[goref-sportscrape-status]: https://pkg.go.dev/badge/github.com/lightning-dabbler/sportscrape.svg
[release-status]: https://img.shields.io/github/v/release/lightning-dabbler/sportscrape?display_name=tag&sort=semver (Latest Release)
[releases]: https://github.com/lightning-dabbler/sportscrape/releases (Releases)
