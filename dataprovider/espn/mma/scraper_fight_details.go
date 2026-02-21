package mma

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/model"
	"github.com/lightning-dabbler/sportscrape/scraper"
)

// https://www.espn.com/mma/fightcenter/_/id/600040033/league/ufc
const ESPNMMAEventURL = "https://www.espn.com/mma/fightcenter/_/id/%s/league/%s"

type ESPNMMAFightDetailsScraper struct {
	scraper.BaseScraper
	League string //ufc or PFL
}

func (e ESPNMMAFightDetailsScraper) Init() {
	e.BaseScraper.Init()
	if e.League != "pfl" && e.League != "ufc" {
		log.Fatalln("League must be either pfl or ufc")
	}
}

func (e ESPNMMAFightDetailsScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.FightDetails] {

	jsonRetriever := scraper.BaseJsonScraper[jsonresponse.ESPNEventData]{}

	url := fmt.Sprintf(ESPNMMAEventURL, matchup.EventID, e.League)
	doc, err := e.RetrieveDocument(url, network.Headers{}, "html")
	if err != nil {
		return sportscrape.EventDataOutput[model.FightDetails]{
			Error: err,
		}
	}

	data := &jsonresponse.ESPNEventData{}

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		text := s.Text()

		if strings.Contains(text, "window['__espnfitt__']=") {
			parts := strings.SplitAfter(text, "window['__espnfitt__']=")
			payload := []byte(parts[1][0 : len(parts[1])-1])
			result, err := jsonRetriever.HydrateModel(payload)
			if err == nil {
				data = result
			}
		}
	})

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

func (e ESPNMMAFightDetailsScraper) Feed() sportscrape.Feed {
	switch e.League {
	case "ufc":
		return sportscrape.ESPNUFCFightDetails
	case "pfl":
		return sportscrape.ESPNPFLFightDetails
	default:
		return ""
	}
}

func (e ESPNMMAFightDetailsScraper) Provider() sportscrape.Provider {
	return sportscrape.ESPNMMA
}
