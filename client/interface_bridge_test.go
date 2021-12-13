package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateInterfaceBridgeTestObjects() (*InterfaceBridge, error) {
	c := NewClient(GetCredentialsFromEnvVar())
	bridge := new(InterfaceBridge)
	bridge.Name = "test_bridge"
	res, err := c.CreateInterfaceBridge(bridge)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func TestCreateInterfaceBridge(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	bridge := new(InterfaceBridge)
	bridge.Name = "test_bridge"
	res, err := c.CreateInterfaceBridge(bridge)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the to have an id")
	assert.Equal(t, bridge.Name, res.Name)
	err = c.DeleteInterfaceBridge(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadInterfaceBridge(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	bridge, err := CreateInterfaceBridgeTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	res, err := c.ReadInterfaceBridge(bridge.ID)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.ID, bridge.ID)
	assert.Equal(t, bridge.Name, res.Name)
	err = c.DeleteInterfaceBridge(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestUpdateInterfaceBridge(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	bridge, err := CreateInterfaceBridgeTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	bridge_id := bridge.ID
	new_bridge := InterfaceBridge{}
	new_bridge.Name = "test_bridge_foo"
	res, err := c.UpdateInterfaceBridge(bridge_id, &new_bridge)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.Name, new_bridge.Name)
	err = c.DeleteInterfaceBridge(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}
