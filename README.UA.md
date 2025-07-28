# privatbank-api-go
API Приватбанку (на Go). [EN](README.md)

В РОЗРОБЦІ.

Доступні API методи:
- api.[GetCurrencyHistory](api_currency.go#L123)       – Отримання історії курсів валют ([demo](./examples/presentation.go#L180))
- api.[GetSettingsStatement](api_statements.go#L166)   – Отримання серверних дат ([demo](./examples/presentation.go#L36))
- api.[GetBalance](api_statements.go#L206)             – Отримати баланс рахунку за останній підсумковий день ([demo](./examples/presentation.go#L164))
- api.[GetBalanceAt](api_statements.go#L252)           – Отримати баланс рахунку за конкретно вказаний день ([demo](./examples/presentation.go#L110))
- api.[GetInterimBalances](api_statements.go#L310)     – Отримати проміжні дані балансу ([demo](./examples/presentation.go#L146))
- api.[GetBalancesAt](api_statements.go#L342)          – Отримання балансу за певний інтервал ([demo](./examples/presentation.go#L127))
- api.[GetTransactionsAt](api_statements.go#L377)      – Отримання транзакцій за певний інтервал ([demo](./examples/presentation.go#L52))
- api.[GetInterimTransactions](api_statements.go#L401) – Отримання транзакцій проміжних даних ([demo](./examples/presentation.go#L72))
- api.[GetFinalTransactions](api_statements.go#L427)   – Отримати транзакції за останній підсумковий день ([demo](./examples/presentation.go#L90))
- api.[GetReceipt](api_payment.go#L32)                 – Отримання платіжної інструкції (квитанції) у PDF форматі. ([demo](./examples/presentation.go#L310))
- api.[GetMultipleReceipts](api_payment.go#L76)        – Отримання кількох платіжних інструкцій (квитанцій) у PDF форматі. ([demo1](./examples/presentation.go#L345), [demo2](./examples/presentation.go#L395))

Дивіться повну демонстрацію API у [прикладах](./examples/).

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
		// Тип кодування відповідей з серверу
		Encoding: "utf8",
		Token:    "YOUR-PERSONAL-PRIVATBANK-API-TOKEN",
		Logger:   io.MultiWriter(os.Stdout, logFile),
	})

	var (
        // Ваш банківський рахунок у форматі IBAN
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
