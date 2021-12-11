package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateVLAN(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newVlan := new(VLAN)
	newVlan.Name = "VLAN_TEST"
	newVlan.Interface = "bridge"
	newVlan.VlanID = "900"
	newVlan.Disabled = "true"
	res, err := c.CreateVLAN(newVlan)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the VLAN to have an id")
	assert.Equal(t, newVlan.Name, res.Name)
	assert.Equal(t, newVlan.Interface, res.Interface)
	assert.Equal(t, newVlan.VlanID, res.VlanID)
	assert.Equal(t, newVlan.Disabled, res.Disabled)
	err = c.DeleteVLAN(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadVLAN(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newVlan := new(VLAN)
	newVlan.Name = "VLAN_TEST"
	newVlan.Interface = "bridge"
	newVlan.VlanID = "900"
	newVlan.Disabled = "true"
	res, err := c.CreateVLAN(newVlan)
	assert.Nil(t, err, "expecting a nil error")
	vlan_get, err := c.ReadVLAN(res.ID)
	assert.Equal(t, newVlan.Name, vlan_get.Name)
	assert.Equal(t, newVlan.Interface, vlan_get.Interface)
	assert.Equal(t, newVlan.VlanID, vlan_get.VlanID)
	assert.Equal(t, newVlan.Disabled, vlan_get.Disabled)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	err = c.DeleteVLAN(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}
func TestUpdateVLAN(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newVlan := new(VLAN)
	newVlan.Name = "VLAN_TEST"
	newVlan.Interface = "bridge"
	newVlan.VlanID = "900"
	newVlan.Disabled = "true"
	res, err := c.CreateVLAN(newVlan)
	assert.Nil(t, err, "expecting a nil error")
	newVlan.VlanID = "901"
	resp, err := c.UpdateVLAN(res.ID, newVlan)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, newVlan.Name, resp.Name)
	assert.Equal(t, newVlan.Interface, resp.Interface)
	assert.Equal(t, newVlan.VlanID, resp.VlanID)
	assert.Equal(t, newVlan.Disabled, resp.Disabled)
	err = c.DeleteVLAN(resp.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestDeleteVLAN(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newVlan := new(VLAN)
	newVlan.Name = "VLAN_TEST"
	newVlan.Interface = "bridge"
	newVlan.VlanID = "900"
	newVlan.Disabled = "true"
	res, err := c.CreateVLAN(newVlan)
	assert.Nil(t, err, "expecting a nil error")
	err = c.DeleteVLAN(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}
