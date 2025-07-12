package scraper

import "github.com/lightning-dabbler/sportscrape"

type MatchupScraper interface {
	Scrape() sportscrape.MatchupOutput
	Init()
	Feed() sportscrape.Feed
	Provider() sportscrape.Provider
}

type EventDataScraper interface {
	Scrape(matchup interface{}) sportscrape.EventDataOutput
	Feed() sportscrape.Feed
	Provider() sportscrape.Provider
	Init()
}
