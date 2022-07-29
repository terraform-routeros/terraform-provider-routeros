package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateInterfaceVrrpTestObjects() (*InterfaceVrrp, error) {
	c := NewClient(GetCredentialsFromEnvVar())
	interface_vrrp := new(InterfaceVrrp)
	interface_vrrp.Name = "test_vrrp"
	interface_vrrp.Interface = "ether1"
	res, err := c.CreateInterfaceVrrp(interface_vrrp)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func TestCreateInterfaceVrrp(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	interface_vrrp := new(InterfaceVrrp)
	interface_vrrp.Name = "test_vrrp"
	interface_vrrp.Interface = "ether1"
	res, err := c.CreateInterfaceVrrp(interface_vrrp)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the to have an id")
	assert.Equal(t, interface_vrrp.Name, res.Name)
	assert.Equal(t, interface_vrrp.Interface, res.Interface)
	err = c.DeleteInterfaceVrrp(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadInterfaceVrrp(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	interface_vrrp, err := CreateInterfaceVrrpTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	res, err := c.ReadInterfaceVrrp(interface_vrrp.ID)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.ID, interface_vrrp.ID)
	assert.Equal(t, interface_vrrp.Name, res.Name)
	assert.Equal(t, interface_vrrp.Interface, res.Interface)
	err = c.DeleteInterfaceVrrp(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestUpdateInterfaceVrrp(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	interface_vrrp, err := CreateInterfaceVrrpTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	interface_vrrp_id := interface_vrrp.ID
	new_interface_vrrp := InterfaceVrrp{}
	new_interface_vrrp.Name = "test_vrrp2"
	new_interface_vrrp.Interface = "ether1"
	res, err := c.UpdateInterfaceVrrp(interface_vrrp_id, &new_interface_vrrp)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.Interface, new_interface_vrrp.Interface)
	err = c.DeleteInterfaceVrrp(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestDeleteInterfaceVrrp(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	interface_vrrp, err := CreateInterfaceVrrpTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	err = c.DeleteInterfaceVrrp(interface_vrrp)
	assert.Nil(t, err, "expecting a nil error on delete")
}
