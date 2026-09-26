package nhl

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/lightning-dabbler/sportscrape/util/request"
)

// api-web.nhle.com rate limits bursts of requests with a 429 (and a Retry-After header in seconds),
// so fetchJSON retries failed fetches. DefaultFetchAttempts is applied when a scraper doesn't
// set FetchAttempts explicitly (e.g. via the CLI flags).
const (
	DefaultFetchAttempts = 3
)

// Fetcher holds the retry configuration shared by all NHL scrapers
type Fetcher struct {
	// FetchAttempts is the max number of times to fetch a url before giving up.
	// <= 0 falls back to DefaultFetchAttempts.
	FetchAttempts int
	// FetchRetryBackoff is the delay between retry attempts when the response has no Retry-After header.
	// <= 0 retries without a delay.
	FetchRetryBackoff time.Duration
	// Timeout is the timeout for each request. <= 0 falls back to request.DefaultGetTimeout.
	Timeout time.Duration
}

// fetchJSON retrieves the json response at url and unmarshals it into T.
// Fetches that fail with a network error, a 429 or a 5xx status are retried up to attempts times,
// waiting the response's Retry-After (in seconds) when present, otherwise backoff.
// attempts <= 0 is treated as DefaultFetchAttempts; backoff <= 0 retries without a delay (a Retry-After is still honored).
func fetchJSON[T any](url string, fetcher Fetcher) (T, error) {
	var payload T
	attempts := fetcher.FetchAttempts
	if attempts <= 0 {
		attempts = DefaultFetchAttempts
	}
	backoff := fetcher.FetchRetryBackoff
	if backoff < 0 {
		backoff = 0
	}
	var body []byte
	var lastErr error
	for attempt := 1; attempt <= attempts; attempt++ {
		var retryable bool
		var retryAfter time.Duration
		body, retryable, retryAfter, lastErr = fetchBody(url, fetcher.Timeout)
		if lastErr == nil {
			break
		}
		if !retryable {
			return payload, lastErr
		}
		wait := backoff
		if retryAfter > 0 {
			wait = retryAfter
		}
		if attempt < attempts {
			log.Printf("Attempt %d/%d fetching %s failed (%v); retrying in %s\n", attempt, attempts, url, lastErr, wait)
			time.Sleep(wait)
		}
	}
	if lastErr != nil {
		return payload, lastErr
	}
	err := json.Unmarshal(body, &payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

// fetchBody performs a single GET request for url.
// Returns whether a failed request is retryable (network error, 429 or 5xx)
// and the response's Retry-After (0 when absent or invalid).
func fetchBody(url string, timeout time.Duration) ([]byte, bool, time.Duration, error) {
	response, err := request.GetWithTimeout(url, timeout)
	if err != nil {
		return nil, true, 0, err
	}
	defer response.Body.Close()
	switch {
	case response.StatusCode == http.StatusOK:
		body, err := io.ReadAll(response.Body)
		return body, true, 0, err
	case response.StatusCode == http.StatusTooManyRequests:
		var retryAfter time.Duration
		if seconds, err := strconv.Atoi(response.Header.Get("Retry-After")); err == nil && seconds > 0 {
			retryAfter = time.Duration(seconds) * time.Second
		}
		return nil, true, retryAfter, fmt.Errorf("Request to '%s' received a %s status", url, response.Status)
	default:
		retryable := response.StatusCode >= 500
		return nil, retryable, 0, fmt.Errorf("Request to '%s' received a %s status", url, response.Status)
	}
}
