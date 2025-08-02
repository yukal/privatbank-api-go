package privatbank

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// ..............................
// Currency exchange rate

type ResponseCacheInfo struct {
	FromCache  bool   `json:"from_cache"`
	CacheTime  string `json:"cache_time"`  // (optional) 1526645093874
	ServerTime string `json:"server_time"` // (optional) 1526645134001
}

type ResponseCurrency struct {
	CacheInfo ResponseCacheInfo `json:"cache_info"`
	USD       CurrencySaleBuy   `json:"USD"`
	EUR       CurrencySaleBuy   `json:"EUR"`
	// Data      []CurrencyItem    `json:"history"`
}

type ResponseCurrencyHistory struct {
	CacheInfo ResponseCacheInfo `json:"cache_info"`
	Data      struct {
		SessionState any                   `json:"sessionState"`
		History      []CurrencyHistoryItem `json:"history"`
	} `json:"data"`
}

type CurrencySaleBuy struct {
	Sale CurrencyItem `json:"S"`
	Buy  CurrencyItem `json:"B"`
}

type CurrencyItem struct {
	Date      string `json:"date"`
	Rate      string `json:"rate"`
	RateDelta string `json:"rate_delta"`
	NbuRate   string `json:"nbuRate"`
}

type CurrencyHistoryItem struct {
	Date          string `json:"date"`
	CurrencyCode  string `json:"currencyCode"`
	NbuRate       string `json:"nbuRate"`
	RateSale      string `json:"rate_s"`
	RateSaleDelta string `json:"rate_s_delta"`
	RateBuy       string `json:"rate_b"`
	RateBuyDelta  string `json:"rate_b_delta"`
}

// Get currency exchange rate
//
//	{
//	   "cache_info": {
//	           "from_cache": true,
//	           "cache_time": 1526645093874,  // (optional)
//	           "server_time": 1526645134001  // (optional)
//	   },
//	   "USD": {
//	           "B": {
//	                   "date": "30-07-2025 10:00:49",
//	                   "rate": "41.5900000",
//	                   "rate_delta": "0.0600000",
//	                   "nbuRate": "41.7886000"
//	           },
//	           "S": {
//	                   "date": "30-07-2025 10:00:49",
//	                   "rate": "41.9900000",
//	                   "rate_delta": "0.0600000",
//	                   "nbuRate": "41.7886000"
//	           }
//	   },
//	   "EUR": {
//	           "B": {
//	                   "date": "30-07-2025 10:00:49",
//	                   "rate": "48.0550000",
//	                   "rate_delta": "0.2550000",
//	                   "nbuRate": "48.2240000"
//	           },
//	           "S": {
//	                   "date": "30-07-2025 10:00:49",
//	                   "rate": "48.5550000",
//	                   "rate_delta": "-0.1450000",
//	                   "nbuRate": "48.2240000"
//	           }
//	   }
//	}
//
// Where:
//
//	from_cache  – indicates if data is cached;
//	cache_time  – cache time (in milliseconds);
//	server_time – current server time (in milliseconds);
//	B           – buy;
//	S           – sell;
//	date        – exchange rate date;
//	rate        – exchange rate;
//	rate_delta  – exchange rate delta;
//	nbuRate     – rate of the National Bank of Ukraine (NBU).
func (api *API) GetCurrency() (data ResponseWrapper[ResponseCurrency], err error) {
	var (
		responseData ResponseCurrency
		resp         *http.Response
		body         []byte
	)

	// apiURL := buildApiURL("/proxy/currency/", url.Values{})
	apiURL := URL_API_CORPORATE + "/proxy/currency"

	if resp, err = api.httpAgent.Get(apiURL); err != nil {
		return
	}

	defer resp.Body.Close()

	if body, err = io.ReadAll(resp.Body); err != nil {
		return
	}

	if err = json.Unmarshal(body, &responseData); err != nil {
		return
	}

	data = ResponseWrapper[ResponseCurrency]{
		Response: resp,
		RawBody:  body,
		Payload:  responseData,
	}

	return
}

// Get currency exchange rate history
//
//	startDate   DD-MM-YYYY
//	endDate     DD-MM-YYYY
//
//	startDate, endDate – start and end dates of the period (no more than 15 days).
//
// Example response:
//
//	{
//	    "cache_info": {
//	        "from_cache": false
//	    },
//	    "data": {
//	        "sessionState": null,
//	        "history": [
//	            {
//	                "date": "16-05-2025",
//	                "currencyCode": "EUR",
//	                "nbuRate": "46.3831000",
//	                "rate_s": "46.7950000",
//	                "rate_s_delta": "-0.0500000",
//	                "rate_b": "46.2950000",
//	                "rate_b_delta": "-0.0500000"
//	            },
//	            {
//	                "date": "16-05-2025",
//	                "currencyCode": "GBP",
//	                "nbuRate": "55.0360000",
//	                "rate_s": "55.8200000",
//	                "rate_s_delta": "0.1450000",
//	                "rate_b": "54.8200000",
//	                "rate_b_delta": "0.1450000"
//	            },
//	            ...
//	        ]
//	    }
//	}
func (api *API) GetCurrencyHistory(startDate, endDate string) (data ResponseWrapper[ResponseCurrencyHistory], err error) {
	var (
		responseData ResponseCurrencyHistory
		resp         *http.Response
		body         []byte
	)

	params := make(url.Values, 2)
	params.Add("startDate", startDate)
	params.Add("endDate", endDate)

	apiURL := buildApiURL("/proxy/currency/history", params)

	if resp, err = api.httpAgent.Get(apiURL); err != nil {
		return
	}

	defer resp.Body.Close()

	if body, err = io.ReadAll(resp.Body); err != nil {
		return
	}

	if err = json.Unmarshal(body, &responseData); err != nil {
		return
	}

	data = ResponseWrapper[ResponseCurrencyHistory]{
		Response: resp,
		RawBody:  body,
		Payload:  responseData,
	}

	return
}
