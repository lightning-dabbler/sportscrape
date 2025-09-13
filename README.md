# sportscrape
A Go package for collecting and transforming sports statistics from various sources into standardized formats.

[![Deploy][sportscrape-ci-status]][sportscrape-ci]
[![Go Report Card][go-report-status]][go-report]
[![Go Reference][goref-sportscrape-status]][goref-sportscrape]
[![Releases][release-status]][releases]

## Installation
```console
go get github.com/lightning-dabbler/sportscrape
```

## Quick start
Retrieve and output `2025-02-20` NBA matchups from https://foxsports.com
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/runner"
)

func main() {
	date := "2025-02-20"
	// define matchup scraper
	matchupScraper := foxsports.NewMatchupScraper(
		foxsports.MatchupScraperLeague(foxsports.NBA),
		foxsports.MatchupScraperSegmenter(&foxsports.GeneralSegmenter{Date: date}),
	)
	// define matchup runner
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupScraper),
	)
	// Retrieve matchups
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}
	// Output each matchup as pretty json
	for _, matchup := range matchups {
		jsonBytes, err := json.MarshalIndent(matchup, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}
```

## Usage
- [basketball-reference.com NBA scrape examples](dataprovider/basketballreferencenba/example_test.go)
- [baseball-reference.com MLB scrape examples](dataprovider/baseballreferencemlb/example_test.go)
- [foxsports.com scraping examples](dataprovider/foxsports/example_test.go)
- [baseballsavant.mlb.com scraping examples](dataprovider/baseballsavantmlb/example_test.go)
- [ESPN MMA scraping examples](dataprovider/espn/mma/example_test.go)

## Data providers

| Source                           | League | Feed                      |      Periods Available       |                                  Data Model                                  |	Deprecated	| Point-in-time|
|----------------------------------|--------|---------------------------|:----------------------------:|:----------------------------------------------------------------------------:|:---------------------:|:------------:|
| https://basketball-reference.com | NBA    | Matchup                   |             Full             |        [model](dataprovider/basketballreferencenba/model/matchup.go)         |		|✅|
| https://basketball-reference.com | NBA    | Basic box score stats     | H1, H2, Q1, Q2, Q3, Q4, Full | [model](dataprovider/basketballreferencenba/model/basic_box_score_stats.go)  |		|✅|
| https://basketball-reference.com | NBA    | Advanced box score stats  |             Full             |  [model](dataprovider/basketballreferencenba/model/adv_box_score_stats.go)   ||✅|
| https://baseball-reference.com   | MLB    | Matchup                   |             Full             |         [model](dataprovider/baseballreferencemlb/model/matchup.go)          ||✅|
| https://baseball-reference.com   | MLB    | Batting box score stats   |             Full             | [model](dataprovider/baseballreferencemlb/model/batting_box_score_stats.go)  ||✅|
| https://baseball-reference.com   | MLB    | Pitching box score stats  |             Full             | [model](dataprovider/baseballreferencemlb/model/pitching_box_score_stats.go) ||✅|
| https://www.foxsports.com		      | NBA	   | Matchup				               |        Live, Full			         |               [model](dataprovider/foxsports/model/matchup.go)               ||✅|
| https://www.foxsports.com		      | NBA	   | Box score stats		         |        Live, Full			         |         [model](dataprovider/foxsports/model/nba_box_score_stats.go)         ||✅|
| https://www.foxsports.com		      | WNBA	  | Matchup				               |        Live, Full			         |               [model](dataprovider/foxsports/model/matchup.go)               ||✅|
| https://www.foxsports.com		      | WNBA	  | Box score stats		         |        Live, Full			         |         [model](dataprovider/foxsports/model/nba_box_score_stats.go)         ||✅|
| https://www.foxsports.com		      | MLB	   | Matchup				               |        Live, Full			         |               [model](dataprovider/foxsports/model/matchup.go)               ||✅|
| https://www.foxsports.com		      | MLB	   | Batting Box score stats   |        Live, Full			         |     [model](dataprovider/foxsports/model/mlb_batting_box_score_stats.go)     ||✅|
| https://www.foxsports.com		      | MLB	   | Pitching Box score stats  |        Live, Full			         |    [model](dataprovider/foxsports/model/mlb_pitching_box_score_stats.goo)    ||✅|
| https://www.foxsports.com		      | MLB	   | Probable starting pitcher |           Full			            |    [model](dataprovider/foxsports/model/mlb_probable_starting_pitcher.go)    ||✅|
| https://www.foxsports.com		      | MLB	   | Betting Odds Money line   |           Full			            |         [model](dataprovider/foxsports/model/mlb_odds_money_line.go)         ||✅|
| https://www.foxsports.com		      | MLB	   | Betting Odds Total        |           Full			            |           [model](dataprovider/foxsports/model/mlb_odds_total.go)            ||✅|
| https://www.foxsports.com		      | NCAAB	 | Matchup				               |        Live, Full			         |               [model](dataprovider/foxsports/model/matchup.go)               ||✅|
| https://www.foxsports.com		      | NFL	   | Matchup				               |        Live, Full			         |               [model](dataprovider/foxsports/model/matchup.go)               ||✅|
| https://baseballsavant.mlb.com		 | MLB	   | Matchup				               |        Live, Full			         |           [model](dataprovider/baseballsavantmlb/model/matchup.go)           ||✅|
| https://baseballsavant.mlb.com		 | MLB	   | Batting box score stats   |        Live, Full			         |      [model](dataprovider/baseballsavantmlb/model/batting_box_score.go)      ||✅|
| https://baseballsavant.mlb.com		 | MLB	   | Pitching box score stats  |        Live, Full			         |     [model](dataprovider/baseballsavantmlb/model/pitching_box_score.go)      ||✅|
| https://baseballsavant.mlb.com		 | MLB	   | Fielding box score stats  |        Live, Full			         |     [model](dataprovider/baseballsavantmlb/model/fielding_box_score.go)      ||✅|
| https://baseballsavant.mlb.com		 | MLB	   | Play by play              |        Live, Full			         |        [model](dataprovider/baseballsavantmlb/model/play_by_play.go)         ||✅|
| https://www.espn.com/mma/     		 | MMMA	  | Stats, Odds, Results      |           Full			            |               [model](dataprovider/espn/mma/model/matchup.go)                ||✅|

## Supported Formats
File formats the constructed data models support on export and import.
|Format|Export|Import|Go Package|
|:------|:----:|:-----:|:-----|
|Parquet|✅|✅|[xitongsys/parquet-go](https://pkg.go.dev/github.com/xitongsys/parquet-go)|
|JSON|✅|✅|[encoding/json](https://pkg.go.dev/encoding/json)|

## Development
### Prerequisites
Go 1.24 or higher

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
[go-report]: https://goreportcard.com/report/github.com/lightning-dabbler/sportscrape (Go report)
[go-report-status]: https://goreportcard.com/badge/github.com/lightning-dabbler/sportscrape (Go report Badge)
