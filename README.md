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
