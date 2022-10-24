package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateIpv6AddressTestObjects() (*IPv6Address, error) {
	c := NewClient(GetCredentialsFromEnvVar())
	ipaddr := new(IPv6Address)
	ipaddr.Address = "fd55::1/64"
	ipaddr.Interface = "bridge"
	ipaddr.Disabled = "yes"
	res, err := c.CreateIPv6Address(ipaddr)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func TestGetIpv6Address(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	ipaddr, err := CreateIpv6AddressTestObjects()
	assert.Nil(t, err, "expecting nil error")
	res, err := c.GetIPv6Address(ipaddr.ID)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	assert.Equal(t, res.Address, ipaddr.Address)
	err = c.DeleteIPv6Address(ipaddr.ID)
	assert.Nil(t, err, "Expecting a nil error")
}

func TestCreateIpv6Address(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	ipaddr := new(IPv6Address)
	ipaddr.Address = "fd55::2/64"
	ipaddr.Interface = "bridge"
	ipaddr.Disabled = "yes"
	res, err := c.CreateIPv6Address(ipaddr)
	assert.Nil(t, err, "Expecting a nil error")
	err = c.DeleteIPv6Address(res.ID)
	assert.Nil(t, err, "Expecting a nil error")
}

func TestDeleteIpv6Address(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	ipaddr, err := CreateIpv6AddressTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	err = c.DeleteIPv6Address(ipaddr.ID)
	assert.Nil(t, err, "Expecting a nil error")
}
