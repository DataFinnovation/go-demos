package apidemos

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// FactsStringQuery runs a single query-string query and
// returns parsed results
func FactsStringQuery(query, token string, nResults int) *SearchResults {
	urlStub := "facts/textquery"
	return stringQuery(urlStub, token, query, nResults)
}

// DocumentsStringQuery runs a single query-string query and
// returns parsed results
func DocumentsStringQuery(query, token string, nResults int) *SearchResults {
	urlStub := "documents/textquery"
	return stringQuery(urlStub, token, query, nResults)
}

func stringQuery(urlStub, token, query string, nResults int) *SearchResults {
	queryDict := url.Values{}
	queryDict.Set("querystring", query)
	queryDict.Set("simplequery", "false")
	queryDict.Set("maxresult", strconv.Itoa(nResults))
	res := Get(urlStub, token, queryDict)
	var fr SearchResults
	json.Unmarshal(res, &fr)
	return &fr
}

// eof
