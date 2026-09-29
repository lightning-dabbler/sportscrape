package jsonresponse

// PlayerStatistics is the response of https://api.nfl.com/football/v2/stats/live/player-statistics/{game_id}
type PlayerStatistics struct {
	GameID   string       `json:"gameId"`
	HomeTeam TeamStatsSet `json:"homeTeam"`
	AwayTeam TeamStatsSet `json:"awayTeam"`
}

type TeamStatsSet struct {
	TeamID  string        `json:"teamId"`
	Players []PlayerStats `json:"players"`
}

type PlayerStats struct {
	GSISPlayerID           string `json:"gsisPlayerId"`
	GSISPlayerJerseyNumber string `json:"gsisPlayerJerseyNumber"`
	GSISPlayerName         string `json:"gsisPlayerName"`
	PersonID               string `json:"personId"`

	DefensiveFumblesForced                 int32   `json:"defensiveFumblesForced"`
	DefensiveFumblesRecovered              int32   `json:"defensiveFumblesRecovered"`
	DefensiveInterceptions                 int32   `json:"defensiveInterceptions"`
	DefensiveMiscellaneousFumblesForced    int32   `json:"defensiveMiscellaneousFumblesForced"`
	DefensiveMiscellaneousFumblesRecovered int32   `json:"defensiveMiscellaneousFumblesRecovered"`
	DefensiveMiscellaneousTackles          float32 `json:"defensiveMiscellaneousTackles"`
	DefensiveMiscellaneousTacklesAssists   int32   `json:"defensiveMiscellaneousTacklesAssists"`
	DefensivePassesDefended                int32   `json:"defensivePassesDefended"`
	DefensiveQuarterbackHits               int32   `json:"defensiveQuarterbackHits"`
	DefensiveSacks                         float32 `json:"defensiveSacks"`
	DefensiveSackYards                     float32 `json:"defensiveSackYards"`
	DefensiveSafeties                      int32   `json:"defensiveSafeties"`
	DefensiveSpecialTeamsFumblesForced     int32   `json:"defensiveSpecialTeamsFumblesForced"`
	DefensiveSpecialTeamsFumblesRecovered  int32   `json:"defensiveSpecialTeamsFumblesRecovered"`
	DefensiveSpecialTeamsTackles           float32 `json:"defensiveSpecialTeamsTackles"`
	DefensiveSpecialTeamsTacklesAssists    int32   `json:"defensiveSpecialTeamsTacklesAssists"`
	DefensiveSpecialTeamsBlocks            int32   `json:"defensiveSpecialTeamsBlocks"`
	DefensiveTackles                       float32 `json:"defensiveTackles"`
	DefensiveTacklesAssists                int32   `json:"defensiveTacklesAssists"`
	DefensiveTacklesCombined               float32 `json:"defensiveTacklesCombined"`
	DefensiveTacklesForLoss                float32 `json:"defensiveTacklesForLoss"`
	DefensiveTacklesForLossYards           float32 `json:"defensiveTacklesForLossYards"`

	ExtraPointsAttempted int32 `json:"extraPointsAttempted"`
	ExtraPointsMade      int32 `json:"extraPointsMade"`
	ExtraPointsMissed    int32 `json:"extraPointsMissed"`
	ExtraPointsBlocked   int32 `json:"extraPointsBlocked"`

	FieldGoalsAttempted     int32   `json:"fieldGoalsAttempted"`
	FieldGoalsAverageLength float32 `json:"fieldGoalsAverageLength"`
	FieldGoalsBlocked       int32   `json:"fieldGoalsBlocked"`
	FieldGoalsLongestMade   int32   `json:"fieldGoalsLongestMade"`
	FieldGoalsMade          int32   `json:"fieldGoalsMade"`
	FieldGoalsMissed        int32   `json:"fieldGoalsMissed"`
	FieldGoalsTotalYards    int32   `json:"fieldGoalsTotalYards"`

	Fumbles                               int32 `json:"fumbles"`
	FumblesForced                         int32 `json:"fumblesForced"`
	FumblesLost                           int32 `json:"fumblesLost"`
	FumblesRecoveredInEndZoneForTouchdown int32 `json:"fumblesRecoveredInEndZoneForTouchdown"`
	FumblesOpponentRecoveries             int32 `json:"fumblesOpponentRecoveries"`
	FumblesOpponentRecoveryTouchdowns     int32 `json:"fumblesOpponentRecoveryTouchdowns"`
	FumblesOpponentRecoveryYards          int32 `json:"fumblesOpponentRecoveryYards"`
	FumblesOutOfBounds                    int32 `json:"fumblesOutOfBounds"`
	FumblesOwnRecoveries                  int32 `json:"fumblesOwnRecoveries"`
	FumblesOwnRecoveryTouchdowns          int32 `json:"fumblesOwnRecoveryTouchdowns"`
	FumblesOwnRecoveryYards               int32 `json:"fumblesOwnRecoveryYards"`

	Interceptions                 int32 `json:"interceptions"`
	InterceptionsLong             int32 `json:"interceptionsLong"`
	InterceptionsLongestTouchdown int32 `json:"interceptionsLongestTouchdown"`
	InterceptionsTouchdowns       int32 `json:"interceptionsTouchdowns"`
	InterceptionsYards            int32 `json:"interceptionsYards"`

	Kickoffs            int32 `json:"kickoffs"`
	KickoffsInside20    int32 `json:"kickoffsInside20"`
	KickoffsOutOfBounds int32 `json:"kickoffsOutOfBounds"`
	KickoffsReturnYards int32 `json:"kickoffsReturnYards"`
	KickoffsToEndZone   int32 `json:"kickoffsToEndZone"`
	KickoffsTouchbacks  int32 `json:"kickoffsTouchbacks"`
	KickoffsYards       int32 `json:"kickoffsYards"`

	KickReturns                 int32   `json:"kickReturns"`
	KickReturnsFairCatches      int32   `json:"kickReturnsFairCatches"`
	KickReturnsLongest          int32   `json:"kickReturnsLongest"`
	KickReturnsLongestTouchdown int32   `json:"kickReturnsLongestTouchdown"`
	KickReturnsTouchdowns       int32   `json:"kickReturnsTouchdowns"`
	KickReturnsYards            int32   `json:"kickReturnsYards"`
	KickReturnsYardsAverage     float32 `json:"kickReturnsYardsAverage"`

	PassingAttempts             int32   `json:"passingAttempts"`
	PassingCompletions          int32   `json:"passingCompletions"`
	PassingCompletionPercent    float32 `json:"passingCompletionPercent"`
	PassingInterceptions        int32   `json:"passingInterceptions"`
	PassingLong                 int32   `json:"passingLong"`
	PassingLongestTouchdownPass int32   `json:"passingLongestTouchdownPass"`
	PassingRating               float32 `json:"passingRating"`
	PassingSackYardsLost        float32 `json:"passingSackYardsLost"`
	PassingTimesSacked          int32   `json:"passingTimesSacked"`
	PassingTouchdowns           int32   `json:"passingTouchdowns"`
	PassingYards                int32   `json:"passingYards"`
	PassingYardsAverage         float32 `json:"passingYardsAverage"`

	Punts                  int32   `json:"punts"`
	PuntsBlocked           int32   `json:"puntsBlocked"`
	PuntsInside20          int32   `json:"puntsInside20"`
	PuntsLongest           int32   `json:"puntsLongest"`
	PuntsReturnYards       int32   `json:"puntsReturnYards"`
	PuntsTouchbacks        int32   `json:"puntsTouchbacks"`
	PuntsYards             int32   `json:"puntsYards"`
	PuntsYardsAverageGross float32 `json:"puntsYardsAverageGross"`
	PuntsYardsAverageNet   float32 `json:"puntsYardsAverageNet"`

	PuntReturns                 int32   `json:"puntReturns"`
	PuntReturnsFairCatches      int32   `json:"puntReturnsFairCatches"`
	PuntReturnsLongest          int32   `json:"puntReturnsLongest"`
	PuntReturnsLongestTouchdown int32   `json:"puntReturnsLongestTouchdown"`
	PuntReturnsTouchdowns       int32   `json:"puntReturnsTouchdowns"`
	PuntReturnsYards            int32   `json:"puntReturnsYards"`
	PuntReturnsYardsAverage     float32 `json:"puntReturnsYardsAverage"`

	Receptions                 int32   `json:"receptions"`
	ReceptionsAverage          float32 `json:"receptionsAverage"`
	ReceptionsLong             int32   `json:"receptionsLong"`
	ReceptionsLongestTouchdown int32   `json:"receptionsLongestTouchdown"`
	ReceptionsPassTarget       int32   `json:"receptionsPassTarget"`
	ReceptionsTouchdowns       int32   `json:"receptionsTouchdowns"`
	ReceptionsYards            int32   `json:"receptionsYards"`
	ReceptionsYardsAfterCatch  int32   `json:"receptionsYardsAfterCatch"`

	RushingAttempts         int32   `json:"rushingAttempts"`
	RushingAverage          float32 `json:"rushingAverage"`
	RushingLong             int32   `json:"rushingLong"`
	RushingLongestTouchdown int32   `json:"rushingLongestTouchdown"`
	RushingTouchdowns       int32   `json:"rushingTouchdowns"`
	RushingYards            int32   `json:"rushingYards"`

	TwoPointDefensiveAttempts  int32 `json:"twoPointDefensiveAttempts"`
	TwoPointDefensiveSuccesses int32 `json:"twoPointDefensiveSuccesses"`
	TwoPointPassingAttempts    int32 `json:"twoPointPassingAttempts"`
	TwoPointPassingSuccesses   int32 `json:"twoPointPassingSuccesses"`
	TwoPointReceptionAttempts  int32 `json:"twoPointReceptionAttempts"`
	TwoPointReceptionSuccesses int32 `json:"twoPointReceptionSuccesses"`
	TwoPointRushingAttempts    int32 `json:"twoPointRushingAttempts"`
	TwoPointRushingSuccesses   int32 `json:"twoPointRushingSuccesses"`
}
