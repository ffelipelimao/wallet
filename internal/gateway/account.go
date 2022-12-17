package gateway

import "github.com/ffelipelimao/walletcore/internal/entities"

type AccountGateway interface {
	Save(account *entities.Account) error
	Get(id string) (*entities.Account, error)
}
