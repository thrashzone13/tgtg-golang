package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const baseUrl = "https://apptoogoodtogo.com/api"

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) Request(method string, endpoint string, payload interface{}, response interface{}, accessToken *string) error {
	url := baseUrl + endpoint

	var body []byte
	if payload != nil {
		var err error
		body, err = json.Marshal(payload)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("ser-Agent", "TGTG/24.6.1 Dalvik/2.1.0 (Linux; Android 12; SM-G920V Build/MMB29K)")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("CAccept", "application/json")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("Accept-Encoding", "gzip")

	if accessToken != nil {
		req.Header.Set("Authorization", "Bearer "+*accessToken)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return errors.New("API call failed with status code " + resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(response)
}
