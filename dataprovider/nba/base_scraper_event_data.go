package nba

import (
	"fmt"
	"log"
	"net/url"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
)

type BaseEventDataScraper struct {
	Scraper
	Period       Period
	FeedType     FeedType
	BoxScoreType BoxScoreType
}

func (beds *BaseEventDataScraper) Init() {
	beds.Scraper.Init()
	if beds.FeedType.Undefined() {
		log.Fatalln("FeedType is a required argument")
	}

	switch beds.FeedType {
	case BoxScore:
		if beds.BoxScoreType.Undefined() {
			log.Fatalln("BoxScoreType is a required argument when FeedType is BoxScore")
		}
		switch beds.BoxScoreType {
		case Traditional, Advanced, Misc, Scoring, Usage, FourFactors:
			if beds.Period.Undefined() {
				log.Printf("Warning: Period is unset for BoxScore FeedType (%s)... defaulting to %s\n", beds.BoxScoreType.Type(), Full.Period())
				beds.Period = Full
			}
		}
	case PlayByPlay:
		if !beds.BoxScoreType.Undefined() {
			log.Println("Warning: BoxScoreType argument will be ignored for PlayByPlay FeedType")
		}
		if beds.Period != Full {
			log.Println("Setting Period to Full for PlayByPlay FeedType")
			beds.Period = Full
		}
	}
}

func (beds BaseEventDataScraper) ConstructContext(matchup model.Matchup) sportscrape.EventDataContext {
	return sportscrape.EventDataContext{
		AwayTeam:  matchup.AwayTeam,
		AwayID:    matchup.AwayTeamID,
		HomeTeam:  matchup.HomeTeam,
		HomeID:    matchup.HomeTeamID,
		EventTime: matchup.EventTime,
		EventID:   matchup.EventID,
	}
}

func (beds BaseEventDataScraper) URL(share_url string) (string, error) {
	if share_url == "" {
		return "", fmt.Errorf("share_url should not be empty")
	}
	URL, err := url.Parse(share_url)
	if err != nil {
		return "", err
	}
	var urlstr string
	switch beds.FeedType {
	case BoxScore:
		joinedURL := URL.JoinPath(BoxScore.Type())
		queryValues := joinedURL.Query()
		switch beds.BoxScoreType {
		case Traditional, Advanced, Misc, Scoring, Usage, FourFactors:
			queryValues.Add("period", beds.Period.Period())
		}
		if beds.BoxScoreType != Live {
			queryValues.Add("type", beds.BoxScoreType.Type())
		}
		joinedURL.RawQuery = queryValues.Encode()
		urlstr = joinedURL.String()

	case PlayByPlay:
		joinedURL := URL.JoinPath(PlayByPlay.Type())
		queryValues := joinedURL.Query()
		queryValues.Add("period", beds.Period.Period())
		joinedURL.RawQuery = queryValues.Encode()
		urlstr = joinedURL.String()
	}

	if urlstr == "" {
		return "", fmt.Errorf("result urlstr should not be empty")
	}
	return urlstr, nil
}

func (beds BaseEventDataScraper) PeriodBasedBoxScoreDataAvailable(period int32, gameStatus int32) bool {
	// Game is Final
	if gameStatus == int32(3) {
		switch beds.Period {
		case AllOT:
			if period > 4 {
				return true
			}
		default:
			return true
		}
	}
	return false
}

func (beds BaseEventDataScraper) NonPeriodBasedBoxScoreDataAvailable(gameStatus int32) bool {
	// Game is Final
	if gameStatus == int32(3) {
		return true
	}
	return false
}

func (beds BaseEventDataScraper) LiveBoxScoreDataAvailable(gameStatus int32) bool {
	// Game is ongoing
	if gameStatus == int32(2) {
		return true
	}
	return false
}
