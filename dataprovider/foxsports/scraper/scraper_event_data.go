package scraper

import (
	"io"
	"log"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/util/request"
)

type EventDataScraper struct {
	// League - The league of interest to fetch matchups data
	League foxsports.League
	// Params - URL Query parameters
	Params map[string]string
}

func (e *EventDataScraper) Init() {
	// Ensure League is set
	if e.League.Undefined() {
		log.Fatalln("League is a required argument for foxsports EventDataScraper")
	}
	// Params
	if e.Params == nil {
		e.Params = map[string]string{}
	}
	e.League.SetParams(e.Params)
}

func (e *EventDataScraper) ConstructEventDataURL(eventID int64) (string, error) {
	url, err := e.League.V1EventDataURL(eventID)
	if err != nil {
		return "", err
	}
	queryValues := url.Query()
	for k, v := range e.Params {
		queryValues.Add(k, v)
	}
	url.RawQuery = queryValues.Encode()
	return url.String(), nil
}

func (e *EventDataScraper) ConstructMatchupComparisonURL(eventID int64) (string, error) {
	url, err := e.League.V1MatchupComparisonURL(eventID)
	if err != nil {
		return "", err
	}
	queryValues := url.Query()
	for k, v := range e.Params {
		queryValues.Add(k, v)
	}
	url.RawQuery = queryValues.Encode()
	return url.String(), nil
}

func (e *EventDataScraper) FetchData(url string) ([]byte, error) {
	response, err := request.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

func (e EventDataScraper) Provider() sportscrape.Provider {
	return sportscrape.FS
}
