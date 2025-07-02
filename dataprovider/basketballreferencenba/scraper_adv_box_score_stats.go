package basketballreferencenba

import (
	"fmt"
	"log"
	"time"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreferencenba/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/lightning-dabbler/sportscrape/util/sportsreference"
	"github.com/xitongsys/parquet-go/types"
)

const (
	// Selector for Advanced Box Score stats tables
	advBoxScoreSelector = `table[id$='game-advanced']`
)

// advBoxScoreStarterHeaders represents the headers in sequential order for the starter team members
var advBoxScoreStarterHeaders sportsreference.Headers = sportsreference.Headers{
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
var advBoxScoreReservesHeaders sportsreference.Headers = sportsreference.Headers{
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
type AdvBoxScoreOption func(*AdvBoxScoreScraper)

// WithAdvBoxScoreTimeout sets the timeout duration for advanced box score runner
func WithAdvBoxScoreTimeout(timeout time.Duration) AdvBoxScoreOption {
	return func(s *AdvBoxScoreScraper) {
		s.Timeout = timeout
	}
}

// WithAdvBoxScoreDebug enables or disables debug mode for advanced box score runner
func WithAdvBoxScoreDebug(debug bool) AdvBoxScoreOption {
	return func(s *AdvBoxScoreScraper) {
		s.Debug = debug
	}
}

// NewAdvBoxScoreScraper creates a new AdvBoxScoreScraper with the provided options
func NewAdvBoxScoreScraper(options ...AdvBoxScoreOption) *AdvBoxScoreScraper {
	s := &AdvBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}
	s.Init()

	return s
}

// AdvBoxScoreScraper specialized Runner for retrieving NBA advanced box score statistics
// with support for concurrent processing.
type AdvBoxScoreScraper struct {
	EventDataScraper
}

func (s *AdvBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.BasketballReferenceNBAAdvBoxScore
}

// Scrape retrieves NBA advanced box score statistics for a single matchup.
func (abs *AdvBoxScoreScraper) Scrape(matchup interface{}) sportscrape.EventDataOutput {
	matchupModel := matchup.(model.NBAMatchup)
	context := abs.ConstructContext(matchupModel)
	output := sportscrape.EventDataOutput{
		Context: context,
	}
	url := matchupModel.BoxScoreLink
	PullTimestamp := time.Now().UTC()
	start := time.Now().UTC()
	var advNBABoxScoreStats []interface{}
	log.Println("Scraping Advanced Box Score: " + url)
	doc, err := abs.RetrieveDocument(url, networkHeaders, contentReadySelector)
	if err != nil {
		output.Error = err
		return output
	}
	doc.Find(advBoxScoreSelector).EachWithBreak(func(i int, s *goquery.Selection) bool {
		var starterHeader string
		var reserveHeader string
		s.Find(boxScoreStarterHeaders).EachWithBreak(func(idx int, s *goquery.Selection) bool {
			starterHeader = util.CleanTextDatum(s.Text())
			expectedHeader := advBoxScoreStarterHeaders[idx]
			if starterHeader != expectedHeader {
				err = fmt.Errorf("starter header '%s' at position %d does not equal expected header '%s' @ %s", starterHeader, idx, expectedHeader, url)
				output.Error = err
				return false
			}
			return true
		})
		if output.Error != nil {
			return false
		}

		s.Find(boxScoreStatsRecordsSelector).EachWithBreak(func(j int, s *goquery.Selection) bool {
			var boxScoreStats model.NBAAdvBoxScoreStats
			if j < 5 || j > 5 {
				boxScoreStats.PullTimestamp = PullTimestamp
				boxScoreStats.PullTimestampParquet = types.TimeToTIMESTAMP_MILLIS(PullTimestamp, true)
				boxScoreStats.EventID = matchupModel.EventID
				if i == 0 {
					boxScoreStats.Team = matchupModel.AwayTeam
					boxScoreStats.TeamID = matchupModel.AwayTeamID
					boxScoreStats.Opponent = matchupModel.HomeTeam
					boxScoreStats.OpponentID = matchupModel.HomeTeamID
				} else {
					boxScoreStats.Team = matchupModel.HomeTeam
					boxScoreStats.TeamID = matchupModel.HomeTeamID
					boxScoreStats.Opponent = matchupModel.AwayTeam
					boxScoreStats.OpponentID = matchupModel.AwayTeamID
				}
				boxScoreStats.EventDate = matchupModel.EventDate
				boxScoreStats.EventDateParquet = util.TimeToDays(matchupModel.EventDate)
				if j < 5 {
					boxScoreStats.Starter = true
				} else {
					boxScoreStats.Starter = false
				}
				boxScoreStats.PlayerLink = sportsreference.BasketballRefURL + util.CleanTextDatum(s.Find(boxScorePlayerLinkSelector).AttrOr("href", ""))
				boxScoreStats.Player = util.CleanTextDatum(s.Find(boxScorePlayerSelector).Text())
				playerID, err := sportsreference.PlayerID(boxScoreStats.PlayerLink)
				if err != nil {
					output.Error = err
					return false
				}
				boxScoreStats.PlayerID = playerID
				minutesPlayed := util.CleanTextDatum(s.Find("td:nth-child(2)").Text())
				if len(minutesPlayed) > 0 && unicode.IsDigit(rune(minutesPlayed[0])) {
					totalMinutes, err := transformMinutesPlayed(minutesPlayed)
					if err != nil {
						output.Error = err
						return false
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
				s.Find(boxScoreReserveHeaders).EachWithBreak(func(idx int, s *goquery.Selection) bool {

					reserveHeader = util.CleanTextDatum(s.Text())
					expectedHeader := advBoxScoreReservesHeaders[idx]
					if reserveHeader != expectedHeader {
						err = fmt.Errorf("reserve header '%s' at position %d does not equal expected header '%s' @ %s", reserveHeader, idx, expectedHeader, url)
						output.Error = err
						return false
					}
					return true
				})
			}
			if output.Error != nil {
				return false
			}
			return true
		})
		return true
	})

	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of %s Completed in %s\n", url, diff)
	output.Output = advNBABoxScoreStats
	return output
}
