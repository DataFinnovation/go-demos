package main

import (
	"fmt"

	"github.com/DataFinnovation/go-demos/apidemos/access"
)

func main() {
	token := access.NewDFAccessDefaultScopes()
	queryString := `{ "query" : { "term" : { "companyname" : "microsoft"}}}`
	res := token.FactsDSLQuery(queryString, 100)
	fmt.Println(res.TotalHits)
	fmt.Println(len(res.Hits))
	res = token.DocumentsDSLQuery(queryString, 100)
	fmt.Println(res.TotalHits)
	fmt.Println(len(res.Hits))
}

// eof
