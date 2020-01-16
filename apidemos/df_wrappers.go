package apidemos

import (
	"encoding/json"
	"net/url"
	"strconv"
)

func parseRawResult(rawResult []byte) *SearchResults {
	var sr SearchResults
	json.Unmarshal(rawResult, &sr)
	return &sr
}

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
	return parseRawResult(res)
}

// FactsDSLQuery runs a dsl query against the fact database
func FactsDSLQuery(dslQueryString, token string, nResults int) *SearchResults {
	return dslQuery("facts/dslquery", token, dslQueryString, nResults)
}

// DocumentsDSLQuery runs a dsl query against the fact database
func DocumentsDSLQuery(dslQueryString, token string, nResults int) *SearchResults {
	return dslQuery("documents/dslquery", token, dslQueryString, nResults)
}

func dslQuery(urlStub, token, queryString string, nResults int) *SearchResults {
	fullString := `{"dslquery":` + queryString + `}`
	queryDict := url.Values{}
	queryDict.Set("maxresult", strconv.Itoa(nResults))
	res := Post(urlStub, token, fullString, queryDict)
	return parseRawResult(res)
}

// eof
