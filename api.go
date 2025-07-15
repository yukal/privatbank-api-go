package privatbank

import (
	"io"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

const _API = "https://acp.privatbank.ua/api"

type API struct {
	Logger    io.Writer
	httpAgent *apiHttpAgent
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
		httpAgent: newAPIHttpAgent(options.Token, options.Encoding),
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
