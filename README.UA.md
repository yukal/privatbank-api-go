# privatbank-api-go

[![PrivatBank](.github/brand/PrivatBank.svg)](https://privatbank.ua/)

[EN](README.md)

Ця бібліотека створена на мові [Go](https://go.dev/) для взаємодії з [ПриватБанк API](https://api.privatbank.ua/) згідно офіційної документації та підтримує як публічну так і для юридичних осіб версії віддаленого інтерфейсу. Вона не являється офіційною розробкою та надбанням ПриватБанк а є особистим внесоком в [open-source](https://github.com/open-source)!

Бібліотека буде корисною у спрощенні інтеграції з сервісами банку, для автоматизації фінансових операцій та роботи з рахунками, транзакціями, валютними курсами, електронними документами тощо. Вона орієнтована на розробників, які бажають використовувати можливості API ПриватБанку у власних проєктах на Go, та підтримує розширення функціоналу відповідно до потреб спільноти.

[![Go Reference](.github/badges/badge-goref.svg)](https://pkg.go.dev/github.com/yukal/privatbank-api-go)

## Встановлення

```bash
# отримати останню версію
go get github.com/yukal/privatbank-api-go

# отримати конкретну версію
go get github.com/yukal/privatbank-api-go@v0.15.2
```

### Демо

```go
package main

import (
	"fmt"
	"io"
	"os"

	// імпорт за допомогою псевдоніма
	pb "github.com/yukal/privatbank-api-go"

	// або:
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
	// Використовуйте статичні методи для взаємодії з публічним API ПриватБанку
	rw, err := pb.GetCurrency(5)
	if err != nil {
		p.writeError(err)
	}

	// ..............................
	// Використовуйте NewAPI для створення інтерфейсу взаємодії 
	// з корпоративним API ПриватБанку
	api := pb.NewAPI(pb.APIOptions{
		// Тип кодування відповідей з серверу
		Encoding: "utf8",
		Token:    "YOUR-PERSONAL-PRIVATBANK-API-TOKEN",
		Logger:   io.MultiWriter(os.Stdout, logFile),
	})

	var (
		// Ваш банківський рахунок у форматі IBAN
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

## Взаємодія з публічним API

- pb.[GetCurrency](api_public.go#L63) ([demo](./examples/presentation.go#L42))
- pb.[GetCurrencyHistoryAt](api_public.go#L147) ([demo](./examples/presentation.go#L52))
- [НЕ ДОДАНО] Оплата частинами
- [НЕ ДОДАНО] Статус оброблення документів та платежу за міжнародними документарними операціями

## Взаємодія з корпоративним API

#### Отримання балансів і транзакцій за рахунками

- api.[GetSettingsStatement](api_statements.go#L166) ([demo](./examples/presentation.go#L65))
- api.[GetBalance](api_statements.go#L204) ([demo](./examples/presentation.go#L193))
- api.[GetBalanceAt](api_statements.go#L248) ([demo](./examples/presentation.go#L139))
- api.[GetBalancesAt](api_statements.go#L349) ([demo](./examples/presentation.go#L156))
- api.[GetInterimBalances](api_statements.go#L312) ([demo](./examples/presentation.go#L175))
- api.[GetTransactionsAt](api_statements.go#L389) ([demo](./examples/presentation.go#L81))
- api.[GetInterimTransactions](api_statements.go#L428) ([demo](./examples/presentation.go#L101))
- api.[GetFinalTransactions](api_statements.go#L462) ([demo](./examples/presentation.go#L119))

#### Робота в режимі групи ПП

- [НЕ ДОДАНО] Отримання списку клієнтів, які входять до групи ПП (тільки якщо «Автоклієнт» створено для групи ПП)

#### Отримання курсів валют

- api.[GetCurrency](api_currency.go#L106) ([demo](./examples/presentation.go#L209))
- api.[GetCurrencyHistory](api_currency.go#L179) ([demo](./examples/presentation.go#L223))

#### Створення платежу

- [НЕ ДОДАНО] Завантаження підписаного платежу

#### Електронний документообіг (в т.ч. Інвойсинг)

- [НЕ ДОДАНО] Журнал документів
- [НЕ ДОДАНО] Завантаження документів XML (без ЕЦП) (в т.ч. це ІНВОЙСИНГ)
- [НЕ ДОДАНО] Завантаження документів XML з одночасним надсиланням контрагентові з ЕЦП (в т.ч. це ІНВОЙСИНГ)
- [НЕ ДОДАНО] Завантаження документів PDF (без ЕЦП)
- [НЕ ДОДАНО] Завантаження документів PDF до Base64 (без ЕЦП)
- [НЕ ДОДАНО] Завантаження документів з ЕЦП з одночасним надсиланням контрагентові (оригінал документа повинен бути в журналі)
- [НЕ ДОДАНО] Надсилання контрагентові завантаженого непідписаного документа
- [НЕ ДОДАНО] Створення платежу на основі рахунка-фактури
- [НЕ ДОДАНО] Видалення документа
- [НЕ ДОДАНО] Отримання XML-документа
- [НЕ ДОДАНО] Отримання Base64-документа (для підписання)
- [НЕ ДОДАНО] Отримання PDF-документа
- [НЕ ДОДАНО] Отримання інформації щодо ЕЦП на документі
- [НЕ ДОДАНО] Отримання документа з ЕЦП (використовується для XML-документів) у форматі .p7s

#### Електронна звітність

- [НЕ ДОДАНО]

#### Зарплатний проєкт

- [НЕ ДОДАНО] Отримання списку доступних груп
- [НЕ ДОДАНО] Список одержувачів у групі
- [НЕ ДОДАНО] Додавання нового співробітника до групи SALARY/STUDENT
- [НЕ ДОДАНО] Робота з відомостями
- [НЕ ДОДАНО] Заголовок пакета maspay
- [НЕ ДОДАНО] Вміст пакета maspay
- [НЕ ДОДАНО] Додавання отримувача у відомість
- [НЕ ДОДАНО] Видалення отримувача з відомості
- [НЕ ДОДАНО] Надсилання відомості maspay на перевірку
- [НЕ ДОДАНО] Створення нової відомості maspay

#### Розрахункові листи

- [НЕ ДОДАНО]

#### Корпоративні картки

- [НЕ ДОДАНО] Зведена інформація за всіма корпораціями для вибраного підприємства
- [НЕ ДОДАНО] Отримання списку карток за конкретною корпорацією
- [НЕ ДОДАНО] Виписка за групою карток
- [НЕ ДОДАНО] Виписка за карткою

#### Сервіс для реєстрації та перевірки контрагентів

- [НЕ ДОДАНО]

#### Отримання квитанцій по платежам

- api.[GetReceipt](api_payment.go#L30) ([demo](./examples/presentation.go#L348))
- api.[GetMultipleReceipts](api_payment.go#L72) ([demo1](./examples/presentation.go#L383), [demo2](./examples/presentation.go#L433))

#### Інструкція. Як додати співробітника, надати йому дозволи та отримати КЕП

- [НЕ ДОДАНО]

#### Контакти відповідального співробітника

- [НЕ ДОДАНО]

Дивіться повну демонстрацію API у [прикладах](./examples/).
