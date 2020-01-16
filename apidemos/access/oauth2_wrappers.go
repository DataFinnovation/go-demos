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

func tokenURL() string {
	return Getenv("DF_TOKEN_URL", "https://apiauth.dfnapp.com/oauth2/token")
}

func apiURLStub() string {
	return Getenv("DF_API_URL_STUB", "https://clientapi.dfnapp.com/")
}

// GetTokenDefaultScopes uses the default scope
func GetTokenDefaultScopes() *oauth2bearer.BearerToken {
	return GetToken([]string{"clientapi/basicsearch", "clientapi/advancedsearch"})
}

// GetToken returns a single oauth2 token
func GetToken(scopes []string) *oauth2bearer.BearerToken {
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

func newRequest(method, url string, body io.Reader, token *oauth2bearer.BearerToken) *http.Request {
	req, err := token.NewHTTPRequest(method, url, body)
	if err != nil {
		log.Panic("error creating request")
	}
	req.Header.Set("x-api-key", os.Getenv("DF_API_KEY"))
	req.Header.Set("Content-Type", "application/json")
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

func doRequest(method, url string, body io.Reader, token *oauth2bearer.BearerToken) []byte {
	req := newRequest(method, url, body, token)
	return runRequest(req)
}

// Get does a token-authenticated get
func Get(urlStub string, token *oauth2bearer.BearerToken, queryDict url.Values) []byte {
	queryURL := apiURLStub() + urlStub + "?" + queryDict.Encode()
	return doRequest("GET", queryURL, nil, token)
}

// Post does a token-authenticated get
func Post(urlStub, jsonString string, token *oauth2bearer.BearerToken, queryDict url.Values) []byte {
	queryURL := apiURLStub() + urlStub
	return doRequest("POST", queryURL, bytes.NewBuffer([]byte(jsonString)), token)
}

// eof
