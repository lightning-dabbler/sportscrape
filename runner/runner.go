package runner

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/scraper"
)

// EventDataRunnerOption
type EventDataRunnerOption func(*EventDataRunner)

// EventDataRunnerConcurrency option
func EventDataRunnerConcurrency(concurrency int) EventDataRunnerOption {
	return func(r *EventDataRunner) {
		r.Concurrency = concurrency
	}
}

// EventDataRunnerScraper option
func EventDataRunnerScraper(scraper scraper.EventDataScraper) EventDataRunnerOption {
	return func(r *EventDataRunner) {
		r.Scraper = scraper
	}
}

// NewEventDataRunner Instantiates a new EventDataRunner
func NewEventDataRunner(options ...EventDataRunnerOption) *EventDataRunner {
	r := &EventDataRunner{}
	// Default
	r.Concurrency = 1
	// Apply all options
	for _, option := range options {
		option(r)
	}

	r.Scraper.Init()
	return r
}

type EventDataRunner struct {
	Concurrency int
	Scraper     scraper.EventDataScraper
}

// Deprecated is a deprecation check for the feed/provider
func (t *EventDataRunner) Deprecated() bool {
	provider := t.Scraper.Provider()
	feed := t.Scraper.Feed()
	if provider.Deprecated() {
		return true
	}
	return feed.Deprecated()
}

func (t *EventDataRunner) Run(matchups ...interface{}) ([]interface{}, error) {
	if t.Deprecated() {
		return nil, t.Scraper.Feed().Deprecation()
	}
	t.Scraper.Init()
	start := time.Now().UTC()
	concurrency := t.Concurrency
	if concurrency <= 0 {
		concurrency = runtime.NumCPU()
	}

	var wg sync.WaitGroup
	workerMatchups := make(chan interface{}, concurrency)
	eventData := make(chan sportscrape.EventDataOutput, len(matchups))

	// Start worker goroutines
	for i := 0; i < cap(workerMatchups); i++ {
		go t.Worker(&wg, workerMatchups, eventData)
	}

	// Send matchups to workers
	matchupsCount := len(matchups)
	log.Printf("Processing %d matchups", matchupsCount)
	for _, matchup := range matchups {
		wg.Add(1)
		workerMatchups <- matchup
	}

	wg.Wait()
	close(workerMatchups)
	close(eventData)

	// Collect results
	var errors int
	var output []interface{}
	for ow := range eventData {
		if ow.Error != nil {
			errors += 1
			log.Println(fmt.Errorf("issue Scraping %v (%s vs %s) at url: '%s': %w", ow.Context.EventID, ow.Context.AwayTeam, ow.Context.HomeTeam, ow.Context.URL, ow.Error))
			continue
		} else {
			log.Printf("%v (%s vs %s) scraped for %v record(s) for %s at url: %s\n", ow.Context.EventID, ow.Context.AwayTeam, ow.Context.HomeTeam, len(ow.Output), t.Scraper.Feed(), ow.Context.URL)
		}
		output = append(output, ow.Output...)
	}
	outputCount := len(output)
	if errors != 0 {
		return nil, fmt.Errorf("error: %d/%d events errored out", errors, matchupsCount)
	}
	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of %s with %d record(s) completed in %s\n", t.Scraper.Feed(), outputCount, diff)

	return output, nil
}

func (t *EventDataRunner) Worker(wg *sync.WaitGroup, workerMatchups <-chan interface{}, eventData chan<- sportscrape.EventDataOutput) {
	for matchup := range workerMatchups {
		ow := t.Scraper.Scrape(matchup)
		eventData <- ow
		wg.Done()
	}
}

// MatchupRunnerOption
type MatchupRunnerOption func(*MatchupRunner)

// MatchupRunnerScraper option
func MatchupRunnerScraper(scraper scraper.MatchupScraper) MatchupRunnerOption {
	return func(r *MatchupRunner) {
		r.Scraper = scraper
	}
}

// NewMatchupRunner Instantiates a new MatchupRunner
func NewMatchupRunner(options ...MatchupRunnerOption) *MatchupRunner {
	r := &MatchupRunner{}
	// Apply all options
	for _, option := range options {
		option(r)
	}

	r.Scraper.Init()
	return r
}

// MatchupRunner is a general matchup runner for scraping NBA, MLB, NCAAB, etc. matchup data.
type MatchupRunner struct {
	// Scraper
	Scraper scraper.MatchupScraper
}

// Deprecated is a deprecation check for the feed/provider
func (r *MatchupRunner) Deprecated() bool {
	provider := r.Scraper.Provider()
	feed := r.Scraper.Feed()
	if provider.Deprecated() {
		return true
	}
	return feed.Deprecated()
}

// Run gets all matchups
func (r *MatchupRunner) Run() ([]interface{}, error) {
	if r.Deprecated() {
		return nil, r.Scraper.Feed().Deprecation()
	}
	r.Scraper.Init()
	start := time.Now().UTC()
	ou := r.Scraper.Scrape()
	if ou.Error != nil {
		return nil, ou.Error
	}
	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of %s with %d record(s) completed in %s\n", r.Scraper.Feed(), len(ou.Output), diff)
	if ou.Context.Skips != 0 {
		log.Printf("WARNING: %d event(s) skipped\n", ou.Context.Skips)
	}
	if ou.Context.Errors != 0 {
		return ou.Output, fmt.Errorf("error: %d event(s) errored out", ou.Context.Errors)
	}

	return ou.Output, nil
}
