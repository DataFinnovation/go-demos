package access

import "encoding/json"

// SearchResults contains the full requests from a search query
type SearchResults struct {
	TotalHits int     `json:"totalHits"`
	MaxScore  float64 `json:"maxScore"`
	Hits      []Hit   `json:"hits"`
}

// Hit covers a single search hit
type Hit struct {
	Type   string    `json:"type"`
	ID     string    `json:"id"`
	Score  float64   `json:"score"`
	Source HitRecord `json:"source"`
}

// HitRecord has the content within a single hit
type HitRecord struct {
	FieldTag      string `json:"fieldtag"`
	Source        string `json:"filingsource"`
	ReportID      string `json:"reportid"`
	FieldValue    string `json:"fieldvalue"`
	Units         string `json:"units"`
	FieldName     string `json:"fieldname"`
	StartDate     string `json:"startdate"`
	EndDate       string `json:"enddate"`
	Key           string `json:"key"`
	FilingTime    string `json:"filingtime"`
	RetrievalTime string `json:"retrievaltime"`
	FilerID       string `json:"filerid"`
	Context       string `json:"context"`
	CompanyName   string `json:"companyname"`
}

func parseRawResult(rawResult []byte) *SearchResults {
	var sr SearchResults
	json.Unmarshal(rawResult, &sr)
	return &sr
}

// eof
