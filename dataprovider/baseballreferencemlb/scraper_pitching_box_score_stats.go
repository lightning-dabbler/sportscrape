package baseballreferencemlb

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreferencemlb/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/lightning-dabbler/sportscrape/util/sportsreference"
	"github.com/xitongsys/parquet-go/types"
)

var pitchingBoxScoreHeaders sportsreference.Headers = sportsreference.Headers{
	"Pitching",
	"IP",
	"H",
	"R",
	"ER",
	"BB",
	"SO",
	"HR",
	"ERA",
	"BF",
	"Pit",
	"Str",
	"Ctct",
	"StS",
	"StL",
	"GB",
	"FB",
	"LD",
	"Unk",
	"GSc",
	"IR",
	"IS",
	"WPA",
	"aLI",
	"cWPA",
	"acLI",
	"RE24",
}

// PitchingBoxScoreOption defines a configuration option for the pitching box score scraper
type PitchingBoxScoreOption func(*PitchingBoxScoreScraper)

// WithPitchingBoxScoreTimeout sets the timeout duration for pitching box score scraper
func WithPitchingBoxScoreTimeout(timeout time.Duration) PitchingBoxScoreOption {
	return func(bsr *PitchingBoxScoreScraper) {
		bsr.Timeout = timeout
	}
}

// WithPitchingBoxScoreDebug enables or disables debug mode for the pitching box score scraper
func WithPitchingBoxScoreDebug(debug bool) PitchingBoxScoreOption {
	return func(bsr *PitchingBoxScoreScraper) {
		bsr.Debug = debug
	}
}

// NewPitchingBoxScoreScraper creates a new PitchingBoxScoreScraper with the provided options
func NewPitchingBoxScoreScraper(options ...PitchingBoxScoreOption) *PitchingBoxScoreScraper {
	bsr := &PitchingBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(bsr)
	}
	bsr.Init()

	return bsr
}

// PitchingBoxScoreScraper specialized Runner for retrieving MLB pitching box score statistics
// with support for concurrent processing.
type PitchingBoxScoreScraper struct {
	EventDataScraper
}

func (s PitchingBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.BaseballReferenceMLBPitchingBoxScore
}

