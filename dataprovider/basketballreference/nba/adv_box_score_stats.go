package nba

import (
	"fmt"
	"log"
	"time"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference"
	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference/nba/model"
	"github.com/lightning-dabbler/sportscrape/util"
	sportsreferenceutil "github.com/lightning-dabbler/sportscrape/util/sportsreference"
	"github.com/xitongsys/parquet-go/types"
)

const (
	// Selector for Advanced Box Score stats tables
	advBoxScoreSelector = `table[id$='game-advanced']`
)

// advBoxScoreStarterHeaders represents the headers in sequential order for the starter team members
var advBoxScoreStarterHeaders sportsreferenceutil.Headers = sportsreferenceutil.Headers{
	"Starters",
	"MP",
	"TS%",
	"eFG%",
	"3PAr",
	"FTr",
	"ORB%",
	"DRB%",
	"TRB%",
	"AST%",
	"STL%",
	"BLK%",
	"TOV%",
	"USG%",
	"ORtg",
	"DRtg",
	"BPM",
}

// advBoxScoreReservesHeaders represents the headers in sequential order for the reserve team members
var advBoxScoreReservesHeaders sportsreferenceutil.Headers = sportsreferenceutil.Headers{
	"Reserves",
	"MP",
	"TS%",
	"eFG%",
	"3PAr",
	"FTr",
	"ORB%",
	"DRB%",
	"TRB%",
	"AST%",
	"STL%",
	"BLK%",
	"TOV%",
	"USG%",
	"ORtg",
	"DRtg",
	"BPM",
}

// AdvBoxScoreOption defines a configuration option for advanced box score runners
type AdvBoxScoreOption func(*AdvBoxScoreRunner)

// WithAdvBoxScoreTimeout sets the timeout duration for advanced box score runner
func WithAdvBoxScoreTimeout(timeout time.Duration) AdvBoxScoreOption {
	return func(absr *AdvBoxScoreRunner) {
		absr.Timeout = timeout
	}
}

// WithAdvBoxScoreDebug enables or disables debug mode for advanced box score runner
func WithAdvBoxScoreDebug(debug bool) AdvBoxScoreOption {
	return func(absr *AdvBoxScoreRunner) {
		absr.Debug = debug
	}
}

// WithAdvBoxScoreConcurrency sets the number of concurrent workers
func WithAdvBoxScoreConcurrency(n int) AdvBoxScoreOption {
	return func(absr *AdvBoxScoreRunner) {
		absr.Concurrency = n
	}
}

// NewAdvBoxScoreRunner creates a new AdvBoxScoreRunner with the provided options
func NewAdvBoxScoreRunner(options ...AdvBoxScoreOption) *AdvBoxScoreRunner {
	absr := &AdvBoxScoreRunner{}
	absr.Processor = absr

	// Apply all options
	for _, option := range options {
		option(absr)
	}

	return absr
}

// AdvBoxScoreRunner specialized Runner for retrieving NBA advanced box score statistics
// with support for concurrent processing.
type AdvBoxScoreRunner struct {
	sportsreferenceutil.BoxScoreRunner
}

