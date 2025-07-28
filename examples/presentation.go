package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/yukal/privatbank"
)

type ApiPresentation struct {
	api         *privatbank.API
	bankAccount string
	EOL         string
}

var eol string = "\n"

func init() {
	if runtime.GOOS == "windows" {
		eol = "\r\n"
	}
}

func NewPresentation(api *privatbank.API) *ApiPresentation {
	return &ApiPresentation{
		bankAccount: "UA991112220000026001234567890",
		api:         api,
	}
}

func (p *ApiPresentation) GetStatementsSettings() {
	var (
		data privatbank.ResponseWrapper[privatbank.SettingsStatement]
		err  error
	)

	if data, err = p.api.GetSettingsStatement(); err != nil {
		p.writeError(err)
		return
	}

	p.dump(data.Payload)
}

// Transactions

func (p *ApiPresentation) GetStatementsTransactionsAt() {
	var (
		limit     uint16 = 100
		startDate        = "28-04-2025"
		endDate          = "30-04-2025"

		data []privatbank.TransactionStatement
		err  error
	)

	if data, err = p.api.GetTransactionsAt(
		p.bankAccount, startDate, endDate, limit); err != nil {

		p.writeError(err)
		return
	}

	p.dump(data)
}

func (p *ApiPresentation) GetStatementsTransactionsInterim() {
	var (
		limit uint16 = 100

		items []privatbank.TransactionStatement
		err   error
	)

	if items, err = p.api.GetInterimTransactions(
		p.bankAccount, limit); err != nil {

		p.writeError(err)
		return
	}

	p.dump(items)
}

func (p *ApiPresentation) GetStatementsTransactionsFinal() {
	var (
		limit uint16 = 100

		items []privatbank.TransactionStatement
		err   error
	)

	if items, err = p.api.GetFinalTransactions(
		p.bankAccount, limit); err != nil {

		p.writeError(err)
		return
	}

	p.dump(items)
}

// Balance

func (p *ApiPresentation) GetStatementsBalanceAt() {
	var (
		// date = "01-05-2025"
		date = "31-05-2025"

		data privatbank.ResponseWrapper[privatbank.BalanceStatement]
		err  error
	)

	if data, err = p.api.GetBalanceAt(p.bankAccount, date); err != nil {
		p.writeError(err)
		return
	}

	p.dump(data.Payload)
}

func (p *ApiPresentation) GetStatementsBalancesAt() {
	var (
		limit     uint16 = 100
		startDate        = "28-04-2025"
		endDate          = "30-04-2025"

		items []privatbank.BalanceStatement
		err   error
	)

	if items, err = p.api.GetBalancesAt(
		p.bankAccount, startDate, endDate, limit); err != nil {
		p.writeError(err)
		return
	}

	p.dump(items)
}

func (p *ApiPresentation) GetStatementsBalancesInterim() {
	var (
		limit uint16 = 100

		items []privatbank.BalanceStatement
		err   error
	)

	if items, err = p.api.GetInterimBalances(
		p.bankAccount, limit); err != nil {

		p.writeError(err)
		return
	}

	p.dump(items)
}

func (p *ApiPresentation) GetStatementsBalanceFinal() {
	var (
		data privatbank.ResponseWrapper[privatbank.BalanceStatement]
		err  error
	)

	if data, err = p.api.GetBalance(p.bankAccount); err != nil {
		p.writeError(err)
		return
	}

	p.dump(data.Payload)
}

// Currency

func (p *ApiPresentation) GetCurrencyHistory() {
	var (
		startDate = "14-05-2025"
		endDate   = "16-05-2025"

		data privatbank.ResponseWrapper[privatbank.ResponseCurrencyHistory]
		err  error
	)

	if data, err = p.api.GetCurrencyHistory(startDate, endDate); err != nil {
		p.writeError(err)
		return
	}

	p.dump(data.Payload)
}

// Payment - UNDER CONSTRUCTION

func (p *ApiPresentation) GetPaymentInfo() {
	var (
		paymentRef = "1234567890123456789"

		resp *http.Response
		err  error
	)

	if resp, err = p.api.GetPaymentInfo(paymentRef); err != nil {
		p.writeError(err)
		return
	}

	defer resp.Body.Close()

	if err = p.processBody(resp.Body); err != nil {
		p.writeError(err)
	}
}

// ...............................................................................
// ...............................................................................

// Journal

func (p *ApiPresentation) GetJournalInbox() {
	var (
		dateBegin = "01-06-2025"
		dateEnd   = "30-06-2025"

		resp *http.Response
		err  error
	)

	if resp, err = p.api.GetJournalInbox(dateBegin, dateEnd); err != nil {
		p.writeError(err)
		return
	}

	defer resp.Body.Close()

	if err = p.processBody(resp.Body); err != nil {
		p.writeError(err)
	}
}

func (p *ApiPresentation) GetJournalOutbox() {
	var (
		dateBegin = "01-06-2025"
		dateEnd   = "30-06-2025"

		resp *http.Response
		err  error
	)

	if resp, err = p.api.GetJournalOutbox(dateBegin, dateEnd); err != nil {
		p.writeError(err)
		return
	}

	defer resp.Body.Close()

	if err = p.processBody(resp.Body); err != nil {
		p.writeError(err)
	}
}

func (p *ApiPresentation) GetJournalAll() {
	var (
		dateBegin = "01-06-2025"
		dateEnd   = "30-06-2025"

		resp *http.Response
		err  error
	)

	if resp, err = p.api.GetJournalAll(dateBegin, dateEnd); err != nil {
		p.writeError(err)
		return
	}

	defer resp.Body.Close()

	if err = p.processBody(resp.Body); err != nil {
		p.writeError(err)
	}
}

