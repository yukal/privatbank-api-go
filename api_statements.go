package privatbank

import (
	"fmt"
	"net/http"
)

// Отримання серверних дат.
// Якщо значення phase відмінне від WRK, то в цей період запити до API можуть повертатися з помилками.
//
//	GET /api/statements/settings
//
// Приклад відповіді:
//
//	{
//	  "status": "SUCCESS",
//	  "type": "settings",
//	  "settings": {
//	    "phase": "WRK",
//	    "dates_without_oper_day": // дні без опер. днів
//	      "01.01.2018 00:00:00",
//	      "30.12.2018 00:00:00",
//	      "31.12.2018 00:00:00",
//	      "01.01.2019 00:00:00",
//	      "29.12.2019 00:00:00",
//	      "30.12.2019 00:00:00",
//	      "31.12.2019 00:00:00",
//	      "01.01.2020 00:00:00"
//	    ],
//	    "today": "30.03.2020 00:00:00", // дата поточного опер. дня (проміжна виписка)
//	    "lastday": "29.03.2020 00:00:00", // дата минулого опер. дня (проміжна виписка)
//	    "work_balance": "N", // чи проходять регламент завдання, N – можна робити запити, Y – запити не робити
//	    "server_date_time": "30.03.2020 12:03:51",
//	    "date_final_statement": "28.03.2020 00:00:00" // дата, включно з якої є підсумкова виписка
//	  }
//	}
func (a *API) GetStatementsSettings() (resp *http.Response, err error) {
	if resp, err = a.agent.requestGet("/statements/settings"); err != nil {
		return
	}

	a.logResponse(resp)

	return
}

// Отримання серверних дат за певний інтервал для отримання балансів.
//
//	acc        - номер банківського рахунку
//	startDate  - ДД-ММ-РРРР - дата початку (обов’язковий параметр)
//	endDate    - ДД-ММ-РРРР - дата закінчення (необов’язковий параметр)
//	followId   - ID наступної пачки з відповіді (необов’язковий параметр)
//	limit      - кількість записів у пачці (за замовчуванням 20), максимальне значення - 500, рекомендується використовувати не більше 100
func (a *API) GetStatementsBalance(acc, startDate, endDate string, limit uint16) (resp *http.Response, err error) {
	apiURL := fmt.Sprintf("/statements/balance?acc=%s&startDate=%s&endDate=%s&limit=%d",
		acc,
		startDate,
		endDate,
		limit,
	)

	if resp, err = a.agent.requestGet(apiURL); err != nil {
		return
	}

	a.logResponse(resp)

	return
}

// Отримання серверних дат за певний інтервал для отримання транзакцій.
//
//	acc        - номер банківського рахунку
//	startDate  - ДД-ММ-РРРР - дата початку (обов’язковий параметр)
//	endDate    - ДД-ММ-РРРР - дата закінчення (необов’язковий параметр)
//	followId   - ID наступної пачки з відповіді (необов’язковий параметр)
//	limit      - кількість записів у пачці (за замовчуванням 20), максимальне значення - 500, рекомендується використовувати не більше 100
func (a *API) GetStatementsTransactions(acc, startDate, endDate string, limit uint16) (resp *http.Response, err error) {
	strURL := fmt.Sprintf("/statements/transactions?acc=%s&startDate=%s&endDate=%s&limit=%d",
		acc,
		startDate,
		endDate,
		limit,
	)

	if resp, err = a.agent.requestGet(strURL); err != nil {
		return
	}

	a.logResponse(resp)

	return
}

// Отримати проміжні дані – з lastday по today
func (a *API) GetStatementsInterimBalance(acc, startDate, endDate string, limit uint16) (resp *http.Response, err error) {
	apiURL := fmt.Sprintf("/statements/balance/interim?acc=%s&startDate=%s&endDate=%s&limit=%d",
		acc,
		startDate,
		endDate,
		limit,
	)

	if resp, err = a.agent.requestGet(apiURL); err != nil {
		return
	}

	a.logResponse(resp)

	return
}
