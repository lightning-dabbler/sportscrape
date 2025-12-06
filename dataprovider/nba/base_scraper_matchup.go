package nba

import (
	"log"
	"net/url"

	"github.com/lightning-dabbler/sportscrape/util"
)

const (
	BaseURL = "https://www.nba.com/games"
)

type BaseMatchupScraper struct {
	Scraper
	Date string
}

func (bms BaseMatchupScraper) Init() {
	bms.Scraper.Init()
	if bms.Date == "" {
		log.Fatalln("Date is a required argument")
	}
	// Validate Date in the form YYYY-MM-DD
	_, err := util.DateStrToTime(bms.Date)
	if err != nil {
		log.Fatalln(err)
	}
}

func (bms BaseMatchupScraper) URL() (string, error) {
	URL, err := url.Parse(BaseURL)
	if err != nil {
		return "", err
	}
	queryValues := URL.Query()
	queryValues.Add("date", bms.Date)
	URL.RawQuery = queryValues.Encode()
	return URL.String(), nil
}
