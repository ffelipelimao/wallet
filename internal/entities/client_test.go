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
	client, _ := NewClient("test", "fakemail")

	err := client.Update("test2", "fakemail2")
	assert.NotNil(t, err)

	assert.Equal(t, "test2", client.Name)
}

func TestFailUpdatesClientEmptyArgs(t *testing.T) {
	client, _ := NewClient("test", "fakemail")

	err := client.Update("", "")
	assert.NotNil(t, err)
}
