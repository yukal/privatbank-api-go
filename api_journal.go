package privatbank

import (
	"bytes"
	"fmt"
	"net/http"
)

// Розрахункові листи.

// Отримання журналу зарплатних проєктів.
//
// Відповідь:
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
	if resp, err = a.agent.requestGet("/paysheets/journal"); err != nil {
		return
	}

	a.logResponse(resp)

	return
}

func (a *API) GetJournalInbox(dateBegin, dateEnd string) (resp *http.Response, err error) {
	body := bytes.NewBuffer(
		[]byte(
			fmt.Sprintf(`{"dateBegin":"%s", "dateEnd": "%s", "limit": "3"}`,
				dateBegin,
				dateEnd,
			),
		),
	)

	if resp, err = a.agent.requestPost("/proxy/edoc/journal/inbox", body); err != nil {
		return
	}

	a.logResponse(resp)

	return
}

func (a *API) GetJournalOutbox(dateBegin, dateEnd string) (resp *http.Response, err error) {
	body := bytes.NewBuffer(
		[]byte(
			fmt.Sprintf(`{"dateBegin":"%s", "dateEnd": "%s", "limit": "3"}`,
				dateBegin,
				dateEnd,
			),
		),
	)

	if resp, err = a.agent.requestPost("/proxy/edoc/journal/outbox", body); err != nil {
		return
	}

	a.logResponse(resp)

	return
}

func (a *API) GetJournalAll(dateBegin, dateEnd string) (resp *http.Response, err error) {
	body := bytes.NewBuffer(
		[]byte(
			fmt.Sprintf(`{"dateBegin":"%s", "dateEnd": "%s", "limit": "3"}`,
				dateBegin,
				dateEnd,
			),
		),
	)

	if resp, err = a.agent.requestPost("/proxy/edoc/journal/all", body); err != nil {
		return
	}

	a.logResponse(resp)

	return
}
