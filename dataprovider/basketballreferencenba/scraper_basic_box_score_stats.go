package basketballreferencenba

import (
	"fmt"
	"log"
	"time"
	"unicode"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreferencenba/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/lightning-dabbler/sportscrape/util/sportsreference"
	"github.com/xitongsys/parquet-go/types"

	"github.com/PuerkitoBio/goquery"
)

// Period enum for different stat periods of basic NBA basic box score stats
type Period int

const (
	Undefined Period = iota
	Full
	Q1
	Q2
	Q3
	Q4
	H1
	H2
)

func (p Period) String() string {
	switch p {
	case Full:
		return "Full"
	case Q1:
		return "Q1"
	case Q2:
		return "Q2"
	case Q3:
		return "Q3"
	case Q4:
		return "Q4"
	case H1:
		return "H1"
	case H2:
		return "H2"
	default:
		return "Undefined"
	}
}

func (p Period) TableSelector() string {
	baseSelector := `table[id$='%s-basic']`
	switch p {
	case Full:
		return fmt.Sprintf(baseSelector, "game")
	case Q1:
		return fmt.Sprintf(baseSelector, "q1")
	case Q2:
		return fmt.Sprintf(baseSelector, "q2")
	case Q3:
		return fmt.Sprintf(baseSelector, "q3")
	case Q4:
		return fmt.Sprintf(baseSelector, "q4")
	case H1:
		return fmt.Sprintf(baseSelector, "h1")
	case H2:
		return fmt.Sprintf(baseSelector, "h2")
	default:
		return "Undefined"
	}
}

func (p Period) Undefined() bool {
	switch p {
	case Full, Q1, Q2, Q3, Q4, H1, H2:
		return false
	default:
		return true
	}
}

