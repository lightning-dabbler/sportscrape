package foxsports

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/lightning-dabbler/sportscrape/util/request"
	"github.com/xitongsys/parquet-go/types"
)

// MatchupScraperOption defines a configuration option for the scraper
type MatchupScraperOption func(*MatchupScraper)

// MatchupScraperLeague sets the League option
func MatchupScraperLeague(league League) MatchupScraperOption {
	return func(s *MatchupScraper) {
		s.League = league
	}
}

// MatchupScraperParams sets the Params option
func MatchupScraperParams(params map[string]string) MatchupScraperOption {
	return func(s *MatchupScraper) {
		s.Params = params
	}
}

// MatchupScraperSegmenter sets the Segmenter option
func MatchupScraperSegmenter(segmenter Segmenter) MatchupScraperOption {
	return func(s *MatchupScraper) {
		s.Segmenter = segmenter
	}
}

// NewMatchupScraper creates a new MatchupScraper with the provided options
func NewMatchupScraper(options ...MatchupScraperOption) *MatchupScraper {
	s := &MatchupScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}
	s.Init()

	return s
}

type MatchupScraper struct {
	// League - The league of interest to fetch matchups data
	League League
	// Params - URL Query parameters
	Params map[string]string
	// Segmenter - The interface for constructing segment IDs
	Segmenter Segmenter
	// segmentID - The base subdirectory in url used to fetch the point-in-time dataset
	segmentID string
	// pullTimestamp - approximate timestamp for when the request to fetch matchups was made
	pullTimestamp time.Time
}

func (s *MatchupScraper) Init() {
	// Ensure Segmenter is set
	if s.Segmenter == nil {
		log.Fatalln("Segmenter is a required argument for foxsports MatchupScraper")
	}
	// Ensure League is set
	if s.League.Undefined() {
		log.Fatalln("League is a required argument for foxsports MatchupScraper")
	}
	// Params
	if s.Params == nil {
		s.Params = map[string]string{}
	}
	s.League.SetParams(s.Params)
}
func (s MatchupScraper) Provider() sportscrape.Provider {
	return sportscrape.FS
}

func (s *MatchupScraper) Feed() sportscrape.Feed {
	switch s.League {
	case NBA:
		return sportscrape.FSNBAMatchup
	case MLB:
		return sportscrape.FSMLBMatchup
	case NFL:
		return sportscrape.FSNFLMatchup
	case NCAAB:
		return sportscrape.FSNCAAMatchup
	}
	return sportscrape.FSNBAMatchup
}

// ConstructFullURL constructs the full url (query params included) to retrieve matchup data
func (s *MatchupScraper) ConstructFullURL() (string, error) {
	if s.segmentID == "" {
		segmentID, err := s.Segmenter.GetSegmentID()
		if err != nil {
			return "", err
		}
		s.segmentID = segmentID
	}
	url, err := s.League.V1MatchupURL(s.segmentID)
	if err != nil {
		return "", err
	}
	queryValues := url.Query()
	for k, v := range s.Params {
		queryValues.Add(k, v)
	}
	url.RawQuery = queryValues.Encode()
	return url.String(), nil
}

// FetchMatchups fetches matchups data
//
// Paramater:
//   - url: The URL to fetch matchups from
//
// Returns the JSON struct and optional error
func (s *MatchupScraper) FetchMatchups(url string) (jsonresponse.Matchup, error) {
	var responsePayload jsonresponse.Matchup
	response, err := request.Get(url)
	if err != nil {
		return responsePayload, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return responsePayload, err
	}
	err = json.Unmarshal(body, &responsePayload)
	if err != nil {
		return responsePayload, err
	}
	return responsePayload, nil
}

