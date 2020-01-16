package main

import (
	"fmt"

	"github.com/DataFinnovation/go-demos/apidemos/access"
)

func main() {
	token := access.GetTokenDefaultScopes()
	queryString := `{ "query" : { "term" : { "companyname" : "microsoft"}}}`
	res := access.FactsDSLQuery(queryString, token, 100)
	fmt.Println(res.TotalHits)
	fmt.Println(len(res.Hits))
	res = access.DocumentsDSLQuery(queryString, token, 100)
	fmt.Println(res.TotalHits)
	fmt.Println(len(res.Hits))
}

// eof
