package privatbankapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type SettingsStatement struct {
	DateFinalStatement  string   `json:"date_final_statement"`   // дата, включно з якої є підсумкова виписка
	DatesWithoutOperDay []string `json:"dates_without_oper_day"` // дні без операційних днів
	LastDay             string   `json:"lastday"`                // дата минулого операційного дня
	Phase               string   `json:"phase"`                  // фаза роботи (WRK, NOP, ...)
	ServerDateTime      string   `json:"server_date_time"`       // дата та час сервера
	Today               string   `json:"today"`                  // дата поточного операційного дня
	WorkBalance         string   `json:"work_balance"`           // чи проходять регламент завдання, N – можна робити запити, Y – запити не робити
}

type BalanceStatement struct {
	Account        string `json:"acc"`               // рахунок "UACCXXXXX000002600XXXXXXXXXX"
	AccountName    string `json:"nameACC"`           // найменування рахунку (ВЕСЕЛКА ОСББ)
	Atp            string `json:"atp"`               // ?
	BalanceIn      string `json:"balanceIn"`         // вхідний баланс (109064.22)
	BalanceInEq    string `json:"balanceInEq"`       // вхідний баланс в національній валюті (109064.22)
	BalanceOut     string `json:"balanceOut"`        // вихідний баланс (109536.71)
	BalanceOutEq   string `json:"balanceOutEq"`      // вихідний баланс в національній валюті (109536.71)
	BgfIBrnm       string `json:"bgfIBrnm"`          // бранч, що залучив контрагента
	AccountBranch  string `json:"brnm"`              // бранч рахунку (VCSZ)
	Currency       string `json:"currency"`          // валюта рахунку (UAH)
	DateCloseAcc   string `json:"date_close_acc"`    // дата закриття рахунку (01.01.1900 00:00:00)
	DateOpenAccReg string `json:"date_open_acc_reg"` // дата відкриття рахунку (30.01.2023 12:14:34)
	DateOpenAccSys string `json:"date_open_acc_sys"` // дата відкриття рахунку в системі (30.01.2023 00:00:00)
	DPD            string `json:"dpd"`               // дата останнього руху за рахунком (12.07.2025 00:00:00)
	Flmn           string `json:"flmn"`              // (VC/DN)
	IsFinalBal     bool   `json:"is_final_bal"`      // ?
	State          string `json:"state"`             // стан рахунку? (a - активний?, c - закритий?, d - заблокований?)
	TurnoverCred   string `json:"turnoverCred"`      // оборот, кредит (472.49)
	TurnoverCredEq string `json:"turnoverCredEq"`    // оборот, кредит (екв. у нац. валюті) (472.49)
	TurnoverDebt   string `json:"turnoverDebt"`      // оборот, дебет (0.00)
	TurnoverDebtEq string `json:"turnoverDebtEq"`    // оборот, дебет (екв. у нац. валюті) (0.00)
}

type TransactionStatement struct {
	// Receiver

	RecvEDRPOU  string `json:"AUT_MY_CRF"`      // ЄДРПОУ одержувача
	RecvMFO     string `json:"AUT_MY_MFO"`      // МФО одержувача
	RecvAcc     string `json:"AUT_MY_ACC"`      // рахунок одержувача
	RecvName    string `json:"AUT_MY_NAM"`      // назва одержувача
	RecvMFOName string `json:"AUT_MY_MFO_NAME"` // банк одержувача
	RecvMFOCity string `json:"AUT_MY_MFO_CITY"` // назва міста банку

	// Contragent

	CntrEDRPOU  string `json:"AUT_CNTR_CRF"`      // ЄДРПОУ контрагента
	CntrMFO     string `json:"AUT_CNTR_MFO"`      // МФО контрагента
	CntrAcc     string `json:"AUT_CNTR_ACC"`      // рахунок контрагента
	CntrName    string `json:"AUT_CNTR_NAM"`      // назва контрагента
	CntrMFOName string `json:"AUT_CNTR_MFO_NAME"` // назва банку контрагента
	CntrMFOCity string `json:"AUT_CNTR_MFO_CITY"` // назва міста банку

	// Other parameters

	Currency        string `json:"CCY"`                      // валюта
	Reality         string `json:"FL_REAL"`                  // ознака реальності проведення (r,i)
	Status          string `json:"PR_PR"`                    // стан p - проводиться, t - сторнована, r - проведена, n - забракована
	DocumentType    string `json:"DOC_TYP"`                  // тип пл. документа
	DocumentNum     string `json:"NUM_DOC"`                  // номер документа
	ClientDate      string `json:"DAT_KL"`                   // клієнтська дата
	CurrDate        string `json:"DAT_OD"`                   // дата валютування
	OSND            string `json:"OSND"`                     // підстава платежу
	Sum1            string `json:"SUM"`                      // сума
	Sum2            string `json:"SUM_E"`                    // сума в національній валюті (грн)
	RefAct          string `json:"REF"`                      // референс проведення
	RefN            string `json:"REFN"`                     // № з/п всередині проведення
	Time            string `json:"TIM_P"`                    // час проведення
	DateTime        string `json:"DATE_TIME_DAT_OD_TIM_P"`   // дата та час проведення
	TransactionId   string `json:"ID"`                       // ID транзакції
	TransactionType string `json:"TRANTYPE"`                 // тип транзакції дебет/кредит (D, C)
	PaymentRef      string `json:"DLR"`                      // референс платежу сервісу, через який створювали платіж (payment_pack_ref - у разі створення платежу через АPI «Автоклієнт»)
	TTransactionId  string `json:"TECHNICAL_TRANSACTION_ID"` // "4140766673_online"
}

