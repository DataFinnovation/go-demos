package access

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/oauth2/clientcredentials"
)

const defaultTokenURL = "https://apiauth.dfnapp.com/oauth2/token"
const defaultAPIStub = "https://clientapi.dfnapp.com/"

const defaultScope1 = "clientapi/basicsearch"
const defaultScope2 = "clientapi/advancedsearch"

// DFAccess is the top-level DF API access object
// It is meant to be opaque -- these just get passed into
// the data access wrapper functions to sort out authentication.
type DFAccess struct {
	client *http.Client
	config clientcredentials.Config
}

// NewDFAccessDefaultScopes builds a new DFAccess object with
// the default scopes
func NewDFAccessDefaultScopes() *DFAccess {
	return NewDFAccess([]string{defaultScope1, defaultScope2})
}

// NewDFAccess builds a DFAccess object given a list of scopes
func NewDFAccess(scopes []string) *DFAccess {
	var dfa DFAccess

	conf := clientcredentials.Config{
		ClientID:     os.Getenv("DF_CLIENT_ID"),
		ClientSecret: os.Getenv("DF_CLIENT_SECRET"),
		TokenURL:     tokenURL(),
		Scopes:       scopes,
	}
	dfa.config = conf
	dfa.client = conf.Client(context.Background())

	return &dfa
}

func (dfa *DFAccess) newRequest(method, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Set("x-api-key", os.Getenv("DF_API_KEY"))
	req.Header.Set("Content-Type", "application/json")
	return req
}

func (dfa *DFAccess) runRequest(req *http.Request) []byte {
	client := dfa.client
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
	queryURL := apiURLStub() + urlStub + "?" + queryDict.Encode()
	return dfa.doRequest("POST", queryURL, bytes.NewBuffer([]byte(jsonString)))
}

func apiURLStub() string {
	return Getenv("DF_API_URL_STUB", defaultAPIStub)
}

func tokenURL() string {
	return Getenv("DF_TOKEN_URL", defaultTokenURL)
}

// eof
