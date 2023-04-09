package create_transaction

import (
	"context"
	"testing"

	"github.com/ffelipelimao/walletcore/internal/entities"
	"github.com/ffelipelimao/walletcore/internal/event"
	"github.com/ffelipelimao/walletcore/internal/usecase/mocks"
	"github.com/ffelipelimao/walletcore/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entities.NewClient("test", "fakemail")
	account1 := entities.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entities.NewClient("test2", "fakemail2")
	account2 := entities.NewAccount(client2)
	account2.Credit(1000)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	cm := &mocks.AccountGatewayMock{}
	cm.On("Get", account1.ID).Return(account1, nil)
	cm.On("Get", account2.ID).Return(account2, nil)

	tm := &mocks.TransactionGatewayMock{}
	tm.On("Save", mock.Anything).Return(nil)

	dispatcher := events.NewEventDispatcher()
	event := event.NewTransactionCreated()

	uc := NewCreateTransactionUseCase(mockUow, dispatcher, event)

	input := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        float64(100),
	}

	ctx := context.Background()

	output, err := uc.Execute(ctx, input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)

}
