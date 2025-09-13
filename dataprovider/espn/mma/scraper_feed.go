package mma

import (
	"fmt"

	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/model"
	"github.com/lightning-dabbler/sportscrape/util/request"
)

const ESPNMMAEventsFeedURL = "https://www.espn.com/mma/schedule/_/year/%s?_xhr=pageContent"

type espnScraperFeed struct {
	request.JsonRetriever[model.ESPNMMAFeed]
}

func (e espnScraperFeed) Scrape(year string) (model *model.ESPNMMAFeed, err error) {
	url := fmt.Sprintf(ESPNMMAEventsFeedURL, year)
	model, err = e.RetrieveModel(url)
	return
}
