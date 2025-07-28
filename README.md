# privatbank-api-go
PrivatBank API based on Go.

UNDER CONSTRUCTION.

Available API methods:
- api.[GetSettingsStatement](api_statements.go#L166)   – Отримання серверних дат ([demo]())
- api.[GetBalance](api_statements.go#L206)             – Отримати баланс рахунку за останній підсумковий день ([demo]())
- api.[GetBalanceAt](api_statements.go#L252)           – Отримати баланс рахунку за конкретно вказаний день ([demo]())
- api.[GetInterimBalances](api_statements.go#L310)     – Отримати проміжні дані балансу ([demo]())
- api.[GetBalancesAt](api_statements.go#L342)          – Отримання балансу за певний інтервал ([demo]())
- api.[GetTransactionsAt](api_statements.go#L377)      – Отримання транзакцій за певний інтервал ([demo]())
- api.[GetInterimTransactions](api_statements.go#L401) – Отримання транзакцій проміжних даних ([demo]())
- api.[GetFinalTransactions](api_statements.go#L427)   – Отримати транзакції за останній підсумковий день ([demo]())

See full API demonstration at [examples](./examples/).

```go
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/yukal/privatbank"
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

	api := privatbank.NewAPI(privatbank.APIOptions{
		Encoding: "utf8",
		Token:    "YOUR-PERSONAL-PRIVATBANK-API-TOKEN",
		Logger:   io.MultiWriter(os.Stdout, logFile),
	})

	var (
        // YOUR IBAN (bank account)
		account = "UAXXNNNNNN0000029001234567890"
		data    privatbank.ResponseWrapper[privatbank.BalanceStatement]
	)

	if data, err = api.GetBalance(account); err != nil {
		fmt.Fprintf(api.Logger, "unable get balance: %v", err)
		return
	}

	// data.Response
	// data.RawBody
	// data.Payload

	fmt.Printf("%s %s %s\b",
		data.Payload.AccountName,
		data.Payload.BalanceInEq,
		data.Payload.BalanceOutEq,
	)
}
```
