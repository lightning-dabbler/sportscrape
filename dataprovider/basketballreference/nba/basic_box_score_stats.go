package nba

import (
	"fmt"
	"log"
	"time"
	"unicode"

	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference"
	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference/nba/model"
	"github.com/lightning-dabbler/sportscrape/util"
	sportsreferenceutil "github.com/lightning-dabbler/sportscrape/util/sportsreference"

	"github.com/PuerkitoBio/goquery"
)

const (
	// Selector for Basic Box Score stats tables
	basicBoxScoreSelector = `table[id$='game-basic']`
)

// basicBoxScoreStarterHeaders represents the headers in sequential order for the starter team members
var basicBoxScoreStarterHeaders sportsreferenceutil.Headers = sportsreferenceutil.Headers{
	"Starters",
	"MP",
	"FG",
	"FGA",
	"FG%",
	"3P",
	"3PA",
	"3P%",
	"FT",
	"FTA",
	"FT%",
	"ORB",
	"DRB",
	"TRB",
	"AST",
	"STL",
	"BLK",
	"TOV",
	"PF",
	"PTS",
	"GmSc",
	"+/-",
}

// basicBoxScoreReservesHeaders represents the headers in sequential order for the reserve team members
var basicBoxScoreReservesHeaders sportsreferenceutil.Headers = sportsreferenceutil.Headers{
	"Reserves",
	"MP",
	"FG",
	"FGA",
	"FG%",
	"3P",
	"3PA",
	"3P%",
	"FT",
	"FTA",
	"FT%",
	"ORB",
	"DRB",
	"TRB",
	"AST",
	"STL",
	"BLK",
	"TOV",
	"PF",
	"PTS",
	"GmSc",
	"+/-",
}

// BasicBoxScoreOption defines a configuration option for basic box score runners
type BasicBoxScoreOption func(*BasicBoxScoreRunner)

// WithBasicBoxScoreTimeout sets the timeout duration for basic box score runner
func WithBasicBoxScoreTimeout(timeout time.Duration) BasicBoxScoreOption {
	return func(bsr *BasicBoxScoreRunner) {
		bsr.Timeout = timeout
	}
}

// WithBasicBoxScoreDebug enables or disables debug mode for basic box score runner
func WithBasicBoxScoreDebug(debug bool) BasicBoxScoreOption {
	return func(bsr *BasicBoxScoreRunner) {
		bsr.Debug = debug
	}
}

// WithBasicBoxScoreConcurrency sets the number of concurrent workers
func WithBasicBoxScoreConcurrency(n int) BasicBoxScoreOption {
	return func(bsr *BasicBoxScoreRunner) {
		bsr.Concurrency = n
	}
}

// NewBasicBoxScoreRunner creates a new BasicBoxScoreRunner with the provided options
func NewBasicBoxScoreRunner(options ...BasicBoxScoreOption) *BasicBoxScoreRunner {
	bsr := &BasicBoxScoreRunner{}
	bsr.Processor = bsr

	// Apply all options
	for _, option := range options {
		option(bsr)
	}

	return bsr
}

// BasicBoxScoreRunner specialized Runner for retrieving NBA basic box score statistics
// with support for concurrent processing.
type BasicBoxScoreRunner struct {
	sportsreferenceutil.BoxScoreRunner
}