func (p *ApiPresentation) GetPaysheetsJournal() {
	var (
		resp *http.Response
		err  error
	)

	if resp, err = p.api.GetPaysheetsJournal(); err != nil {
		p.writeError(err)
		return
	}

	defer resp.Body.Close()

	if err = p.processBody(resp.Body); err != nil {
		p.writeError(err)
	}
}

// Receipt

// Получити квитанцію по референсу
// refn - номер квитанції, якщо не вказано, то буде використано "P"

func (p *ApiPresentation) GetReceipt() {
	var (
		filename string
		resp     *http.Response
		body     []byte
		err      error
	)

	reference := "ABCDE4FABC47D7"
	refn := "P"

	if resp, err = p.api.GetReceipt(p.bankAccount, reference, refn); err != nil {
		p.writeError(err)
		return
	}

	defer resp.Body.Close()

	if filename, err = privatbank.ExtractFilenameFromContentDisposition(resp.Header); err != nil {
		filename = "receipt-" + reference + ".pdf"
		p.writeError(err)
		// return
	}

	if body, err = io.ReadAll(resp.Body); err != nil {
		p.writeError(err)
		return
	}

	if err = os.WriteFile(filename, body, 0644); err != nil {
		p.writeError(err)
		return
	}
}

func (p *ApiPresentation) GetMultipleReceiptsOf2() {
	var (
		perPage uint8 = 2
		payload       = []map[string]string{
			{
				"account":   p.bankAccount,
				"reference": "ABCDE4FABC47D7",
				"refn":      "P",
			},
			{
				"account":   p.bankAccount,
				"reference": "ABCDE4FABC47D8",
				"refn":      "1",
			},
		}

		filename string
		resp     *http.Response
		body     []byte
		err      error
	)

	if resp, err = p.api.GetMultipleReceipts(payload, perPage); err != nil {
		p.writeError(err)
		return
	}

	defer resp.Body.Close()

	if filename, err = privatbank.ExtractFilenameFromContentDisposition(resp.Header); err != nil {
		// time.RFC3339:     2006-01-02T15:04:05Z07:00
		// time.RFC3339Nano: 2006-01-02T15:04:05.999999999Z07:00
		fname := time.Now().Format("20060102T150405999999999Z0700")

		filename = "receipt-" + fname + ".pdf"
		p.writeError(err)
		// return
	}

	if body, err = io.ReadAll(resp.Body); err != nil {
		p.writeError(err)
		return
	}

	if err = os.WriteFile(filename, body, 0644); err != nil {
		p.writeError(err)
		return
	}
}

func (p *ApiPresentation) GetMultipleReceiptsOf4() {
	var (
		perPage uint8 = 4
		payload       = []map[string]string{
			{
				"account":   p.bankAccount,
				"reference": "ABCDE4FABC47D7",
				"refn":      "P",
			},
			{
				"account":   p.bankAccount,
				"reference": "ABCDE4FABC47D8",
				"refn":      "1",
			},
			{
				"account":   p.bankAccount,
				"reference": "ABCDE4FABC47D9",
				"refn":      "1",
			},
			{
				"account":   p.bankAccount,
				"reference": "ABCDE4FABC47DA",
				"refn":      "1",
			},
		}

		filename string
		resp     *http.Response
		body     []byte
		err      error
	)

	if resp, err = p.api.GetMultipleReceipts(payload, perPage); err != nil {
		p.writeError(err)
		return
	}

	defer resp.Body.Close()

	if filename, err = privatbank.ExtractFilenameFromContentDisposition(resp.Header); err != nil {
		// time.RFC3339:     2006-01-02T15:04:05Z07:00
		// time.RFC3339Nano: 2006-01-02T15:04:05.999999999Z07:00
		fname := time.Now().Format("20060102T150405999999999Z0700")

		filename = "receipt-" + fname + ".pdf"
		p.writeError(err)
		// return
	}

	if body, err = io.ReadAll(resp.Body); err != nil {
		p.writeError(err)
		return
	}

	if err = os.WriteFile(filename, body, 0644); err != nil {
		p.writeError(err)
		return
	}
}

// ...............................................................................
// Helper methods

func (p *ApiPresentation) processBody(responseBody io.ReadCloser) error {
	if body, err := io.ReadAll(responseBody); err != nil {
		p.writeError(err)
		return err
	} else {
		p.writeBody(body)
	}

	return nil
}

func (p *ApiPresentation) writeBody(body []byte) {
	buf := append([]byte("Body"+eol), body...)
	buf = append(buf, []byte(eol+eol)...)

	p.api.Logger.Write(buf)
}

func (p *ApiPresentation) dump(data any) {
	var (
		bytes []byte
		err   error
	)

	if bytes, err = json.MarshalIndent(data, "", "  "); err != nil {
		fmt.Fprintf(p.api.Logger, "Data:\n%+v\n%+v\n", string(bytes), err)
		return
	}

	// fmt.Fprintf(p.api.Logger, "%+v\n", data)
	fmt.Fprintf(p.api.Logger, "Body\n%s\n\n\n", string(bytes))
}

func (p *ApiPresentation) writeError(err error) {
	buf := append(
		[]byte("Error\n"),
		[]byte(err.Error()+eol+eol)...,
	)

	p.api.Logger.Write(buf)
}
