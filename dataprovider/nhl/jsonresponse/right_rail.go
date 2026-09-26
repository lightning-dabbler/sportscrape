package jsonresponse

// RightRail is the json response from https://api-web.nhle.com/v1/gamecenter/{game_id}/right-rail
type RightRail struct {
	// Linescore is absent before the game starts
	Linescore *Linescore `json:"linescore"`
	// ShotsByPeriod is absent before the game starts
	ShotsByPeriod []PeriodTally `json:"shotsByPeriod"`
}

type Linescore struct {
	ByPeriod []PeriodTally `json:"byPeriod"`
}

// PeriodTally is an away/home count for a period (goals in the linescore, shots on goal in shotsByPeriod)
type PeriodTally struct {
	PeriodDescriptor PeriodDescriptor `json:"periodDescriptor"`
	Away             int32            `json:"away"`
	Home             int32            `json:"home"`
}
