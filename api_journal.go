package privatbank

import (
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
