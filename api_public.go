// ..............................
// Public API

package privatbank

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

type PubRespCurrencyItem struct {
	Currency     string          `json:"ccy"`
	BaseCurrency string          `json:"base_ccy"`
	Buy          decimal.Decimal `json:"buy"`
	Sale         decimal.Decimal `json:"sale"`
}

type PubRespCurrencyHistory struct {
	Date             time.Time
	DateStr          string             `json:"date"`
	BankName         string             `json:"bank"`
	BaseCurrency     int                `json:"baseCurrency"`
	BaseCurrencyCode string             `json:"baseCurrencyLit"`
	ExchangeRates    []ExchangeRateItem `json:"exchangeRate"`
}

type ExchangeRateItem struct {
	BaseCurrency   string          `json:"baseCurrency"`   // "UAH",
	Currency       string          `json:"currency"`       // "AUD",
	SaleRateNB     decimal.Decimal `json:"saleRateNB"`     // 26.6097000,
	PurchaseRateNB decimal.Decimal `json:"purchaseRateNB"` // 26.6097000
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

// Archive of PrivatBank and NBU currency exchange rates.
//
// The API allows you to get information about PrivatBank and NBU cash exchange rates
// for a selected date. The archive stores data for the last 4 years.
//
// Arguments:
//
//	date [string] - "YYYY-MM-DD" // "2025-04-01"
//
// Example response:
//
//	{
//		"date": "30.04.2025",
//		"bank": "PB",
//		"baseCurrency": 980,
//		"baseCurrencyLit": "UAH",
//		"exchangeRate": [
//			{
//				"baseCurrency": "UAH",
//				"currency": "UAH",
//				"saleRateNB": 1.0000000,
//				"purchaseRateNB": 1.0000000
//			},
//			{
//				"baseCurrency": "UAH",
//				"currency": "GBP",
//				"saleRateNB": 55.6364000,
//				"purchaseRateNB": 55.6364000,
//				"saleRate": 55.6600000,
//				"purchaseRate": 54.8500000
//			},
//			...
//		]
//	}
//
// Response parameters:
//
//	baseCurrency    Base currency
//	currency        Transaction currency
//	saleRateNB      NBU* sale rate (National Bank of Ukraine)
//	purchaseRateNB  NBU purchase rate
//	saleRate        PrivatBank sale rate
//	purchaseRate    PrivatBank purchase rate
func GetCurrencyHistoryAt(date string) (data ResponseWrapper[PubRespCurrencyHistory], err error) {
	var (
		respData PubRespCurrencyHistory
		resp     *http.Response
		req      *http.Request
		body     []byte
	)

	formattedDate := date[8:10] + "." + date[5:7] + "." + date[:4]

	apiURL := URL_API_PUBLIC + "/exchange_rates?json&date=" + formattedDate
	// date.Format("02.01.2006")

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

	if err = json.Unmarshal(body, &respData); err != nil {
		return
	}

	if respData.Date, err = time.ParseInLocation(
		DateOnly, respData.DateStr, Location); err != nil {
		return
	}

	data = ResponseWrapper[PubRespCurrencyHistory]{
		Response: resp,
		RawBody:  body,
		Payload:  respData,
	}

	return
}
