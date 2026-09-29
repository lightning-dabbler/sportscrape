package nfl

import (
	"fmt"
	"log"
	"sync"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
)

// DefaultPersonNames is shared by every Fetcher without PersonNames, so a person is fetched once per run across scrapers
var DefaultPersonNames = NewPersonNames()

// PersonNames resolves and caches person IDs to full names (displayName e.g. Dak Prescott). Safe for concurrent use.
// Failed lookups are cached too, so a person whose name can't be retrieved is only requested (and logged) once.
// That includes failures that were transient (e.g. a network error or token failure that outlasted the retries):
// the person gets the abbreviated name for the rest of the run.
type PersonNames struct {
	mu     sync.Mutex
	names  map[string]string
	failed map[string]error
}

// NewPersonNames creates an empty PersonNames cache
func NewPersonNames() *PersonNames {
	return &PersonNames{names: make(map[string]string), failed: make(map[string]error)}
}

// Name returns the person's full name, fetching https://api.nfl.com/football/v2/persons/{person_id} when it isn't cached.
// Returns the cached error when an earlier lookup for the person failed.
func (c *PersonNames) Name(personID string, fetcher Fetcher) (string, error) {
	c.mu.Lock()
	name, exists := c.names[personID]
	failure := c.failed[personID]
	c.mu.Unlock()
	if exists {
		return name, nil
	}
	if failure != nil {
		return "", failure
	}
	name, err := fetchPersonName(personID, fetcher)
	c.mu.Lock()
	defer c.mu.Unlock()
	if err != nil {
		if _, recorded := c.failed[personID]; !recorded {
			log.Printf("full name for person %s unavailable, falling back to the abbreviated name: %v\n", personID, err)
		}
		c.failed[personID] = err
		return "", err
	}
	c.names[personID] = name
	return name, nil
}

// fetchPersonName retrieves the person's displayName from https://api.nfl.com/football/v2/persons/{person_id}
func fetchPersonName(personID string, fetcher Fetcher) (string, error) {
	person, err := fetchJSON[jsonresponse.Person](ConstructPersonURL(personID), fetcher)
	if err != nil {
		return "", err
	}
	if person.DisplayName == "" {
		return "", fmt.Errorf("no displayName for person %s", personID)
	}
	return person.DisplayName, nil
}

// playerName returns the player's full name through the fetcher's PersonNames cache.
// Falls back to shortName (the abbreviated name e.g. D.Prescott) when the full name can't be retrieved (logged once per person by PersonNames).
func (f Fetcher) playerName(personID string, shortName string) string {
	if personID == "" {
		return shortName
	}
	name, err := f.personNames().Name(personID, f)
	if err != nil {
		return shortName
	}
	return name
}
