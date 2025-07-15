package privatbank

import (
	"fmt"
	"net/http"
)

// ..............................
// Отримання курсів валют

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
func (a *API) GetCurrencyHistory(startDate, endDate string) (resp *http.Response, err error) {
	strURL := fmt.Sprintf("/proxy/currency/history?startDate=%s&endDate=%s",
		startDate,
		endDate,
	)

	if resp, err = a.httpAgent.Get(strURL); err != nil {
		return
	}

	a.logResponse(resp)

	return
}