type ResponseSettingsStatement struct {
	Status string            `json:"status"`
	Type   string            `json:"type"`
	Data   SettingsStatement `json:"settings"`
}

type ResponseBalanceStatement struct {
	Status        string             `json:"status"`
	Type          string             `json:"type"`
	ExistNextPage bool               `json:"exist_next_page"`
	NextPageId    string             `json:"next_page_id"`
	Data          []BalanceStatement `json:"balances"`
}

type ResponseTransactionStatement struct {
	Status        string                 `json:"status"`
	Type          string                 `json:"type"`
	ExistNextPage bool                   `json:"exist_next_page"`
	NextPageId    string                 `json:"next_page_id"`
	Data          []TransactionStatement `json:"transactions"`
}

// ResponseBalanceStatement

func (r ResponseBalanceStatement) GetPayloadData() []BalanceStatement {
	return r.Data
}

func (r ResponseBalanceStatement) GetMetaData() ResponseMetaData {
	return ResponseMetaData{
		Status:        r.Status,
		Type:          r.Type,
		ExistNextPage: r.ExistNextPage,
		NextPageId:    r.NextPageId,
	}
}

// ResponseTransactionStatement

func (r ResponseTransactionStatement) GetPayloadData() []TransactionStatement {
	return r.Data
}

func (r ResponseTransactionStatement) GetMetaData() ResponseMetaData {
	return ResponseMetaData{
		Status:        r.Status,
		Type:          r.Type,
		ExistNextPage: r.ExistNextPage,
		NextPageId:    r.NextPageId,
	}

}

// Get server dates.
// If the value of phase is not WRK, API requests during this period may return errors.
//
//	GET /api/statements/settings
//
// Example response:
//
//	{
//		"status": "SUCCESS",
//		"type": "settings",
//		"settings": {
//			"phase": "WRK",
//			"dates_without_oper_day": [
//				"01.01.2018 00:00:00",
//				"30.12.2018 00:00:00",
//				...
//				"01.01.2020 00:00:00"
//			],
//			"today": "30.03.2020 00:00:00", // current operational day (interim statement)
//			"lastday": "29.03.2020 00:00:00", // previous operational day (interim statement)
//			"work_balance": "N", // whether regulatory tasks are running, N – requests allowed, Y – requests not allowed
//			"server_date_time": "30.03.2020 12:03:51",
//			"date_final_statement": "28.03.2020 00:00:00" // date of the final statement
//		}
//	}
func (api *API) GetSettingsStatement() (settings ResponseWrapper[SettingsStatement], err error) {
	var (
		responseData ResponseSettingsStatement
		httpResponse *http.Response
		body         []byte
	)

	apiURL := URL_API_CORPORATE + "/statements/settings"

	if httpResponse, err = api.httpAgent.Get(apiURL); err != nil {
		return
	}

	defer httpResponse.Body.Close()
	api.logResponse(httpResponse)

	if body, err = io.ReadAll(httpResponse.Body); err != nil {
		return
	}

	if err = json.Unmarshal(body, &responseData); err != nil {
		return
	}

	if responseData.Status != RESPONSE_SUCCESS {
		fmt.Fprintf(api.Logger, "Error getting statements settings: %s\n", body)
		err = fmt.Errorf("error getting statements settings: %s", responseData.Status)
		return
	}

	settings = ResponseWrapper[SettingsStatement]{
		Response: httpResponse,
		RawBody:  body,
		Payload:  responseData.Data,
	}

	return
}

