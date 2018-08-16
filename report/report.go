package report

import (
	"encoding/json"
	"time"
)

type Report struct {
	Timestamp time.Time `json:"@timestamp"`
	CspReport struct {
		DocumentURI       string `json:"document-uri"`
		Referrer          string `json:"referrer"`
		BlockedURI        string `json:"blocked-uri"`
		ViolatedDirective string `json:"violated-directive"`
		OriginalPolicy    string `json:"original-policy"`
	} `json:"csp-report"`
}

func parse(data []byte) (Report, error) {
	var report Report

	if err := json.Unmarshal(data, &report); err != nil {
		return report, err
	}

	return report, nil
}
