package create_account

import (
	"testing"

	"github.com/ffelipelimao/walletcore/internal/entities"
	"github.com/ffelipelimao/walletcore/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entities.NewClient("test", "fakemail")

	cm := &mocks.ClientGatewayMock{}
	cm.On("Get", client.ID).Return(client, nil)

	am := &mocks.AccountGatewayMock{}
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
