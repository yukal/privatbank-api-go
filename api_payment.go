package privatbank

import (
	"bytes"
	"fmt"
	"net/http"
)

// Завантаження підписаного платежу

// Отримання інформації по платежу
func (a *API) GetPaymentInfo(paymentRef string) (resp *http.Response, err error) {
	if resp, err = a.agent.requestGet("/proxy/payment/get?ref=" + paymentRef); err != nil {
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

	if resp, err = a.agent.requestPostOctet("/paysheets/print_receipt", body); err != nil {
		return
	}

	a.logResponse(resp)

	return
}
