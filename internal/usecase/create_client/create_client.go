package createclient

import (
	"time"

	"github.com/ffelipelimao/walletcore/internal/entities"
	"github.com/ffelipelimao/walletcore/internal/gateway"
)

type CreateClientInputDTO struct {
	Name  string
	Email string
}

type CreateClientOutputDTO struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateClientClientUseCase struct {
	ClientGateway gateway.ClientGateway
}

func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *CreateClientClientUseCase {
	return &CreateClientClientUseCase{
		ClientGateway: clientGateway,
	}
}

func (c *CreateClientClientUseCase) Execute(input CreateClientInputDTO) (*CreateClientOutputDTO, error) {
	client, err := entities.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	err = c.ClientGateway.Save(client)
	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}, nil

}
