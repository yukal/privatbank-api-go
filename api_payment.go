package privatbank

import (
	"io"
	"net/http"
	"net/url"
)

// Завантаження підписаного платежу

// Отримання інформації по платежу
func (a *API) GetPaymentInfo(paymentRef string) (resp *http.Response, err error) {
	params := make(url.Values, 1)
	params.Add("ref", paymentRef)

	apiURL := buildApiURL("/proxy/payment/get", params)

	if resp, err = a.httpAgent.Get(apiURL); err != nil {
		return
	}

	a.logResponse(resp)

	return
}

// Отримання платіжної інструкції (квитанції) у PDF форматі.
//
//	account   - рахунок по якому платіж був проведений
//	reference - референс проведеного платежу (в транзакціях це параметр REF)
//	refn      - додатковий референс платежу (в транзакціях це параметр REFN)
func (a *API) GetReceipt(account, reference, refn string) (resp *http.Response, err error) {
	var payload io.Reader

	apiURL := API_URL + "/paysheets/print_receipt"
	payloadData := map[string]any{
		"transactions": []map[string]string{
			{
				"account":   account,
				"reference": reference,
				"refn":      refn,
			},
		},

		"perPage": 1,
	}

	if payload, err = toJSONReader(payloadData); err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Accept": "application/octet-stream",
	}

	if resp, err = a.httpAgent.Post(
		apiURL, payload, headers); err != nil {
		return
	}

	a.logResponse(resp)

	return
}
