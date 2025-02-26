# sportscrape
A Go package for collecting and transforming sports statistics from various sources into standardized formats.

## Installation
```console
go get github.com/lightning-dabbler/sportscrape
```

## Usage
```go
package main

import (
    "fmt"
    "time"
    
    "github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference/nba"
)

func main() {
    // Get Feb 20, 2025's matchups
    date := "2025-02-20"
    matchups := nba.GetMatchups(date)
    
    // Get basic box score stats with 5 concurrent requests
    stats := nba.GetBasicBoxScoreStats(5, matchups...)
    
    fmt.Printf("Found %d matchups with %d player stat lines\n", len(matchups), len(stats))
}
```

## Data providers

| Source                           | League | Feed                  | Function
|----------------------------------|--------|-----------------------|-----------------------|
| https://basketball-reference.com | NBA    | Matchup               | `GetMatchups(date)`   |
| https://basketball-reference.com | NBA    | Basic box score stats | `GetBasicBoxScoreStats(concurrency, matchups...)` |

## License
MIT
