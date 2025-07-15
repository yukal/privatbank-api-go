package privatbank

import (
	"io"
	"net/http"
	"strings"
	"time"
)

const _USER_AGENT = "golang-http-req"

type HttpAgent struct {
	token    string
	encoding string

	client *http.Client
	// req    *http.Request
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
	req.Header.Add("User-Agent", _USER_AGENT)
	req.Header.Add("Content-Type", "application/json;charset="+a.encoding)
	req.Header.Add("token", a.token)
}

func (a *HttpAgent) Get(path string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, _API+path, nil)

	if err != nil {
		return nil, err
	}

	a.setBasicHeaders(req)
	return a.client.Do(req)
}

func (a *HttpAgent) Post(path string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, _API+path, body)

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
