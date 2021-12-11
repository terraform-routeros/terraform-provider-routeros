package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateVLAN(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newVlan := new(VLAN)
	newVlan.Name = "VLAN_TEST"
	newVlan.Interface = "bridge1"
	newVlan.VlanID = "900"
	newVlan.Disabled = "true"
	res, err := c.CreateVLAN(newVlan)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
}

func TestGetVLAN(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	res, err := c.GetIPAddresses("*4")
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}
