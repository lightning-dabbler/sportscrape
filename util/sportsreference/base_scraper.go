package sportsreference

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/lightning-dabbler/sportscrape/util/request"
)

// Runner provides base functionality for retrieving documents from web pages.
// It encapsulates configuration for document retrieval operations.
type Runner struct {
	// Timeout is the context timeout for Chrome operations
	Timeout time.Duration
	// Debug enables verbose logging of Chrome operations when true
	Debug bool
}

// RetrieveDocument fetches and parses a web document from the specified URL.
// It uses a DocumentRetriever to load the page and parse it into a goquery Document.
//
// Parameters:
//   - url: The web page URL to retrieve
//   - networkHeaders: HTTP headers to include with the request
//   - waitReadySelector: CSS selector of element that indicates page is fully loaded
//
// Returns:
//   - *goquery.Document: The parsed document
//   - error: Any error encountered during fetching or parsing
func (runner *Runner) RetrieveDocument(url string, networkHeaders network.Headers, waitReadySelector string) (*goquery.Document, error) {
	var retrieveOptions []request.RetrieverOption
	// Only add the timeout option if it's non-zero
	if runner.Timeout != 0 {
		retrieveOptions = append(retrieveOptions, request.WithTimeout(runner.Timeout))
	}

	// Add debug option
	if runner.Debug {
		retrieveOptions = append(retrieveOptions, request.WithDebug(runner.Debug))
	}

	// Create the document retriever with the collected options
	dr := request.NewDocumentRetriever(retrieveOptions...)
	return dr.RetrieveDocument(url, networkHeaders, waitReadySelector)
}

// MatchupRunner specialized Runner for retrieving matchup information.
type MatchupRunner struct {
	Runner
}

// BaseScraper provides base functionality for retrieving documents from web pages.
// It encapsulates configuration for document retrieval operations.
type BaseScraper struct {
	// Timeout is the context timeout for Chrome operations
	Timeout time.Duration
	// Debug enables verbose logging of Chrome operations when true
	Debug bool
}

func (s BaseScraper) Init() {
	if s.Timeout == 0 {
		log.Fatalln("Timeout needs to be > 0")
	}
}

// RetrieveDocument fetches and parses a web document from the specified URL.
// It uses a DocumentRetriever to load the page and parse it into a goquery Document.
//
// Parameters:
//   - url: The web page URL to retrieve
//   - networkHeaders: HTTP headers to include with the request
//   - waitReadySelector: CSS selector of element that indicates page is fully loaded
//
// Returns:
//   - *goquery.Document: The parsed document
//   - error: Any error encountered during fetching or parsing
func (s *BaseScraper) RetrieveDocument(url string, networkHeaders network.Headers, waitReadySelector string) (*goquery.Document, error) {
	var retrieveOptions []request.RetrieverOption
	// Only add the timeout option if it's non-zero
	if s.Timeout != 0 {
		retrieveOptions = append(retrieveOptions, request.WithTimeout(s.Timeout))
	}

	// Add debug option
	if s.Debug {
		retrieveOptions = append(retrieveOptions, request.WithDebug(s.Debug))
	}

	// Create the document retriever with the collected options
	dr := request.NewDocumentRetriever(retrieveOptions...)
	return dr.RetrieveDocument(url, networkHeaders, waitReadySelector)
}

// GetMatchups retrieves matchups for the specified date.
//
// Parameter:
//   - date: The date for which to retrieve matchups
//
// Returns a slice of matchup data as interface{} values
func (matchupRunner *MatchupRunner) GetMatchups(date string) []interface{} {
	return []interface{}{}
}

// BoxScoreProcessor defines the interface for processing box scores
type BoxScoreProcessor interface {
	GetSegmentBoxScoreStats(matchup interface{}) []interface{}
}

// BoxScoreRunner specialized Runner for retrieving box score statistics
// with support for concurrent processing.
type BoxScoreRunner struct {
	Runner
	// Concurrency sets the maximum number of concurrent workers to request documents
	Concurrency int
	// Processor defines the interface for processing box scores
	Processor BoxScoreProcessor
}

// GetBoxScoresStats retrieves box score statistics for the provided matchups
// using concurrent workers for improved performance.
//
// Parameter:
//   - matchups: A variadic list of matchups to process
//
// Returns a slice containing all box score statistics as interface{} values
func (boxScoreRunner *BoxScoreRunner) GetBoxScoresStats(matchups ...interface{}) []interface{} {
	// Set default concurrency if not specified
	concurrency := boxScoreRunner.Concurrency
	if concurrency <= 0 {
		concurrency = runtime.NumCPU()
	}

	var wg sync.WaitGroup
	workerMatchups := make(chan interface{}, concurrency)
	BoxScoreStats := make(chan []interface{}, len(matchups))

	// Start worker goroutines
	for i := 0; i < cap(workerMatchups); i++ {
		go boxScoreRunner.Worker(&wg, workerMatchups, BoxScoreStats)
	}

	// Send matchups to workers
	for _, matchup := range matchups {
		wg.Add(1)
		workerMatchups <- matchup
	}

	wg.Wait()
	close(workerMatchups)
	close(BoxScoreStats)

	// Collect results
	var allBoxScoreStats []interface{}
	for boxScoreStats := range BoxScoreStats {
		allBoxScoreStats = append(allBoxScoreStats, boxScoreStats...)
	}

	return allBoxScoreStats
}

// Worker processes matchups from the input channel and sends resulting box score
// statistics to the output channel. Designed to be run as a goroutine.
//
// Parameters:
//   - wg: WaitGroup for synchronizing worker completion
//   - workerMatchups: Channel providing matchups to process
//   - boxScoreStats: Channel for sending processed box score statistics
func (boxScoreRunner *BoxScoreRunner) Worker(wg *sync.WaitGroup, workerMatchups <-chan interface{}, boxScoreStats chan<- []interface{}) {
	for matchup := range workerMatchups {
		boxScoreStats <- boxScoreRunner.Processor.GetSegmentBoxScoreStats(matchup)
		wg.Done()
	}
}
