package main

import (
	"io"
	"os"

	"github.com/yukal/privatbank"
)

const TOKEN = "YOUR-PERSONAL-PRIVATBANK-API-TOKEN"

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
		Token:    TOKEN,
		Encoding: "utf8",
		Logger:   io.MultiWriter(os.Stdout, logFile),
	})

	pb := NewPresentation(api)

	// ............................
	// Statement

	pb.GetStatementsSettings()

	// pb.GetStatementsBalanceAt()
	// pb.GetStatementsBalancesAt()
	// pb.GetStatementsBalancesInterim()
	// pb.GetStatementsBalanceFinal()

	// pb.GetStatementsTransactionsAt()
	// pb.GetStatementsTransactionsInterim()
	// pb.GetStatementsTransactionsFinal()

	// ............................
	// Currency

	// pb.GetCurrencyHistory()

	// ............................
	// Payment

	// pb.GetPaymentInfo()

	// ............................
	// Journal

	// pb.GetJournalInbox()
	// pb.GetJournalOutbox()
	// pb.GetJournalAll()
	// pb.GetPaysheetsJournal()

	// ............................
	// Other

	// pb.GetReceipt()
	// pb.GetMultipleReceiptsOf2()
	// pb.GetMultipleReceiptsOf4()
}
