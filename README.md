# privatbank-api-go

[![PrivatBank](.github/brand/PrivatBank.svg)](https://privatbank.ua/en)

[UA](README.UA.md)

This library is written in [Go](https://go.dev/) for interacting with the [PrivatBank API](https://api.privatbank.ua/) according to the official documentation. It supports both public and corporate versions of the remote interface. The library is not an official product of PrivatBank, but a personal contribution to [open-source](https://github.com/open-source).

The library simplifies integration with bank services, automates financial operations, and enables working with accounts, transactions, currency rates, electronic documents, and more. It is intended for developers who want to use PrivatBank API capabilities in their Go projects and supports extending functionality based on community needs.

[![Go Reference](.github/badges/badge-goref.svg)](https://pkg.go.dev/github.com/yukal/privatbank-api-go)

## Install

```bash
# get latest version
go get github.com/yukal/privatbank-api-go

# get a specific version
go get github.com/yukal/privatbank-api-go@v0.16.0
```

### Demo

```go
package main

import (
	"fmt"
	"io"
	"os"

	// import using alias
	pb "github.com/yukal/privatbank-api-go"

	// or:
	// . "github.com/yukal/privatbank-api-go"
)

func main() {
	var (
		logFile *os.File
		err     error
	)

	if logFile, err = os.OpenFile("api.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644); err != nil {

		panic(err)
	}

	defer logFile.Close()

	// ..............................
	// Use static methods to interact with public PrivatBank API
	rw, err := pb.GetCurrency(5)
	if err != nil {
		p.writeError(err)
	}

	// ..............................
	// Use NewAPI to create interface to interact with
	// corporate PrivatBank API
	api := pb.NewAPI(pb.APIOptions{
		// Encoding type of responses from the server
		Encoding: "utf8",
		Token:    "YOUR-PERSONAL-PRIVATBANK-API-TOKEN",
		Logger:   io.MultiWriter(os.Stdout, logFile),
	})

	var (
		// Your bank account in IBAN format
		account = "UAXXNNNNNN0000029001234567890"
		rw    pb.ResponseWrapper[pb.BalanceStatement]
	)

	if rw, err = api.GetBalance(account); err != nil {
		fmt.Fprintf(api.Logger, "unable get balance: %v", err)
		return
	}

	// rw.Response
	// rw.RawBody
	// rw.Payload

	fmt.Printf("%s %s %s\n",
		rw.Payload.AccountName,
		rw.Payload.BalanceInEq,
		rw.Payload.BalanceOutEq,
	)
}
```

## Interaction with public API

- pb.[GetCurrency](api_public.go#L63) ([demo](./examples/presentation.go#L42))
- pb.[GetCurrencyHistoryAt](api_public.go#L147) ([demo](./examples/presentation.go#L52))
- [NOT IMPLEMENTED] Payment in installments
- [NOT IMPLEMENTED] Status of document and payment processing for international documentary operations

## Interaction with corporate API

#### Obtaining account balances and transactions

- api.[GetSettingsStatement](api_statements.go#L166) ([demo](./examples/presentation.go#L65))
- api.[GetBalance](api_statements.go#L204) ([demo](./examples/presentation.go#L193))
- api.[GetBalanceAt](api_statements.go#L248) ([demo](./examples/presentation.go#L139))
- api.[GetBalancesAt](api_statements.go#L349) ([demo](./examples/presentation.go#L156))
- api.[GetInterimBalances](api_statements.go#L312) ([demo](./examples/presentation.go#L175))
- api.[GetTransactionsAt](api_statements.go#L389) ([demo](./examples/presentation.go#L81))
- api.[GetInterimTransactions](api_statements.go#L428) ([demo](./examples/presentation.go#L101))
- api.[GetFinalTransactions](api_statements.go#L462) ([demo](./examples/presentation.go#L119))

#### Working in PP group mode

- [NOT IMPLEMENTED] Getting a list of clients that are part of a PP group (only if an "Autoclient" is created for a PP group)

#### Get currency rates

- api.[GetCurrency](api_currency.go#L106) ([demo](./examples/presentation.go#L209))
- api.[GetCurrencyHistory](api_currency.go#L179) ([demo](./examples/presentation.go#L223))

#### Creating a Payment

- [NOT IMPLEMENTED] Uploading a signed payment

#### Electronic Document Management (including Invoicing)

- [NOT IMPLEMENTED] Document journal
- [NOT IMPLEMENTED] Uploading XML documents (without digital signature) (including invoicing)
- [NOT IMPLEMENTED] Uploading XML documents with simultaneous sending to the counterparty with digital signature (including invoicing)
- [NOT IMPLEMENTED] Uploading PDF documents (without digital signature)
- [NOT IMPLEMENTED] Uploading PDF documents as Base64 (without digital signature)
- [NOT IMPLEMENTED] Uploading documents with digital signature with simultaneous sending to the counterparty (the original document must be in the journal)
- [NOT IMPLEMENTED] Sending an uploaded unsigned document to the counterparty
- [NOT IMPLEMENTED] Creating a payment based on an invoice
- [NOT IMPLEMENTED] Deleting a document
- [NOT IMPLEMENTED] Getting an XML document
- [NOT IMPLEMENTED] Getting a Base64 document (for signing)
- [NOT IMPLEMENTED] Getting a PDF document
- [NOT IMPLEMENTED] Getting information about the digital signature on a document
- [NOT IMPLEMENTED] Getting a document with digital signature (used for XML documents) in .p7s format

#### Electronic Reporting

- [NOT IMPLEMENTED]

#### Salary Project

- [NOT IMPLEMENTED] Getting a list of available groups
- [NOT IMPLEMENTED] List of recipients in a group
- [NOT IMPLEMENTED] Adding a new employee to the SALARY/STUDENT group
- [NOT IMPLEMENTED] Working with payroll statements
- [NOT IMPLEMENTED] Maspay package header
- [NOT IMPLEMENTED] Maspay package content
- [NOT IMPLEMENTED] Adding a recipient to the payroll statement
- [NOT IMPLEMENTED] Removing a recipient from the payroll statement
- [NOT IMPLEMENTED] Sending the maspay statement for verification
- [NOT IMPLEMENTED] Creating a new maspay statement

#### Payslips

- [NOT IMPLEMENTED]

#### Corporate Cards

- [NOT IMPLEMENTED] Summary information for all corporations for the selected enterprise
- [NOT IMPLEMENTED] Getting a list of cards for a specific corporation
- [NOT IMPLEMENTED] Statement for a group of cards
- [NOT IMPLEMENTED] Statement for a card

#### Service for Registration and Verification of Counterparties

- [NOT IMPLEMENTED]

#### Getting Payment Receipts

- api.[GetReceipt](api_payment.go#L30) ([demo](./examples/presentation.go#L348))
- api.[GetMultipleReceipts](api_payment.go#L72) ([demo1](./examples/presentation.go#L383), [demo2](./examples/presentation.go#L433))

#### Instructions. How to add an employee, grant permissions, and obtain a Qualified Electronic Signature (QES)

- [NOT IMPLEMENTED]

#### Responsible Employee Contacts

- [NOT IMPLEMENTED]

See full API demonstration at [examples](./examples/).
