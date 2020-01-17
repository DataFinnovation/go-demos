package access

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/DataFinnovation/go-demos/apidemos/oauth2bearer"
)

// DFAccess is the top-level DF API access object
// It is meant to be opaque -- these just get passed into
// the data access wrapper functions to sort out authentication.
type DFAccess struct {
	token            *oauth2bearer.BearerToken
	requestHeaderMap map[string]string
}

// NewDFAccessDefaultScopes builds a new DFAccess object with
// the default scopes
func NewDFAccessDefaultScopes() *DFAccess {
	return NewDFAccess([]string{"clientapi/basicsearch", "clientapi/advancedsearch"})
}

// NewDFAccess builds a DFAccess object given a list of scopes
func NewDFAccess(scopes []string) *DFAccess {
	var dfa DFAccess
	dfa.token = getToken(scopes)
	dfa.requestHeaderMap = map[string]string{
		"x-api-key":    os.Getenv("DF_API_KEY"),
		"Content-Type": "application/json",
	}
	return &dfa
}

func tokenURL() string {
	return Getenv("DF_TOKEN_URL", "https://apiauth.dfnapp.com/oauth2/token")
}

func apiURLStub() string {
	return Getenv("DF_API_URL_STUB", "https://clientapi.dfnapp.com/")
}

// GetToken returns a single oauth2 token
func getToken(scopes []string) *oauth2bearer.BearerToken {
	// note: there is no sensible default for these parameters
	clientID := os.Getenv("DF_CLIENT_ID")
	clientSecret := os.Getenv("DF_CLIENT_SECRET")
	creds := oauth2bearer.ClientCredentials{ClientID: clientID, ClientSecret: clientSecret}
	source := oauth2bearer.TokenSource{
		Credentials: creds,
		ScopesList:  scopes,
		URL:         tokenURL(),
	}
	return oauth2bearer.GetNewBearerToken(source)
}

func newRequest(method, url string, body io.Reader, dfa *DFAccess) *http.Request {
	req, err := dfa.token.NewHTTPRequest(method, url, body)
	if err != nil {
		log.Panic("error creating request")
	}
	oauth2bearer.SetHeaders(req, dfa.requestHeaderMap)
	return req
}

func runRequest(req *http.Request) []byte {
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyResult, err := ioutil.ReadAll(resp.Body)
	return bodyResult
}

func doRequest(method, url string, body io.Reader, dfa *DFAccess) []byte {
	req := newRequest(method, url, body, dfa)
	return runRequest(req)
}

// Get does an authenticated get
func Get(urlStub string, dfa *DFAccess, queryDict url.Values) []byte {
	queryURL := apiURLStub() + urlStub + "?" + queryDict.Encode()
	return doRequest("GET", queryURL, nil, dfa)
}

// Post does an authenticated post
func Post(urlStub, jsonString string, dfa *DFAccess, queryDict url.Values) []byte {
	queryURL := apiURLStub() + urlStub
	return doRequest("POST", queryURL, bytes.NewBuffer([]byte(jsonString)), dfa)
}

// eof
