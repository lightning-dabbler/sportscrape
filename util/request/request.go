package request

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

// Get performs a GET request
// Returns an http response
func Get(url string) (*http.Response, error) {
	fmt.Printf("Fetching from %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP Error at %s: %w", url, err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Request to '%s' received a %s status", url, resp.Status)
	}
	return resp, nil
}

// DocumentRetriever
type DocumentRetriever struct {
	timeout        time.Duration
	debug          bool
	ChromeRun      func(ctx context.Context, actions ...chromedp.Action) error
	DocumentReader func(r io.Reader) (*goquery.Document, error)
}

// NewDocumentRetriever accepts a timeout
// returns DocumentRetriever
func NewDocumentRetriever(timeout time.Duration) *DocumentRetriever {
	return &DocumentRetriever{
		timeout:        timeout,
		ChromeRun:      chromedp.Run,
		DocumentReader: goquery.NewDocumentFromReader,
	}
}

// RetrieveDocument accepts a url, network headers, and selector to indicate when the page is ready
// It uses chromium to fetch a web page from the url
// Returns *goquery.Document and error
func (dr *DocumentRetriever) RetrieveDocument(url string, networkHeaders network.Headers, waitReadySelector string) (*goquery.Document, error) {
	var contextOptions []chromedp.ContextOption

	if dr.debug {
		contextOptions = append(contextOptions, chromedp.WithDebugf(log.Printf))
	}
	ctx, cancel := chromedp.NewContext(context.Background(), contextOptions...)

	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, dr.timeout)
	defer cancel()
	fmt.Println("Retrieving document from: " + url)
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
