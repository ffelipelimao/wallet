package createtransaction

import (
	"github.com/ffelipelimao/walletcore/internal/entities"
	"github.com/ffelipelimao/walletcore/internal/gateway"
	"github.com/ffelipelimao/walletcore/pkg/events"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	AccountGateway     gateway.AccountGateway
	TransactionGateway gateway.TransactionGateway
	EventDispatcher    events.EventDispatcherInterface
	TransactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(
	accountGateway gateway.AccountGateway,
	transactionGateway gateway.TransactionGateway,
	eventDispatcher events.EventDispatcherInterface,
	transactionCreated events.EventInterface,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
	}
}

func (c *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := c.AccountGateway.Get(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := c.AccountGateway.Get(input.AccountIDTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entities.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	err = c.TransactionGateway.Save(transaction)
	if err != nil {
		return nil, err
	}

	output := &CreateTransactionOutputDTO{ID: transaction.ID}

	c.TransactionCreated.SetPayload(output)
	c.EventDispatcher.Dispatch(c.TransactionCreated)
	if err != nil {
		return nil, err
	}

	return output, nil
}
