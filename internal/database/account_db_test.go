package database

import (
	"database/sql"
	"testing"

	"github.com/ffelipelimao/walletcore/internal/entities"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entities.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db

	db.Exec("create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")

	s.accountDB = NewAccountDB(db)
	s.client, _ = entities.NewClient("fake", "test@test")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("drop table clients")
	s.db.Exec("drop table accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entities.NewAccount(s.client)

	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestGet() {

	s.db.Exec("INSERT INTO clients (id,name,email,created_at) values (?,?,?,?)", s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt)

	account := entities.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)

	accDB, err := s.accountDB.Get(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accDB.ID)
}
