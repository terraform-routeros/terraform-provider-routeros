package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateInterfaceBridgePortTestObjects() (*InterfaceBridgePort, error) {
	c := NewClient(GetCredentialsFromEnvVar())
	bridge_port := new(InterfaceBridgePort)
	bridge_port.Bridge = "bridge"
	bridge_port.Interface = "ether1"
	bridge_port.Disabled = "false"
	bridge_port.Pvid = "10"
	res, err := c.CreateInterfaceBridgePort(bridge_port)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func TestCreateInterfaceBridgePort(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	bridge_port := new(InterfaceBridgePort)
	bridge_port.Bridge = "bridge"
	bridge_port.Interface = "ether1"
	bridge_port.Disabled = "false"
	bridge_port.Pvid = "10"
	res, err := c.CreateInterfaceBridgePort(bridge_port)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the to have an id")
	assert.Equal(t, bridge_port.Bridge, res.Bridge)
	assert.Equal(t, bridge_port.Interface, res.Interface)
	assert.Equal(t, bridge_port.Disabled, res.Disabled)
	assert.Equal(t, bridge_port.Pvid, res.Pvid)
	err = c.DeleteInterfaceBridgePort(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadInterfaceBridgePort(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	bridge_port, err := CreateInterfaceBridgePortTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	res, err := c.ReadInterfaceBridgePort(bridge_port.ID)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.ID, bridge_port.ID)
	assert.Equal(t, bridge_port.Bridge, res.Bridge)
	assert.Equal(t, bridge_port.Interface, res.Interface)
	assert.Equal(t, bridge_port.Disabled, res.Disabled)
	assert.Equal(t, bridge_port.Pvid, res.Pvid)
	err = c.DeleteInterfaceBridgePort(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestUpdateInterfaceBridgePort(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	bridge_port, err := CreateInterfaceBridgePortTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	bridge_port_id := bridge_port.ID
	new_bridge_port := InterfaceBridgePort{}
	new_bridge_port.Pvid = "40"
	res, err := c.UpdateInterfaceBridgePort(bridge_port_id, &new_bridge_port)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.Pvid, new_bridge_port.Pvid)
	err = c.DeleteInterfaceBridgePort(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}
