package apidemos

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type rawTokenData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func tokenURL() string {
	return Getenv("DF_TOKEN_URL", "https://apiauth.dfnapp.com/oauth2/token")
}

func apiURLStub() string {
	return Getenv("DF_API_URL_STUB", "https://clientapi.dfnapp.com/")
}

// GetTokenDefaultScopes uses the default scope
func GetTokenDefaultScopes() string {
	return GetToken("clientapi/basicsearch clientapi/advancedsearch'")
}

// GetToken returns a single oauth2 token
func GetToken(scopes string) string {
	var username string = os.Getenv("DF_CLIENT_ID")
	var passwd string = os.Getenv("DF_CLIENT_SECRET")
	client := &http.Client{}

	queryData := url.Values{}
	queryData.Set("grant_type", "client_credentials")
	queryData.Set("scope", scopes)
	req, err := http.NewRequest("POST", tokenURL(), strings.NewReader(queryData.Encode()))

	req.SetBasicAuth(username, passwd)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	var rtd rawTokenData
	json.Unmarshal(bodyText, &rtd)
	return rtd.AccessToken
}

// Get ...
func Get(urlStub, token string, queryDict url.Values) []byte {
	queryURL := apiURLStub() + urlStub + "?" + queryDict.Encode()
	req, err := http.NewRequest("GET", queryURL, nil)

	authString := "Bearer " + token
	req.Header.Add("x-api-key", os.Getenv("DF_API_KEY"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authString)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyResult, err := ioutil.ReadAll(resp.Body)
	return bodyResult
}

// eof
