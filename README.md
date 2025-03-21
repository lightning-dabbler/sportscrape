# sportscrape
A Go package for collecting and transforming sports statistics from various sources into standardized formats.

[![All Tests][sportscrape-ci-status]][sportscrape-ci]
[![Go Report Card][go-report-status]][go-report]
[![Go Reference][goref-sportscrape-status]][goref-sportscrape]
[![Releases][release-status]][releases]

## Installation
```console
go get github.com/lightning-dabbler/sportscrape
```

## Quick start
Retrieve and output `2025-02-20` NBA matchups from https://basketball-reference.com
```go
package main

import (
	"fmt"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference/nba"
	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference/nba/model"
)

func main() {
	date := "2025-02-20"
	// Instantiate MatchupRunner
	runner := nba.NewMatchupRunner(
		nba.WithMatchupTimeout(2 * time.Minute),
	)
	// Retrieve NBA matchups associated with date
	matchups := runner.GetMatchups(date)
	for _, matchup := range matchups {
		fmt.Printf("%#v\n", matchup.(model.NBAMatchup))
	}
}
```

## Usage
- [basketball-reference.com NBA scrape examples](dataprovider/basketballreference/nba/example_test.go)
- [baseball-reference.com MLB scrape examples](dataprovider/baseballreference/mlb/example_test.go)

## Data providers

| Source                           | League | Feed                  | Requires Chromium|
|----------------------------------|--------|-----------------------|:----------------:|
| https://basketball-reference.com | NBA    | Matchup               |[x]|
| https://basketball-reference.com | NBA    | Basic box score stats |[x]|
| https://basketball-reference.com | NBA    | Advanced box score stats|[x]|
| https://baseball-reference.com   | MLB    | Matchup                |[x]|
| https://baseball-reference.com   | MLB    | Batting box score stats|[x]|
| https://baseball-reference.com   | MLB    | Pitching box score stats|[x]|

## Development
### Prerequisites
- Go 1.24 or higher

OR

- Docker
    - In which case refer to the Makefile at root to build and shell into a development docker container

### Testing
This project is using [mockery](https://github.com/vektra/mockery) to mock interfaces.

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

[sportscrape-ci]: https://github.com/lightning-dabbler/sportscrape/actions/workflows/all-tests.yml (Test CI)
[sportscrape-ci-status]: https://github.com/lightning-dabbler/sportscrape/actions/workflows/all-tests.yml/badge.svg (Test CI)
[goref-sportscrape]: https://pkg.go.dev/github.com/lightning-dabbler/sportscrape
[goref-sportscrape-status]: https://pkg.go.dev/badge/github.com/lightning-dabbler/sportscrape.svg
[release-status]: https://img.shields.io/github/v/release/lightning-dabbler/sportscrape?display_name=tag&sort=semver (Latest Release)
[releases]: https://github.com/lightning-dabbler/sportscrape/releases (Releases)
[go-report]: https://goreportcard.com/report/github.com/lightning-dabbler/sportscrape (Go report)
[go-report-status]: https://goreportcard.com/badge/github.com/lightning-dabbler/sportscrape (Go report Badge)

