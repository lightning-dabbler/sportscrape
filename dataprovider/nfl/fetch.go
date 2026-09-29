package nfl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/lightning-dabbler/sportscrape/util/request"
)

// DefaultFetchAttempts is applied when a scraper doesn't set FetchAttempts explicitly (e.g. via the CLI flags).
// Retry-After handling below is standard HTTP behavior but is unverified against api.nfl.com: no 429 has been observed
// and responses carry no Retry-After or rate limit headers, so retries typically wait FetchRetryBackoff.
const (
	DefaultFetchAttempts = 3
)

// Fetcher holds the retry and auth configuration shared by all NFL scrapers
type Fetcher struct {
	// FetchAttempts is the max number of attempts per request before giving up: a data request and the token request
	// before it each get up to this many attempts. <= 0 falls back to DefaultFetchAttempts.
	FetchAttempts int
	// FetchRetryBackoff is the delay between retry attempts when the response has no Retry-After header.
	// <= 0 retries without a delay.
	FetchRetryBackoff time.Duration
	// Timeout is the timeout for each request. <= 0 falls back to request.DefaultGetTimeout.
	Timeout time.Duration
	// TokenProvider mints the api.nfl.com bearer token. nil falls back to DefaultTokenProvider.
	TokenProvider *TokenProvider
	// PersonNames caches player full names. nil falls back to DefaultPersonNames.
	PersonNames *PersonNames
}

func (f Fetcher) personNames() *PersonNames {
	if f.PersonNames != nil {
		return f.PersonNames
	}
	return DefaultPersonNames
}

func (f Fetcher) tokenProvider() *TokenProvider {
	if f.TokenProvider != nil {
		return f.TokenProvider
	}
	return DefaultTokenProvider
}

// StatusError is returned when a request receives a non 200 status
type StatusError struct {
	URL        string
	StatusCode int
	Status     string
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("Request to '%s' received a %s status", e.URL, e.Status)
}

// attemptsAndBackoff returns the fetcher's max attempts per request (<= 0 falls back to DefaultFetchAttempts)
// and the delay between retry attempts (< 0 is treated as no delay)
func (f Fetcher) attemptsAndBackoff() (int, time.Duration) {
	attempts := f.FetchAttempts
	if attempts <= 0 {
		attempts = DefaultFetchAttempts
	}
	return attempts, max(f.FetchRetryBackoff, 0)
}

// token returns an api.nfl.com bearer token from tokens. The token request gets its own attempts (separate from the data request's):
// failures with a network error, a 429 or a 5xx status are retried up to FetchAttempts times, waiting FetchRetryBackoff;
// any other status (e.g. 403 for invalid client credentials) is returned right away.
func (f Fetcher) token(tokens *TokenProvider) (string, error) {
	attempts, backoff := f.attemptsAndBackoff()
	var err error
	for attempt := 1; attempt <= attempts; attempt++ {
		var token string
		token, err = tokens.Token(f.Timeout)
		if err == nil {
			return token, nil
		}
		if !retryableTokenError(err) {
			break
		}
		if attempt < attempts {
			log.Printf("Attempt %d/%d minting api.nfl.com token failed (%v); retrying in %s\n", attempt, attempts, err, backoff)
			time.Sleep(backoff)
		}
	}
	return "", fmt.Errorf("api.nfl.com token: %w", err)
}

// fetchJSON retrieves the json response at url and unmarshals it into T.
// Fetches that fail with a network error, a 401 (the token is re-minted), a 429 or a 5xx status are retried up to attempts times,
// waiting the response's Retry-After (in seconds) when present, otherwise backoff.
// attempts <= 0 is treated as DefaultFetchAttempts; backoff <= 0 retries without a delay (a Retry-After is still honored).
// The token request before each attempt has its own attempts (see Fetcher.token); a token failure is returned right away.
// A non 200 status is returned as a *StatusError.
func fetchJSON[T any](url string, fetcher Fetcher) (T, error) {
	var payload T
	attempts, backoff := fetcher.attemptsAndBackoff()
	tokens := fetcher.tokenProvider()
	var body []byte
	var lastErr error
	for attempt := 1; attempt <= attempts; attempt++ {
		token, err := fetcher.token(tokens)
		if err != nil {
			return payload, err
		}
		var retryable bool
		var retryAfter time.Duration
		var statusCode int
		body, statusCode, retryable, retryAfter, lastErr = fetchBody(url, token, fetcher.Timeout)
		if lastErr == nil {
			break
		}
		if statusCode == http.StatusUnauthorized {
			tokens.Invalidate(token)
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

// retryableTokenError reports whether a failed token request is worth retrying: network and response errors, 429 and 5xx are;
// any other status (e.g. 403 for invalid client credentials) isn't
func retryableTokenError(err error) bool {
	var statusErr *StatusError
	if errors.As(err, &statusErr) {
		return statusErr.StatusCode == http.StatusTooManyRequests || statusErr.StatusCode >= 500
	}
	return true
}

// fetchBody performs a single authorized GET request for url.
// Returns the status code (0 on a network error), whether a failed request is retryable (network error, 401, 429 or 5xx)
// and the response's Retry-After (0 when absent or invalid).
func fetchBody(url string, token string, timeout time.Duration) ([]byte, int, bool, time.Duration, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, false, 0, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	log.Printf("Fetching from %s\n", url)
	response, err := request.NewClient(timeout).Do(req)
	if err != nil {
		return nil, 0, true, 0, fmt.Errorf("HTTP Error at %s: %w", url, err)
	}
	defer response.Body.Close()
	statusErr := &StatusError{URL: url, StatusCode: response.StatusCode, Status: response.Status}
	switch response.StatusCode {
	case http.StatusOK:
		body, err := io.ReadAll(response.Body)
		return body, response.StatusCode, true, 0, err
	case http.StatusTooManyRequests:
		var retryAfter time.Duration
		if seconds, err := strconv.Atoi(response.Header.Get("Retry-After")); err == nil && seconds > 0 {
			retryAfter = time.Duration(seconds) * time.Second
		}
		return nil, response.StatusCode, true, retryAfter, statusErr
	default:
		retryable := response.StatusCode == http.StatusUnauthorized || response.StatusCode >= 500
		return nil, response.StatusCode, retryable, 0, statusErr
	}
}
