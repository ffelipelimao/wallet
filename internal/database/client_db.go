package database

import (
	"database/sql"

	"github.com/ffelipelimao/walletcore/internal/entities"
)

type ClientDB struct {
	DB *sql.DB
}

func NewClientDB(db *sql.DB) *ClientDB {
	return &ClientDB{
		DB: db,
	}
}

func (c *ClientDB) Get(id string) (*entities.Client, error) {
	client := &entities.Client{}

	stmt, err := c.DB.Prepare("SELECT id,name,email,created_at FROM clients WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err := row.Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt); err != nil {
		return nil, err
	}
	return client, nil
}

func (c *ClientDB) Save(client *entities.Client) error {

	stmt, err := c.DB.Prepare("INSERT INTO clients (id,name,email,created_at) values (?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(client.ID, client.Name, client.Email, client.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
