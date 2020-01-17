package main

import (
	"fmt"

	"github.com/DataFinnovation/go-benfords/benfords"
	"github.com/DataFinnovation/go-demos/apidemos/access"
)

func main() {
	accessToken := access.NewDFAccessDefaultScopes()
	theQueryString := `filingsource:"US SEC" AND companyname:"bank" AND enddate:[2018-01-01 TO 2018-12-31]`
	factResults := access.FactsStringQuery(theQueryString, accessToken, 5000)
	valueStrings := make([]string, factResults.TotalHits)
	for i, h := range factResults.Hits {
		valueStrings[i] = h.Source.FieldValue
	}
	b := benfords.Benfords{Base: 10}
	realizedDist, nObs := benfords.ComputeLeadDigitDistributionFromStrings(valueStrings, 10)
	fmt.Println("n obs: ", nObs)
	pvalue := b.ChiSquarePValue(realizedDist)
	cgStat := b.ChoGainesStat(nObs, realizedDist)
	leemisStat := b.LeemisStat(nObs, realizedDist)
	fmt.Println("chi-square p value: ", pvalue)
	fmt.Println("cho-gaines: ", cgStat)
	fmt.Println("leemis: ", leemisStat)
}
