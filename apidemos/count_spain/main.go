package main

import (
	"fmt"
	"time"

	"github.com/DataFinnovation/go-demos/apidemos/access"
)

func genDate(year int, half int) (time.Time, time.Time) {
	loc := time.Now().Location()
	sd := time.Date(year, 1, 1, 0, 0, 0, 0, loc)
	ed := time.Date(year, 6, 30, 0, 0, 0, 0, loc)
	if half == 2 {
		sd = time.Date(year, 7, 1, 0, 0, 0, 0, loc)
		ed = time.Date(year, 12, 31, 0, 0, 0, 0, loc)
	}
	return sd, ed
}

func main() {

	dfa := access.NewDFAccessDefaultScopes()
	total := 0

	now := time.Now()
	thisYear := now.Year()

	for year := 2005; year <= thisYear; year++ {
		for half := 1; half <= 2; half++ {
			startDate, endDate := genDate(year, half)
			aStr := `filingsource:"Spain CNMV" AND startdate:[`
			queryString := fmt.Sprintf("%s %d-%02d-%02d TO %d-%02d-%02d ]", aStr,
				startDate.Year(), startDate.Month(), startDate.Day(),
				endDate.Year(), endDate.Month(), endDate.Day())
			res := dfa.DocumentsStringQuery(queryString, 1)
			total += res.TotalHits
			fmt.Println(year, half, " : ", res.TotalHits)
		}
	}
}
