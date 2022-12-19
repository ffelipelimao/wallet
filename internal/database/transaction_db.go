package database

import (
	"database/sql"

	"github.com/ffelipelimao/walletcore/internal/entities"
)

type TransactionDB struct {
	DB *sql.DB
}

func NewTransactionDB(db *sql.DB) *TransactionDB {
	return &TransactionDB{
		DB: db,
	}
}

func (a *TransactionDB) Save(transaction *entities.Transaction) error {

	stmt, err := a.DB.Prepare("INSERT INTO transactions (id,accounts_id_from,accounts_id_to,amount,created_at) values (?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(transaction.ID, transaction.AccountFrom.ID, transaction.AccountTo.ID, transaction.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
