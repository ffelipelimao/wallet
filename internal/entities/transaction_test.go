package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("test1", "fakemail")
	acc1 := NewAccount(client1)

	client2, _ := NewClient("test1", "fakemail")
	acc2 := NewAccount(client2)

	acc1.Credit(1000)
	acc2.Credit(1000)

	transaction, err := NewTransaction(acc1, acc2, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)

	assert.Equal(t, float64(1100), acc2.Balance)
	assert.Equal(t, float64(900), acc1.Balance)
}

func TestCreateTransactionWithInsufficientBalance(t *testing.T) {
	client1, _ := NewClient("test1", "fakemail")
	acc1 := NewAccount(client1)

	client2, _ := NewClient("test1", "fakemail")
	acc2 := NewAccount(client2)

	acc1.Credit(1000)
	acc2.Credit(1000)

	transaction, err := NewTransaction(acc1, acc2, 2000)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)

	assert.Equal(t, float64(1000), acc2.Balance)
	assert.Equal(t, float64(1000), acc1.Balance)

	assert.Error(t, err, "insufficient balance")
}
