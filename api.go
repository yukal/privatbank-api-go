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
)

const API_URL = "https://acp.privatbank.ua/api"

type API struct {
	Logger    io.Writer
	httpAgent *HttpAgent
	_EOL      string
}

type APIOptions struct {
	Token    string
	Encoding string
	Logger   io.Writer
}

func NewAPI(options APIOptions) *API {
	eol := "\n"

	if runtime.GOOS == "windows" {
		eol = "\r\n"
	}

	return &API{
		httpAgent: NewHttpAgent(options.Token, options.Encoding),
		Logger:    options.Logger,
		_EOL:      eol,
	}
}

func (a *API) logResponse(resp *http.Response) {
	text := resp.Status + " " +
		resp.Proto + " " +
		resp.Request.Method + " " +
		resp.Request.URL.String() + " [" +
		strconv.FormatInt(resp.ContentLength, 10) + "b]" + a._EOL +
		"Header" + a._EOL

	for key, item := range resp.Header {
		text += "  " + key + ": " + strings.Join(item, "; ") + a._EOL
	}

	if cookies := resp.Cookies(); len(cookies) > 0 {
		text += a._EOL + "Cookies"

		for _, item := range resp.Cookies() {
			text += "  " + item.String() + a._EOL
		}
	}

	a.Logger.Write([]byte(text + a._EOL))
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
