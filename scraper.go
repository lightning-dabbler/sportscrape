package sportscrape

type MatchupScraper interface {
	Scrape() MatchupOutput
	Init()
	Feed() Feed
	Provider() Provider
}

type EventDataScraper interface {
	Scrape(matchup interface{}) EventDataOutput
	Feed() Feed
	Provider() Provider
	Init()
}
