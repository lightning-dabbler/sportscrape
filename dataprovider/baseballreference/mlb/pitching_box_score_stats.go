package mlb

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreference"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreference/mlb/model"
	"github.com/lightning-dabbler/sportscrape/util"
	sportsreferenceutil "github.com/lightning-dabbler/sportscrape/util/sportsreference"
	"github.com/xitongsys/parquet-go/types"
)

var pitchingBoxScoreHeaders sportsreferenceutil.Headers = sportsreferenceutil.Headers{
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

// PitchingBoxScoreOption defines a configuration option for the pitching box score runner
type PitchingBoxScoreOption func(*PitchingBoxScoreRunner)

// WithPitchingBoxScoreTimeout sets the timeout duration for pitching box score runner
func WithPitchingBoxScoreTimeout(timeout time.Duration) PitchingBoxScoreOption {
	return func(bsr *PitchingBoxScoreRunner) {
		bsr.Timeout = timeout
	}
}

// WithPitchingBoxScoreDebug enables or disables debug mode for the pitching box score runner
func WithPitchingBoxScoreDebug(debug bool) PitchingBoxScoreOption {
	return func(bsr *PitchingBoxScoreRunner) {
		bsr.Debug = debug
	}
}

// WithPitchingBoxScoreConcurrency sets the number of concurrent workers
func WithPitchingBoxScoreConcurrency(n int) PitchingBoxScoreOption {
	return func(bsr *PitchingBoxScoreRunner) {
		bsr.Concurrency = n
	}
}

// NewPitchingBoxScoreRunner creates a new PitchingBoxScoreRunner with the provided options
func NewPitchingBoxScoreRunner(options ...PitchingBoxScoreOption) *PitchingBoxScoreRunner {
	bsr := &PitchingBoxScoreRunner{}
	bsr.Processor = bsr
	//

	// Apply all options
	for _, option := range options {
		option(bsr)
	}

	return bsr
}

// PitchingBoxScoreRunner specialized Runner for retrieving MLB pitching box score statistics
// with support for concurrent processing.
type PitchingBoxScoreRunner struct {
	sportsreferenceutil.BoxScoreRunner
}

// GetSegmentBoxScoreStats retrieves MLB pitching box score statistics for a single matchup.
//
// Parameter:
//   - matchup: The MLB matchup for which to retrieve pitching box score statistics
//
// Returns a slice of MLB pitching box score statistics as interface{} values
func (boxScoreRunner *PitchingBoxScoreRunner) GetSegmentBoxScoreStats(matchup interface{}) []interface{} {
	matchupModel := matchup.(model.MLBMatchup)
	url := matchupModel.BoxScoreLink
	PullTimestamp := time.Now().UTC()
	start := time.Now().UTC()
	var boxScoreStats []interface{}
	log.Println("Scraping pitching Box Score: " + url)
	doc, err := boxScoreRunner.RetrieveDocument(url, networkHeaders, waitReadyBoxScoreContentSelector)
	if err != nil {
		log.Fatalln(err)
	}
	homeTeamSelector := generateStatTableSelector(matchupModel.HomeTeam, Pitching)
	awayTeamSelector := generateStatTableSelector(matchupModel.AwayTeam, Pitching)
	// Home team pitching stat box
	homeStatBox := doc.Find(homeTeamSelector)
	// Away team pitching stat box
	awayStatBox := doc.Find(awayTeamSelector)
	// Validate headers and their positions for each team's stat box
	homeStatBox.Find(headersSelector).Each(func(idx int, s *goquery.Selection) {
		header := util.CleanTextDatum(s.Text())
		expectedHeader := pitchingBoxScoreHeaders[idx]
		if header != expectedHeader {
			log.Fatalf("Home team header '%s' at position %d does not equal expected header '%s' @ %s\n", header, idx, expectedHeader, url)
		}
	})

	awayStatBox.Find(headersSelector).Each(func(idx int, s *goquery.Selection) {
		header := util.CleanTextDatum(s.Text())
		expectedHeader := pitchingBoxScoreHeaders[idx]
		if header != expectedHeader {
			log.Fatalf("Away team header '%s' at position %d does not equal expected header '%s' @ %s\n", header, idx, expectedHeader, url)
		}
	})

	// Parse records - home team
	homeStatBox.Find(recordSelector).Each(func(idx int, s *goquery.Selection) {
		statline := parsePitchingBoxScore(s)
		statline.PitchingOrder = int32(idx + 1)
		statline.PullTimestamp = PullTimestamp
		statline.EventID = matchupModel.EventID
		statline.Team = matchupModel.HomeTeam
		statline.Opponent = matchupModel.AwayTeam
		statline.EventDate = matchupModel.EventDate
		statline.PullTimestampParquet = types.TimeToTIMESTAMP_MILLIS(PullTimestamp, true)
		statline.EventDateParquet = util.TimeToDays(matchupModel.EventDate)
		boxScoreStats = append(boxScoreStats, statline)
	})
	// Parse records - away team
	awayStatBox.Find(recordSelector).Each(func(idx int, s *goquery.Selection) {
		statline := parsePitchingBoxScore(s)
		statline.PitchingOrder = int32(idx + 1)
		statline.PullTimestamp = PullTimestamp
		statline.EventID = matchupModel.EventID
		statline.Team = matchupModel.AwayTeam
		statline.Opponent = matchupModel.HomeTeam
		statline.EventDate = matchupModel.EventDate
		statline.PullTimestampParquet = types.TimeToTIMESTAMP_MILLIS(PullTimestamp, true)
		statline.EventDateParquet = util.TimeToDays(matchupModel.EventDate)
		boxScoreStats = append(boxScoreStats, statline)
	})

	if len(boxScoreStats) == 0 {
		log.Printf("No Data Scraped @ %s\n", url)
	} else {
		diff := time.Now().UTC().Sub(start)
		log.Printf("Scraping of %s Completed in %s\n", url, diff)
	}
	return boxScoreStats
}

// parsePitchingBoxScore parses a player's batting statline
//
// Parameters:
//   - s: goquery.Selection representing a player's batting statline
//
// Returns model.MLBPitchingBoxScoreStats containing parsed statistics
func parsePitchingBoxScore(s *goquery.Selection) model.MLBPitchingBoxScoreStats {
	var statline model.MLBPitchingBoxScoreStats
	// Player, PlayerLink, & PlayerID
	player := s.Find(playerSelector)
	statline.PlayerLink = baseballreference.URL + util.CleanTextDatum(player.AttrOr("href", ""))
	statline.Player = util.CleanTextDatum(player.Text())
	playerID, err := sportsreferenceutil.PlayerID(statline.PlayerLink)
	if err != nil {
		log.Fatalln(err)
	}
	statline.PlayerID = playerID

	// InningsPitched
	inningsPitchedText := util.CleanTextDatum(s.Find("td:nth-child(2)").Text())
	inningsPitched, err := util.TextToFloat32(inningsPitchedText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for inningsPitchedText to float32 - %w", inningsPitchedText, err))
	}
	statline.InningsPitched = inningsPitched

	// HitsAllowed
	hitsAllowedText := util.CleanTextDatum(s.Find("td:nth-child(3)").Text())
	hitsAllowed, err := util.TextToInt32(hitsAllowedText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for hitsAllowedText to int - %w", hitsAllowedText, err))
	}
	statline.HitsAllowed = hitsAllowed

	// RunsAllowed
	runsAllowedText := util.CleanTextDatum(s.Find("td:nth-child(4)").Text())
	runsAllowed, err := util.TextToInt32(runsAllowedText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for runsAllowedText to int - %w", runsAllowedText, err))
	}
	statline.RunsAllowed = runsAllowed

	// EarnedRunsAllowed
	earnedRunsAllowedText := util.CleanTextDatum(s.Find("td:nth-child(5)").Text())
	earnedRunsAllowed, err := util.TextToInt32(earnedRunsAllowedText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for earnedRunsAllowedText to int - %w", earnedRunsAllowedText, err))
	}
	statline.EarnedRunsAllowed = earnedRunsAllowed

	// Walks
	walksText := util.CleanTextDatum(s.Find("td:nth-child(6)").Text())
	walks, err := util.TextToInt32(walksText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for walksText to int - %w", walksText, err))
	}
	statline.Walks = walks

	// Strikeouts
	strikeoutsText := util.CleanTextDatum(s.Find("td:nth-child(7)").Text())
	strikeouts, err := util.TextToInt32(strikeoutsText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for strikeoutsText to int - %w", strikeoutsText, err))
	}
	statline.Strikeouts = strikeouts

	// HomeRunsAllowed
	homeRunsAllowedText := util.CleanTextDatum(s.Find("td:nth-child(8)").Text())
	homeRunsAllowed, err := util.TextToInt32(homeRunsAllowedText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for homeRunsAllowedText to int - %w", homeRunsAllowedText, err))
	}
	statline.HomeRunsAllowed = homeRunsAllowed

	// EarnedRunAverage
	earnedRunAverageText := util.CleanTextDatum(s.Find("td:nth-child(9)").Text())
	earnedRunAverage, err := util.TextToFloat32(earnedRunAverageText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for earnedRunAverageText to float32 - %w", earnedRunAverageText, err))
	}
	statline.EarnedRunAverage = earnedRunAverage

	// BattersFaced
	battersFacedText := util.CleanTextDatum(s.Find("td:nth-child(10)").Text())
	battersFaced, err := util.TextToInt32(battersFacedText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for battersFacedText to int - %w", battersFacedText, err))
	}
	statline.BattersFaced = battersFaced

	// PitchesPerPlateAppearance
	pitchesPerPlateAppearanceText := util.CleanTextDatum(s.Find("td:nth-child(11)").Text())
	pitchesPerPlateAppearance, err := util.TextToInt32(pitchesPerPlateAppearanceText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for pitchesPerPlateAppearanceText to int - %w", pitchesPerPlateAppearanceText, err))
	}
	statline.PitchesPerPlateAppearance = pitchesPerPlateAppearance

	// Strikes
	strikesText := util.CleanTextDatum(s.Find("td:nth-child(12)").Text())
	strikes, err := util.TextToInt32(strikesText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for strikesText to int - %w", strikesText, err))
	}
	statline.Strikes = strikes

	// StrikesByContact
	strikesByContactText := util.CleanTextDatum(s.Find("td:nth-child(13)").Text())
	strikesByContact, err := util.TextToInt32(strikesByContactText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for strikesByContactText to int - %w", strikesByContactText, err))
	}
	statline.StrikesByContact = strikesByContact

	// StrikesSwinging
	strikesSwingingText := util.CleanTextDatum(s.Find("td:nth-child(14)").Text())
	strikesSwinging, err := util.TextToInt32(strikesSwingingText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for strikesSwingingText to int - %w", strikesSwingingText, err))
	}
	statline.StrikesSwinging = strikesSwinging

	// StrikesLooking
	strikesLookingText := util.CleanTextDatum(s.Find("td:nth-child(15)").Text())
	strikesLooking, err := util.TextToInt32(strikesLookingText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for strikesLookingText to int - %w", strikesLookingText, err))
	}
	statline.StrikesLooking = strikesLooking

	// GroundBalls
	groundBallsText := util.CleanTextDatum(s.Find("td:nth-child(16)").Text())
	groundBalls, err := util.TextToInt32(groundBallsText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for groundBallsText to int - %w", groundBallsText, err))
	}
	statline.GroundBalls = groundBalls

	// FlyBalls
	flyBallsText := util.CleanTextDatum(s.Find("td:nth-child(17)").Text())
	flyBalls, err := util.TextToInt32(flyBallsText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for flyBallsText to int - %w", flyBallsText, err))
	}
	statline.FlyBalls = flyBalls

	// LineDrives
	lineDrivesText := util.CleanTextDatum(s.Find("td:nth-child(18)").Text())
	lineDrives, err := util.TextToInt32(lineDrivesText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for lineDrivesText to int - %w", lineDrivesText, err))
	}
	statline.LineDrives = lineDrives

	// UnknownBattedBallType
	unknownBattedBallTypeText := util.CleanTextDatum(s.Find("td:nth-child(19)").Text())
	unknownBattedBallType, err := util.TextToInt32(unknownBattedBallTypeText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for unknownBattedBallTypeText to int - %w", unknownBattedBallTypeText, err))
	}
	statline.UnknownBattedBallType = unknownBattedBallType

	// GameScore
	gameScoreText := util.CleanTextDatum(s.Find("td:nth-child(20)").Text())
	if gameScoreText != "" {
		gameScore, err := util.TextToInt32(gameScoreText)
		if err != nil {
			log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for gameScoreText to int - %w!", gameScoreText, err))
		}
		statline.GameScore = &gameScore
	}

	// InheritedRunners
	inheritedRunnersText := util.CleanTextDatum(s.Find("td:nth-child(21)").Text())
	if inheritedRunnersText != "" {
		inheritedRunners, err := util.TextToInt32(inheritedRunnersText)
		if err != nil {
			log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for inheritedRunnersText to int - %w!", inheritedRunnersText, err))
		}
		statline.InheritedRunners = &inheritedRunners
	}

	// InheritedScore
	inheritedScoreText := util.CleanTextDatum(s.Find("td:nth-child(22)").Text())
	if inheritedScoreText != "" {
		inheritedScore, err := util.TextToInt32(inheritedScoreText)
		if err != nil {
			log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for inheritedScoreText to int - %w!", inheritedScoreText, err))
		}
		statline.InheritedScore = &inheritedScore
	}

	// WinProbabilityAdded
	winProbabilityAddedText := util.CleanTextDatum(s.Find("td:nth-child(23)").Text())
	winProbabilityAdded, err := util.TextToFloat32(winProbabilityAddedText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for winProbabilityAddedText to float32 - %w", winProbabilityAddedText, err))
	}
	statline.WinProbabilityAdded = winProbabilityAdded

	// AverageLeverageIndex
	averageLeverageIndexText := util.CleanTextDatum(s.Find("td:nth-child(24)").Text())
	averageLeverageIndex, err := util.TextToFloat32(averageLeverageIndexText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for averageLeverageIndexText to float32 - %w", averageLeverageIndexText, err))
	}
	statline.AverageLeverageIndex = averageLeverageIndex

	// ChampionshipWinProbabilityAdded
	championshipWinProbabilityAddedText := strings.TrimRight(util.CleanTextDatum(s.Find("td:nth-child(25)").Text()), "%") // -2.92% --> -2.92
	championshipWinProbabilityAdded, err := util.TextToFloat32(championshipWinProbabilityAddedText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for championshipWinProbabilityAddedText to float32 - %w", championshipWinProbabilityAddedText, err))
	}
	statline.ChampionshipWinProbabilityAdded = championshipWinProbabilityAdded

	// AverageChampionshipLeverageIndex
	averageChampionshipLeverageIndexText := util.CleanTextDatum(s.Find("td:nth-child(26)").Text())
	averageChampionshipLeverageIndex, err := util.TextToFloat32(averageChampionshipLeverageIndexText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for averageChampionshipLeverageIndexText to int - %w", averageChampionshipLeverageIndexText, err))
	}
	statline.AverageChampionshipLeverageIndex = averageChampionshipLeverageIndex

	// BaseOutRunsSaved
	baseOutRunsSavedText := util.CleanTextDatum(s.Find("td:nth-child(27)").Text())
	baseOutRunsSaved, err := util.TextToFloat32(baseOutRunsSavedText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for baseOutRunsSavedText to float32 - %w", baseOutRunsSavedText, err))
	}
	statline.BaseOutRunsSaved = baseOutRunsSaved

	return statline
}
