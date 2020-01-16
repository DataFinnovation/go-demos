package main

import (
	"fmt"

	"github.com/DataFinnovation/go-benfords/benfords"
	"github.com/DataFinnovation/go-demos/apidemos"
)

func main() {
	accessToken := apidemos.GetTokenDefaultScopes()
	theQueryString := `filingsource:"US SEC" AND companyname:"bank" AND enddate:[2018-01-01 TO 2018-12-31]`
	factResults := apidemos.FactsStringQuery(theQueryString, accessToken, 5000)
	valueStrings := make([]string, factResults.TotalHits)
	for i, h := range factResults.Hits {
		valueStrings[i] = h.Source.FieldValue
	}
	b := benfords.Benfords{Base: 10}
	realizedDist, nObs := benfords.ComputeLeadDigitDistributionFromStrings(valueStrings, 10)
	fmt.Println("n obs: ", nObs)
	pvalue := b.ChiSquarePValue(realizedDist)
	fmt.Println(pvalue)
}
