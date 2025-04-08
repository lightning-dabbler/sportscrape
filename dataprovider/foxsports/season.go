package foxsports

import "fmt"

type SeasonType int

const (
	REGULARSEASON SeasonType = iota + 1
	POSTSEASON
	PRESEASON
)

// SelectionId outputs the selection id used to fetch NFL matchups based on year, week, and season type
//
// Example: https://api.foxsports.com/bifrost/v1/nfl/scoreboard/segment/2024-4-2?apikey=jE7yBJVRNAwdDesMgTzTXUUSx1It41Fq
func (st SeasonType) SelectionId(year int32, week int32) string {
	return fmt.Sprintf("%d-%d-%d", year, week, st)
}