// Get account balance for the last final day
func (api *API) GetBalance(accout string) (balance ResponseWrapper[BalanceStatement], err error) {
	var (
		responseData ResponseBalanceStatement
		resp         *http.Response
		body         []byte
	)

	apiURL := URL_API_CORPORATE + "/statements/balance/final?limit=1&acc=" +
		url.QueryEscape(accout)

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

	if responseData.Status != RESPONSE_SUCCESS {
		fmt.Fprintf(api.Logger, "Error getting balance: %s\n", body)
		err = fmt.Errorf("error getting balance: %s", responseData.Status)
		return
	}

	if len(responseData.Data) == 0 {
		err = errors.New("no balances found")
		return
	}

	balance = ResponseWrapper[BalanceStatement]{
		Response: resp,
		RawBody:  body,
		Payload:  responseData.Data[len(responseData.Data)-1],
	}

	return
}

// Get account balance for a specific day
func (api *API) GetBalanceAt(account, date string) (data ResponseWrapper[BalanceStatement], err error) {
	var (
		resp         *http.Response
		responseData ResponseBalanceStatement
		body         []byte
	)

	params := make(url.Values)
	params.Add("acc", account)
	params.Add("startDate", date)
	params.Add("endDate", date)
	// params.Add("limit", strconv.FormatUint(uint64(LIMIT_DATA), 10))

	apiURL := URL_API_CORPORATE + "/statements/balance" + "?" + params.Encode()

	if resp, err = api.httpAgent.Get(apiURL); err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("response(%s)", resp.Status)
		return
	}

	defer resp.Body.Close()
	api.logResponse(resp)

	if body, err = io.ReadAll(resp.Body); err != nil {
		err = fmt.Errorf("io.read-all: %v", err)
		return
	}

	if err = json.Unmarshal(body, &responseData); err != nil {
		err = fmt.Errorf("unmarshal: %v; raw body: '%s'", err, body)
		return
	}

	if responseData.Status != RESPONSE_SUCCESS {
		fmt.Fprintf(api.Logger, "Error getting statements settings: %s\n", body)
		err = fmt.Errorf("error getting statements settings: %s", responseData.Status)
		return
	}

	if len(responseData.Data) == 0 {
		err = errors.New("no balances found")
		return
	}

	data = ResponseWrapper[BalanceStatement]{
		Response: resp,
		RawBody:  body,
		Payload:  responseData.Data[0],
	}

	return
}

// Get interim balances – from lastday to today.
//
// Since the method may generate additional requests depending on the limit argument,
// it will return the last successful response object, and RawBody will contain several
// bodies separated by "\r\n\r\n".
// @see fetchWithinMultipleRequests
//
//	account - bank account (IBAN format)
//	limit   - data limit (per request)
func (api *API) GetInterimBalances(account string, limit uint16) (r ResponseWrapper[[]BalanceStatement], err error) {
	apiURL := URL_API_CORPORATE + "/statements/balance/interim"

	params := make(url.Values)
	params.Add("acc", account)

	if limit <= LIMIT_DATA {
		params.Add("limit", strconv.FormatUint(uint64(limit), 10))
	}

	if r, err = fetchWithinMultipleRequests[BalanceStatement, ResponseBalanceStatement](
		api, apiURL, params); err != nil {
		return
	}

	// api.logResponse(r.Response)

	if len(r.Payload) == 0 {
		err = errors.New("no balances found")
		// return
	}

	return
}

