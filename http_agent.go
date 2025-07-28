package privatbank

// Package provides a simple HTTP agent for making requests to an API.
// It includes methods for GET and POST requests, setting headers, and handling responses.

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"strings"
	"time"
)

const USER_AGENT = "golang-http-req"

type HttpAgent struct {
	token    string
	encoding string

	client *http.Client
	// req    *http.Request
}

// HTTPError represents an HTTP error with status code and message.
type HTTPError struct {
	StatusCode int
	Message    string
}

func (e *HTTPError) Error() string {
	return e.Message
}

func NewHttpAgent(token, encoding string) *HttpAgent {
	return &HttpAgent{
		token:    token,
		encoding: strings.ToLower(encoding),

		client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:       10,
				IdleConnTimeout:    30 * time.Second,
				DisableCompression: true,
			},
			Timeout: 10 * time.Second,
		},
	}
}

func (a *HttpAgent) setBasicHeaders(req *http.Request) {
	req.Header.Add("User-Agent", USER_AGENT)
	req.Header.Add("Content-Type", "application/json;charset="+a.encoding)
	req.Header.Add("token", a.token)
}

func (a *HttpAgent) Get(path string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, err
	}

	a.setBasicHeaders(req)
	return a.client.Do(req)
}

func (a *HttpAgent) Post(path string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, path, body)

	if err != nil {
		return nil, err
	}

	a.setBasicHeaders(req)

	if len(headers) > 0 {
		for key, val := range headers {
			req.Header.Set(key, val)
		}
	}

	return a.client.Do(req)
}

func ExtractFilenameFromContentDisposition(header http.Header) (filename string, err error) {
	var params map[string]string

	if _, params, err = mime.ParseMediaType(
		header.Get("Content-Disposition")); err != nil {
		fmt.Printf("Error parsing Content-Disposition: %v\n", err)
		return
	}

	filename = params["filename"]
	return
}
