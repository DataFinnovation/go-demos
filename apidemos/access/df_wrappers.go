package access

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/DataFinnovation/go-demos/apidemos/oauth2bearer"
)

func parseRawResult(rawResult []byte) *SearchResults {
	var sr SearchResults
	json.Unmarshal(rawResult, &sr)
	return &sr
}

// FactsStringQuery runs a single query-string query and
// returns parsed results
func FactsStringQuery(query string, token *oauth2bearer.BearerToken, nResults int) *SearchResults {
	urlStub := "facts/textquery"
	return stringQuery(urlStub, query, token, nResults)
}

// DocumentsStringQuery runs a single query-string query and
// returns parsed results
func DocumentsStringQuery(query string, token *oauth2bearer.BearerToken, nResults int) *SearchResults {
	urlStub := "documents/textquery"
	return stringQuery(urlStub, query, token, nResults)
}

func stringQuery(urlStub, query string, token *oauth2bearer.BearerToken, nResults int) *SearchResults {
	queryDict := url.Values{}
	queryDict.Set("querystring", query)
	queryDict.Set("simplequery", "false")
	queryDict.Set("maxresult", strconv.Itoa(nResults))
	res := Get(urlStub, token, queryDict)
	return parseRawResult(res)
}

// FactsDSLQuery runs a dsl query against the fact database
func FactsDSLQuery(dslQueryString string, token *oauth2bearer.BearerToken, nResults int) *SearchResults {
	return dslQuery("facts/dslquery", dslQueryString, token, nResults)
}

// DocumentsDSLQuery runs a dsl query against the fact database
func DocumentsDSLQuery(dslQueryString string, token *oauth2bearer.BearerToken, nResults int) *SearchResults {
	return dslQuery("documents/dslquery", dslQueryString, token, nResults)
}

func dslQuery(urlStub, queryString string, token *oauth2bearer.BearerToken, nResults int) *SearchResults {
	fullString := `{"dslquery":` + queryString + `}`
	queryDict := url.Values{}
	queryDict.Set("maxresult", strconv.Itoa(nResults))
	res := Post(urlStub, fullString, token, queryDict)
	return parseRawResult(res)
}

// eof
