package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("test", "fakemail")

	assert.Nil(t, err)
	assert.Equal(t, "test", client.Name)
}

func TestFailCreateNewClientEmptyArgs(t *testing.T) {
	client, err := NewClient("", "")

	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("John Doe Update", "j@j.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe Update", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestFailUpdatesClientEmptyArgs(t *testing.T) {
	client, _ := NewClient("test", "fakemail")

	err := client.Update("", "")
	assert.NotNil(t, err)
}
