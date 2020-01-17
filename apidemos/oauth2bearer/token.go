// Package oauth2bearer provides a simple set of wrappers for oauth2
// bearer tokens
package oauth2bearer

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// refreshMarginSeconds controls how many seconds before the expected expiry
// of a bearer token we decide to refresh it
const refreshMarginSeconds = 10

// BearerToken is the blob of data returned, in JSON, when
// you request a new bearer token
type BearerToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	expiryTime  time.Time
	source      *BearerTokenSource
}

// ClientCredentials is an id/secret pair of credentials for
// authentication
type ClientCredentials struct {
	ClientID     string
	ClientSecret string
}

// BearerTokenSource describes a place to get tokens
type BearerTokenSource struct {
	Credentials ClientCredentials
	ScopesList  []string
	URL         string
}

// Token retrieves a new fresh bearer token from the
// requested source
func (source *BearerTokenSource) Token() (*BearerToken, error) {
	scopeString := strings.Join(source.ScopesList, " ")

	queryData := url.Values{}
	queryData.Set("grant_type", "client_credentials")
	queryData.Set("scope", scopeString)

	req, err := http.NewRequest("POST", source.URL, strings.NewReader(queryData.Encode()))
	if err != nil {
		log.Panic("error getting bearer token")
	}

	req.SetBasicAuth(source.Credentials.ClientID, source.Credentials.ClientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyRaw, err := ioutil.ReadAll(resp.Body)
	var token BearerToken
	json.Unmarshal(bodyRaw, &token)
	token.expiryTime = time.Now().Add(time.Second*time.Duration(token.ExpiresIn) -
		time.Second*time.Duration(refreshMarginSeconds))
	token.source = source
	return &token, nil
}

// Type returns Bearer for this token
func (t *BearerToken) Type() string {
	return "Bearer"
}

// FreshAuthString returns a valid, properly refreshed auth string
func (t *BearerToken) FreshAuthString() string {
	t.Refresh()
	authString := "Bearer " + t.AccessToken
	return authString
}

// SetAuthHeader puts this bearer token's authorization into the request
func (t *BearerToken) SetAuthHeader(req *http.Request) {
	req.Header.Set("Authorization", t.FreshAuthString())
}

// Valid indicates whether this token is still valid
// or should be refreshed
func (t *BearerToken) Valid() bool {
	if t.expiryTime.Before(time.Now()) {
		return false
	}
	return true
}

func (t *BearerToken) doRefresh() {
	nt, _ := t.source.Token()
	if t.TokenType != nt.TokenType {
		log.Panic("Somehow token type changed")
	}
	t.AccessToken = nt.AccessToken
	t.ExpiresIn = nt.ExpiresIn
	t.expiryTime = nt.expiryTime
}

// Refresh refreshes the token if it looks necessary
func (t *BearerToken) Refresh() {
	if !t.Valid() {
		t.doRefresh()
	}
}

// NewHTTPRequest returns a new authenticated http request
func (t *BearerToken) NewHTTPRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// NewAuthenticatedClient builds a new client with auth built-in
func (t *BearerToken) NewAuthenticatedClient() *AuthenticatedClient {
	return t.source.NewAuthenticatedClient()
}

// eof
