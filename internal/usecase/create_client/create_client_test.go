package createclient

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

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)

	uc := NewCreateClientUseCase(m)

	output, err := uc.Execute(CreateClientInputDTO{
		Name:  "name",
		Email: "mail",
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "name", output.Name)
	assert.Equal(t, "mail", output.Email)

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
