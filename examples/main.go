package main

import (
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
		Token:    mustLoadToken(".data/.token"),
		Encoding: "utf8",
	})

	p := NewPresentationWrapper(api,
		io.MultiWriter(os.Stdout, logFile))

	// Public API

	p.PublicGetCurrency()
	// p.PublicGetCurrencyHistoryAt()

	// ............................
	// Statement

	// p.GetStatementsSettings()

	// p.GetStatementsBalanceAt()
	// p.GetStatementsBalancesAt()
	// p.GetStatementsBalancesInterim()
	// p.GetStatementsBalanceFinal()

	// p.GetStatementsTransactionsAt()
	// p.GetStatementsTransactionsInterim()
	// p.GetStatementsTransactionsFinal()

	// ............................
	// Currency

	// p.GetCurrency()
	// p.GetCurrencyHistory()

	// ............................
	// Journal

	// p.GetJournalInbox()
	// p.GetJournalOutbox()
	// p.GetJournalAll()
	// p.GetPaysheetsJournal()

	// ............................
	// Other

	// p.GetReceipt()
	// p.GetMultipleReceiptsOf2()
	// p.GetMultipleReceiptsOf4()
}

func mustLoadToken(fpath string) string {
	var (
		err  error
		body []byte
	)

	if body, err = os.ReadFile(fpath); err != nil {
		os.Stderr.WriteString("unable to load token: " + err.Error() + "\r\n")
		os.Exit(1)
	}

	return string(body)
}
