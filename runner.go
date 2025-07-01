package sportscrape

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

type EventDataOutput struct {
	Error   error
	Output  []interface{}
	Context EventDataContext
}

type EventDataContext struct {
	PullTimestamp time.Time
	EventTime     time.Time
	EventID       int64
	URL           string
	AwayID        int64
	AwayTeam      string
	HomeID        int64
	HomeTeam      string
}

// EventDataRunnerOption
type EventDataRunnerOption func(*EventDataRunner)

// EventDataRunnerConcurrency option
func EventDataRunnerConcurrency(concurrency int) EventDataRunnerOption {
	return func(r *EventDataRunner) {
		r.Concurrency = concurrency
	}
}

// EventDataRunnerScraper option
func EventDataRunnerScraper(scraper EventDataScraper) EventDataRunnerOption {
	return func(r *EventDataRunner) {
		r.Scraper = scraper
	}
}

// NewEventDataRunner Instantiates a new EventDataRunner
func NewEventDataRunner(options ...EventDataRunnerOption) *EventDataRunner {
	r := &EventDataRunner{}
	// Apply all options
	for _, option := range options {
		option(r)
	}

	r.Scraper.Init()
	return r
}

type EventDataRunner struct {
	Concurrency int
	Scraper     EventDataScraper
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
	start := time.Now().UTC()
	concurrency := t.Concurrency
	if concurrency <= 0 {
		concurrency = runtime.NumCPU()
	}

	var wg sync.WaitGroup
	workerMatchups := make(chan interface{}, concurrency)
	eventData := make(chan EventDataOutput, len(matchups))

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
			log.Println(fmt.Errorf("issue Scraping %d (%s vs %s) at url: '%s': %w", ow.Context.EventID, ow.Context.AwayTeam, ow.Context.HomeTeam, ow.Context.URL, ow.Error))
			continue
		} else {
			log.Printf("%d (%s vs %s) scraped for %d record(s) for %s at url: %s\n", ow.Context.EventID, ow.Context.AwayTeam, ow.Context.HomeTeam, len(ow.Output), t.Scraper.Feed(), ow.Context.URL)
		}
		output = append(output, ow.Output...)
	}
	outputCount := len(output)
	if errors != 0 {
		log.Printf("WARNING: %d/%d events errored out\n", errors, matchupsCount)
	}
	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of %s with %d record(s) completed in %s\n", t.Scraper.Feed(), outputCount, diff)

	return output, nil
}

func (t *EventDataRunner) Worker(wg *sync.WaitGroup, workerMatchups <-chan interface{}, eventData chan<- EventDataOutput) {
	for matchup := range workerMatchups {
		ow := t.Scraper.Scrape(matchup)
		eventData <- ow
		wg.Done()
	}
}

type MatchupOutput struct {
	Output  []interface{}
	Context MatchupContext
}

type MatchupContext struct {
	Errors int
	Skips  int
	Total  int
}

// MatchupRunnerOption
type MatchupRunnerOption func(*MatchupRunner)

// MatchupRunnerScraper option
func MatchupRunnerScraper(scraper MatchupScraper) MatchupRunnerOption {
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
	Scraper MatchupScraper
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
	start := time.Now().UTC()
	ou := r.Scraper.Scrape()
	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of %s with %d record(s) completed in %s\n", r.Scraper.Feed(), len(ou.Output), diff)
	if ou.Context.Errors != 0 {
		log.Printf("WARNING: %d events errored out\n", ou.Context.Errors)
	}

	if ou.Context.Skips != 0 {
		log.Printf("WARNING: %d events skipped\n", ou.Context.Skips)
	}
	return ou.Output, nil
}
