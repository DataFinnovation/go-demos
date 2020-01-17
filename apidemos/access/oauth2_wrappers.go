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

const defaultTokenURL = "https://apiauth.dfnapp.com/oauth2/token"
const defaultAPIStub = "https://clientapi.dfnapp.com/"
const defaultScope1 = "clientapi/basicsearch"
const defaultScope2 = "clientapi/advancedsearch"

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
	return NewDFAccess([]string{defaultScope1, defaultScope2})
}

// NewDFAccess builds a DFAccess object given a list of scopes
func NewDFAccess(scopes []string) *DFAccess {
	var dfa DFAccess
	newToken, err := getToken(scopes)
	if err != nil {
		log.Panic("Error getting token in NewDFAccess")
	}
	dfa.token = newToken
	dfa.requestHeaderMap = map[string]string{
		"x-api-key":    os.Getenv("DF_API_KEY"),
		"Content-Type": "application/json",
	}
	return &dfa
}

func tokenURL() string {
	return Getenv("DF_TOKEN_URL", defaultTokenURL)
}

func apiURLStub() string {
	return Getenv("DF_API_URL_STUB", defaultAPIStub)
}

// GetToken returns a single oauth2 token
func getToken(scopes []string) (*oauth2bearer.BearerToken, error) {
	// note: there is no sensible default for these parameters
	clientID := os.Getenv("DF_CLIENT_ID")
	clientSecret := os.Getenv("DF_CLIENT_SECRET")
	creds := oauth2bearer.ClientCredentials{ClientID: clientID, ClientSecret: clientSecret}
	source := oauth2bearer.BearerTokenSource{
		Credentials: creds,
		ScopesList:  scopes,
		URL:         tokenURL(),
	}
	return source.Token()
}

func (dfa *DFAccess) newRequest(method, url string, body io.Reader) *http.Request {
	req, err := dfa.token.NewHTTPRequest(method, url, body)
	if err != nil {
		log.Panic("error creating request")
	}
	oauth2bearer.SetHeaders(req, dfa.requestHeaderMap)
	return req
}

func (dfa *DFAccess) runRequest(req *http.Request) []byte {
	client := dfa.token.NewAuthenticatedClient()
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	bodyResult, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	return bodyResult
}

func (dfa *DFAccess) doRequest(method, url string, body io.Reader) []byte {
	req := dfa.newRequest(method, url, body)
	return dfa.runRequest(req)
}

// Get does an authenticated get
func (dfa *DFAccess) Get(urlStub string, queryDict url.Values) []byte {
	queryURL := apiURLStub() + urlStub + "?" + queryDict.Encode()
	return dfa.doRequest("GET", queryURL, nil)
}

// Post does an authenticated post
func (dfa *DFAccess) Post(urlStub, jsonString string, queryDict url.Values) []byte {
	queryURL := apiURLStub() + urlStub
	return dfa.doRequest("POST", queryURL, bytes.NewBuffer([]byte(jsonString)))
}

// eof
