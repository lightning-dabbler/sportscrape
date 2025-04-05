package foxsports

import (
	"fmt"
	"net/url"
)

type League int

const (
	NBA League = iota
	MLB
	NCAAB
)
const (
	API               = "https://api.foxsports.com"
	BifrostEndpointV1 = API + "/bifrost/v1"
	APIKey            = "jE7yBJVRNAwdDesMgTzTXUUSx1It41Fq"
)

// LeaguePath ouputs the fox sports provided league subsection representation
func (l League) LeaguePath() string {
	switch l {
	case NBA:
		return "nba"
	case MLB:
		return "mlb"
	case NCAAB:
		return "cbk"
	default:
		return "?"
	}
}

// V1MatchupURL generates the full path for matchup based on league and date string
//
// Parameter:
//   - date: date string in the format YYYYMMDD
//
// Returns the parsed URL for the path for a matchup
func (l League) V1MatchupURL(date string) (*url.URL, error) {
	scoreboardPath := "scoreboard/segment/" + date
	switch l {
	case NBA:
		return url.Parse(fmt.Sprintf("%s/%s/%s", BifrostEndpointV1, l.LeaguePath(), scoreboardPath))
	case MLB:
		return url.Parse(fmt.Sprintf("%s/%s/%s", BifrostEndpointV1, l.LeaguePath(), scoreboardPath))
	case NCAAB:
		return url.Parse(fmt.Sprintf("%s/%s/%s", BifrostEndpointV1, l.LeaguePath(), scoreboardPath))
	default:
		return nil, fmt.Errorf("Unknown League identified: %#v", l)
	}
}
