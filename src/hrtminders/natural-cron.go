package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

type Response struct {
	Cron string `json:"cron"`
}

func convertToCron(naturalCron string) (string, bool) {
	// Convert natural language cron to cron expression
	apiUrl := ENVIRONMENT.CronApiUrl

	params := url.Values{
		"schedule": {naturalCron},
	}

	reqUrl := apiUrl + "/?" + params.Encode()
	resp, err := http.Get(reqUrl)
	if err != nil {
		log.Fatalf("Error calling cron api: %v", err)
		return "", false
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
		return "", false
	}

	return string(body), true
}
