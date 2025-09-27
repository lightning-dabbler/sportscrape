package mma

import (
	"fmt"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/scraper"
)

const ESPNMMAEventsFeedURL = "https://www.espn.com/mma/schedule/_/year/%s?_xhr=pageContent"

type ESPNMMAMatchupScraper struct {
	scraper.BaseJsonScraper[jsonresponse.ESPNMMASchedule]
	Year string
}

func (m ESPNMMAMatchupScraper) Scrape() sportscrape.MatchupOutput {
	url := fmt.Sprintf(ESPNMMAEventsFeedURL, m.Year)

	jsonModel, err := m.RetrieveModel(url)

	if err != nil {
		return sportscrape.MatchupOutput{
			Context: sportscrape.MatchupContext{
				Errors: 1,
				Skips:  1,
			},
			Error: err,
		}
	}
	jsonModel.PullTime = time.Now()
	matchups := jsonModel.GetScrapableMatchup()
	output := make([]interface{}, 0, len(matchups))
	for _, matchup := range matchups {
		output = append(output, matchup)
	}

	return sportscrape.MatchupOutput{
		Context: sportscrape.MatchupContext{
			Errors: 0,
		},
		Output: output,
		Error:  err,
	}
}

func (m ESPNMMAMatchupScraper) Feed() sportscrape.Feed {
	return sportscrape.ESPNMMAMatchups
}

func (m ESPNMMAMatchupScraper) Provider() sportscrape.Provider {
	return sportscrape.ESPNMMA
}
