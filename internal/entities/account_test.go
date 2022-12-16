package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("test", "fakemail")
	acc := NewAccount(client)

	assert.NotNil(t, acc)
	assert.Equal(t, client.ID, acc.Client.ID)
}

func TestBlankAccount(t *testing.T) {
	acc := NewAccount(nil)
	assert.Nil(t, acc)
}

func TestCreditAcc(t *testing.T) {
	client, _ := NewClient("test", "fakemail")
	acc := NewAccount(client)

	acc.Credit(100)

	assert.Equal(t, float64(100), acc.Balance)
}

func TestDebitAcc(t *testing.T) {
	client, _ := NewClient("test", "fakemail")
	acc := NewAccount(client)

	acc.Credit(100)
	acc.Debit(50)

	assert.Equal(t, float64(50), acc.Balance)
}
