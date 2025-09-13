package mma

import (
	"sync"

	"github.com/lightning-dabbler/sportscrape"
)

type ScraperMatchups struct {
	feedScraper  espnScraperFeed
	eventScraper espnEventDataScraper
	years        []string
	concurrency  int
}

func NewScraperMatchups(concurrency int, years []string) *ScraperMatchups {
	return &ScraperMatchups{
		concurrency:  concurrency,
		feedScraper:  espnScraperFeed{},
		eventScraper: espnEventDataScraper{},
		years:        years,
	}
}

func (s *ScraperMatchups) Scrape() sportscrape.MatchupOutput {
	output := sportscrape.MatchupOutput{
		Error:  nil,
		Output: make([]interface{}, 0, 100),
		Context: sportscrape.MatchupContext{
			Errors: 0,
			Skips:  0,
		},
	}
	for _, year := range s.years {
		feed, err := s.feedScraper.Scrape(year)
		if err != nil {
			output.Error = err
			return output
		}

		events := feed.FilterScrapeableEvents()

		sem := make(chan struct{}, s.concurrency)
		var wg sync.WaitGroup
		lock := sync.Mutex{}

		for _, event := range events {
			sem <- struct{}{}
			wg.Add(1)
			go func(id string) {
				defer wg.Done()
				data, err := s.eventScraper.Scrape(id)
				lock.Lock()
				if err != nil {
					output.Context.Errors++
				} else {
					for _, m := range data.GetMatchups() {
						output.Output = append(output.Output, m)
					}
				}
				lock.Unlock()
				<-sem
			}(event.ID)
		}
		wg.Wait()

	}
	return output
}

func (s *ScraperMatchups) Init() {}

func (s *ScraperMatchups) Feed() sportscrape.Feed {
	return sportscrape.ESPNMMAMatchups
}

func (s *ScraperMatchups) Provider() sportscrape.Provider {
	return sportscrape.ESPNMMA
}
