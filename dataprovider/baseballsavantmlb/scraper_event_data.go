package baseballsavantmlb

import (
	"encoding/json"
	"io"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/model"
	"github.com/lightning-dabbler/sportscrape/util/request"
)

type EventDataScraper struct{}

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
	response, err := request.Get(url)
	if err != nil {
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
