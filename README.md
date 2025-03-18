# sportscrape
A Go package for collecting and transforming sports statistics from various sources into standardized formats.

## Installation
```console
go get github.com/lightning-dabbler/sportscrape
```

## Usage
- [basketball-reference.com NBA scrape examples](dataprovider/basketballreference/nba/example_test.go)
- [baseball-reference.com NBA scrape examples](dataprovider/baseballreference/mlb/example_test.go)

## Data providers

| Source                           | League | Feed                  |
|----------------------------------|--------|-----------------------|
| https://basketball-reference.com | NBA    | Matchup               |
| https://basketball-reference.com | NBA    | Basic box score stats |
| https://basketball-reference.com | NBA    | Advanced box score stats|
| https://baseball-reference.com   | MLB    | Matchup                |
| https://baseball-reference.com   | MLB    | Batting box score stats|
| https://baseball-reference.com   | MLB    | Pitching box score stats|

## License
MIT
