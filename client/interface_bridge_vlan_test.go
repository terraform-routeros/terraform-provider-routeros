package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateInterfaceBridgeVlanTestObjects() (*InterfaceBridgeVlan, error) {
	c := NewClient(GetCredentialsFromEnvVar())
	bridge_vlan := new(InterfaceBridgeVlan)
	bridge_vlan.Bridge = "bridge"
	bridge_vlan.Untagged = "ether1"
	bridge_vlan.Tagged = "bridge"
	bridge_vlan.Disabled = "false"
	bridge_vlan.VlanIds = "200"
	res, err := c.CreateInterfaceBridgeVlan(bridge_vlan)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func TestCreateInterfaceBridgeVlan(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	bridge_vlan := new(InterfaceBridgeVlan)
	bridge_vlan.Bridge = "bridge"
	bridge_vlan.Untagged = "ether1"
	bridge_vlan.Tagged = "bridge"
	bridge_vlan.Disabled = "false"
	bridge_vlan.VlanIds = "200"
	res, err := c.CreateInterfaceBridgeVlan(bridge_vlan)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the to have an id")
	assert.Equal(t, bridge_vlan.Bridge, res.Bridge)
	assert.Equal(t, bridge_vlan.Untagged, res.Untagged)
	assert.Equal(t, bridge_vlan.Tagged, res.Tagged)
	assert.Equal(t, bridge_vlan.Disabled, res.Disabled)
	assert.Equal(t, bridge_vlan.VlanIds, res.VlanIds)
	err = c.DeleteInterfaceBridgeVlan(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadInterfaceBridgeVlan(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	bridge_vlan, err := CreateInterfaceBridgeVlanTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	res, err := c.ReadInterfaceBridgeVlan(bridge_vlan.ID)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.ID, bridge_vlan.ID)
	assert.Equal(t, bridge_vlan.Bridge, res.Bridge)
	assert.Equal(t, bridge_vlan.Untagged, res.Untagged)
	assert.Equal(t, bridge_vlan.Tagged, res.Tagged)
	assert.Equal(t, bridge_vlan.Disabled, res.Disabled)
	assert.Equal(t, bridge_vlan.VlanIds, res.VlanIds)
	err = c.DeleteInterfaceBridgeVlan(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestUpdateInterfaceBridgeVlan(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	bridge_vlan, err := CreateInterfaceBridgeVlanTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	bridge_vlan_id := bridge_vlan.ID
	new_bridge_vlan := InterfaceBridgeVlan{}
	new_bridge_vlan.VlanIds = "201"
	res, err := c.UpdateInterfaceBridgeVlan(bridge_vlan_id, &new_bridge_vlan)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.VlanIds, new_bridge_vlan.VlanIds)
	err = c.DeleteInterfaceBridgeVlan(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}