// Scrape retrieves MLB pitching box score statistics for a single matchup.
func (s *PitchingBoxScoreScraper) Scrape(matchup interface{}) sportscrape.EventDataOutput {
	matchupModel := matchup.(model.MLBMatchup)
	context := s.ConstructContext(matchupModel)
	output := sportscrape.EventDataOutput{
		Context: context,
	}

	url := matchupModel.BoxScoreLink
	PullTimestamp := time.Now().UTC()
	start := time.Now().UTC()
	var boxScoreStats []interface{}
	log.Println("Scraping pitching Box Score: " + url)
	doc, err := s.RetrieveDocument(url, networkHeaders, contentReadySelector)
	if err != nil {
		output.Error = err
		return output
	}
	homeTeamSelector := generateStatTableSelector(matchupModel.HomeTeam, Pitching)
	awayTeamSelector := generateStatTableSelector(matchupModel.AwayTeam, Pitching)
	// Home team pitching stat box
	homeStatBox := doc.Find(homeTeamSelector)
	// Away team pitching stat box
	awayStatBox := doc.Find(awayTeamSelector)
	// Validate headers and their positions for each team's stat box
	homeStatBox.Find(headersSelector).EachWithBreak(func(idx int, s *goquery.Selection) bool {
		header := util.CleanTextDatum(s.Text())
		expectedHeader := pitchingBoxScoreHeaders[idx]
		if header != expectedHeader {
			err = fmt.Errorf("home team header '%s' at position %d does not equal expected header '%s' @ %s", header, idx, expectedHeader, url)
			output.Error = err
			return false
		}
		return true
	})

	if output.Error != nil {
		return output
	}

	awayStatBox.Find(headersSelector).EachWithBreak(func(idx int, s *goquery.Selection) bool {
		header := util.CleanTextDatum(s.Text())
		expectedHeader := pitchingBoxScoreHeaders[idx]
		if header != expectedHeader {
			err = fmt.Errorf("away team header '%s' at position %d does not equal expected header '%s' @ %s", header, idx, expectedHeader, url)
			output.Error = err
			return false
		}
		return true
	})

	if output.Error != nil {
		return output
	}

	// Parse records - home team
	homeStatBox.Find(recordSelector).EachWithBreak(func(idx int, s *goquery.Selection) bool {
		statline, err := parsePitchingBoxScore(s)
		if err != nil {
			output.Error = err
			return false
		}
		statline.PitchingOrder = int32(idx + 1)
		statline.PullTimestamp = PullTimestamp
		statline.EventID = matchupModel.EventID
		statline.Team = matchupModel.HomeTeam
		statline.TeamID = matchupModel.HomeTeamID
		statline.OpponentID = matchupModel.AwayTeamID
		statline.Opponent = matchupModel.AwayTeam
		statline.EventDate = matchupModel.EventDate
		statline.PullTimestampParquet = types.TimeToTIMESTAMP_MILLIS(PullTimestamp, true)
		statline.EventDateParquet = util.TimeToDays(matchupModel.EventDate)
		boxScoreStats = append(boxScoreStats, statline)
		return true
	})
	if output.Error != nil {
		return output
	}
	// Parse records - away team
	awayStatBox.Find(recordSelector).EachWithBreak(func(idx int, s *goquery.Selection) bool {
		statline, err := parsePitchingBoxScore(s)
		if err != nil {
			output.Error = err
			return false
		}
		statline.PitchingOrder = int32(idx + 1)
		statline.PullTimestamp = PullTimestamp
		statline.EventID = matchupModel.EventID
		statline.Team = matchupModel.AwayTeam
		statline.TeamID = matchupModel.AwayTeamID
		statline.Opponent = matchupModel.HomeTeam
		statline.OpponentID = matchupModel.HomeTeamID
		statline.EventDate = matchupModel.EventDate
		statline.PullTimestampParquet = types.TimeToTIMESTAMP_MILLIS(PullTimestamp, true)
		statline.EventDateParquet = util.TimeToDays(matchupModel.EventDate)
		boxScoreStats = append(boxScoreStats, statline)
		return true
	})
	if output.Error != nil {
		return output
	}

	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of %s Completed in %s\n", url, diff)
	output.Output = boxScoreStats
	return output
}

