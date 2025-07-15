package privatbank

import (
	"bytes"
	"fmt"
	"net/http"
)

// Завантаження підписаного платежу

// Отримання інформації по платежу
func (a *API) GetPaymentInfo(paymentRef string) (resp *http.Response, err error) {
	if resp, err = a.httpAgent.Get("/proxy/payment/get?ref=" + paymentRef); err != nil {
		return
	}

	a.logResponse(resp)

	return
}

func (a *API) GetReceipt(account, reference, refn string) (resp *http.Response, err error) {
	body := bytes.NewBuffer(
		[]byte(
			fmt.Sprintf(`{"transactions":[{"account": "%s", "reference": "%s", "refn": "%s"}],"perPage": 4}`,
				account,
				reference,
				refn,
			),
		),
	)

	headers := map[string]string{
		"Accept": "application/octet-stream",
	}

	if resp, err = a.httpAgent.Post("/paysheets/print_receipt", body, headers); err != nil {
		return
	}

	a.logResponse(resp)

	return
}
