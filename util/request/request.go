package request

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

// DefaultGetTimeout is the timeout applied to GET requests when no (or a non-positive) timeout is given
const DefaultGetTimeout = 120 * time.Second

// NewClient returns an http client whose requests time out after timeout.
// timeout <= 0 is treated as DefaultGetTimeout.
func NewClient(timeout time.Duration) *http.Client {
	if timeout <= 0 {
		timeout = DefaultGetTimeout
	}
	return &http.Client{Timeout: timeout}
}

// Get performs a GET request that times out after DefaultGetTimeout
// Returns an http response, or an error if the request failed or the response status is not 200
func Get(url string) (*http.Response, error) {
	resp, err := GetWithTimeout(url, 0)
	if err != nil {
		return nil, err
	}
	if err := CheckStatus(url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetWithTimeout performs a GET request that times out after timeout (<= 0 is treated as DefaultGetTimeout)
// Returns the http response whatever its status code; evaluating the status (e.g. via CheckStatus) is up to the caller
func GetWithTimeout(url string, timeout time.Duration) (*http.Response, error) {
	log.Printf("Fetching from %s\n", url)
	resp, err := NewClient(timeout).Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP Error at %s: %w", url, err)
	}
	return resp, nil
}

// CheckStatus returns an error, and closes the response body, if resp's status is not 200
func CheckStatus(url string, resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("Request to '%s' received a %s status", url, resp.Status)
	}
	return nil
}

// DocumentRetriever
type DocumentRetriever struct {
	// timeout is the context timeout for ChromeRun operations
	Timeout time.Duration
	// debug enables verbose logging of Chrome operations when true
	Debug bool
	// ChromeRun executes Chrome DevTools Protocol actions in a given context
	ChromeRun func(ctx context.Context, actions ...chromedp.Action) error
	// DocumentReader parses HTML content into a goquery Document
	DocumentReader func(r io.Reader) (*goquery.Document, error)
}

// RetrieverOption defines a function that configures a DocumentRetriever.
// This is part of the functional options pattern for configuration.
type RetrieverOption func(*DocumentRetriever)

// WithTimeout returns a RetrieverOption that sets the timeout duration for document retrieval operations.
//
// Parameter:
//   - timeout: The duration after which document retrieval operations will be cancelled
func WithTimeout(timeout time.Duration) RetrieverOption {
	return func(dr *DocumentRetriever) {
		dr.Timeout = timeout
	}
}

// WithDebug returns a RetrieverOption that enables or disables debug logging.
//
// Parameter:
//   - debug: When true, enables verbose logging of Chrome operations
func WithDebug(debug bool) RetrieverOption {
	return func(dr *DocumentRetriever) {
		dr.Debug = debug
	}
}

// NewDocumentRetriever creates a new DocumentRetriever with default settings,
// which can be overridden by the provided options.
//
// By default, the retriever uses:
//   - 1 minute timeout
//   - Debug logging disabled
//   - Standard chromedp.Run for Chrome operations
//   - Standard goquery document parser
//
// Parameters:
//   - options: A variadic list of RetrieverOption functions that modify the default configuration
//
// Returns a pointer to the configured DocumentRetriever
func NewDocumentRetriever(options ...RetrieverOption) *DocumentRetriever {
	// Set defaults
	dr := &DocumentRetriever{
		Timeout:        1 * time.Minute,
		Debug:          false,
		ChromeRun:      chromedp.Run,
		DocumentReader: goquery.NewDocumentFromReader,
	}

	// Apply options
	for _, option := range options {
		option(dr)
	}

	return dr
}

// RetrieveDocument fetches and parses a web document from the specified URL.
// It uses headless Chrome to load the page, waits for the element matching waitReadySelector
// to be ready, and returns the parsed document.
//
// Parameters:
//   - url: The web page URL to retrieve
//   - networkHeaders: HTTP headers to include with the request
//   - waitReadySelector: CSS selector of element that indicates page is fully loaded
//
// Returns:
//   - *goquery.Document: The parsed document
//   - error: Any error encountered during fetching or parsing
func (dr *DocumentRetriever) RetrieveDocument(url string, networkHeaders network.Headers, waitReadySelector string) (*goquery.Document, error) {
	var contextOptions []chromedp.ContextOption

	if dr.Debug {
		contextOptions = append(contextOptions, chromedp.WithDebugf(log.Printf))
	}
	ctx, cancel := chromedp.NewContext(context.Background(), contextOptions...)

	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, dr.Timeout)
	defer cancel()
	slog.Info("Retrieving document", "url", url)
	var outer string
	if err := dr.ChromeRun(ctx,
		network.Enable(),
		network.SetExtraHTTPHeaders(networkHeaders),
		chromedp.Navigate(url),
		chromedp.WaitReady(waitReadySelector),
		chromedp.OuterHTML(waitReadySelector, &outer, chromedp.ByQuery),
	); err != nil {
		return nil, fmt.Errorf("Error fetching document from %s: %w", url, err)
	}
	myReader := strings.NewReader(outer)
	doc, err := dr.DocumentReader(myReader)
	if err != nil {
		return nil, fmt.Errorf("Error reading document from %s: %w", url, err)
	}
	return doc, nil
}
