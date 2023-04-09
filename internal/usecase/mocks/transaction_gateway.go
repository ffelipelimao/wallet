package mocks

import (
	"github.com/ffelipelimao/walletcore/internal/entities"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Save(transaction *entities.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}