// basicBoxScoreStarterHeaders represents the headers in sequential order for the starter team members
var basicBoxScoreStarterHeaders sportsreference.Headers = sportsreference.Headers{
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
var basicBoxScoreReservesHeaders sportsreference.Headers = sportsreference.Headers{
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
type BasicBoxScoreOption func(*BasicBoxScoreScraper)

// WithBasicBoxScoreTimeout sets the timeout duration for basic box score runner
func WithBasicBoxScoreTimeout(timeout time.Duration) BasicBoxScoreOption {
	return func(s *BasicBoxScoreScraper) {
		s.Timeout = timeout
	}
}

// WithBasicBoxScoreDebug enables or disables debug mode for basic box score runner
func WithBasicBoxScoreDebug(debug bool) BasicBoxScoreOption {
	return func(s *BasicBoxScoreScraper) {
		s.Debug = debug
	}
}

// WithBasicBoxScorePeriod sets the period of data to scrape
func WithBasicBoxScorePeriod(period Period) BasicBoxScoreOption {
	return func(s *BasicBoxScoreScraper) {
		s.Period = period
	}
}

// NewBasicBoxScoreScraper creates a new BasicBoxScoreScraper with the provided options
func NewBasicBoxScoreScraper(options ...BasicBoxScoreOption) *BasicBoxScoreScraper {
	s := &BasicBoxScoreScraper{}
	// default period
	s.Period = Full

	// Apply all options
	for _, option := range options {
		option(s)
	}
	s.Init()

	return s
}

// BasicBoxScoreScraper specialized Runner for retrieving NBA basic box score statistics
// with support for concurrent processing.
type BasicBoxScoreScraper struct {
	EventDataScraper
	Period Period
}

func (s *BasicBoxScoreScraper) Init() {
	if s.Period.Undefined() {
		log.Fatalln("Period is a required argument for basketball reference nba BasicBoxScoreScraper")
	}
}

func (s *BasicBoxScoreScraper) Feed() sportscrape.Feed {
	switch s.Period {
	case Full:
		return sportscrape.BasketballReferenceNBABoxScore
	case Q1:
		return sportscrape.BasketballReferenceNBABoxScoreQ1
	case Q2:
		return sportscrape.BasketballReferenceNBABoxScoreQ2
	case Q3:
		return sportscrape.BasketballReferenceNBABoxScoreQ3
	case Q4:
		return sportscrape.BasketballReferenceNBABoxScoreQ4
	case H1:
		return sportscrape.BasketballReferenceNBABoxScoreH1
	case H2:
		return sportscrape.BasketballReferenceNBABoxScoreH2
	}
	return sportscrape.BasketballReferenceNBABoxScore
}

// Scrape retrieves NBA basic box score statistics for a single matchup.
func (bs *BasicBoxScoreScraper) Scrape(matchup interface{}) sportscrape.EventDataOutput {
	matchupModel := matchup.(model.NBAMatchup)
	context := bs.ConstructContext(matchupModel)
	output := sportscrape.EventDataOutput{
		Context: context,
	}
	url := matchupModel.BoxScoreLink
	PullTimestamp := time.Now().UTC()
	start := time.Now().UTC()
	var basicNBABoxScoreStats []interface{}
	log.Printf("Scraping %s Basic Box Score: %s\n", bs.Period.String(), url)
	doc, err := bs.RetrieveDocument(url, networkHeaders, contentReadySelector)
	if err != nil {
		output.Error = err
		return output
	}
	doc.Find(bs.Period.TableSelector()).EachWithBreak(func(i int, s *goquery.Selection) bool {
		var starterHeader string
		var reserveHeader string
		s.Find(boxScoreStarterHeaders).EachWithBreak(func(idx int, s *goquery.Selection) bool {
			starterHeader = util.CleanTextDatum(s.Text())
			expectedHeader := basicBoxScoreStarterHeaders[idx]
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
			var boxScoreStats model.NBABasicBoxScoreStats
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

					fieldGoalsMadeText := util.CleanTextDatum(s.Find("td:nth-child(3)").Text())
					fieldGoalsMade, err := util.TextToInt32(fieldGoalsMadeText)
					if fieldGoalsMadeText == "" {
						fieldGoalsMade = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for fieldGoalsMadeText to Int - %w; defaulting to 0.", fieldGoalsMadeText, err))
						fieldGoalsMade = 0
					}
					fieldGoalAttemptsText := util.CleanTextDatum(s.Find("td:nth-child(4)").Text())
					fieldGoalAttempts, err := util.TextToInt32(fieldGoalAttemptsText)
					if fieldGoalAttemptsText == "" {
						fieldGoalAttempts = 0
					} else if err != nil {
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
							err = fmt.Errorf("can't convert '%s' for rawFieldGoalPercentage to Float64 - %w", rawFieldGoalPercentage, err)
							output.Error = err
							return false
						}
					}

					threePointsMadeText := util.CleanTextDatum(s.Find("td:nth-child(6)").Text())
					threePointsMade, err := util.TextToInt32(threePointsMadeText)
					if threePointsMadeText == "" {
						threePointsMade = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for threePointsMadeText to Int - %w; defaulting to 0.", threePointsMadeText, err))
						threePointsMade = 0
					}
					threePointAttemptsText := util.CleanTextDatum(s.Find("td:nth-child(7)").Text())
					threePointAttempts, err := util.TextToInt32(threePointAttemptsText)
					if threePointAttemptsText == "" {
						threePointAttempts = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for threePointAttemptsText to Int - %w; defaulting to 0.", threePointAttemptsText, err))
						threePointAttempts = 0
					}
					rawthreePointPercentage := util.CleanTextDatum(s.Find("td:nth-child(8)").Text())
					var threePointPercentage float32

					if rawthreePointPercentage == "" {
						threePointPercentage = 0
					} else {
						threePointPercentage, err = util.TextToFloat32(rawthreePointPercentage)
						if err != nil {
							err = fmt.Errorf("can't convert '%s' for rawthreePointPercentage to Float64 - %w", rawthreePointPercentage, err)
							output.Error = err
							return false
						}
					}

					freeThrowsMadeText := util.CleanTextDatum(s.Find("td:nth-child(9)").Text())
					freeThrowsMade, err := util.TextToInt32(freeThrowsMadeText)
					if freeThrowsMadeText == "" {
						freeThrowsMade = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for freeThrowsMadeText to Int - %w; defaulting to 0.", freeThrowsMadeText, err))
						freeThrowsMade = 0
					}
					freeThrowAttemptsText := util.CleanTextDatum(s.Find("td:nth-child(10)").Text())
					freeThrowAttempts, err := util.TextToInt32(freeThrowAttemptsText)
					if freeThrowAttemptsText == "" {
						freeThrowAttempts = 0
					} else if err != nil {
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
							err = fmt.Errorf("error: can't convert '%s' for rawfreeThrowPercentage to Float64 - %w", rawfreeThrowPercentage, err)
							output.Error = err
							return false
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
					boxScoreStats.OffensiveRebounds, err = util.TextToInt32(OffensiveReboundsText)
					if OffensiveReboundsText == "" {
						boxScoreStats.OffensiveRebounds = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for OffensiveReboundsText to Int - %w; defaulting to 0.", OffensiveReboundsText, err))
						boxScoreStats.OffensiveRebounds = 0
					}

					DefensiveReboundsText := util.CleanTextDatum(s.Find("td:nth-child(13)").Text())
					boxScoreStats.DefensiveRebounds, err = util.TextToInt32(DefensiveReboundsText)
					if DefensiveReboundsText == "" {
						boxScoreStats.DefensiveRebounds = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for DefensiveReboundsText to Int - %w; defaulting to 0.", DefensiveReboundsText, err))
						boxScoreStats.DefensiveRebounds = 0
					}

					TotalReboundsText := util.CleanTextDatum(s.Find("td:nth-child(14)").Text())
					boxScoreStats.TotalRebounds, err = util.TextToInt32(TotalReboundsText)
					if TotalReboundsText == "" {
						boxScoreStats.TotalRebounds = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for TotalReboundsText to Int - %w; defaulting to 0.", TotalReboundsText, err))
						boxScoreStats.TotalRebounds = 0
					}

					AssistsText := util.CleanTextDatum(s.Find("td:nth-child(15)").Text())
					boxScoreStats.Assists, err = util.TextToInt32(AssistsText)
					if AssistsText == "" {
						boxScoreStats.Assists = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for AssistsText to Int - %w; defaulting to 0.", AssistsText, err))
						boxScoreStats.Assists = 0
					}

					StealsText := util.CleanTextDatum(s.Find("td:nth-child(16)").Text())
					boxScoreStats.Steals, err = util.TextToInt32(StealsText)
					if StealsText == "" {
						boxScoreStats.Steals = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for StealsText to Int - %w; defaulting to 0.", StealsText, err))
						boxScoreStats.Steals = 0
					}

					BlocksText := util.CleanTextDatum(s.Find("td:nth-child(17)").Text())
					boxScoreStats.Blocks, err = util.TextToInt32(BlocksText)
					if BlocksText == "" {
						boxScoreStats.Blocks = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for BlocksText to Int - %w; defaulting to 0.", BlocksText, err))
						boxScoreStats.Blocks = 0
					}

					TurnoversText := util.CleanTextDatum(s.Find("td:nth-child(18)").Text())
					boxScoreStats.Turnovers, err = util.TextToInt32(TurnoversText)
					if TurnoversText == "" {
						boxScoreStats.Turnovers = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for TurnoversText to In - %w; defaulting to 0.", TurnoversText, err))
						boxScoreStats.Turnovers = 0
					}

					PersonalFoulsText := util.CleanTextDatum(s.Find("td:nth-child(19)").Text())
					boxScoreStats.PersonalFouls, err = util.TextToInt32(PersonalFoulsText)
					if PersonalFoulsText == "" {
						boxScoreStats.PersonalFouls = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for PersonalFoulsText to Int - %w; defaulting to 0.", PersonalFoulsText, err))
						boxScoreStats.PersonalFouls = 0
					}

					PointsText := util.CleanTextDatum(s.Find("td:nth-child(20)").Text())
					boxScoreStats.Points, err = util.TextToInt32(PointsText)
					if PointsText == "" {
						boxScoreStats.Points = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for PointsText to Int - %w; defaulting to 0.", PointsText, err))
						boxScoreStats.Points = 0
					}

					GameScoreText := util.CleanTextDatum(s.Find("td:nth-child(21)").Text())
					boxScoreStats.GameScore, err = util.TextToFloat32(GameScoreText)
					if GameScoreText == "" {
						boxScoreStats.GameScore = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for GameScoreText to Float64 - %w; defaulting to 0.", GameScoreText, err))
						boxScoreStats.GameScore = 0
					}

					PlusMinusText := util.CleanTextDatum(s.Find("td:nth-child(22)").Text())
					boxScoreStats.PlusMinus, err = util.TextToInt32(PlusMinusText)
					if PlusMinusText == "" {
						boxScoreStats.PlusMinus = 0
					} else if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for PlusMinusText to Int - %w; defaulting to 0.", PlusMinusText, err))
						boxScoreStats.PlusMinus = 0
					}
				} else {
					boxScoreStats.MinutesPlayed = 0
				}
				basicNBABoxScoreStats = append(basicNBABoxScoreStats, boxScoreStats)
			} else {
				s.Find(boxScoreReserveHeaders).EachWithBreak(func(idx int, s *goquery.Selection) bool {
					reserveHeader = util.CleanTextDatum(s.Text())
					expectedHeader := basicBoxScoreReservesHeaders[idx]
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
	output.Output = basicNBABoxScoreStats
	return output
}
