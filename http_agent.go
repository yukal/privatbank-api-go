package privatbank

import (
	"io"
	"net/http"
	"strings"
	"time"
)

const _USER_AGENT = "golang-http-req"

type apiHttpAgent struct {
	token    string
	encoding string

	client *http.Client
	// req    *http.Request
}

func newAPIHttpAgent(token, encoding string) *apiHttpAgent {
	return &apiHttpAgent{
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

func (a *apiHttpAgent) setBasicHeaders(req *http.Request) {
	req.Header.Add("User-Agent", _USER_AGENT)
	req.Header.Add("Content-Type", "application/json;charset="+a.encoding)
	req.Header.Add("token", a.token)
}

func (a *apiHttpAgent) Get(path string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, _API+path, nil)

	if err != nil {
		return nil, err
	}

	a.setBasicHeaders(req)
	return a.client.Do(req)
}

func (a *apiHttpAgent) Post(path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, _API+path, body)

	if err != nil {
		return nil, err
	}

	a.setBasicHeaders(req)
	return a.client.Do(req)
}

func (a *apiHttpAgent) PostOctet(path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, _API+path, body)

	if err != nil {
		return nil, err
	}

	a.setBasicHeaders(req)
	req.Header.Set("Accept", "application/octet-stream")

	return a.client.Do(req)
}
