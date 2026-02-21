package runner

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/scraper"
)

// EventDataRunnerConfig

type EventDataRunnerConfig[M, E any] struct {
	Concurrency int
	Scraper     scraper.EventDataScraper[M, E]
}

func NewEventDataRunner[M, E any](config EventDataRunnerConfig[M, E]) *EventDataRunner[M, E] {
	if config.Concurrency <= 0 {
		config.Concurrency = 1
	}
	r := &EventDataRunner[M, E]{
		Concurrency: config.Concurrency,
		Scraper:     config.Scraper,
	}
	r.Scraper.Init()
	return r
}

type EventDataRunner[M, E any] struct {
	Concurrency int
	Scraper     scraper.EventDataScraper[M, E]
}

// Deprecated is a deprecation check for the feed/provider
func (t *EventDataRunner[M, E]) Deprecated() bool {
	provider := t.Scraper.Provider()
	feed := t.Scraper.Feed()
	if provider.Deprecated() {
		return true
	}
	return feed.Deprecated()
}

func (t *EventDataRunner[M, E]) Run(matchups []M) ([]E, error) {
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
	workerMatchups := make(chan M, concurrency)
	eventData := make(chan sportscrape.EventDataOutput[E], len(matchups))

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
	var errCnt int
	var output []E
	var outputErr error
	for ow := range eventData {
		if ow.Error != nil {
			errCnt += 1
			outputErr = errors.Join(outputErr, fmt.Errorf("issue Scraping %v (%s vs %s) at url: '%s': %w", ow.Context.EventID, ow.Context.AwayTeam, ow.Context.HomeTeam, ow.Context.URL, ow.Error))
			continue
		} else {
			log.Printf("%v (%s vs %s) scraped for %v record(s) for %s at url: %s\n", ow.Context.EventID, ow.Context.AwayTeam, ow.Context.HomeTeam, len(ow.Output), t.Scraper.Feed(), ow.Context.URL)
		}
		output = append(output, ow.Output...)
	}
	outputCount := len(output)
	if errCnt != 0 {
		log.Printf("error: %d/%d events errored out\n", errCnt, matchupsCount)
		return nil, outputErr
	}
	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of %s with %d record(s) completed in %s\n", t.Scraper.Feed(), outputCount, diff)

	return output, nil
}

func (t *EventDataRunner[M, E]) Worker(wg *sync.WaitGroup, workerMatchups <-chan M, eventData chan<- sportscrape.EventDataOutput[E]) {
	for matchup := range workerMatchups {
		ow := t.Scraper.Scrape(matchup)
		eventData <- ow
		wg.Done()
	}
}

// MatchupRunnerConfig
type MatchupRunnerConfig[M any] struct {
	Scraper scraper.MatchupScraper[M]
}

// NewMatchupRunner Instantiates a new MatchupRunner
func NewMatchupRunner[M any](config MatchupRunnerConfig[M]) *MatchupRunner[M] {
	r := &MatchupRunner[M]{
		Scraper: config.Scraper,
	}
	r.Scraper.Init()
	return r
}

// MatchupRunner is a general matchup runner for scraping NBA, MLB, NCAAB, etc. matchup data.
type MatchupRunner[M any] struct {
	// Scraper
	Scraper scraper.MatchupScraper[M]
}

// Deprecated is a deprecation check for the feed/provider
func (r *MatchupRunner[M]) Deprecated() bool {
	provider := r.Scraper.Provider()
	feed := r.Scraper.Feed()
	if provider.Deprecated() {
		return true
	}
	return feed.Deprecated()
}

// Run gets all matchups
func (r *MatchupRunner[M]) Run() ([]M, error) {
	if r.Deprecated() {
		return nil, r.Scraper.Feed().Deprecation()
	}
	r.Scraper.Init()
	start := time.Now().UTC()
	ou := r.Scraper.Scrape()
	if ou.Context.Errors != 0 {
		log.Printf("error: %d event(s) errored out\n", ou.Context.Errors)
	}
	if ou.Error != nil {
		return nil, ou.Error
	}
	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of %s with %d record(s) completed in %s\n", r.Scraper.Feed(), len(ou.Output), diff)
	if ou.Context.Skips != 0 {
		log.Printf("WARNING: %d event(s) skipped\n", ou.Context.Skips)
	}
	return ou.Output, nil
}
