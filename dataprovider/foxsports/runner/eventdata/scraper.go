package eventdata

import (
	"io"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/util/request"
)

type Scraper interface {
	Scrape(matchup interface{}) OutputWrapper
	SetParams()
}

type EventDataScraper struct {
	Name string
	// League - The league of interest to fetch matchups data
	League foxsports.League
	// Params - URL Query parameters
	Params map[string]string
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

func (e *EventDataScraper) FetchEventData(url string) ([]byte, error) {
	response, err := request.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

func (e *EventDataScraper) SetParams() {
	e.League.SetParams(e.Params)
}
