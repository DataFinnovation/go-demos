package main

import (
	"fmt"
	"time"

	"github.com/DataFinnovation/go-demos/apidemos/access"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const tooManyDays = 14

func main() {
	sources := []string{"US SEC", "UK CH", "Chile SVS",
		"Peru SMV", "Japan EDINET", "Taiwan TWSE",
		"Indonesia IDX", "Spain CNMV", "Korea FSS",
		"India BSE", "US SEC Non-XBRL"}

	dfa := access.NewDFAccessDefaultScopes()
	p := message.NewPrinter(language.English)
	total := 0
	for _, s := range sources {
		queryString :=
			`{ "query" : { "term" : { "filingsource" : "` + s + `"}},` +
				`"sort" : {"retrievaltime" : "desc"}` + `}`
		res := dfa.DocumentsDSLQuery(queryString, 1)
		total += res.TotalHits
		lastArrivalTimeString := res.Hits[0].Source.RetrievalTime
		lastArrivalTime, _ := time.Parse(time.RFC3339, lastArrivalTimeString+"Z")
		daysAgo := time.Since(lastArrivalTime).Hours() / 24
		suffix := ""
		if daysAgo > tooManyDays {
			suffix = fmt.Sprintf(" ---> LAST FILING %.0f DAYS AGO!", daysAgo)
		}
		p.Printf("%s : %d %s\n", s, res.TotalHits, suffix)
	}
	p.Printf("TOTAL : %d\n", total)
}
