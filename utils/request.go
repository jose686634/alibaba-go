package utils

import (
	"bytes"
	"context"
	"net/http"
	"time"
)

type HTTPClient struct {
	client  *http.Client
	headers map[string]string
	timeout time.Duration
}

// NewHTTPClient creates a new instance of HTTPClient
func NewHTTPClient(timeout time.Duration, headers map[string]string) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
		headers: headers,
		timeout: timeout,
	}
}

// SetHeaders sets the request headers
func (hc *HTTPClient) SetHeaders(headers map[string]string) {
	hc.headers = headers
}

// UpdateHeaders sets the request headers
func (hc *HTTPClient) UpdateHeaders(key, value string) {
	hc.headers[key] = value
}

// Get makes a GET request
func (hc *HTTPClient) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return hc.doRequest(req)
}

// Post makes a POST request
func (hc *HTTPClient) Post(url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	return hc.doRequest(req)
}

// DoRequest executes the HTTP request
func (hc *HTTPClient) doRequest(req *http.Request) (*http.Response, error) {
	// Set headers
	for key, value := range hc.headers {
		req.Header.Set(key, value)
	}

	// Set timeout context
	ctx, cancel := context.WithTimeout(context.Background(), hc.timeout)
	defer cancel()
	req = req.WithContext(ctx)

	resp, err := hc.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
