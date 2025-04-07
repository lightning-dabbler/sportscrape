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
	NFL
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
	case NFL:
		return "nfl"
	default:
		return "?"
	}
}

// V1MatchupURL generates the full path for matchup based on league and date string
//
// Parameter:
//   - selectionId: date string in the format YYYYMMDD (NBA, MLB, or NCAAB) or selection id in the format \d{4}-\d{1,2}-\d{1} representing year, week, and season type (NFL)
//
// Returns the parsed URL for the path for a matchup
func (l League) V1MatchupURL(selectionId string) (*url.URL, error) {
	scoreboardPath := "scoreboard/segment/" + selectionId
	switch l {
	case NBA:
		return url.Parse(fmt.Sprintf("%s/%s/%s", BifrostEndpointV1, l.LeaguePath(), scoreboardPath))
	case MLB:
		return url.Parse(fmt.Sprintf("%s/%s/%s", BifrostEndpointV1, l.LeaguePath(), scoreboardPath))
	case NCAAB:
		return url.Parse(fmt.Sprintf("%s/%s/%s", BifrostEndpointV1, l.LeaguePath(), scoreboardPath))
	case NFL:
		return url.Parse(fmt.Sprintf("%s/%s/%s", BifrostEndpointV1, l.LeaguePath(), scoreboardPath))
	default:
		return nil, fmt.Errorf("Unknown League identified: %#v", l)
	}
}

// String outputs the string abbreviation of the league
func (l League) String() string {
	switch l {
	case NBA:
		return "NBA"
	case MLB:
		return "MLB"
	case NCAAB:
		return "NCAAB"
	case NFL:
		return "NFL"
	default:
		return "?"
	}
}

// SetParams updates params with relevant query parameters
func (l League) SetParams(params map[string]string) {
	_, exists := params["apikey"]
	if !exists {
		params["apikey"] = APIKey
	}
	_, exists = params["groupId"]
	if l == NCAAB && !exists {
		params["groupId"] = "2" // All D1 matchups
	}
}
