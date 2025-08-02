package privatbankapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	URL_API_CORPORATE = "https://acp.privatbank.ua/api"

	// Response Status Codes

	RESPONSE_SUCCESS = "SUCCESS"
	RESPONSE_ERROR   = "ERROR"

	// Limitation per Request
	LIMIT_DATA = 100
)

type API struct {
	httpAgent     *HttpAgent
	timeoutPerReq time.Duration
}

type APIOptions struct {
	Token    string
	Encoding string

	// Timeout per request
	//  - default 200 milliseconds
	//  - minimum 200 milliseconds
	TimeoutPerReq time.Duration
}

type ResponseWrapper[TData any] struct {
	Response *http.Response
	RawBody  []byte `json:"rawBody"`
	Payload  TData  `json:"payload"`
}

type ResponseDataStatement[TResponsePayloadData any] interface {
	// GetStatus() string
	// GetExistNextPage() bool
	// GetNextPageId() string
	GetMetaData() ResponseMetaData
	GetPayloadData() []TResponsePayloadData
}

type ResponseMetaData struct {
	Status        string `json:"status"`
	Type          string `json:"type"`
	ExistNextPage bool   `json:"exist_next_page"`
	NextPageId    string `json:"next_page_id"`
}

var eol string = "\n"

func NewAPI(options APIOptions) *API {
	if options.Token == "" {
		panic("API token is required")
	}

	if options.Encoding == "" {
		if runtime.GOOS == "windows" {
			options.Encoding = "windows-1251"
		} else {
			options.Encoding = "utf8"
		}
	}

	if options.Encoding != "utf8" && options.Encoding != "windows-1251" {
		panic("Unsupported encoding: " + options.Encoding)
	}

	if runtime.GOOS == "windows" {
		eol = "\r\n"
	}

	return &API{
		httpAgent:     NewHttpAgent(options.Token, options.Encoding),
		timeoutPerReq: max(200*time.Millisecond, options.TimeoutPerReq),
	}
}

func buildApiURL(apiPath string, queryParams url.Values) string {
	return URL_API_CORPORATE + apiPath + "?" + queryParams.Encode()
}

func toJSONReader(payload any) (r *bytes.Reader, err error) {
	var buf []byte

	if buf, err = json.Marshal(payload); err != nil {
		return nil, err
	}

	return bytes.NewReader(buf), nil
}

func fetchWithinMultipleRequests[

	TExactData BalanceStatement | TransactionStatement,
	TRespData ResponseDataStatement[TExactData],

](a *API, apiPath string, params url.Values) (rw ResponseWrapper[[]TExactData], err error) {
	rw.Payload = make([]TExactData, 0)

	for next := true; next; {
		var (
			body []byte
			data TRespData
		)

		fullApiURL := apiPath + "?" + params.Encode()

		if rw.Response, err = a.httpAgent.Get(fullApiURL); err != nil {
			return
		}

		defer rw.Response.Body.Close()

		if body, err = io.ReadAll(rw.Response.Body); err != nil {
			return
		}

		body = bytes.TrimSpace(body)

		if err = json.Unmarshal(body, &data); err != nil {
			return
		}

		if len(rw.RawBody) == 0 {
			rw.RawBody = body
		} else {
			bodyWithSeparator := append([]byte("\r\n\r\n"), body...)
			rw.RawBody = append(rw.RawBody, bodyWithSeparator...)
		}

		// Використовуємо методи інтерфейсу для доступу до полів
		rw.Payload = append(rw.Payload, data.GetPayloadData()...)
		server := data.GetMetaData()

		if server.NextPageId != "" {
			params.Set("followId", server.NextPageId)
		} else {
			params.Del("followId")
		}

		if next = server.Status == RESPONSE_SUCCESS &&
			server.ExistNextPage; next {

			// Затримка перед наступним запитом
			<-time.After(a.timeoutPerReq)
		}
	}

	return
}