// Get balances for a specific interval.
//
// Since the request may generate additional requests depending on the limit argument,
// it will return the last successful response object, and RawBody will contain several
// bodies separated by "\r\n\r\n".
// @see fetchWithinMultipleRequests
//
//	acc        - bank account number
//	startDate  - DD-MM-YYYY - start date (required)
//	endDate    - DD-MM-YYYY - end date (optional)
//	followId   - next batch ID from response (optional)
//	limit      - number of records per batch (default 20), max 500, recommended no more than 100
func (api *API) GetBalancesAt(account, startDate, endDate string, limit uint16) (r ResponseWrapper[[]BalanceStatement], err error) {
	apiURL := URL_API_CORPORATE + "/statements/balance"

	params := make(url.Values)
	params.Add("acc", account)
	params.Add("startDate", startDate)

	if endDate != "" {
		params.Add("endDate", endDate)
	}

	if limit <= LIMIT_DATA {
		params.Add("limit", strconv.FormatUint(uint64(limit), 10))
	}

	if r, err = fetchWithinMultipleRequests[BalanceStatement, ResponseBalanceStatement](
		api, apiURL, params); err != nil {
		return
	}

	if len(r.Payload) == 0 {
		err = errors.New("no balances found")
		// return
	}

	return
}

// Get transactions for a specific interval.
//
// Since the method may generate additional requests depending on the limit argument,
// it will return the last successful response object, and RawBody will contain several
// bodies separated by "\r\n\r\n".
// @see fetchWithinMultipleRequests
//
//	acc        - bank account number
//	startDate  - DD-MM-YYYY - start date (required)
//	endDate    - DD-MM-YYYY - end date (optional)
//	followId   - next batch ID from response (optional)
//	limit      - number of records per batch (default 20), max 500, recommended no more than 100
func (api *API) GetTransactionsAt(account, startDate, endDate string, limit uint16) (r ResponseWrapper[[]TransactionStatement], err error) {
	apiURL := URL_API_CORPORATE + "/statements/transactions"

	params := make(url.Values)
	params.Add("acc", account)
	params.Add("startDate", startDate)

	if endDate != "" {
		params.Add("endDate", endDate)
	}

	if limit <= LIMIT_DATA {
		params.Add("limit", strconv.FormatUint(uint64(limit), 10))
	}

	if r, err = fetchWithinMultipleRequests[TransactionStatement, ResponseTransactionStatement](
		api, apiURL, params); err != nil {
		return
	}

	// api.logResponse(r.Response)

	if len(r.Payload) == 0 {
		err = errors.New("no transactions found")
		// return
	}

	return
}

// Get interim transactions (from lastday to today)
//
// Since the request may generate additional requests depending on the limit argument,
// it will return the last successful response object, and RawBody will contain several
// bodies separated by "\r\n\r\n".
// @see fetchWithinMultipleRequests
//
//	account - bank account (IBAN format)
//	limit   - data limit (per request)
func (api *API) GetInterimTransactions(account string, limit uint16) (r ResponseWrapper[[]TransactionStatement], err error) {
	apiURL := URL_API_CORPORATE + "/statements/transactions/interim"

	params := make(url.Values)
	params.Add("acc", account)

	if limit <= LIMIT_DATA {
		params.Add("limit", strconv.FormatUint(uint64(limit), 10))
	}

	if r, err = fetchWithinMultipleRequests[TransactionStatement, ResponseTransactionStatement](
		api, apiURL, params); err != nil {
		return
	}

	// api.logResponse(r.Response)

	if len(r.Payload) == 0 {
		err = errors.New("no transactions found")
		// return
	}

	return
}

// Get transactions for the last final day
//
// Since the method may generate additional requests depending on the limit argument,
// it will return the last successful response object, and RawBody will contain several
// bodies separated by "\r\n\r\n".
// @see fetchWithinMultipleRequests
//
//	account - bank account (IBAN format)
//	limit   - data limit (per request)
func (api *API) GetFinalTransactions(account string, limit uint16) (r ResponseWrapper[[]TransactionStatement], err error) {
	apiURL := URL_API_CORPORATE + "/statements/transactions/final"

	params := make(url.Values)
	params.Add("acc", account)

	if limit <= LIMIT_DATA {
		params.Add("limit", strconv.FormatUint(uint64(limit), 10))
	}

	r, err = fetchWithinMultipleRequests[TransactionStatement, ResponseTransactionStatement](
		api, apiURL, params)

	// api.logResponse(r.Response)

	if len(r.Payload) == 0 {
		err = errors.New("no transactions found")
		// return
	}

	return
}
