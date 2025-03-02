# sportscrape
A Go package for collecting and transforming sports statistics from various sources into standardized formats.

## Installation
```console
go get github.com/lightning-dabbler/sportscrape
```

## Usage
- [basketball-reference.com NBA scrape examples](dataprovider/basketballreference/nba/example_test.go)

## Data providers

| Source                           | League | Feed                  | Function
|----------------------------------|--------|-----------------------|-----------------------|
| https://basketball-reference.com | NBA    | Matchup               | `GetMatchups(date)`   |
| https://basketball-reference.com | NBA    | Basic box score stats | `GetBasicBoxScoreStats(concurrency, matchups...)` |
| https://basketball-reference.com | NBA    | Advanced box score stats | `GetAdvBoxScoreStats(concurrency, matchups...)` |

## License
MIT
