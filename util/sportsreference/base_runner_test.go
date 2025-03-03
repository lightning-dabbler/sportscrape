//go:build unit

package sportsreferenceutil

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

// Simple implementation of BoxScoreProcessor for testing
type TestBoxScoreProcessor struct{}

func (p *TestBoxScoreProcessor) GetSegmentBoxScoreStats(matchup interface{}) []interface{} {
	// Simple implementation that returns the matchup in a slice
	return []interface{}{matchup}
}

// TestMatchupRunnerStructure verifies MatchupRunner embeds Runner properly
func TestMatchupRunnerStructure(t *testing.T) {
	// Verify MatchupRunner embeds Runner
	runner := MatchupRunner{
		Runner: Runner{
			Timeout: 5 * time.Second,
			Debug:   true,
		},
	}

	// Check that Runner fields are accessible
	if runner.Timeout != 5*time.Second {
		t.Errorf("Expected Timeout to be 5s, got %v", runner.Timeout)
	}

	if !runner.Debug {
		t.Errorf("Expected Debug to be true, got false")
	}

	// Since GetMatchups is likely a placeholder meant to be overridden,
	// we'll just verify it exists and returns a non-nil value
	result := runner.GetMatchups("2023-01-01")
	if result == nil {
		t.Errorf("GetMatchups should return non-nil, even if empty")
	}
}

// TestBoxScoreRunnerGetBoxScoresStats tests the BoxScoreRunner.GetBoxScoresStats method
func TestBoxScoreRunnerGetBoxScoresStats(t *testing.T) {
	processor := &TestBoxScoreProcessor{}

	testCases := []struct {
		name        string
		concurrency int
		matchups    []interface{}
		expected    int
	}{
		{
			name:        "No matchups",
			concurrency: 2,
			matchups:    []interface{}{},
			expected:    0,
		},
		{
			name:        "Single matchup",
			concurrency: 2,
			matchups:    []interface{}{1},
			expected:    1,
		},
		{
			name:        "Multiple matchups",
			concurrency: 2,
			matchups:    []interface{}{1, 2, 3},
			expected:    3,
		},
		{
			name:        "Default concurrency",
			concurrency: 0, // Should use runtime.NumCPU()
			matchups:    []interface{}{1, 2, 3, 4},
			expected:    4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			runner := &BoxScoreRunner{
				Runner: Runner{
					Timeout: 1 * time.Second,
				},
				Concurrency: tc.concurrency,
				Processor:   processor,
			}

			results := runner.GetBoxScoresStats(tc.matchups...)

			if len(results) != tc.expected {
				t.Errorf("Expected %d results, got %d", tc.expected, len(results))
			}

			// Verify each matchup was processed correctly
			for i, matchup := range tc.matchups {
				if i < len(results) && !reflect.DeepEqual(results[i], matchup) {
					t.Errorf("Result %d doesn't match input: expected %v, got %v", i, matchup, results[i])
				}
			}
		})
	}
}

// TestBoxScoreRunnerWorker tests the BoxScoreRunner.Worker method directly
func TestBoxScoreRunnerWorker(t *testing.T) {
	processor := &TestBoxScoreProcessor{}
	runner := &BoxScoreRunner{
		Processor: processor,
	}

	// Test data
	matchups := []interface{}{1, "test", map[string]string{"key": "value"}}

	// Setup channels and wait group
	var wg sync.WaitGroup
	workerMatchups := make(chan interface{}, len(matchups))
	boxScoreStats := make(chan []interface{}, len(matchups))

	// Start worker
	go runner.Worker(&wg, workerMatchups, boxScoreStats)

	// Send matchups to worker
	for _, m := range matchups {
		wg.Add(1)
		workerMatchups <- m
	}

	// Wait for processing to complete
	wg.Wait()
	close(workerMatchups)
	close(boxScoreStats)

	// Collect results
	var results []interface{}
	for stats := range boxScoreStats {
		results = append(results, stats...)
	}

	// Verify results
	if len(results) != len(matchups) {
		t.Errorf("Expected %d results, got %d", len(matchups), len(results))
	}

	// Verify each result matches its input
	for i, m := range matchups {
		if i < len(results) && !reflect.DeepEqual(results[i], m) {
			t.Errorf("Result %d doesn't match input: expected %v, got %v", i, m, results[i])
		}
	}
}

// TestRunnerStructureAndOptions verifies the Runner structure and its configuration options
func TestRunnerStructureAndOptions(t *testing.T) {
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
			runner := Runner{
				Timeout: tc.timeout,
				Debug:   tc.debug,
			}

			if runner.Timeout != tc.timeout {
				t.Errorf("Expected Timeout to be %v, got %v", tc.timeout, runner.Timeout)
			}

			if runner.Debug != tc.debug {
				t.Errorf("Expected Debug to be %v, got %v", tc.debug, runner.Debug)
			}
		})
	}
}
