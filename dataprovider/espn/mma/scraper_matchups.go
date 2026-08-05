package mma

import (
	"fmt"
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/model"
	"github.com/lightning-dabbler/sportscrape/scraper"
)

// const ESPNMMAEventsFeedURL = "https://www.espn.com/mma/schedule/_/year/%s?_xhr=pageContent"
const ESPNMMAEventsFeedURL = "https://www.espn.com/mma/schedule/_/year/%s/league/%s"

type ESPNMMAMatchupScraper struct {
	scraper.BaseDocumentScraper
	Year   string
	League string
	// FetchAttempts is the number of times to retry fetching the schedule
	// page if ESPN serves a bot-check interstitial instead of real content.
	// <= 0 falls back to DefaultFetchAttempts.
	FetchAttempts int
	// FetchRetryBackoff is the delay between retry attempts.
	// <= 0 falls back to DefaultFetchRetryBackoff.
	FetchRetryBackoff time.Duration
}

func (m *ESPNMMAMatchupScraper) Init() {
	if m.Year == "" {
		log.Fatalln("Year is a required argument")
	}
	if m.League == "" {
		log.Fatalln("League is a required argument")
	}
	if m.League != "pfl" && m.League != "ufc" {
		log.Fatalln("League must be either pfl or ufc")
	}
	m.BaseDocumentScraper.Init()
}

func (m *ESPNMMAMatchupScraper) Scrape() sportscrape.MatchupOutput[model.Matchup] {
	url := fmt.Sprintf(ESPNMMAEventsFeedURL, m.Year, m.League)

	payload, err := fetchESPNFittPayload(m.FetchDoc, url, "html", m.FetchAttempts, m.FetchRetryBackoff)
	if err != nil {
		return sportscrape.MatchupOutput[model.Matchup]{
			Context: sportscrape.MatchupContext{
				Errors: 1,
				Skips:  0,
			}, Error: err,
		}
	}

	jsonRetriever := scraper.BaseJsonScraper[jsonresponse.ESPNMMASchedule]{}
	data, err := jsonRetriever.HydrateModel(payload)
	if err != nil {
		return sportscrape.MatchupOutput[model.Matchup]{
			Context: sportscrape.MatchupContext{
				Errors: 1,
				Skips:  1,
			},
			Error: fmt.Errorf("could not unmarshall schedule data: %w", err),
		}
	}
	data.PullTime = time.Now()
	matchups := data.GetScrapableMatchup()
	output := make([]model.Matchup, 0, len(matchups))
	output = append(output, matchups...)

	return sportscrape.MatchupOutput[model.Matchup]{
		Context: sportscrape.MatchupContext{
			Errors: 0,
		},
		Output: output,
	}
}

func (m *ESPNMMAMatchupScraper) Feed() sportscrape.Feed {
	if m.League == "ufc" {
		return sportscrape.ESPNUFCMatchups
	} else if m.League == "pfl" {
		return sportscrape.ESPNPFLMatchups
	}
	return ""
}

func (m *ESPNMMAMatchupScraper) Provider() sportscrape.Provider {
	return sportscrape.ESPNMMA
}
