package main

import (
	"fmt"

	"github.com/DataFinnovation/go-demos/apidemos"
)

func main() {
	token := apidemos.GetTokenDefaultScopes()
	queryString := `{ "query" : { "term" : { "companyname" : "microsoft"}}}`
	res := apidemos.FactsDSLQuery(queryString, token, 100)
	fmt.Println(res.TotalHits)
	fmt.Println(len(res.Hits))
	res = apidemos.DocumentsDSLQuery(queryString, token, 100)
	fmt.Println(res.TotalHits)
	fmt.Println(len(res.Hits))
}

// eof
