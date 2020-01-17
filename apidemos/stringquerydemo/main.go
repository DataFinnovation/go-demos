package main

import (
	"fmt"

	"github.com/DataFinnovation/go-demos/apidemos/access"
)

func main() {
	scopes := []string{"clientapi/basicsearch"}
	accessToken := access.NewDFAccess(scopes)

	theQueryString := `filingsource:"Korea FSS" AND fieldname:(hedge OR (foreign AND exchange) OR (interest AND rate))`
	factResults := accessToken.FactsStringQuery(theQueryString, 100)
	fmt.Println(factResults.TotalHits)

	theDocQueryStrig := `filingsource:"US SEC" AND companyname:amazon`
	docResults := accessToken.DocumentsStringQuery(theDocQueryStrig, 100)
	fmt.Println(docResults.TotalHits)
}
