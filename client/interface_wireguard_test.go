package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateInterfaceWireguardTestObjects() (*InterfaceWireguard, error) {
	c := NewClient(GetCredentialsFromEnvVar())
	interface_wireguard := new(InterfaceWireguard)
	interface_wireguard.Name = "test_wg_interface"
	interface_wireguard.ListenPort = "13231"
	res, err := c.CreateInterfaceWireguard(interface_wireguard)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func TestCreateInterfaceWireguard(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	interface_wireguard := new(InterfaceWireguard)
	interface_wireguard.Name = "test_wg_interface"
	interface_wireguard.ListenPort = "13231"
	res, err := c.CreateInterfaceWireguard(interface_wireguard)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the to have an id")
	assert.Equal(t, interface_wireguard.Name, res.Name)
	assert.Equal(t, interface_wireguard.ListenPort, res.ListenPort)
	err = c.DeleteInterfaceWireguard(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadInterfaceWireguard(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	interface_wireguard, err := CreateInterfaceWireguardTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	res, err := c.ReadInterfaceWireguard(interface_wireguard.ID)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.ID, interface_wireguard.ID)
	assert.Equal(t, interface_wireguard.Name, res.Name)
	assert.Equal(t, interface_wireguard.ListenPort, res.ListenPort)
	err = c.DeleteInterfaceWireguard(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestUpdateInterfaceWireguard(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	interface_wireguard, err := CreateInterfaceWireguardTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	interface_wireguard_id := interface_wireguard.ID
	new_interface_wireguard := InterfaceWireguard{}
	new_interface_wireguard.ListenPort = "51821"
	res, err := c.UpdateInterfaceWireguard(interface_wireguard_id, &new_interface_wireguard)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.ListenPort, new_interface_wireguard.ListenPort)
	err = c.DeleteInterfaceWireguard(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}
