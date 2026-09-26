package jsonresponse

// BoxScore is the json response from https://api-web.nhle.com/v1/gamecenter/{game_id}/boxscore
type BoxScore struct {
	ID        int64  `json:"id"`
	GameState string `json:"gameState"`
	// PlayerByGameStats is absent before the game starts
	PlayerByGameStats *PlayerByGameStats `json:"playerByGameStats"`
}

type PlayerByGameStats struct {
	AwayTeam TeamPlayers `json:"awayTeam"`
	HomeTeam TeamPlayers `json:"homeTeam"`
}

type TeamPlayers struct {
	Forwards []Skater `json:"forwards"`
	Defense  []Skater `json:"defense"`
	Goalies  []Goalie `json:"goalies"`
}

type Skater struct {
	PlayerID      int64 `json:"playerId"`
	SweaterNumber int32 `json:"sweaterNumber"`
	// Name e.g. J. Huberdeau
	Name           LocalizedString `json:"name"`
	Position       string          `json:"position"`
	Goals          int32           `json:"goals"`
	Assists        int32           `json:"assists"`
	Points         int32           `json:"points"`
	PlusMinus      int32           `json:"plusMinus"`
	PIM            int32           `json:"pim"`
	Hits           int32           `json:"hits"`
	PowerPlayGoals int32           `json:"powerPlayGoals"`
	SOG            int32           `json:"sog"`
	// FaceoffWinningPctg is absent for some players (e.g. every defenseman in game 2026010045)
	FaceoffWinningPctg *float32 `json:"faceoffWinningPctg"`
	// TOI e.g. 18:45
	TOI          string `json:"toi"`
	BlockedShots int32  `json:"blockedShots"`
	Shifts       int32  `json:"shifts"`
	Giveaways    int32  `json:"giveaways"`
	Takeaways    int32  `json:"takeaways"`
}

type Goalie struct {
	PlayerID      int64           `json:"playerId"`
	SweaterNumber int32           `json:"sweaterNumber"`
	Name          LocalizedString `json:"name"`
	Position      string          `json:"position"`
	// saves/shots e.g. 26/28
	EvenStrengthShotsAgainst string `json:"evenStrengthShotsAgainst"`
	PowerPlayShotsAgainst    string `json:"powerPlayShotsAgainst"`
	ShorthandedShotsAgainst  string `json:"shorthandedShotsAgainst"`
	SaveShotsAgainst         string `json:"saveShotsAgainst"`
	// SavePctg is absent when no shots were faced
	SavePctg                 *float32 `json:"savePctg"`
	EvenStrengthGoalsAgainst int32    `json:"evenStrengthGoalsAgainst"`
	PowerPlayGoalsAgainst    int32    `json:"powerPlayGoalsAgainst"`
	ShorthandedGoalsAgainst  int32    `json:"shorthandedGoalsAgainst"`
	PIM                      int32    `json:"pim"`
	GoalsAgainst             int32    `json:"goalsAgainst"`
	TOI                      string   `json:"toi"`
	Starter                  bool     `json:"starter"`
	// Decision is absent when the goalie did not receive a decision
	Decision     *string `json:"decision"`
	ShotsAgainst int32   `json:"shotsAgainst"`
	Saves        int32   `json:"saves"`
}
