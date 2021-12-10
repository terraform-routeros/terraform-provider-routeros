package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIpAddress(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	res, err := c.GetIPAddresses("*4")
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestCreateIpAddress(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	ipaddr := new(IPAddress)
	ipaddr.Address = "10.1.71.5/24"
	ipaddr.Interface = "VLAN_71_BACKUP_WAN"
	ipaddr.Network = "10.1.71.0"
	ipaddr.Disabled = "yes"
	_, err := c.CreateIPAddress(ipaddr)
	assert.Nil(t, err, "Expecting a nil error")
}

func TestDeleteIpAddress(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	id := "*60"
	err := c.DeleteIPAddress(id)
	assert.Nil(t, err, "Expecting a nil error")
}
