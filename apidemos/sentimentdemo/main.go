package main

import (
	"log"

	"github.com/DataFinnovation/go-demos/apidemos/access"
	"github.com/jonreiter/govader"
)

func main() {
	accessToken := access.NewDFAccessDefaultScopes()

	theQueryString := `filingsource:"US SEC Non-XBRL" AND companyname:amazon AND enddate:[2018-01-01 TO 2019-12-31] AND reporttype:"10-K"`
	factResults := accessToken.FactsStringQuery(theQueryString, 5000)

	analyzer := govader.NewSentimentIntensityAnalyzer()

	for _, hit := range factResults.Hits {
		theContent := hit.Source.FieldValue
		res := analyzer.PolarityScores(theContent)
		log.Print(res)
	}

}
