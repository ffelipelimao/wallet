package create_account

import (
	"github.com/ffelipelimao/walletcore/internal/entities"
	"github.com/ffelipelimao/walletcore/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(clientGateway gateway.ClientGateway, accountGateway gateway.AccountGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		ClientGateway:  clientGateway,
		AccountGateway: accountGateway,
	}
}

func (c *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := c.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}

	account := entities.NewAccount(client)

	err = c.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{
		ID: account.ID,
	}, nil

}
