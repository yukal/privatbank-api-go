package privatbank

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
	API_URL = "https://acp.privatbank.ua/api"

	// Response Status Codes

	RESPONSE_SUCCESS = "SUCCESS"
	RESPONSE_ERROR   = "ERROR"

	// Limitation per Request
	LIMIT_DATA = 100
)

type API struct {
	Logger    io.Writer
	httpAgent *HttpAgent

	timeoutPerReq time.Duration
}

type APIOptions struct {
	Token    string
	Encoding string
	Logger   io.Writer

	// Timeout per request
	//  - default 500 milliseconds
	//  - minimum 100 milliseconds
	//  - maximum 1 second
	TimeoutPerReq time.Duration
}

type ResponseDataStatement[T any] interface {
	// GetStatus() string
	// GetExistNextPage() bool
	// GetNextPageId() string
	GetServerData() ResponseServerData
	GetClientData() []T
}

type ResponseServerData struct {
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

	logger := io.Discard
	// options.Logger = io.Discard
	// options.Logger = os.Stdout

	if options.Logger != nil {
		logger = options.Logger
	}

	if runtime.GOOS == "windows" {
		eol = "\r\n"
	}

	return &API{
		Logger:        logger,
		httpAgent:     NewHttpAgent(options.Token, options.Encoding),
		timeoutPerReq: max(200*time.Millisecond, options.TimeoutPerReq),
	}
}

func (a *API) logResponse(resp *http.Response) {
	text := resp.Status + " " +
		resp.Proto + " " +
		resp.Request.Method + " " +
		resp.Request.URL.String() + " [" +
		strconv.FormatInt(resp.ContentLength, 10) + "b]" + eol +
		"Header" + eol

	for key, item := range resp.Header {
		text += "  " + key + ": " + strings.Join(item, "; ") + eol
	}

	if cookies := resp.Cookies(); len(cookies) > 0 {
		text += eol + "Cookies"

		for _, item := range resp.Cookies() {
			text += "  " + item.String() + eol
		}
	}

	a.Logger.Write([]byte(text + eol))
}

func buildApiURL(apiPath string, queryParams url.Values) string {
	return API_URL + apiPath + "?" + queryParams.Encode()
}

func toJSONReader(payload any) (r *bytes.Reader, err error) {
	var buf []byte

	if buf, err = json.Marshal(payload); err != nil {
		return nil, err
	}

	return bytes.NewReader(buf), nil
}

func fetchWithinMultipleRequests[

	ExactData BalanceStatement | TransactionStatement,
	RespData ResponseDataStatement[ExactData],

](a *API, apiPath string, params url.Values) (items []ExactData, err error) {
	items = make([]ExactData, 0)
	next := true

	for next {
		var (
			resp *http.Response
			body []byte
			data RespData
		)

		fullApiURL := apiPath + "?" + params.Encode()

		if resp, err = a.httpAgent.Get(fullApiURL); err != nil {
			return
		}

		defer resp.Body.Close()
		a.logResponse(resp)

		if body, err = io.ReadAll(resp.Body); err != nil {
			return
		}

		if err = json.Unmarshal(body, &data); err != nil {
			return
		}

		// Використовуємо методи інтерфейсу для доступу до полів
		items = append(items, data.GetClientData()...)
		server := data.GetServerData()

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
