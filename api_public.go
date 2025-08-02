// ..............................
// Public API

package privatbankapi

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/shopspring/decimal"
)

type PubRespCurrencyItem struct {
	Currency     string          `json:"ccy"`
	BaseCurrency string          `json:"base_ccy"`
	Buy          decimal.Decimal `json:"buy"`
	Sale         decimal.Decimal `json:"sale"`
}

// Get currency exchange rate.
// Available currencies: USD, EUR
//
//	[
//		{
//			"ccy": "EUR",
//			"base_ccy": "UAH",
//			"buy": "47.20000",
//			"sale": "48.20000"
//		},
//		{
//			"ccy": "USD",
//			"base_ccy": "UAH",
//			"buy": "41.28000",
//			"sale": "41.88000"
//		}
//	]
//
// Response parameters:
//
//	ccy        Currency code (currency codes reference: https://en.wikipedia.org/wiki/List_of_circulating_currencies)
//	base_ccy   National currency code
//	buy        Buy rate
//	sale       Sell rate
func GetCurrency(cid uint8) (data ResponseWrapper[[]PubRespCurrencyItem], err error) {
	var (
		responseData []PubRespCurrencyItem
		resp         *http.Response
		req          *http.Request
		body         []byte
	)

	apiURL := URL_API_PUBLIC + "/pubinfo?json&exchange&coursid=" +
		strconv.FormatUint(uint64(cid), 10)

	if req, err = http.NewRequest(http.MethodGet, apiURL, nil); err != nil {
		return
	}

	req.Header.Add("User-Agent", USER_AGENT)
	req.Header.Add("Content-Type", "application/json;charset=utf8")

	if resp, err = client.Do(req); err != nil {
		return
	}

	defer resp.Body.Close()

	if body, err = io.ReadAll(resp.Body); err != nil {
		return
	}

	if err = json.Unmarshal(body, &responseData); err != nil {
		return
	}

	data = ResponseWrapper[[]PubRespCurrencyItem]{
		Response: resp,
		RawBody:  body,
		Payload:  responseData,
	}

	return
}
