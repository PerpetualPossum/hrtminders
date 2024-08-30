package main

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Response struct {
	Cron string `json:"cron"`
}

func convertToCron(naturalCron string) (string, error) {
	// Convert natural language cron to cron expression
	apiUrl := ENVIRONMENT.CronApiUrl

	params := url.Values{
		"schedule": {naturalCron},
	}

	reqUrl := apiUrl + "/?" + params.Encode()
	resp, err := http.Get(reqUrl)
	if err != nil {
		return "", err
	} else if resp.StatusCode != 200 {
		return "", errors.New("failed to convert natural cron to cron expression")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	print("\n" + string(body) + "\n" + strconv.Itoa((resp.StatusCode)))

	return string(body), nil
}
