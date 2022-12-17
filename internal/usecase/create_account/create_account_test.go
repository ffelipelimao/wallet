package createaccount

import (
	"testing"

	"github.com/ffelipelimao/walletcore/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entities.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entities.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

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

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entities.NewClient("test", "fakemail")

	cm := &ClientGatewayMock{}
	cm.On("Get", client.ID).Return(client, nil)

	am := &AccountGatewayMock{}
	am.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(cm, am)

	output, err := uc.Execute(CreateAccountInputDTO{
		ClientID: client.ID,
	})

	assert.Nil(t, err)
	assert.NotNil(t, output.ID)

	cm.AssertExpectations(t)
	am.AssertExpectations(t)

	am.AssertNumberOfCalls(t, "Save", 1)
	cm.AssertNumberOfCalls(t, "Get", 1)
}
