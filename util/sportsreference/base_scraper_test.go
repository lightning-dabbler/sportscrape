//go:build unit

package sportsreference

import (
	"testing"
	"time"
)

// TestBaseScraperStructureAndOptions verifies the BaseScraper structure and its configuration options
func TestBaseScraperStructureAndOptions(t *testing.T) {
	testCases := []struct {
		name    string
		timeout time.Duration
		debug   bool
	}{
		{"Default", 0, false},
		{"WithTimeout", 5 * time.Second, false},
		{"WithDebug", 0, true},
		{"WithTimeoutAndDebug", 5 * time.Second, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			scraper := BaseScraper{
				Timeout: tc.timeout,
				Debug:   tc.debug,
			}

			if scraper.Timeout != tc.timeout {
				t.Errorf("Expected Timeout to be %v, got %v", tc.timeout, runner.Timeout)
			}

			if scraper.Debug != tc.debug {
				t.Errorf("Expected Debug to be %v, got %v", tc.debug, runner.Debug)
			}
		})
	}
}
