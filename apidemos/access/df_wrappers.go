package access

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
func FactsStringQuery(query string, dfa *DFAccess, nResults int) *SearchResults {
	urlStub := "facts/textquery"
	return stringQuery(urlStub, query, dfa, nResults)
}

// DocumentsStringQuery runs a single query-string query and
// returns parsed results
func DocumentsStringQuery(query string, dfa *DFAccess, nResults int) *SearchResults {
	urlStub := "documents/textquery"
	return stringQuery(urlStub, query, dfa, nResults)
}

func stringQuery(urlStub, query string, dfa *DFAccess, nResults int) *SearchResults {
	queryDict := url.Values{}
	queryDict.Set("querystring", query)
	queryDict.Set("simplequery", "false")
	queryDict.Set("maxresult", strconv.Itoa(nResults))
	res := Get(urlStub, dfa, queryDict)
	return parseRawResult(res)
}

// FactsDSLQuery runs a dsl query against the fact database
func FactsDSLQuery(dslQueryString string, dfa *DFAccess, nResults int) *SearchResults {
	return dslQuery("facts/dslquery", dslQueryString, dfa, nResults)
}

// DocumentsDSLQuery runs a dsl query against the fact database
func DocumentsDSLQuery(dslQueryString string, dfa *DFAccess, nResults int) *SearchResults {
	return dslQuery("documents/dslquery", dslQueryString, dfa, nResults)
}

func dslQuery(urlStub, queryString string, dfa *DFAccess, nResults int) *SearchResults {
	fullString := `{"dslquery":` + queryString + `}`
	queryDict := url.Values{}
	queryDict.Set("maxresult", strconv.Itoa(nResults))
	res := Post(urlStub, fullString, dfa, queryDict)
	return parseRawResult(res)
}

// eof
