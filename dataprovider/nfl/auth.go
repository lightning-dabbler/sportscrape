package nfl

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/util/request"
)

// nfl.com's public web (desktop) client credentials, used to mint anonymous api.nfl.com tokens.
// Found in https://p.nfltags.com/nfl/live/NflUmdComponents.NFLToken.js
const (
	DefaultClientKey    = "4cFUW6DmwJpzT9L7LrG3qRAcABG5s04g"
	DefaultClientSecret = "CZuvCL49d9OwfGsR"
)

// Environment variables overriding the nfl.com client credentials (e.g. if nfl.com rotates them).
// Both must be set to take effect.
const (
	ClientKeyEnv    = "SPORTSCRAPE_NFL_DOT_COM_CLIENT_KEY"
	ClientSecretEnv = "SPORTSCRAPE_NFL_DOT_COM_CLIENT_SECRET"
)

// tokenExpiryBuffer is how long before its expiry a token is considered expired
const tokenExpiryBuffer = time.Minute

// DefaultTokenProvider is shared by every Fetcher without a TokenProvider, so a token is minted once and reused across scrapers.
// Its credentials come from ClientKeyEnv/ClientSecretEnv, or DefaultClientKey/DefaultClientSecret.
var DefaultTokenProvider = NewTokenProvider("", "")

// TokenProvider mints and caches the bearer token required by api.nfl.com. Safe for concurrent use.
type TokenProvider struct {
	// ClientKey and ClientSecret are used when both are set; otherwise the credentials come from
	// ClientKeyEnv/ClientSecretEnv when both are set, otherwise DefaultClientKey/DefaultClientSecret
	ClientKey    string
	ClientSecret string

	mu          sync.Mutex
	deviceID    string
	accessToken string
	expiresAt   time.Time
}

// NewTokenProvider creates a TokenProvider for the given nfl.com client credentials
func NewTokenProvider(clientKey, clientSecret string) *TokenProvider {
	return &TokenProvider{ClientKey: clientKey, ClientSecret: clientSecret}
}

// credentials returns the client key and secret to mint tokens with, always as a pair from a single source:
// the provider's ClientKey/ClientSecret, else ClientKeyEnv/ClientSecretEnv, else the defaults
func (p *TokenProvider) credentials() (string, string) {
	if p.ClientKey != "" && p.ClientSecret != "" {
		return p.ClientKey, p.ClientSecret
	}
	key, secret := os.Getenv(ClientKeyEnv), os.Getenv(ClientSecretEnv)
	if key != "" && secret != "" {
		return key, secret
	}
	return DefaultClientKey, DefaultClientSecret
}

// Token returns the cached access token, minting a new one when there is none or it is about to expire.
// timeout is the timeout for the token request (<= 0 falls back to request.DefaultGetTimeout).
func (p *TokenProvider) Token(timeout time.Duration) (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.accessToken != "" && time.Now().Before(p.expiresAt.Add(-tokenExpiryBuffer)) {
		return p.accessToken, nil
	}
	if p.deviceID == "" {
		deviceID, err := newDeviceID()
		if err != nil {
			return "", err
		}
		p.deviceID = deviceID
	}
	token, err := p.mint(timeout)
	if err != nil {
		return "", err
	}
	p.accessToken = token.AccessToken
	p.expiresAt = time.Unix(token.ExpiresIn, 0)
	return p.accessToken, nil
}

// Invalidate discards token (e.g. after a 401) so the next Token call mints a new one.
// A token other than the cached one is ignored, as it has already been replaced.
func (p *TokenProvider) Invalidate(token string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.accessToken == token {
		p.accessToken = ""
	}
}

// mint requests a new anonymous token
func (p *TokenProvider) mint(timeout time.Duration) (jsonresponse.Token, error) {
	var token jsonresponse.Token
	clientKey, clientSecret := p.credentials()
	body, err := json.Marshal(map[string]any{
		"clientKey":    clientKey,
		"clientSecret": clientSecret,
		"deviceId":     p.deviceID,
		"deviceInfo":   "",
		"networkType":  "other",
		"nflClaims":    nil,
	})
	if err != nil {
		return token, err
	}
	url := ConstructTokenURL()
	log.Printf("Fetching token from %s\n", url)
	response, err := request.NewClient(timeout).Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return token, fmt.Errorf("HTTP Error at %s: %w", url, err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return token, &StatusError{URL: url, StatusCode: response.StatusCode, Status: response.Status}
	}
	payload, err := io.ReadAll(response.Body)
	if err != nil {
		return token, err
	}
	if err := json.Unmarshal(payload, &token); err != nil {
		return token, err
	}
	if token.AccessToken == "" {
		return token, fmt.Errorf("no accessToken in response from '%s'", url)
	}
	return token, nil
}

// newDeviceID returns a random (version 4) UUID
func newDeviceID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16]), nil
}
