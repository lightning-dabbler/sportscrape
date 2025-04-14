package eventdata

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

type OutputWrapper struct {
	Error   error
	Output  []interface{}
	Context Context
}

type Context struct {
	PullTimestamp time.Time
	EventTime     time.Time
	EventID       int64
	URL           string
	AwayID        int64
	AwayTeam      string
	HomeID        int64
	HomeTeam      string
}

// RunnerOption
type RunnerOption func(*Runner)

// RunnerName option
func RunnerName(name string) RunnerOption {
	return func(r *Runner) {
		r.Name = name
	}
}

// RunnerConcurrency option
func RunnerConcurrency(concurrency int) RunnerOption {
	return func(r *Runner) {
		r.Concurrency = concurrency
	}
}

// RunnerScraper option
func RunnerScraper(scraper Scraper) RunnerOption {
	return func(r *Runner) {
		r.Scraper = scraper
	}
}

// NewRunner Instantiates a new Runner
func NewRunner(options ...RunnerOption) *Runner {
	r := &Runner{}
	// Apply all options
	for _, option := range options {
		option(r)
	}

	r.Scraper.SetParams()
	return r
}

type Runner struct {
	Name        string
	Concurrency int
	Scraper     Scraper
}

func (t *Runner) RunEventsDataScraper(matchups ...interface{}) []interface{} {
	start := time.Now().UTC()
	concurrency := t.Concurrency
	if concurrency <= 0 {
		concurrency = runtime.NumCPU()
	}

	var wg sync.WaitGroup
	workerMatchups := make(chan interface{}, concurrency)
	eventData := make(chan OutputWrapper, len(matchups))

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
			log.Println(fmt.Errorf("Issue Scraping %d (%s vs %s) at url: '%s': %w", ow.Context.EventID, ow.Context.AwayTeam, ow.Context.HomeTeam, ow.Context.URL, ow.Error))
			continue
		} else {
			log.Printf("%d (%s vs %s) scraped for %d records for %s at url: %s\n", ow.Context.EventID, ow.Context.AwayTeam, ow.Context.HomeTeam, len(ow.Output), t.Name, ow.Context.URL)
		}
		output = append(output, ow.Output...)
	}
	outputCount := len(output)
	if errors != 0 {
		log.Printf("WARNING: %d/%d events errored out\n", errors, matchupsCount)
	}
	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of %s with %d records completed in %s\n", t.Name, outputCount, diff)

	return output
}

func (t *Runner) Worker(wg *sync.WaitGroup, workerMatchups <-chan interface{}, eventData chan<- OutputWrapper) {
	for matchup := range workerMatchups {
		ow := t.Scraper.Scrape(matchup)
		eventData <- ow
		wg.Done()
	}
}
