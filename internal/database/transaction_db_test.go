package database

import (
	"database/sql"
	"testing"

	"github.com/ffelipelimao/walletcore/internal/entities"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	transactionDB *TransactionDB
	client        *entities.Client
	client2       *entities.Client
	accountTo     *entities.Account
	accountFrom   *entities.Account
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db

	db.Exec("create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	db.Exec("create table transactions (id varchar(255), account_from_id varchar(255), account_to_id varchar(255), amount int,created_at date)")

	s.client, err = entities.NewClient("fake", "test@test")
	s.Nil(err)
	s.client2, err = entities.NewClient("fake2", "test@test")
	s.Nil(err)

	//creating accounts
	accountFrom := entities.NewAccount(s.client)
	accountFrom.Balance = 1000

	accountTo := entities.NewAccount(s.client2)
	accountTo.Balance = 1000

	s.transactionDB = NewTransactionDB(db)

}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("drop table clients")
	s.db.Exec("drop table accounts")
	s.db.Exec("drop table transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestSave() {
	transaction, err := entities.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)

	err = s.transactionDB.Save(transaction)
	s.Nil(err)
}
