package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateIpv6NeighborDiscoveryTestObjects() (*IPv6NeighborDiscovery, error) {
	c := NewClient(GetCredentialsFromEnvVar())
	nd := new(IPv6NeighborDiscovery)
	nd.Interface = "bridge"
	nd.Disabled = "yes"
	res, err := c.CreateIPv6NeighborDiscovery(nd)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func TestGetIpv6NeighborDiscovery(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	nd, err := CreateIpv6NeighborDiscoveryTestObjects()
	assert.Nil(t, err, "expecting nil error")
	res, err := c.GetIPv6NeighborDiscovery(nd.ID)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	assert.Equal(t, res.Interface, nd.Interface)
}

func TestCreateIpv6NeighborDiscovery(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	nd := new(IPv6NeighborDiscovery)
	nd.Interface = "bridge"
	nd.Disabled = "yes"
	_, err := c.CreateIPv6NeighborDiscovery(nd)
	assert.Nil(t, err, "Expecting a nil error")
}

func TestDeleteIpv6NeighborDiscovery(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	nd, err := CreateIpv6NeighborDiscoveryTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	err = c.DeleteIPv6NeighborDiscovery(nd.ID)
	assert.Nil(t, err, "Expecting a nil error")
}
