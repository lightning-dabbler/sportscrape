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

// https://www.espn.com/mma/fightcenter/_/id/600040033/league/ufc
const ESPNMMAEventURL = "https://www.espn.com/mma/fightcenter/_/id/%s/league/%s"

type ESPNMMAFightDetailsScraper struct {
	scraper.BaseDocumentScraper
	League string //ufc or PFL
	// FetchAttempts is the number of times to retry fetching the fight card
	// page if ESPN serves a bot-check interstitial instead of real content.
	// <= 0 falls back to DefaultFetchAttempts.
	FetchAttempts int
	// FetchRetryBackoff is the delay between retry attempts.
	// <= 0 falls back to DefaultFetchRetryBackoff.
	FetchRetryBackoff time.Duration
}

func (e *ESPNMMAFightDetailsScraper) Init() {
	if e.League != "pfl" && e.League != "ufc" {
		log.Fatalln("League must be either pfl or ufc")
	}
	e.BaseDocumentScraper.Init()
}

func (e *ESPNMMAFightDetailsScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.FightDetails] {
	url := fmt.Sprintf(ESPNMMAEventURL, matchup.EventID, e.League)

	payload, err := fetchESPNFittPayload(e.FetchDoc, url, "html", e.FetchAttempts, e.FetchRetryBackoff)
	if err != nil {
		return sportscrape.EventDataOutput[model.FightDetails]{
			Error: err,
		}
	}

	jsonRetriever := scraper.BaseJsonScraper[jsonresponse.ESPNEventData]{}
	data, err := jsonRetriever.HydrateModel(payload)
	if err != nil {
		return sportscrape.EventDataOutput[model.FightDetails]{
			Error: fmt.Errorf("could not unmarshall event data: %w", err),
		}
	}

	data.PullTime = time.Now()

	fights := data.GetFightDetails(matchup)

	out := make([]model.FightDetails, 0, len(fights))
	out = append(out, fights...)
	return sportscrape.EventDataOutput[model.FightDetails]{
		Error:  nil,
		Output: out,
		Context: sportscrape.EventDataContext{
			PullTimestamp: data.PullTime,
			EventTime:     matchup.EventTime,
			EventID:       matchup.EventID,
			URL:           url,
			AwayID:        "NA/Multiple",
			AwayTeam:      "NA/Multiple",
			HomeID:        "NA/Multiple",
			HomeTeam:      "NA/Multiple",
		},
	}
}

func (e *ESPNMMAFightDetailsScraper) Feed() sportscrape.Feed {
	switch e.League {
	case "ufc":
		return sportscrape.ESPNUFCFightDetails
	case "pfl":
		return sportscrape.ESPNPFLFightDetails
	default:
		return ""
	}
}

func (e *ESPNMMAFightDetailsScraper) Provider() sportscrape.Provider {
	return sportscrape.ESPNMMA
}
