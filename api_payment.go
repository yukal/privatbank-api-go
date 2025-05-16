package privatbank

import (
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