// GetSegmentBoxScoreStats retrieves NBA advanced box score statistics for a single matchup.
//
// Parameter:
//   - matchup: The NBA matchup for which to retrieve advanced box score statistics
//
// Returns a slice of NBA advanced box score statistics as interface{} values
func (boxScoreRunner *AdvBoxScoreRunner) GetSegmentBoxScoreStats(matchup interface{}) []interface{} {
	matchupModel := matchup.(model.NBAMatchup)
	url := matchupModel.BoxScoreLink
	PullTimestamp := time.Now().UTC()
	start := time.Now().UTC()
	var advNBABoxScoreStats []interface{}
	log.Println("Scraping Advanced Box Score: " + url)
	doc, err := boxScoreRunner.RetrieveDocument(url, networkHeaders, waitReadyBoxScoreContentSelector)
	if err != nil {
		log.Fatalln(err)
	}
	doc.Find(advBoxScoreSelector).Each(func(i int, s *goquery.Selection) {
		var starterHeader string
		var reserveHeader string
		s.Find(boxScoreStarterHeaders).Each(func(idx int, s *goquery.Selection) {
			starterHeader = util.CleanTextDatum(s.Text())
			expectedHeader := advBoxScoreStarterHeaders[idx]
			if starterHeader != expectedHeader {
				log.Fatalf("Starter header '%s' at position %d does not equal expected header '%s' @ %s\n", starterHeader, idx, expectedHeader, url)
			}
		})

		s.Find(boxScoreStatsRecordsSelector).Each(func(j int, s *goquery.Selection) {
			var boxScoreStats model.NBAAdvBoxScoreStats
			if j < 5 || j > 5 {
				boxScoreStats.PullTimestamp = PullTimestamp
				boxScoreStats.PullTimestampParquet = types.TimeToTIMESTAMP_MILLIS(PullTimestamp, true)
				boxScoreStats.EventID = matchupModel.EventID
				if i == 0 {
					boxScoreStats.Team = matchupModel.AwayTeam
					boxScoreStats.Opponent = matchupModel.HomeTeam
				} else {
					boxScoreStats.Team = matchupModel.HomeTeam
					boxScoreStats.Opponent = matchupModel.AwayTeam
				}
				boxScoreStats.EventDate = matchupModel.EventDate
				boxScoreStats.EventDateParquet = util.TimeToDays(matchupModel.EventDate)
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
					// TrueShootingPercentage
					trueShootingPercentageText := util.CleanTextDatum(s.Find("td:nth-child(3)").Text())
					trueShootingPercentage, err := util.TextToFloat32(trueShootingPercentageText)
					if trueShootingPercentageText == "" {
						trueShootingPercentage = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for trueShootingPercentageText to float32 - %w; defaulting to 0.", trueShootingPercentageText, err))
						trueShootingPercentage = 0
					}
					boxScoreStats.TrueShootingPercentage = trueShootingPercentage

					// EffectiveFieldGoalPercentage
					effectiveFieldGoalPercentageText := util.CleanTextDatum(s.Find("td:nth-child(4)").Text())
					effectiveFieldGoalPercentage, err := util.TextToFloat32(effectiveFieldGoalPercentageText)
					if effectiveFieldGoalPercentageText == "" {
						effectiveFieldGoalPercentage = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for effectiveFieldGoalPercentageText to float32 - %w; defaulting to 0.", effectiveFieldGoalPercentageText, err))
						effectiveFieldGoalPercentage = 0
					}
					boxScoreStats.EffectiveFieldGoalPercentage = effectiveFieldGoalPercentage

					// ThreePointAttemptRate
					threePointAttemptRateText := util.CleanTextDatum(s.Find("td:nth-child(5)").Text())
					threePointAttemptRate, err := util.TextToFloat32(threePointAttemptRateText)
					if threePointAttemptRateText == "" {
						threePointAttemptRate = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for threePointAttemptRateText to float32 - %w; defaulting to 0.", threePointAttemptRateText, err))
						threePointAttemptRate = 0
					}
					boxScoreStats.ThreePointAttemptRate = threePointAttemptRate

					// FreeThrowAttemptRate
					freeThrowAttemptRateText := util.CleanTextDatum(s.Find("td:nth-child(6)").Text())
					freeThrowAttemptRate, err := util.TextToFloat32(freeThrowAttemptRateText)
					if freeThrowAttemptRateText == "" {
						freeThrowAttemptRate = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for freeThrowAttemptRateText to float32 - %w; defaulting to 0.", freeThrowAttemptRateText, err))
						freeThrowAttemptRate = 0
					}
					boxScoreStats.FreeThrowAttemptRate = freeThrowAttemptRate

					// OffensiveReboundPercentage
					offensiveReboundPercentageText := util.CleanTextDatum(s.Find("td:nth-child(7)").Text())
					offensiveReboundPercentage, err := util.TextToFloat32(offensiveReboundPercentageText)
					if offensiveReboundPercentageText == "" {
						offensiveReboundPercentage = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for offensiveReboundPercentageText to float32 - %w; defaulting to 0.", offensiveReboundPercentageText, err))
						offensiveReboundPercentage = 0
					}
					boxScoreStats.OffensiveReboundPercentage = offensiveReboundPercentage

					// DefensiveReboundPercentage
					defensiveReboundPercentageText := util.CleanTextDatum(s.Find("td:nth-child(8)").Text())
					defensiveReboundPercentage, err := util.TextToFloat32(defensiveReboundPercentageText)
					if defensiveReboundPercentageText == "" {
						defensiveReboundPercentage = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for defensiveReboundPercentageText to float32 - %w; defaulting to 0.", defensiveReboundPercentageText, err))
						defensiveReboundPercentage = 0
					}
					boxScoreStats.DefensiveReboundPercentage = defensiveReboundPercentage

					// TotalReboundPercentage
					totalReboundPercentageText := util.CleanTextDatum(s.Find("td:nth-child(9)").Text())
					totalReboundPercentage, err := util.TextToFloat32(totalReboundPercentageText)
					if totalReboundPercentageText == "" {
						totalReboundPercentage = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for totalReboundPercentageText to float32 - %w; defaulting to 0.", totalReboundPercentageText, err))
						totalReboundPercentage = 0
					}
					boxScoreStats.TotalReboundPercentage = totalReboundPercentage

					// AssistPercentage
					assistPercentageText := util.CleanTextDatum(s.Find("td:nth-child(10)").Text())
					assistPercentage, err := util.TextToFloat32(assistPercentageText)
					if assistPercentageText == "" {
						assistPercentage = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for assistPercentageText to float32 - %w; defaulting to 0.", assistPercentageText, err))
						assistPercentage = 0
					}
					boxScoreStats.AssistPercentage = assistPercentage

					// StealPercentage
					stealPercentageText := util.CleanTextDatum(s.Find("td:nth-child(11)").Text())
					stealPercentage, err := util.TextToFloat32(stealPercentageText)
					if stealPercentageText == "" {
						stealPercentage = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for stealPercentageText to float32 - %w; defaulting to 0.", stealPercentageText, err))
						stealPercentage = 0
					}
					boxScoreStats.StealPercentage = stealPercentage

					// BlockPercentage
					blockPercentageText := util.CleanTextDatum(s.Find("td:nth-child(12)").Text())
					blockPercentage, err := util.TextToFloat32(blockPercentageText)
					if blockPercentageText == "" {
						blockPercentage = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for blockPercentageText to float32 - %w; defaulting to 0.", blockPercentageText, err))
						blockPercentage = 0
					}
					boxScoreStats.BlockPercentage = blockPercentage

					// TurnoverPercentage
					turnoverPercentageText := util.CleanTextDatum(s.Find("td:nth-child(13)").Text())
					turnoverPercentage, err := util.TextToFloat32(turnoverPercentageText)
					if turnoverPercentageText == "" {
						turnoverPercentage = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for turnoverPercentageText to float32 - %w; defaulting to 0.", turnoverPercentageText, err))
						turnoverPercentage = 0
					}
					boxScoreStats.TurnoverPercentage = turnoverPercentage

					// UsagePercentage
					usagePercentageText := util.CleanTextDatum(s.Find("td:nth-child(14)").Text())
					usagePercentage, err := util.TextToFloat32(usagePercentageText)
					if usagePercentageText == "" {
						usagePercentage = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for usagePercentageText to float32 - %w; defaulting to 0.", usagePercentageText, err))
						usagePercentage = 0
					}
					boxScoreStats.UsagePercentage = usagePercentage

					// OffensiveRating
					offensiveRatingText := util.CleanTextDatum(s.Find("td:nth-child(15)").Text())
					offensiveRating, err := util.TextToInt32(offensiveRatingText)
					if offensiveRatingText == "" {
						offensiveRating = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for offensiveRatingText to float32 - %w; defaulting to 0.", offensiveRatingText, err))
						offensiveRating = 0
					}
					boxScoreStats.OffensiveRating = offensiveRating

					// DefensiveRating
					defensiveRatingText := util.CleanTextDatum(s.Find("td:nth-child(16)").Text())
					defensiveRating, err := util.TextToInt32(defensiveRatingText)
					if defensiveRatingText == "" {
						defensiveRating = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for defensiveRatingText to float32 - %w; defaulting to 0.", defensiveRatingText, err))
						defensiveRating = 0
					}
					boxScoreStats.DefensiveRating = defensiveRating

					// BoxPlusMinus
					boxPlusMinusText := util.CleanTextDatum(s.Find("td:nth-child(17)").Text())
					boxPlusMinus, err := util.TextToFloat32(boxPlusMinusText)
					if boxPlusMinusText == "" {
						boxPlusMinus = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for boxPlusMinusText to float32 - %w; defaulting to 0.", boxPlusMinusText, err))
						boxPlusMinus = 0
					}
					boxScoreStats.BoxPlusMinus = boxPlusMinus

				} else {
					boxScoreStats.MinutesPlayed = 0
				}

				advNBABoxScoreStats = append(advNBABoxScoreStats, boxScoreStats)
			} else {
				s.Find(boxScoreReserveHeaders).Each(func(idx int, s *goquery.Selection) {

					reserveHeader = util.CleanTextDatum(s.Text())
					expectedHeader := advBoxScoreReservesHeaders[idx]
					if reserveHeader != expectedHeader {
						log.Fatalf("Reserve header '%s' at position %d does not equal expected header '%s' @ %s\n", reserveHeader, idx, expectedHeader, url)
					}
				})
			}
		})
	})
	if len(advNBABoxScoreStats) == 0 {
		log.Printf("No Data Scraped @ %s\n", url)
	} else {
		diff := time.Now().UTC().Sub(start)
		log.Printf("Scraping of %s Completed in %s\n", url, diff)
	}

	return advNBABoxScoreStats
}
