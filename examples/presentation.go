package main

import (
	"io"
	"net/http"
	"os"
	"runtime"

	"github.com/yukal/privatbank"
)

type ApiPresentation struct {
	api         *privatbank.API
	bankAccount string
	_EOL        string
}

func NewApiPresentation(api *privatbank.API) *ApiPresentation {
	eol := "\n"

	if runtime.GOOS == "windows" {
		eol = "\r\n"
	}

	return &ApiPresentation{
		bankAccount: "UA991112220000026001234567890",
		api:         api,
		_EOL:        eol,
	}
}

func (p *ApiPresentation) GetStatementsSettings() (err error) {
	var resp *http.Response

	if resp, err = p.api.GetStatementsSettings(); err != nil {
		p.writeError(err)
		return
	}

	defer resp.Body.Close()
	err = p.processBody(resp.Body)

	return
}

func (p *ApiPresentation) GetStatementsBalance() (err error) {
	var resp *http.Response

	startDate := "28-04-2025"
	endDate := "30-04-2025"
	limit := uint16(5)

	if resp, err = p.api.GetStatementsBalance(
		p.bankAccount, startDate, endDate, limit); err != nil {

		p.writeError(err)
		return
	}

	defer resp.Body.Close()
	err = p.processBody(resp.Body)

	return
}

func (p *ApiPresentation) GetStatementsTransactions() (err error) {
	var resp *http.Response

	startDate := "28-04-2025"
	endDate := "30-04-2025"
	limit := uint16(5)

	if resp, err = p.api.GetStatementsTransactions(
		p.bankAccount, startDate, endDate, limit); err != nil {

		p.writeError(err)
		return
	}

	defer resp.Body.Close()
	err = p.processBody(resp.Body)

	return
}

func (p *ApiPresentation) GetStatementsInterimBalance() (err error) {
	var resp *http.Response

	startDate := "28-04-2025"
	endDate := "30-04-2025"
	limit := uint16(5)

	if resp, err = p.api.GetStatementsInterimBalance(
		p.bankAccount, startDate, endDate, limit); err != nil {

		p.writeError(err)
		return
	}

	defer resp.Body.Close()
	err = p.processBody(resp.Body)

	return
}

func (p *ApiPresentation) GetStatementsInterimTransactions() (err error) {
	var resp *http.Response

	startDate := "28-04-2025"
	endDate := "30-04-2025"
	limit := uint16(5)

	if resp, err = p.api.GetStatementsInterimTransactions(
		p.bankAccount, startDate, endDate, limit); err != nil {

		p.writeError(err)
		return
	}

	defer resp.Body.Close()
	err = p.processBody(resp.Body)

	return
}

func (p *ApiPresentation) GetStatementsFinalBalance() (err error) {
	var resp *http.Response

	startDate := "16-05-2025"
	limit := uint16(5)

	if resp, err = p.api.GetStatementsFinalBalance(
		p.bankAccount, startDate, limit); err != nil {

		p.writeError(err)
		return
	}

	defer resp.Body.Close()
	err = p.processBody(resp.Body)

	return
}

func (p *ApiPresentation) GetStatementsFinalTransactions() (err error) {
	var resp *http.Response

	startDate := "28-04-2025"
	endDate := "31-05-2025"
	limit := uint16(5)

	if resp, err = p.api.GetStatementsFinalTransactions(
		p.bankAccount, startDate, endDate, limit); err != nil {

		p.writeError(err)
		return
	}

	defer resp.Body.Close()
	err = p.processBody(resp.Body)

	return
}

func (p *ApiPresentation) GetCurrencyHistory() (err error) {
	var resp *http.Response

	startDate := "14-05-2025"
	endDate := "16-05-2025"

	if resp, err = p.api.GetCurrencyHistory(startDate, endDate); err != nil {
		p.writeError(err)
		return
	}

	defer resp.Body.Close()
	err = p.processBody(resp.Body)

	return
}

func (p *ApiPresentation) GetPaymentInfo() (err error) {
	var resp *http.Response

	paymentRef := "1234567890123456789"

	if resp, err = p.api.GetPaymentInfo(paymentRef); err != nil {
		p.writeError(err)
		return
	}

	defer resp.Body.Close()
	err = p.processBody(resp.Body)

	return
}

// ...............................................................................
// ...............................................................................

func (p *ApiPresentation) GetJournalInbox() (err error) {
	var resp *http.Response

	dateBegin := "28-04-2025"
	dateEnd := "30-04-2025"

	if resp, err = p.api.GetJournalInbox(dateBegin, dateEnd); err != nil {
		p.writeError(err)
		return
	}

	defer resp.Body.Close()
	err = p.processBody(resp.Body)

	return
}

func (p *ApiPresentation) GetJournalOutbox() (err error) {
	var resp *http.Response

	dateBegin := "28-04-2025"
	dateEnd := "30-04-2025"

	if resp, err = p.api.GetJournalOutbox(dateBegin, dateEnd); err != nil {
		p.writeError(err)
		return
	}

	defer resp.Body.Close()
	err = p.processBody(resp.Body)

	return
}

func (p *ApiPresentation) GetJournalAll() (err error) {
	var resp *http.Response

	dateBegin := "28-04-2025"
	dateEnd := "30-04-2025"

	if resp, err = p.api.GetJournalAll(dateBegin, dateEnd); err != nil {
		p.writeError(err)
		return
	}

	defer resp.Body.Close()
	err = p.processBody(resp.Body)

	return
}

func (p *ApiPresentation) GetPaysheetsJournal() (err error) {
	var resp *http.Response

	if resp, err = p.api.GetPaysheetsJournal(); err != nil {
		p.writeError(err)
		return
	}

	defer resp.Body.Close()
	err = p.processBody(resp.Body)

	return
}

func (p *ApiPresentation) GetReceipt() (err error) {
	var (
		resp *http.Response
		body []byte
	)

	reference := "ABCDE4FABB47D7"
	refn := "P"

	saveAs := "receipt-" + reference + ".pdf"

	if resp, err = p.api.GetReceipt(p.bankAccount, reference, refn); err != nil {
		p.writeError(err)
		return
	}

	defer resp.Body.Close()

	if body, err = io.ReadAll(resp.Body); err != nil {
		p.writeError(err)
		return
	}

	if err = os.WriteFile(saveAs, body, 0644); err != nil {
		p.writeError(err)
		return
	}

	return
}

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
	buf := append([]byte("Body"+p._EOL), body...)
	buf = append(buf, []byte(p._EOL+p._EOL)...)

	p.api.Logger.Write(buf)
}

func (p *ApiPresentation) writeError(err error) {
	buf := append(
		[]byte("Error\n"),
		[]byte(err.Error()+p._EOL+p._EOL)...,
	)

	p.api.Logger.Write(buf)
}
