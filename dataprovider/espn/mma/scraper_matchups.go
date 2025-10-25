package mma

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/scraper"
)

// const ESPNMMAEventsFeedURL = "https://www.espn.com/mma/schedule/_/year/%s?_xhr=pageContent"
const ESPNMMAEventsFeedURL = "https://www.espn.com/mma/schedule/_/year/%s/league/%s"

type ESPNMMAMatchupScraper struct {
	scraper.BaseScraper
	Year   string
	League string
}

func (m ESPNMMAMatchupScraper) Scrape() sportscrape.MatchupOutput {
	url := fmt.Sprintf(ESPNMMAEventsFeedURL, m.Year, m.League)

	doc, err := m.RetrieveDocument(url, network.Headers{}, "html")

	data := &jsonresponse.ESPNMMASchedule{}

	jsonRetriever := scraper.BaseJsonScraper[jsonresponse.ESPNMMASchedule]{}

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

	empty := &jsonresponse.ESPNMMASchedule{}
	if data == empty {
		return sportscrape.MatchupOutput{
			Context: sportscrape.MatchupContext{
				Errors: 1,
				Skips:  1,
			},
			Error: errors.New("could not unmarshall schedule data"),
		}
	}
	data.PullTime = time.Now()
	matchups := data.GetScrapableMatchup()
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
