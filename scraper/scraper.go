package scraper

import "github.com/lightning-dabbler/sportscrape"

type MatchupScraper[M any] interface {
	Scrape() sportscrape.MatchupOutput[M]
	Init()
	Feed() sportscrape.Feed
	Provider() sportscrape.Provider
}

type EventDataScraper[M, E any] interface {
	Scrape(matchup M) sportscrape.EventDataOutput[E]
	Feed() sportscrape.Feed
	Provider() sportscrape.Provider
	Init()
}
