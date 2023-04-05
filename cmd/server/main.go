package main

import (
	"database/sql"
	"fmt"

	"github.com/ffelipelimao/walletcore/internal/database"
	"github.com/ffelipelimao/walletcore/internal/event"
	"github.com/ffelipelimao/walletcore/internal/usecase/create_account"
	"github.com/ffelipelimao/walletcore/internal/usecase/create_client"
	"github.com/ffelipelimao/walletcore/internal/usecase/create_transaction"
	"github.com/ffelipelimao/walletcore/pkg/events"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	//eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(clientDb, accountDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(accountDb, transactionDb, eventDispatcher, transactionCreatedEvent)

}