// ParseMatchup parses a matchup event JSON object into a matchup data model
func (s *MatchupScraper) ParseMatchup(eventPayload jsonresponse.Event) (model.Matchup, error) {
	var matchup model.Matchup
	//
	eventTime, err := util.RFC3339ToTime(eventPayload.EventTime)
	if err != nil {
		return matchup, err
	}
	matchup.EventTimeParquet = types.TimeToTIMESTAMP_MILLIS(eventTime, true)
	matchup.EventTime = eventTime
	matchup.PullTimestampParquet = types.TimeToTIMESTAMP_MILLIS(s.pullTimestamp, true)
	matchup.PullTimestamp = s.pullTimestamp
	event_id, err := util.TextToInt64(eventPayload.EntityLink.Layout.Tokens.ID)
	if err != nil {
		return matchup, err
	}
	matchup.EventID = event_id
	matchup.EventStatus = eventPayload.EventStatus
	matchup.StatusLine = eventPayload.StatusLine
	// Home team
	homeURISplit := strings.Split(eventPayload.HomeTeam.URI, "/")
	matchup.HomeTeamID, err = util.TextToInt64(homeURISplit[len(homeURISplit)-1])
	if err != nil {
		return matchup, err
	}
	matchup.HomeTeamAbbreviation = eventPayload.HomeTeam.NameAbbreviation
	matchup.HomeTeamNameLong = eventPayload.HomeTeam.LongName
	matchup.HomeTeamNameFull = strings.TrimSpace(eventPayload.HomeTeam.FullNamePt1 + " " + eventPayload.HomeTeam.FullNamePt2)
	matchup.HomeRecord = eventPayload.HomeTeam.Record
	matchup.HomeScore = eventPayload.HomeTeam.Score
	if eventPayload.HomeTeam.Rank != nil {
		rank, err := util.TextToInt(*eventPayload.HomeTeam.Rank)
		if err != nil {
			return matchup, err
		}
		rankInt32 := int32(rank)
		matchup.HomeRank = &rankInt32
	}
	// Away team
	awayURISplit := strings.Split(eventPayload.AwayTeam.URI, "/")
	matchup.AwayTeamID, err = util.TextToInt64(awayURISplit[len(awayURISplit)-1])
	if err != nil {
		return matchup, err
	}
	matchup.AwayTeamAbbreviation = eventPayload.AwayTeam.NameAbbreviation
	matchup.AwayTeamNameLong = eventPayload.AwayTeam.LongName
	matchup.AwayTeamNameFull = strings.TrimSpace(eventPayload.AwayTeam.FullNamePt1 + " " + eventPayload.AwayTeam.FullNamePt2)
	matchup.AwayRecord = eventPayload.AwayTeam.Record
	matchup.AwayScore = eventPayload.AwayTeam.Score
	if eventPayload.AwayTeam.Rank != nil {
		rank, err := util.TextToInt(*eventPayload.AwayTeam.Rank)
		if err != nil {
			return matchup, err
		}
		rankInt32 := int32(rank)
		matchup.AwayRank = &rankInt32
	}

	// Loser
	if eventPayload.AwayTeam.IsLoser {
		matchup.Loser = &matchup.AwayTeamID

	} else if eventPayload.HomeTeam.IsLoser {
		matchup.Loser = &matchup.HomeTeamID
	}

	// IsPlayoff
	if eventPayload.IsPlayoff != nil && *eventPayload.IsPlayoff {
		matchup.IsPlayoff = *eventPayload.IsPlayoff
	}

	return matchup, nil
}

// Scrape gets all matchups of a League and segment ID
func (s *MatchupScraper) Scrape() sportscrape.MatchupOutput {
	var matchups []interface{}
	output := sportscrape.MatchupOutput{}
	// Construct full url
	log.Println("Constructing full URL")
	url, err := s.ConstructFullURL()
	if err != nil {
		log.Println("Issue constructing full URL")
		output.Error = err
		return output
	}

	// Fetch matchups data
	s.pullTimestamp = time.Now().UTC()
	log.Printf("Fetching %s Matchups at segment %s: %s\n", s.League.String(), s.segmentID, url)
	jsonPayload, err := s.FetchMatchups(url)
	if err != nil {
		log.Println("Issue fetching matchups")
		output.Error = err
		return output
	}
	events := jsonPayload.SectionList[0].Events
	nEvents := len(events)
	log.Printf("%d event(s) fetched\n", nEvents)

	// Parse response payload
	var errors, skipped int

	log.Println("Parsing Matchups")
	for idx, event := range events {
		if event.IsTba {
			log.Printf("SKIPPING event #%d as it's still TBA\n", idx)
			skipped += 1
			continue
		}
		parsedEvent, err := s.ParseMatchup(event)
		if err != nil {
			log.Println(fmt.Errorf("WARNING: Error identified at event #%d - %w", idx, err))
			errors += 1
			continue
		}
		matchups = append(matchups, parsedEvent)
	}
	output.Output = matchups
	output.Context = sportscrape.MatchupContext{
		Errors: errors,
		Skips:  skipped,
		Total:  len(matchups),
	}
	return output
}
