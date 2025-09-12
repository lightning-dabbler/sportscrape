package mma

import (
	"fmt"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/model"
	"github.com/lightning-dabbler/sportscrape/util/sportsreference"
)

const ESPNMMAEventsFeedURL = "https://www.espn.com/mma/schedule/_/year/%s?_xhr=pageContent"

type ESPNMMAEventsFeedScraper struct {
	sportsreference.JsonScraper[model.ESPNMMAEventsFeed]
	Year string
}

func (e ESPNMMAEventsFeedScraper) Scrape(year string) (model *model.ESPNMMAEventsFeed, err error) {
	url := fmt.Sprintf(ESPNMMAEventsFeedURL, year)
	model, err = e.RetrieveModel(url)
	return
}

func (e ESPNMMAEventsFeedScraper) Feed() sportscrape.Feed {
	//TODO implement me

	return sportscrape.ESPNMMAEvents
}

func (e ESPNMMAEventsFeedScraper) Provider() sportscrape.Provider {
	return sportscrape.ESPNMMA
}
