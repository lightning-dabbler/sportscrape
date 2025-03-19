# sportscrape
A Go package for collecting and transforming sports statistics from various sources into standardized formats.

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

## License
MIT
