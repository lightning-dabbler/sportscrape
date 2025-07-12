package sportscrape

type MatchupOutput struct {
	Error   error
	Output  []interface{}
	Context MatchupContext
}

type EventDataOutput struct {
	Error   error
	Output  []interface{}
	Context EventDataContext
}
