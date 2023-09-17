package utils

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func FetchXML[T any](url string, timeout time.Duration) (T, error) {
	var result T
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("failed to fetch RSS data: unexpected status code: %d", resp.StatusCode)
	}

	rssData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = xml.Unmarshal(rssData, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
