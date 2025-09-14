package mma

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/model"
	"github.com/lightning-dabbler/sportscrape/scraper"
	"github.com/lightning-dabbler/sportscrape/util/request"
)

// https://www.espn.com/mma/fightcenter/_/id/600040033/league/ufc
const ESPNMMAEventURL = "https://www.espn.com/mma/fightcenter/_/id/%s/league/ufc"

type espnEventDataScraper struct {
	scraper.BaseScraper
}

func (e espnEventDataScraper) Scrape(id string) (data *model.ESPNEventData, err error) {
	jsonRetriever := request.JsonRetriever[model.ESPNEventData]{}

	url := fmt.Sprintf(ESPNMMAEventURL, id)
	doc, err := e.RetrieveDocument(url, network.Headers{}, "html")
	if err != nil {
		return nil, err
	}

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
	return
}
