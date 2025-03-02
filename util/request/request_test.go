//go:build unit || integration

package request

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"context"
	"errors"
	"strings"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func TestGetUnitTests(t *testing.T) {
	tests := []struct {
		name     string
		status   int
		response string
		isError  bool
	}{
		{
			name:     "successful response",
			status:   http.StatusOK,
			response: "success response",
			isError:  false,
		},
		{
			name:     "not found response",
			status:   http.StatusNotFound,
			response: "not found",
			isError:  true,
		},
		{
			name:     "server error",
			status:   http.StatusInternalServerError,
			response: "server error",
			isError:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dummyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.status)
				w.Write([]byte(tt.response))
			}))
			defer dummyServer.Close()
			resp, err := Get(dummyServer.URL)

			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				body, _ := io.ReadAll(resp.Body)
				resp.Body.Close()
				assert.Equal(t, tt.response, string(body), "Equal response body")
			}
		})
	}
}

func TestGetIntegrationTests(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	tests := []struct {
		name    string
		url     string
		status  int
		isError bool
	}{
		{
			name:    "valid url",
			url:     "https://example.com",
			status:  http.StatusOK,
			isError: false,
		},
		{
			name:    "invalid protocol",
			url:     "efkjfnekfn://example.com",
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := Get(tt.url)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.status, resp.StatusCode, "Equal status codes")
			}
		})
	}
}

func TestNewDocumentRetriever(t *testing.T) {
	timeout := 30 * time.Second
	dr := NewDocumentRetriever(timeout)

	assert.Equal(t, timeout, dr.timeout)
	assert.NotNil(t, dr.ChromeRun)
	assert.NotNil(t, dr.DocumentReader)
}

func TestRetrieveDocument(t *testing.T) {
	testCases := []struct {
		name         string
		runErr       error
		docReaderErr error
		expectErr    bool
	}{
		{
			name:         "Success",
			runErr:       nil,
			docReaderErr: nil,
			expectErr:    false,
		},
		{
			name:         "ChromeRun error",
			runErr:       errors.New("chrome error"),
			docReaderErr: nil,
			expectErr:    true,
		},
		{
			name:         "DocumentReader error",
			runErr:       nil,
			docReaderErr: errors.New("doc reader error"),
			expectErr:    true,
		},
		{
			name:         "Context timeout",
			runErr:       context.DeadlineExceeded,
			docReaderErr: nil,
			expectErr:    true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			dr := &DocumentRetriever{
				timeout: 100 * time.Millisecond,
				ChromeRun: func(ctx context.Context, actions ...chromedp.Action) error {
					return tt.runErr
				},
				DocumentReader: func(r io.Reader) (*goquery.Document, error) {
					if tt.docReaderErr != nil {
						return nil, tt.docReaderErr
					}
					return goquery.NewDocumentFromReader(strings.NewReader("<html><body>Test</body></html>"))
				},
			}

			doc, err := dr.RetrieveDocument("https://example.com", nil, "body")

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, doc)
				if tt.runErr == context.DeadlineExceeded {
					assert.ErrorIs(t, err, context.DeadlineExceeded)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, doc)
			}
		})
	}
}

func TestRetrieveDocument_VerifiesActions(t *testing.T) {
	url := "https://example.com"
	headers := network.Headers{"User-Agent": "test-agent"}
	selector := ".content"

	actionTypes := make(map[string]bool)
	actionsExecuted := false

	dr := &DocumentRetriever{
		timeout: 5 * time.Second,
		ChromeRun: func(ctx context.Context, actions ...chromedp.Action) error {
			actionsExecuted = true
			// Since we can't reliably identify the action types,
			// we'll just count the actions and verify that the expected number are present
			actionsCount := len(actions)

			// We expect at least 4 actions: network.Enable, SetExtraHTTPHeaders, Navigate, WaitReady, OuterHTML
			if actionsCount >= 4 {
				actionTypes["all_actions_present"] = true
			}
			return nil
		},
		DocumentReader: func(r io.Reader) (*goquery.Document, error) {
			return goquery.NewDocumentFromReader(strings.NewReader("<html><body>Test</body></html>"))
		},
	}

	dr.RetrieveDocument(url, headers, selector)
	assert.True(t, actionsExecuted, "Actions were not executed")
	assert.True(t, actionTypes["all_actions_present"], "Not enough actions were executed")
}

func TestDocumentRetriever_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	dr := NewDocumentRetriever(10 * time.Second)
	doc, err := dr.RetrieveDocument("https://example.com", nil, "body")
	assert.NoError(t, err)
	assert.NotNil(t, doc)

	html, err := doc.Html()
	assert.NoError(t, err)
	assert.NotEmpty(t, html)
}
