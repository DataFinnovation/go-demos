package main

import (
	"fmt"

	"github.com/DataFinnovation/go-demos/apidemos"
)

func main() {
	scopes := "clientapi/basicsearch"
	accessToken := apidemos.GetToken(scopes)

	theQueryString := `filingsource:"Korea FSS" AND fieldname:(hedge OR (foreign AND exchange) OR (interest AND rate))`
	factResults := apidemos.FactsStringQuery(theQueryString, accessToken, 100)
	fmt.Println(factResults.TotalHits)

	theDocQueryStrig := `filingsource:"US SEC" AND companyname:amazon`
	docResults := apidemos.DocumentsStringQuery(theDocQueryStrig, accessToken, 100)
	fmt.Println(docResults.TotalHits)
}