// GetSegmentBoxScoreStats retrieves NBA basic box score statistics for a single matchup.
//
// Parameter:
//   - matchup: The NBA matchup for which to retrieve basic box score statistics
//
// Returns a slice of NBA basic box score statistics as interface{} values
func (boxScoreRunner *BasicBoxScoreRunner) GetSegmentBoxScoreStats(matchup interface{}) []interface{} {
	matchupModel := matchup.(model.NBAMatchup)
	url := matchupModel.BoxScoreLink
	PullTimestamp := time.Now().UTC()
	start := time.Now().UTC()
	var basicNBABoxScoreStats []interface{}
	log.Println("Scraping Basic Box Score: " + url)
	doc, err := boxScoreRunner.RetrieveDocument(url, networkHeaders, waitReadyBoxScoreContentSelector)
	if err != nil {
		log.Fatalln(err)
	}
	doc.Find(basicBoxScoreSelector).Each(func(i int, s *goquery.Selection) {
		var starterHeader string
		var reserveHeader string
		s.Find(boxScoreStarterHeaders).Each(func(idx int, s *goquery.Selection) {
			starterHeader = util.CleanTextDatum(s.Text())
			expectedHeader := basicBoxScoreStarterHeaders[idx]
			if starterHeader != expectedHeader {
				log.Fatalf("Starter header '%s' at position %d does not equal expected header '%s' @ %s\n", starterHeader, idx, expectedHeader, url)
			}
		})

		s.Find(boxScoreStatsRecordsSelector).Each(func(j int, s *goquery.Selection) {
			var boxScoreStats model.NBABasicBoxScoreStats
			if j < 5 || j > 5 {
				boxScoreStats.PullTimestamp = PullTimestamp
				boxScoreStats.EventID = matchupModel.EventID
				if i == 0 {
					boxScoreStats.Team = matchupModel.AwayTeam
					boxScoreStats.Opponent = matchupModel.HomeTeam
				} else {
					boxScoreStats.Team = matchupModel.HomeTeam
					boxScoreStats.Opponent = matchupModel.AwayTeam
				}
				boxScoreStats.EventDate = matchupModel.EventDate
				if j < 5 {
					boxScoreStats.Starter = true
				} else {
					boxScoreStats.Starter = false
				}
				boxScoreStats.PlayerLink = basketballreference.URL + util.CleanTextDatum(s.Find(boxScorePlayerLinkSelector).AttrOr("href", ""))
				boxScoreStats.Player = util.CleanTextDatum(s.Find(boxScorePlayerSelector).Text())
				playerID, err := sportsreferenceutil.PlayerID(boxScoreStats.PlayerLink)
				if err != nil {
					log.Fatalln(err)
				}
				boxScoreStats.PlayerID = playerID
				minutesPlayed := util.CleanTextDatum(s.Find("td:nth-child(2)").Text())
				if len(minutesPlayed) > 0 && unicode.IsDigit(rune(minutesPlayed[0])) {
					totalMinutes, err := transformMinutesPlayed(minutesPlayed)
					if err != nil {
						log.Fatalln(err)
					}
					boxScoreStats.MinutesPlayed = totalMinutes

					fieldGoalsMadeText := util.CleanTextDatum(s.Find("td:nth-child(3)").Text())
					fieldGoalsMade, err := util.TextToInt(fieldGoalsMadeText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for fieldGoalsMadeText to Int - %w; defaulting to 0.", fieldGoalsMadeText, err))
						fieldGoalsMade = 0
					}
					fieldGoalAttemptsText := util.CleanTextDatum(s.Find("td:nth-child(4)").Text())
					fieldGoalAttempts, err := util.TextToInt(fieldGoalAttemptsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for fieldGoalAttemptsText to Int - %w; defaulting to 0.", fieldGoalAttemptsText, err))
						fieldGoalAttempts = 0
					}
					rawFieldGoalPercentage := util.CleanTextDatum(s.Find("td:nth-child(5)").Text())
					var fieldGoalPercentage float32

					if rawFieldGoalPercentage == "" {
						fieldGoalPercentage = 0
					} else {
						fieldGoalPercentage, err = util.TextToFloat32(rawFieldGoalPercentage)
						if err != nil {
							log.Fatalln(fmt.Errorf("Can't convert '%s' for rawFieldGoalPercentage to Float64 - %w", rawFieldGoalPercentage, err))
						}
					}

					threePointsMadeText := util.CleanTextDatum(s.Find("td:nth-child(6)").Text())
					threePointsMade, err := util.TextToInt(threePointsMadeText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for threePointsMadeText to Int - %w; defaulting to 0.", threePointsMadeText, err))
						threePointsMade = 0
					}
					threePointAttemptsText := util.CleanTextDatum(s.Find("td:nth-child(7)").Text())
					threePointAttempts, err := util.TextToInt(threePointAttemptsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for threePointAttemptsText to Int - %w; defaulting to 0.", threePointAttemptsText, err))
						threePointsMade = 0
					}
					rawthreePointPercentage := util.CleanTextDatum(s.Find("td:nth-child(8)").Text())
					var threePointPercentage float32

					if rawthreePointPercentage == "" {
						threePointPercentage = 0
					} else {
						threePointPercentage, err = util.TextToFloat32(rawthreePointPercentage)
						if err != nil {
							log.Fatalln(fmt.Errorf("Can't convert '%s' for rawthreePointPercentage to Float64 - %w", rawthreePointPercentage, err))
						}
					}

					freeThrowsMadeText := util.CleanTextDatum(s.Find("td:nth-child(9)").Text())
					freeThrowsMade, err := util.TextToInt(freeThrowsMadeText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for freeThrowsMadeText to Int - %w; defaulting to 0.", freeThrowsMadeText, err))
						freeThrowsMade = 0
					}
					freeThrowAttemptsText := util.CleanTextDatum(s.Find("td:nth-child(10)").Text())
					freeThrowAttempts, err := util.TextToInt(freeThrowAttemptsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for freeThrowAttemptsText to Int - %w; defaulting to 0.", freeThrowAttemptsText, err))
						freeThrowAttempts = 0
					}
					rawfreeThrowPercentage := util.CleanTextDatum(s.Find("td:nth-child(11)").Text())
					var freeThrowPercentage float32

					if rawfreeThrowPercentage == "" {
						freeThrowPercentage = 0
					} else {
						freeThrowPercentage, err = util.TextToFloat32(rawfreeThrowPercentage)
						if err != nil {
							log.Fatalln(fmt.Errorf("Error: Can't convert '%s' for rawfreeThrowPercentage to Float64 - %w", rawfreeThrowPercentage, err))
						}
					}
					boxScoreStats.FieldGoalsMade = fieldGoalsMade
					boxScoreStats.FieldGoalAttempts = fieldGoalAttempts
					boxScoreStats.FieldGoalPercentage = fieldGoalPercentage
					boxScoreStats.ThreePointsMade = threePointsMade
					boxScoreStats.ThreePointAttempts = threePointAttempts
					boxScoreStats.ThreePointPercentage = threePointPercentage
					boxScoreStats.FreeThrowsMade = freeThrowsMade
					boxScoreStats.FreeThrowAttempts = freeThrowAttempts
					boxScoreStats.FreeThrowPercentage = freeThrowPercentage

					OffensiveReboundsText := util.CleanTextDatum(s.Find("td:nth-child(12)").Text())
					boxScoreStats.OffensiveRebounds, err = util.TextToInt(OffensiveReboundsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for OffensiveReboundsText to Int - %w; defaulting to 0.", OffensiveReboundsText, err))
						boxScoreStats.OffensiveRebounds = 0
					}

					DefensiveReboundsText := util.CleanTextDatum(s.Find("td:nth-child(13)").Text())
					boxScoreStats.DefensiveRebounds, err = util.TextToInt(DefensiveReboundsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for DefensiveReboundsText to Int - %w; defaulting to 0.", DefensiveReboundsText, err))
						boxScoreStats.DefensiveRebounds = 0
					}

					TotalReboundsText := util.CleanTextDatum(s.Find("td:nth-child(14)").Text())
					boxScoreStats.TotalRebounds, err = util.TextToInt(TotalReboundsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for TotalReboundsText to Int - %w; defaulting to 0.", TotalReboundsText, err))
						boxScoreStats.TotalRebounds = 0
					}

					AssistsText := util.CleanTextDatum(s.Find("td:nth-child(15)").Text())
					boxScoreStats.Assists, err = util.TextToInt(AssistsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for AssistsText to Int - %w; defaulting to 0.", AssistsText, err))
						boxScoreStats.Assists = 0
					}

					StealsText := util.CleanTextDatum(s.Find("td:nth-child(16)").Text())
					boxScoreStats.Steals, err = util.TextToInt(StealsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for StealsText to Int - %w; defaulting to 0.", StealsText, err))
						boxScoreStats.Steals = 0
					}

					BlocksText := util.CleanTextDatum(s.Find("td:nth-child(17)").Text())
					boxScoreStats.Blocks, err = util.TextToInt(BlocksText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for BlocksText to Int - %w; defaulting to 0.", BlocksText, err))
						boxScoreStats.Blocks = 0
					}

					TurnoversText := util.CleanTextDatum(s.Find("td:nth-child(18)").Text())
					boxScoreStats.Turnovers, err = util.TextToInt(TurnoversText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for TurnoversText to In - %w; defaulting to 0.", TurnoversText, err))
						boxScoreStats.Turnovers = 0
					}

					PersonalFoulsText := util.CleanTextDatum(s.Find("td:nth-child(19)").Text())
					boxScoreStats.PersonalFouls, err = util.TextToInt(PersonalFoulsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for PersonalFoulsText to Int - %w; defaulting to 0.", PersonalFoulsText, err))
						boxScoreStats.PersonalFouls = 0
					}

					PointsText := util.CleanTextDatum(s.Find("td:nth-child(20)").Text())
					boxScoreStats.Points, err = util.TextToInt(PointsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for PointsText to Int - %w; defaulting to 0.", PointsText, err))
						boxScoreStats.Points = 0
					}

					GameScoreText := util.CleanTextDatum(s.Find("td:nth-child(21)").Text())
					boxScoreStats.GameScore, err = util.TextToFloat32(GameScoreText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for GameScoreText to Float64 - %w; defaulting to 0.", GameScoreText, err))
						boxScoreStats.GameScore = 0
					}

					PlusMinusText := util.CleanTextDatum(s.Find("td:nth-child(22)").Text())
					boxScoreStats.PlusMinus, err = util.TextToInt(PlusMinusText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for PlusMinusText to Int - %w; defaulting to 0.", PlusMinusText, err))
						boxScoreStats.PlusMinus = 0
					}
				} else {
					boxScoreStats.MinutesPlayed = 0
				}
				basicNBABoxScoreStats = append(basicNBABoxScoreStats, boxScoreStats)
			} else {
				s.Find(boxScoreReserveHeaders).Each(func(idx int, s *goquery.Selection) {
					reserveHeader = util.CleanTextDatum(s.Text())
					expectedHeader := basicBoxScoreReservesHeaders[idx]
					if reserveHeader != expectedHeader {
						log.Fatalf("Reserve header '%s' at position %d does not equal expected header '%s' @ %s\n", reserveHeader, idx, expectedHeader, url)
					}
				})
			}

		})
	})
	if len(basicNBABoxScoreStats) == 0 {
		log.Printf("No Data Scraped @ %s\n", url)
	} else {
		diff := time.Now().UTC().Sub(start)
		log.Printf("Scraping of %s Completed in %s\n", url, diff)
	}

	return basicNBABoxScoreStats
}
