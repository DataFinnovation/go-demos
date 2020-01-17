package main

import (
	"fmt"

	"github.com/DataFinnovation/go-demos/apidemos/access"
)

func main() {
	scopes := []string{"clientapi/basicsearch"}
	accessToken := access.NewDFAccess(scopes)

	theQueryString := `filingsource:"Korea FSS" AND fieldname:(hedge OR (foreign AND exchange) OR (interest AND rate))`
	factResults := access.FactsStringQuery(theQueryString, accessToken, 100)
	fmt.Println(factResults.TotalHits)

	theDocQueryStrig := `filingsource:"US SEC" AND companyname:amazon`
	docResults := access.DocumentsStringQuery(theDocQueryStrig, accessToken, 100)
	fmt.Println(docResults.TotalHits)
}
