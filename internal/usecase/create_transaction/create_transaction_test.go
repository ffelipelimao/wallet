package createtransaction

import (
	"testing"

	"github.com/ffelipelimao/walletcore/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Get(id string) (*entities.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Account), args.Error(1)
}

func (m *AccountGatewayMock) Save(account *entities.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Save(transaction *entities.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entities.NewClient("test", "fakemail")
	account1 := entities.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entities.NewClient("test2", "fakemail2")
	account2 := entities.NewAccount(client2)
	account2.Credit(1000)

	cm := &AccountGatewayMock{}
	cm.On("Get", account1.ID).Return(account1, nil)
	cm.On("Get", account2.ID).Return(account2, nil)

	tm := &TransactionGatewayMock{}
	tm.On("Save", mock.Anything).Return(nil)

	uc := NewCreateTransactionUseCase(cm, tm)

	input := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        float64(100),
	}

	output, err := uc.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output.ID)

	cm.AssertExpectations(t)
	tm.AssertExpectations(t)

	tm.AssertNumberOfCalls(t, "Save", 1)
	cm.AssertNumberOfCalls(t, "Get", 2)
}
