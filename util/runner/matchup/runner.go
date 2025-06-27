package matchup

import (
	"log"
	"time"
)

type OutputWrapper struct {
	Output  []interface{}
	Context Context
}

type Context struct {
	Errors int
	Skips  int
	Total  int
}

// RunnerOption
type RunnerOption func(*Runner)

// RunnerName option
func RunnerName(name string) RunnerOption {
	return func(r *Runner) {
		r.Name = name
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

	r.Scraper.Init()
	return r
}

// Runner is a general matchup runner for scraping NBA, MLB, NCAAB, etc. matchup data.
type Runner struct {
	// Name
	Name string
	// Scraper
	Scraper Scraper
}

// RunMatchupsScraper gets all matchups of a League and segment ID
func (r *Runner) RunMatchupsScraper() []interface{} {
	start := time.Now().UTC()
	ou := r.Scraper.Scrape()
	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of %s with %d record(s) completed in %s\n", r.Name, len(ou.Output), diff)
	if ou.Context.Errors != 0 {
		log.Printf("WARNING: %d events errored out\n", ou.Context.Errors)
	}

	if ou.Context.Skips != 0 {
		log.Printf("WARNING: %d events skipped\n", ou.Context.Skips)
	}
	return ou.Output
}
