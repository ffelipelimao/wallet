package gateway

import "github.com/ffelipelimao/walletcore/internal/entities"

type TransactionGateway interface {
	Save(transaction *entities.Transaction) error
}
