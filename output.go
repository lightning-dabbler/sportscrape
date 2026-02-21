package sportscrape

type MatchupOutput[M any] struct {
	Error   error
	Output  []M
	Context MatchupContext
}

type EventDataOutput[E any] struct {
	Error   error
	Output  []E
	Context EventDataContext
}
