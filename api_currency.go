package privatbankapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// ..............................
// Отримання курсів валют

type ResponseCurrencyHistory struct {
	CacheInfo struct {
		FromCache bool `json:"from_cache"`
	} `json:"cache_info"`

	Data struct {
		SessionState any                   `json:"sessionState"`
		History      []CurrencyHistoryItem `json:"history"`
	} `json:"data"`
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

// Отримання історії курсів валют
//
//	startDate   ДД-ММ-РРРР
//	endDate     ДД-ММ-РРРР
//
//	startDate, endDate – дата початку й закінчення періоду (не більше ніж 15 днів).
//
// Приклад відповіді:
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
//
// ...
//
//	{
//	   "cache_info": {
//	           "from_cache": true,
//	           "cache_time": 1526645093874,
//	           "server_time": 1526645134001
//	   },
//	   "USD": {
//	           "B": {
//	                   "date": "18-05-2018 14:00:01",
//	                   "rate": "26.0700000",
//	                   "rate_delta": "-0.0600000",
//	                   "nbuRate": "26.1888140"
//	           },
//	           "S": {
//	                   "date": "18-05-2018 10:04:13",
//	                   "rate": "26.3200000",
//	                   "rate_delta": "-0.0100000",
//	                   "nbuRate": "26.1888140"
//	           }
//	   },
//	   "EUR": {
//	           "B": {
//	                   "date": "18-05-2018 14:00:29",
//	                   "rate": "30.7300000",
//	                   "rate_delta": "-0.0700000",
//	                   "nbuRate": "30.9158950"
//	           },
//	           "S": {
//	                   "date": "18-05-2018 13:45:40",
//	                   "rate": "31.3000000",
//	                   "rate_delta": "0.0000000",
//	                   "nbuRate": "30.9158950"
//	           }
//	   }
//	}
//
// Де:
//
//	from_cache  – ознака кешування даних;
//	cache_time  – час кешування (в мілісекундах);
//	server_time – поточний час сервера (в мілісекундах);
//	B           – купівля;
//	S           – продаж;
//	date        – дата курсу;
//	rate        – курс;
//	rate_delta  – зміна курсу;
//	nbuRate     – курс НБУ.
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
	api.logResponse(resp)

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
