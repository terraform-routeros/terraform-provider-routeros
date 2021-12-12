package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateIPPoolTestObjects() (*IPPool, error) {
	c := NewClient(GetCredentialsFromEnvVar())
	ip_pool := new(IPPool)
	ip_pool.Name = "test_pool"
	ip_pool.Ranges = "10.0.0.100-10.0.0.200"
	res, err := c.CreateIPPool(ip_pool)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func TestCreateIPPool(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	ip_pool := new(IPPool)
	ip_pool.Name = "test_pool"
	ip_pool.Ranges = "10.0.0.100-10.0.0.200"
	res, err := c.CreateIPPool(ip_pool)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the to have an id")
	assert.Equal(t, ip_pool.Name, res.Name)
	assert.Equal(t, ip_pool.Ranges, res.Ranges)
	err = c.DeleteIPPool(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadIPPool(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	ip_pool, err := CreateIPPoolTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	res, err := c.ReadIPPool(ip_pool.ID)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.ID, ip_pool.ID)
	assert.Equal(t, ip_pool.Name, res.Name)
	assert.Equal(t, ip_pool.Ranges, res.Ranges)
	err = c.DeleteIPPool(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestUpdateIPPool(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	ip_pool, err := CreateIPPoolTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	ip_pool_id := ip_pool.ID
	new_ip_pool := IPPool{}
	new_ip_pool.Name = "false"
	res, err := c.UpdateIPPool(ip_pool_id, &new_ip_pool)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.Name, new_ip_pool.Name)
	err = c.DeleteIPPool(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}
