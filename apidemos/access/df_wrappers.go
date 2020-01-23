package access

import (
	"net/url"
	"strconv"
)

// FactsStringQuery runs a single query-string query and
// returns parsed results
func (dfa *DFAccess) FactsStringQuery(query string, nResults int) *SearchResults {
	urlStub := "facts/textquery"
	return dfa.stringQuery(urlStub, query, nResults)
}

// DocumentsStringQuery runs a single query-string query and
// returns parsed results
func (dfa *DFAccess) DocumentsStringQuery(query string, nResults int) *SearchResults {
	urlStub := "documents/textquery"
	return dfa.stringQuery(urlStub, query, nResults)
}

func (dfa *DFAccess) stringQuery(urlStub, query string, nResults int) *SearchResults {
	queryDict := url.Values{}
	queryDict.Set("querystring", query)
	queryDict.Set("simplequery", "false")
	queryDict.Set("maxresult", strconv.Itoa(nResults))
	res := dfa.Get(urlStub, queryDict)
	return parseRawResult(res)
}

// FactsDSLQuery runs a dsl query against the fact database
func (dfa *DFAccess) FactsDSLQuery(dslQueryString string, nResults int) *SearchResults {
	return dfa.dslQuery("facts/dslquery", dslQueryString, nResults)
}

// DocumentsDSLQuery runs a dsl query against the fact database
func (dfa *DFAccess) DocumentsDSLQuery(dslQueryString string, nResults int) *SearchResults {
	return dfa.dslQuery("documents/dslquery", dslQueryString, nResults)
}

func (dfa *DFAccess) dslQuery(urlStub, queryString string, nResults int) *SearchResults {
	fullString := `{"dslquery":` + queryString + `}`
	queryDict := url.Values{}
	queryDict.Set("maxresult", strconv.Itoa(nResults))
	res := dfa.Post(urlStub, fullString, queryDict)
	return parseRawResult(res)
}

// eof
