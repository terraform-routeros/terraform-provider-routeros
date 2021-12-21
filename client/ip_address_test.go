package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateIpAddressTestObjects() (*IPAddress, error) {
	c := NewClient(GetCredentialsFromEnvVar())
	ipaddr := new(IPAddress)
	ipaddr.Address = "192.168.88.1/24"
	ipaddr.Interface = "bridge"
	ipaddr.Network = "192.168.88.0"
	ipaddr.Disabled = "yes"
	res, err := c.CreateIPAddress(ipaddr)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func TestGetIpAddress(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	ipaddr, err := CreateIpAddressTestObjects()
	assert.Nil(t, err, "expecting nil error")
	res, err := c.GetIPAddress(ipaddr.ID)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	assert.Equal(t, res.Address, ipaddr.Address)
}

func TestCreateIpAddress(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	ipaddr := new(IPAddress)
	ipaddr.Address = "192.168.88.1/24"
	ipaddr.Interface = "bridge"
	ipaddr.Network = "192.168.88.0"
	ipaddr.Disabled = "yes"
	_, err := c.CreateIPAddress(ipaddr)
	assert.Nil(t, err, "Expecting a nil error")
}

func TestDeleteIpAddress(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	ipaddr, err := CreateIpAddressTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	err = c.DeleteIPAddress(ipaddr.ID)
	assert.Nil(t, err, "Expecting a nil error")
}
