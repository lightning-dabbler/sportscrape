//go:build unit

package sportsreferenceutil

import (
	"reflect"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	mocksportsreferenceutil "github.com/lightning-dabbler/sportscrape/util/sportsreference/mocks"
)

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

	testCases := []struct {
		name           string
		concurrency    int
		matchups       []interface{}
		expectedResult []int
		expectedLength int
	}{
		{
			name:           "No matchups",
			concurrency:    2,
			matchups:       []interface{}{},
			expectedLength: 0,
		},
		{
			name:           "Single matchup",
			concurrency:    2,
			matchups:       []interface{}{1},
			expectedResult: []int{2, 2},
			expectedLength: 2,
		},
		{
			name:           "Multiple matchups",
			concurrency:    2,
			matchups:       []interface{}{1, 2, 3},
			expectedResult: []int{2, 2, 4, 4, 6, 6},
			expectedLength: 6,
		},
		{
			name:           "Default concurrency",
			concurrency:    0, // Should use runtime.NumCPU()
			matchups:       []interface{}{1, 2, 3, 4},
			expectedResult: []int{2, 2, 4, 4, 6, 6, 8, 8},
			expectedLength: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockprocessor := &mocksportsreferenceutil.MockBoxScoreProcessor{}
			for _, matchup := range tc.matchups {
				mockprocessor.EXPECT().GetSegmentBoxScoreStats(matchup).Return([]interface{}{2 * matchup.(int), 2 * matchup.(int)})
			}

			runner := &BoxScoreRunner{
				Runner: Runner{
					Timeout: 1 * time.Second,
				},
				Concurrency: tc.concurrency,
				Processor:   mockprocessor,
			}

			results := runner.GetBoxScoresStats(tc.matchups...)
			n_results := len(results)
			ints := make([]int, n_results)
			for idx, item := range results {
				ints[idx] = item.(int)
			}
			sort.Ints(ints)

			if n_results != tc.expectedLength {
				t.Errorf("Expected %d results, got %d", tc.expectedLength, n_results)
			}
			if tc.expectedLength != 0 {
				assert.Equal(t, tc.expectedResult, ints, "Equal output")
			}

		})
	}
}

// TestBoxScoreRunnerWorker tests the BoxScoreRunner.Worker method directly
func TestBoxScoreRunnerWorker(t *testing.T) {
	// Test data
	matchups := []interface{}{1, "test", map[string]string{"key": "value"}}

	mockprocessor := &mocksportsreferenceutil.MockBoxScoreProcessor{}

	for _, matchup := range matchups {
		mockprocessor.EXPECT().GetSegmentBoxScoreStats(matchup).Return([]interface{}{matchup})
	}
	runner := &BoxScoreRunner{
		Processor: mockprocessor,
	}

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
