package baseballsavantmlb

import (
	"encoding/json"
	"io"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/model"
	"github.com/lightning-dabbler/sportscrape/util/request"
)

type EventDataScraper struct {
	// Timeout is the request timeout. <= 0 falls back to request.DefaultGetTimeout.
	Timeout time.Duration
}

func (e EventDataScraper) Init() {}

func (e EventDataScraper) ConstructContext(matchup model.Matchup) sportscrape.EventDataContext {
	return sportscrape.EventDataContext{
		AwayTeam:  matchup.AwayTeamName,
		AwayID:    matchup.AwayTeamID,
		HomeTeam:  matchup.HomeTeamName,
		HomeID:    matchup.HomeTeamID,
		EventTime: matchup.EventTime,
		EventID:   matchup.EventID,
	}
}

func (e EventDataScraper) FetchGameFeed(url string) (jsonresponse.GameFeed, error) {
	var responsePayload jsonresponse.GameFeed
	response, err := request.GetWithTimeout(url, e.Timeout)
	if err != nil {
		return responsePayload, err
	}
	if err := request.CheckStatus(url, response); err != nil {
		return responsePayload, err
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return responsePayload, err
	}
	err = json.Unmarshal(responseBody, &responsePayload)
	if err != nil {
		return responsePayload, err
	}
	return responsePayload, nil
}

func (e EventDataScraper) Provider() sportscrape.Provider {
	return sportscrape.BaseballSavant
}

func (e EventDataScraper) FmtID(playerid string) string {
	return "ID" + playerid
}

func (e EventDataScraper) Close() {}
