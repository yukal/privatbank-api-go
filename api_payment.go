package privatbankapi

import (
	"io"
	"net/http"
	"net/url"
)

// Download signed payment

// Get payment information
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

// Get payment instruction (receipt) in PDF format.
//
//	account   - the bank account for which the payment was made
//	reference - reference of the completed payment (in transactions this is the REF parameter)
//	refn      - additional payment reference (in transactions this is the REFN parameter)
func (a *API) GetReceipt(account, reference, refn string) (resp *http.Response, err error) {
	var payload io.Reader

	apiURL := URL_API_CORPORATE + "/paysheets/print_receipt"
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

// Get multiple payment instructions (receipts) in PDF format.
//
// You can get no more than 45 payment instructions per request.
// The response will be a PDF file with payment instructions.
//
//	transactions - array with attributes:
//		- account   - account for which the payment was made
//		- reference - reference of the completed payment (in transactions this is the REF parameter)
//		- refn      - additional payment reference (in transactions this is the REFN parameter)
//	perPage - number of payment instructions per page [1..4]
func (a *API) GetMultipleReceipts(transactions []map[string]string, perPage uint8) (resp *http.Response, err error) {
	var payload io.Reader

	apiURL := URL_API_CORPORATE + "/paysheets/print_receipt"
	payloadData := map[string]any{
		"transactions": transactions,
		"perPage":      min(max(perPage, 1), 4), // min 1 .. max 4
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
