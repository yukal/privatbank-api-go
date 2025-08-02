package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	pb "github.com/yukal/privatbank-api-go"
)

type ApiPresentation struct {
	api         *pb.API
	logger      io.Writer
	bankAccount string
}

var eol string = "\n"

func init() {
	if runtime.GOOS == "windows" {
		eol = "\r\n"
	}
}

func NewPresentationWrapper(api *pb.API, logger io.Writer) *ApiPresentation {
	return &ApiPresentation{
		bankAccount: "UA991112220000026001234567890",
		api:         api,
		logger:      logger,
	}
}

// ......................................................................
// Public API

func (p *ApiPresentation) PublicGetCurrency() {
	rw, err := pb.GetCurrency(5)

	if err != nil {
		p.writeError(err)
	}

	dump(p, rw)
}

func (p *ApiPresentation) PublicGetCurrencyHistoryAt() {
	rw, err := pb.GetCurrencyHistoryAt("2025-04-30")

	if err != nil {
		p.writeError(err)
	}

	dump(p, rw)
}

// ......................................................................
// Corporate API

func (p *ApiPresentation) GetStatementsSettings() {
	var (
		rw  pb.ResponseWrapper[pb.SettingsStatement]
		err error
	)

	if rw, err = p.api.GetSettingsStatement(); err != nil {
		p.writeError(err)
		return
	}

	dump(p, rw)
}

// Transactions

func (p *ApiPresentation) GetStatementsTransactionsAt() {
	var (
		limit     uint16 = 100
		startDate        = "28-04-2025"
		endDate          = "30-04-2025"

		rw  pb.ResponseWrapper[[]pb.TransactionStatement]
		err error
	)

	if rw, err = p.api.GetTransactionsAt(
		p.bankAccount, startDate, endDate, limit); err != nil {

		p.writeError(err)
		return
	}

	dump(p, rw)
}

func (p *ApiPresentation) GetStatementsTransactionsInterim() {
	var (
		limit uint16 = 100

		rw  pb.ResponseWrapper[[]pb.TransactionStatement]
		err error
	)

	if rw, err = p.api.GetInterimTransactions(
		p.bankAccount, limit); err != nil {

		p.writeError(err)
		return
	}

	dump(p, rw)
}

func (p *ApiPresentation) GetStatementsTransactionsFinal() {
	var (
		limit uint16 = 100

		rw  pb.ResponseWrapper[[]pb.TransactionStatement]
		err error
	)

	if rw, err = p.api.GetFinalTransactions(
		p.bankAccount, limit); err != nil {

		p.writeError(err)
		return
	}

	dump(p, rw)
}

// Balance

func (p *ApiPresentation) GetStatementsBalanceAt() {
	var (
		// date = "01-05-2025"
		date = "31-05-2025"

		rw  pb.ResponseWrapper[pb.BalanceStatement]
		err error
	)

	if rw, err = p.api.GetBalanceAt(p.bankAccount, date); err != nil {
		p.writeError(err)
		return
	}

	dump(p, rw)
}

func (p *ApiPresentation) GetStatementsBalancesAt() {
	var (
		limit     uint16 = 100
		startDate        = "28-04-2025"
		endDate          = "30-04-2025"

		rw  pb.ResponseWrapper[[]pb.BalanceStatement]
		err error
	)

	if rw, err = p.api.GetBalancesAt(
		p.bankAccount, startDate, endDate, limit); err != nil {
		p.writeError(err)
		return
	}

	dump(p, rw)
}

func (p *ApiPresentation) GetStatementsBalancesInterim() {
	var (
		limit uint16 = 100

		rw  pb.ResponseWrapper[[]pb.BalanceStatement]
		err error
	)

	if rw, err = p.api.GetInterimBalances(
		p.bankAccount, limit); err != nil {

		p.writeError(err)
		return
	}

	dump(p, rw)
}

func (p *ApiPresentation) GetStatementsBalanceFinal() {
	var (
		rw  pb.ResponseWrapper[pb.BalanceStatement]
		err error
	)

	if rw, err = p.api.GetBalance(p.bankAccount); err != nil {
		p.writeError(err)
		return
	}

	dump(p, rw)
}

// Currency

func (p *ApiPresentation) GetCurrency() {
	var (
		rw  pb.ResponseWrapper[pb.ResponseCurrency]
		err error
	)

	if rw, err = p.api.GetCurrency(); err != nil {
		p.writeError(err)
		return
	}

	dump(p, rw)
}

func (p *ApiPresentation) GetCurrencyHistory() {
	var (
		startDate = "14-05-2025"
		endDate   = "16-05-2025"

		rw  pb.ResponseWrapper[pb.ResponseCurrencyHistory]
		err error
	)

	if rw, err = p.api.GetCurrencyHistory(startDate, endDate); err != nil {
		p.writeError(err)
		return
	}

	dump(p, rw)
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

// ......................................................................
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

	if filename, err = pb.ExtractFilenameFromContentDisposition(resp.Header); err != nil {
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

	if filename, err = pb.ExtractFilenameFromContentDisposition(resp.Header); err != nil {
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

	if filename, err = pb.ExtractFilenameFromContentDisposition(resp.Header); err != nil {
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

func (p *ApiPresentation) writeError(err error) {
	buf := append(
		[]byte("Error\n"),
		[]byte(err.Error()+eol+eol)...,
	)

	p.logger.Write(buf)
}

func (p *ApiPresentation) writeBody(body []byte) {
	buf := append([]byte("Body"+eol), body...)
	buf = append(buf, []byte(eol+eol)...)

	p.logger.Write(buf)
}

func (p *ApiPresentation) logResponse(resp *http.Response) {
	text := resp.Status + " " +
		resp.Proto + " " +
		resp.Request.Method + " " +
		resp.Request.URL.String() + " [" +
		strconv.FormatInt(resp.ContentLength, 10) + "b]" + eol +
		"Header" + eol

	for key, item := range resp.Header {
		text += "  " + key + ": " + strings.Join(item, "; ") + eol
	}

	if cookies := resp.Cookies(); len(cookies) > 0 {
		text += eol + "Cookies"

		for _, item := range resp.Cookies() {
			text += "  " + item.String() + eol
		}
	}

	p.logger.Write([]byte(text + eol))
}

func dump[TRespData any](p *ApiPresentation, data pb.ResponseWrapper[TRespData]) {
	var (
		bytes []byte
		err   error
	)

	if bytes, err = json.MarshalIndent(data.Payload, "", "  "); err != nil {
		fmt.Fprintf(p.logger, "Data:\n%+v\n%+v\n", string(bytes), err)
		return
	}

	p.logResponse(data.Response)

	// fmt.Fprintf(p.logger, "%+v\n", data)
	fmt.Fprintf(p.logger, "Body\n%s\n\n", string(bytes))
}
