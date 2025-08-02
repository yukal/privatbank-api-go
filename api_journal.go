package privatbankapi

import (
	"io"
	"net/http"
)

// Paysheets.

// Get the payroll projects journal.
//
// Response:
//
//	`content-type: application/json;charset=UTF-8`
//	[
//	  {
//	    "uploadTime": "2021-05-20T00:00:00.000",
//	    "project": {
//	      "reference": "12A08000132",
//	      "name": "TEST TEST MCB_TEST"
//	    },
//	    "reportYm": "202103",
//	    "fileType": "PAYSHEET",
//	    "status": "UPLOADED"
//	  },
//	  { ... },
//	  { ... },
//	  ...
//	]
func (a *API) GetPaysheetsJournal() (resp *http.Response, err error) {
	apiURL := URL_API_CORPORATE + "/paysheets/journal"

	if resp, err = a.httpAgent.Get(apiURL); err != nil {
		return
	}

	a.logResponse(resp)

	return
}

func (a *API) GetJournalInbox(dateBegin, dateEnd string) (resp *http.Response, err error) {
	var payload io.Reader

	apiURL := URL_API_CORPORATE + "/proxy/edoc/journal/inbox"
	payloadData := map[string]string{
		"dateBegin": dateBegin,
		"dateEnd":   dateEnd,
	}

	if payload, err = toJSONReader(payloadData); err != nil {
		return nil, err
	}

	if resp, err = a.httpAgent.Post(
		apiURL, payload, nil); err != nil {
		return
	}

	a.logResponse(resp)

	return
}

func (a *API) GetJournalOutbox(dateBegin, dateEnd string) (resp *http.Response, err error) {
	var payload io.Reader

	apiURL := URL_API_CORPORATE + "/proxy/edoc/journal/outbox"
	payloadData := map[string]string{
		"dateBegin": dateBegin,
		"dateEnd":   dateEnd,
		"limit":     "3",
	}

	if payload, err = toJSONReader(payloadData); err != nil {
		return nil, err
	}

	if resp, err = a.httpAgent.Post(
		apiURL, payload, nil); err != nil {
		return
	}

	a.logResponse(resp)

	return
}

func (a *API) GetJournalAll(dateBegin, dateEnd string) (resp *http.Response, err error) {
	var payload io.Reader

	apiURL := URL_API_CORPORATE + "/proxy/edoc/journal/all"
	payloadData := map[string]string{
		"dateBegin": dateBegin,
		"dateEnd":   dateEnd,
		"limit":     "3",
	}

	if payload, err = toJSONReader(payloadData); err != nil {
		return nil, err
	}

	if resp, err = a.httpAgent.Post(
		apiURL, payload, nil); err != nil {
		return
	}

	a.logResponse(resp)

	return
}
