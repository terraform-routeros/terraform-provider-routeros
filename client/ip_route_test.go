package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateIPRouteTestObjects() (*IPRoute, error) {
	c := NewClient(GetCredentialsFromEnvVar())
	ip_route := new(IPRoute)
	ip_route.Distance = "1"
	ip_route.DstAddress = "10.0.0.0/24"
	ip_route.Gateway = "10.1.99.1"
	res, err := c.CreateIPRoute(ip_route)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func TestCreateIPRoute(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	ip_route := new(IPRoute)
	ip_route.Distance = "1"
	ip_route.DstAddress = "10.0.0.0/24"
	ip_route.Gateway = "10.1.99.1"
	res, err := c.CreateIPRoute(ip_route)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the to have an id")
	assert.Equal(t, ip_route.Distance, res.Distance)
	assert.Equal(t, ip_route.DstAddress, res.DstAddress)
	assert.Equal(t, ip_route.Gateway, res.Gateway)
	err = c.DeleteIPRoute(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadIPRoute(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	ip_route, err := CreateIPRouteTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	res, err := c.ReadIPRoute(ip_route.ID)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.ID, ip_route.ID)
	assert.Equal(t, ip_route.Distance, res.Distance)
	assert.Equal(t, ip_route.DstAddress, res.DstAddress)
	assert.Equal(t, ip_route.Gateway, res.Gateway)
	err = c.DeleteIPRoute(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestUpdateIPRoute(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	ip_route, err := CreateIPRouteTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	ip_route_id := ip_route.ID
	new_ip_route := IPRoute{}
	new_ip_route.Gateway = "10.2.99.1"
	res, err := c.UpdateIPRoute(ip_route_id, &new_ip_route)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.Gateway, new_ip_route.Gateway)
	err = c.DeleteIPRoute(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}
