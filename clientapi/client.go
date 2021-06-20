package clientapi

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL         = "http://localhost:8081/"
	defaultRequestEncoding = "application/json"
	// HealthOK means healthcheck is ok
	HealthOK = 200
)

// Client is the API client
type Client struct {
	// HTTP client for requests
	HTTPClient *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// Encoding or the request
	RequestEncoding string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	Health *HealthService
}

type service struct {
	client *Client
}

// CreateClient returns a new API client. If a nil httpClient is
// provided, a new http.Client will be used. Authentication is skipped.
func CreateClient() *Client {
	client := &http.Client{}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{HTTPClient: client, BaseURL: baseURL, RequestEncoding: defaultRequestEncoding}
	c.common.client = c
	c.Health = (*HealthService)(&c.common)
	return c
}

func (c *Client) sendRequest(method string, url string, codeOK int, requestBody []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != codeOK {
		errorString := fmt.Sprintf("client: Failed to perform (%v) at: %v\nresponseCode: %v\nresponseBody: %v", method, url, resp.StatusCode, string(body))
		systemError := errors.New(errorString)
		return nil, systemError
	}

	return body, nil
}
