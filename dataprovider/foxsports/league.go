package foxsports

import (
	"fmt"
	"net/url"
)

type League int

const (
	Undefined League = iota
	NBA
	MLB
	NCAAB
	NFL
	WNBA
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
	case WNBA:
		return "wnba"
	default:
		return "undefined"
	}
}

func (l League) Undefined() bool {
	switch l {
	case NBA, MLB, NCAAB, NFL, WNBA:
		return false
	default:
		return true
	}
}

// V1MatchupURL generates the full path for matchup based on league and segmentID string
//
// Parameter:
//   - segmentID: date string in the format YYYYMMDD (NBA, MLB, or NCAAB) or selection id in the format \d{4}-\d{1,2}-\d{1} representing year, week, and season type (NFL)
//
// Returns the parsed URL for the path for a matchup
func (l League) V1MatchupURL(segmentID string) (*url.URL, error) {
	scoreboardPath := "scoreboard/segment/" + segmentID
	switch l {
	case NBA, MLB, NCAAB, NFL, WNBA:
		return url.Parse(fmt.Sprintf("%s/%s/%s", BifrostEndpointV1, l.LeaguePath(), scoreboardPath))
	default:
		return nil, fmt.Errorf("undefined league: %#v", l)
	}
}

// V1EventDataURL generates the full path for event data based on league and eventID
//
// Parameter:
//   - eventID: The integer identifier for the event
//
// Returns the parsed URL for the path for event data
func (l League) V1EventDataURL(eventID int64) (*url.URL, error) {
	eventPath := fmt.Sprintf("event/%d/data", eventID)
	switch l {
	case NBA, MLB, NCAAB, NFL, WNBA:
		return url.Parse(fmt.Sprintf("%s/%s/%s", BifrostEndpointV1, l.LeaguePath(), eventPath))
	default:
		return nil, fmt.Errorf("undefined League identified: %#v", l)
	}
}

// V1MatchupComparisonURL generates the full path for matchup comparison based on league and eventID
//
// Parameter:
//   - eventID: The integer identifier for the event
//
// Returns the parsed URL for the path for event data
func (l League) V1MatchupComparisonURL(eventID int64) (*url.URL, error) {
	eventPath := fmt.Sprintf("event/%d/matchup", eventID)
	switch l {
	case NBA, MLB, NCAAB, NFL, WNBA:
		return url.Parse(fmt.Sprintf("%s/%s/%s", BifrostEndpointV1, l.LeaguePath(), eventPath))
	default:
		return nil, fmt.Errorf("undefined League identified: %#v", l)
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
	case WNBA:
		return "WNBA"
	default:
		return "?"
	}
}

// SetParams updates params with relevant query parameters
func (l League) SetParams(params map[string]string) {
	_, exists := params["apikey"]
	if !exists {
		params["apikey"] = APIKey // Default
	}
	_, exists = params["groupId"]
	if l == NCAAB && !exists {
		params["groupId"] = "2" // All D1 matchups
	}
}
