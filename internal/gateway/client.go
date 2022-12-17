package gateway

import "github.com/ffelipelimao/walletcore/internal/entities"

type ClientGateway interface {
	Get(id string) (*entities.Client, error)
	Save(client *entities.Client) error
}