// parsePitchingBoxScore parses a player's batting statline
//
// Parameters:
//   - s: goquery.Selection representing a player's batting statline
//
// Returns model.MLBPitchingBoxScoreStats containing parsed statistics
func parsePitchingBoxScore(s *goquery.Selection) (model.MLBPitchingBoxScoreStats, error) {
	var statline model.MLBPitchingBoxScoreStats
	// Player, PlayerLink, & PlayerID
	player := s.Find(playerSelector)
	statline.PlayerLink = sportsreference.BaseballRefURL + util.CleanTextDatum(player.AttrOr("href", ""))
	statline.Player = util.CleanTextDatum(player.Text())
	playerID, err := sportsreference.PlayerID(statline.PlayerLink)
	if err != nil {
		return statline, err
	}
	statline.PlayerID = playerID

	// InningsPitched
	inningsPitchedText := util.CleanTextDatum(s.Find("td:nth-child(2)").Text())
	inningsPitched, err := util.TextToFloat32(inningsPitchedText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for inningsPitchedText to float32 - %w", inningsPitchedText, err)
		return statline, err
	}
	statline.InningsPitched = inningsPitched

	// HitsAllowed
	hitsAllowedText := util.CleanTextDatum(s.Find("td:nth-child(3)").Text())
	hitsAllowed, err := util.TextToInt32(hitsAllowedText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for hitsAllowedText to int - %w", hitsAllowedText, err)
		return statline, err
	}
	statline.HitsAllowed = hitsAllowed

	// RunsAllowed
	runsAllowedText := util.CleanTextDatum(s.Find("td:nth-child(4)").Text())
	runsAllowed, err := util.TextToInt32(runsAllowedText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for runsAllowedText to int - %w", runsAllowedText, err)
		return statline, err
	}
	statline.RunsAllowed = runsAllowed

	// EarnedRunsAllowed
	earnedRunsAllowedText := util.CleanTextDatum(s.Find("td:nth-child(5)").Text())
	earnedRunsAllowed, err := util.TextToInt32(earnedRunsAllowedText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for earnedRunsAllowedText to int - %w", earnedRunsAllowedText, err)
		return statline, err
	}
	statline.EarnedRunsAllowed = earnedRunsAllowed

	// Walks
	walksText := util.CleanTextDatum(s.Find("td:nth-child(6)").Text())
	walks, err := util.TextToInt32(walksText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for walksText to int - %w", walksText, err)
		return statline, err
	}
	statline.Walks = walks

	// Strikeouts
	strikeoutsText := util.CleanTextDatum(s.Find("td:nth-child(7)").Text())
	strikeouts, err := util.TextToInt32(strikeoutsText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for strikeoutsText to int - %w", strikeoutsText, err)
		return statline, err
	}
	statline.Strikeouts = strikeouts

	// HomeRunsAllowed
	homeRunsAllowedText := util.CleanTextDatum(s.Find("td:nth-child(8)").Text())
	homeRunsAllowed, err := util.TextToInt32(homeRunsAllowedText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for homeRunsAllowedText to int - %w", homeRunsAllowedText, err)
		return statline, err
	}
	statline.HomeRunsAllowed = homeRunsAllowed

	// EarnedRunAverage
	earnedRunAverageText := util.CleanTextDatum(s.Find("td:nth-child(9)").Text())
	earnedRunAverage, err := util.TextToFloat32(earnedRunAverageText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for earnedRunAverageText to float32 - %w", earnedRunAverageText, err)
		return statline, err
	}
	statline.EarnedRunAverage = earnedRunAverage

	// BattersFaced
	battersFacedText := util.CleanTextDatum(s.Find("td:nth-child(10)").Text())
	battersFaced, err := util.TextToInt32(battersFacedText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for battersFacedText to int - %w", battersFacedText, err)
		return statline, err
	}
	statline.BattersFaced = battersFaced

	// PitchesPerPlateAppearance
	pitchesPerPlateAppearanceText := util.CleanTextDatum(s.Find("td:nth-child(11)").Text())
	pitchesPerPlateAppearance, err := util.TextToInt32(pitchesPerPlateAppearanceText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for pitchesPerPlateAppearanceText to int - %w", pitchesPerPlateAppearanceText, err)
		return statline, err
	}
	statline.PitchesPerPlateAppearance = pitchesPerPlateAppearance

	// Strikes
	strikesText := util.CleanTextDatum(s.Find("td:nth-child(12)").Text())
	strikes, err := util.TextToInt32(strikesText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for strikesText to int - %w", strikesText, err)
		return statline, err
	}
	statline.Strikes = strikes

	// StrikesByContact
	strikesByContactText := util.CleanTextDatum(s.Find("td:nth-child(13)").Text())
	strikesByContact, err := util.TextToInt32(strikesByContactText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for strikesByContactText to int - %w", strikesByContactText, err)
		return statline, err
	}
	statline.StrikesByContact = strikesByContact

	// StrikesSwinging
	strikesSwingingText := util.CleanTextDatum(s.Find("td:nth-child(14)").Text())
	strikesSwinging, err := util.TextToInt32(strikesSwingingText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for strikesSwingingText to int - %w", strikesSwingingText, err)
		return statline, err
	}
	statline.StrikesSwinging = strikesSwinging

	// StrikesLooking
	strikesLookingText := util.CleanTextDatum(s.Find("td:nth-child(15)").Text())
	strikesLooking, err := util.TextToInt32(strikesLookingText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for strikesLookingText to int - %w", strikesLookingText, err)
		return statline, err
	}
	statline.StrikesLooking = strikesLooking

	// GroundBalls
	groundBallsText := util.CleanTextDatum(s.Find("td:nth-child(16)").Text())
	groundBalls, err := util.TextToInt32(groundBallsText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for groundBallsText to int - %w", groundBallsText, err)
		return statline, err
	}
	statline.GroundBalls = groundBalls

	// FlyBalls
	flyBallsText := util.CleanTextDatum(s.Find("td:nth-child(17)").Text())
	flyBalls, err := util.TextToInt32(flyBallsText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for flyBallsText to int - %w", flyBallsText, err)
		return statline, err
	}
	statline.FlyBalls = flyBalls

	// LineDrives
	lineDrivesText := util.CleanTextDatum(s.Find("td:nth-child(18)").Text())
	lineDrives, err := util.TextToInt32(lineDrivesText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for lineDrivesText to int - %w", lineDrivesText, err)
		return statline, err
	}
	statline.LineDrives = lineDrives

	// UnknownBattedBallType
	unknownBattedBallTypeText := util.CleanTextDatum(s.Find("td:nth-child(19)").Text())
	unknownBattedBallType, err := util.TextToInt32(unknownBattedBallTypeText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for unknownBattedBallTypeText to int - %w", unknownBattedBallTypeText, err)
		return statline, err
	}
	statline.UnknownBattedBallType = unknownBattedBallType

	// GameScore
	gameScoreText := util.CleanTextDatum(s.Find("td:nth-child(20)").Text())
	if gameScoreText != "" {
		gameScore, err := util.TextToInt32(gameScoreText)
		if err != nil {
			err = fmt.Errorf("ERROR: Can't convert '%s' for gameScoreText to int - %w", gameScoreText, err)
			return statline, err
		}
		statline.GameScore = &gameScore
	}

	// InheritedRunners
	inheritedRunnersText := util.CleanTextDatum(s.Find("td:nth-child(21)").Text())
	if inheritedRunnersText != "" {
		inheritedRunners, err := util.TextToInt32(inheritedRunnersText)
		if err != nil {
			err = fmt.Errorf("ERROR: Can't convert '%s' for inheritedRunnersText to int - %w", inheritedRunnersText, err)
			return statline, err
		}
		statline.InheritedRunners = &inheritedRunners
	}

	// InheritedScore
	inheritedScoreText := util.CleanTextDatum(s.Find("td:nth-child(22)").Text())
	if inheritedScoreText != "" {
		inheritedScore, err := util.TextToInt32(inheritedScoreText)
		if err != nil {
			err = fmt.Errorf("ERROR: Can't convert '%s' for inheritedScoreText to int - %w", inheritedScoreText, err)
			return statline, err
		}
		statline.InheritedScore = &inheritedScore
	}

	// WinProbabilityAdded
	winProbabilityAddedText := util.CleanTextDatum(s.Find("td:nth-child(23)").Text())
	winProbabilityAdded, err := util.TextToFloat32(winProbabilityAddedText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for winProbabilityAddedText to float32 - %w", winProbabilityAddedText, err)
		return statline, err
	}
	statline.WinProbabilityAdded = winProbabilityAdded

	// AverageLeverageIndex
	averageLeverageIndexText := util.CleanTextDatum(s.Find("td:nth-child(24)").Text())
	averageLeverageIndex, err := util.TextToFloat32(averageLeverageIndexText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for averageLeverageIndexText to float32 - %w", averageLeverageIndexText, err)
		return statline, err
	}
	statline.AverageLeverageIndex = averageLeverageIndex

	// ChampionshipWinProbabilityAdded
	championshipWinProbabilityAddedText := strings.TrimRight(util.CleanTextDatum(s.Find("td:nth-child(25)").Text()), "%") // -2.92% --> -2.92
	championshipWinProbabilityAdded, err := util.TextToFloat32(championshipWinProbabilityAddedText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for championshipWinProbabilityAddedText to float32 - %w", championshipWinProbabilityAddedText, err)
		return statline, err
	}
	statline.ChampionshipWinProbabilityAdded = championshipWinProbabilityAdded

	// AverageChampionshipLeverageIndex
	averageChampionshipLeverageIndexText := util.CleanTextDatum(s.Find("td:nth-child(26)").Text())
	averageChampionshipLeverageIndex, err := util.TextToFloat32(averageChampionshipLeverageIndexText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for averageChampionshipLeverageIndexText to int - %w", averageChampionshipLeverageIndexText, err)
		return statline, err
	}
	statline.AverageChampionshipLeverageIndex = averageChampionshipLeverageIndex

	// BaseOutRunsSaved
	baseOutRunsSavedText := util.CleanTextDatum(s.Find("td:nth-child(27)").Text())
	baseOutRunsSaved, err := util.TextToFloat32(baseOutRunsSavedText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for baseOutRunsSavedText to float32 - %w", baseOutRunsSavedText, err)
		return statline, err
	}
	statline.BaseOutRunsSaved = baseOutRunsSaved

	return statline, nil
}
