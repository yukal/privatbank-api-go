# privatbank-api-go
PrivatBank API based on Go. [UA](README.UA.md)

[![Go Reference](.github/badges/badge-goref.svg)](https://pkg.go.dev/github.com/yukal/privatbank-api-go)
![Under Construction](.github/badges/badge-underconstruct.svg)


## Install
```bash
# get latest version
go get github.com/yukal/privatbank-api-go

# get a specific version
go get github.com/yukal/privatbank-api-go@v0.15.2
```

Import then:

```bash
# import using alias
import pb "github.com/yukal/privatbank-api-go"

# or:
import (
  . "github.com/yukal/privatbank-api-go"
)
```

Available API methods:
- api.[GetCurrencyHistory](api_currency.go#L123)       – Get currency exchange rate history ([demo](./examples/presentation.go#L180))
- api.[GetSettingsStatement](api_statements.go#L166)   – Get server dates ([demo](./examples/presentation.go#L36))
- api.[GetBalance](api_statements.go#L206)             – Get the account balance for the last closing day ([demo](./examples/presentation.go#L164))
- api.[GetBalanceAt](api_statements.go#L252)           – Get account balance for a specific day ([demo](./examples/presentation.go#L118))
- api.[GetInterimBalances](api_statements.go#L318)     – Get interim balance data ([demo](./examples/presentation.go#L146))
- api.[GetBalancesAt](api_statements.go#L355)          – Get balance for a certain interval ([demo](./examples/presentation.go#L127))
- api.[GetTransactionsAt](api_statements.go#L395)      – Get transactions at a certain interval ([demo](./examples/presentation.go#L52))
- api.[GetInterimTransactions](api_statements.go#L434) – Get intermediate data transactions ([demo](./examples/presentation.go#L72))
- api.[GetFinalTransactions](api_statements.go#L468)   – Get transactions for the last closing day ([demo](./examples/presentation.go#L90))
- api.[GetReceipt](api_payment.go#L32)                 – Receiving a payment instruction (receipt) in PDF format. ([demo](./examples/presentation.go#L310))
- api.[GetMultipleReceipts](api_payment.go#L76)        – Receiving multiple payment instructions (receipts) in PDF format. ([demo1](./examples/presentation.go#L345), [demo2](./examples/presentation.go#L395))

See full API demonstration at [examples](./examples/).

```go
package main

import (
	"fmt"
	"io"
	"os"

	pb "github.com/yukal/privatbank-api-go"
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

	api := pb.NewAPI(pb.APIOptions{
        // Encoding type of responses from the server
		Encoding: "utf8",
		Token:    "YOUR-PERSONAL-PRIVATBANK-API-TOKEN",
		Logger:   io.MultiWriter(os.Stdout, logFile),
	})

	var (
        // Your bank account in IBAN format
		account = "UAXXNNNNNN0000029001234567890"
		data    pb.ResponseWrapper[pb.BalanceStatement]
	)

	if data, err = api.GetBalance(account); err != nil {
		fmt.Fprintf(api.Logger, "unable get balance: %v", err)
		return
	}

	// data.Response
	// data.RawBody
	// data.Payload

	fmt.Printf("%s %s %s\n",
		data.Payload.AccountName,
		data.Payload.BalanceInEq,
		data.Payload.BalanceOutEq,
	)
}
```
