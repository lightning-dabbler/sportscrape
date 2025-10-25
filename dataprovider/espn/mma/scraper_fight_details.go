package mma

import (
	"fmt"
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
const ESPNMMAEventURL = "https://www.espn.com/mma/fightcenter/_/id/%s/league/ufc"

type ESPNMMAFightDetailsScraper struct {
	scraper.BaseScraper
}

func (e ESPNMMAFightDetailsScraper) Scrape(matchup interface{}) sportscrape.EventDataOutput {

	jsonRetriever := scraper.BaseJsonScraper[jsonresponse.ESPNEventData]{}

	m, ok := matchup.(model.Matchup)
	if !ok {
		return sportscrape.EventDataOutput{
			Error: fmt.Errorf("Input is not of type espn.mma.model.Matchup"),
		}
	}

	url := fmt.Sprintf(ESPNMMAEventURL, m.EventID)
	doc, err := e.RetrieveDocument(url, network.Headers{}, "html")
	if err != nil {
		return sportscrape.EventDataOutput{
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

	fights := data.GetFightDetails(m)

	out := make([]interface{}, 0, len(fights))
	for _, fight := range fights {
		out = append(out, fight)
	}

	return sportscrape.EventDataOutput{
		Error:  nil,
		Output: out,
		Context: sportscrape.EventDataContext{
			PullTimestamp: time.Now(),
			EventTime:     m.EventTime,
			EventID:       m.EventID,
			URL:           url,
			AwayID:        "NA/Multiple",
			AwayTeam:      "NA/Multiple",
			HomeID:        "NA/Multiple",
			HomeTeam:      "NA/Multiple",
		},
	}
}

func (e ESPNMMAFightDetailsScraper) Feed() sportscrape.Feed {
	return sportscrape.ESPNMMAFightDetails
}

func (e ESPNMMAFightDetailsScraper) Provider() sportscrape.Provider {
	return sportscrape.ESPNMMA
}
