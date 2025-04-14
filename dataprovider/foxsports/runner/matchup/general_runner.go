package matchup

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/lightning-dabbler/sportscrape/util/request"
	"github.com/xitongsys/parquet-go/types"
)

// MatchupOption defines a configuration option for the general matchup runner
type GeneralMatchupOption func(*GeneralMatchupRunner)

// GeneralMatchupLeague sets the League option for the general matchup runner
func GeneralMatchupLeague(league foxsports.League) GeneralMatchupOption {
	return func(gmr *GeneralMatchupRunner) {
		gmr.League = league
	}
}

// GeneralMatchupParams sets the Params option for the general matchup runner
func GeneralMatchupParams(params map[string]string) GeneralMatchupOption {
	return func(gmr *GeneralMatchupRunner) {
		gmr.Params = params
	}
}

// GeneralMatchupSegmenter sets the Segmenter option for the general matchup runner
func GeneralMatchupSegmenter(segmenter Segmenter) GeneralMatchupOption {
	return func(gmr *GeneralMatchupRunner) {
		gmr.Segmenter = segmenter
	}
}

// NewGeneralMatchupRunner creates a new GeneralMatchupRunner with the provided options
func NewGeneralMatchupRunner(options ...GeneralMatchupOption) *GeneralMatchupRunner {
	gmr := &GeneralMatchupRunner{}

	// Apply all options
	for _, option := range options {
		option(gmr)
	}
	if gmr.Params == nil {
		gmr.Params = map[string]string{}
	}
	gmr.League.SetParams(gmr.Params)

	return gmr
}

// GeneralMatchupRunner is a general matchup runner for scraping NBA, MLB, NCAAB, etc. matchup data.
type GeneralMatchupRunner struct {
	// League - The league of interest to fetch matchups data
	League foxsports.League
	// Params - URL Query parameters
	Params map[string]string
	// Segmenter - The interface for constructing segment IDs
	Segmenter Segmenter
	// segmentID - The base subdirectory in url used to fetch the point-in-time dataset
	segmentID string
	// pullTimestamp - approximate timestamp for when the request to fetch matchups was made
	pullTimestamp time.Time
}

// Segmenter is the interface for constructing Segment IDs
type Segmenter interface {
	// GetSegmentID returns the ID that is concatenated to a League's URL to fetch the relevant point-in-time dataset.
	GetSegmentID() (string, error)
}

// ConstructFullURL constructs the full url (query params included) to retrieve matchup data
func (mr *GeneralMatchupRunner) ConstructFullURL() (string, error) {
	if mr.segmentID == "" {
		segmentID, err := mr.Segmenter.GetSegmentID()
		if err != nil {
			return "", err
		}
		mr.segmentID = segmentID
	}
	url, err := mr.League.V1MatchupURL(mr.segmentID)
	if err != nil {
		return "", err
	}
	queryValues := url.Query()
	for k, v := range mr.Params {
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
func (mr *GeneralMatchupRunner) FetchMatchups(url string) (jsonresponse.Matchup, error) {
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
func (mr *GeneralMatchupRunner) ParseMatchup(eventPayload jsonresponse.Event) (model.Matchup, error) {
	var matchup model.Matchup
	//
	eventTime, err := util.RFC3339ToTime(eventPayload.EventTime)
	if err != nil {
		return matchup, err
	}
	matchup.EventTimeParquet = types.TimeToTIMESTAMP_MILLIS(eventTime, true)
	matchup.EventTime = eventTime
	matchup.PullTimestampParquet = types.TimeToTIMESTAMP_MILLIS(mr.pullTimestamp, true)
	matchup.PullTimestamp = mr.pullTimestamp
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

// GetMatchups gets all matchups of a League and segment ID
func (mr *GeneralMatchupRunner) GetMatchups() []interface{} {
	start := time.Now().UTC()
	var matchups []interface{}
	// Construct full url
	log.Println("Constructing full URL")
	url, err := mr.ConstructFullURL()
	if err != nil {
		log.Println("Issue constructing full URL")
		log.Fatalln(err)
	}

	// Fetch matchups data
	mr.pullTimestamp = time.Now().UTC()
	log.Printf("Fetching %s Matchups at segment %s: %s\n", mr.League.String(), mr.segmentID, url)
	jsonPayload, err := mr.FetchMatchups(url)
	if err != nil {
		log.Println("Issue fetching matchups")
		log.Fatalln(err)
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
		parsedEvent, err := mr.ParseMatchup(event)
		if err != nil {
			log.Println(fmt.Errorf("WARNING: Error identified at event #%d - %w", idx, err))
			errors += 1
			continue
		}
		matchups = append(matchups, parsedEvent)
	}
	if errors != 0 || skipped != 0 {
		log.Printf("%d/%d events were skipped and %d/%d events errored out\n", skipped, nEvents, errors, nEvents)
	}
	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of %s Matchups at segment %s completed in %s: %s\n", mr.League.String(), mr.segmentID, diff, url)
	return matchups
}
