package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCapsManConfiguration(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newConfiguration := new(CapsManConfiguration)
	newConfiguration.Name = "CapsManConfiguration_TEST"
	res, err := c.CreateCapsManConfiguration(newConfiguration)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the CapsManConfiguration to have an id")
	assert.Equal(t, newConfiguration.Name, res.Name)
	err = c.DeleteCapsManConfiguration(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadCapsManConfiguration(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newConfiguration := new(CapsManConfiguration)
	newConfiguration.Name = "CapsManConfiguration_TEST"
	res, err := c.CreateCapsManConfiguration(newConfiguration)
	assert.Nil(t, err, "expecting a nil error")
	configuration_get, err := c.ReadCapsManConfiguration(res.ID)
	assert.Equal(t, newConfiguration.Name, configuration_get.Name)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	err = c.DeleteCapsManConfiguration(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}
func TestUpdateCapsManConfiguration(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newConfiguration := new(CapsManConfiguration)
	newConfiguration.Name = "CapsManConfiguration_TEST"
	res, err := c.CreateCapsManConfiguration(newConfiguration)
	assert.Nil(t, err, "expecting a nil error")
	newConfiguration.Name = "CapsManConfiguration_NEW"
	resp, err := c.UpdateCapsManConfiguration(res.ID, newConfiguration)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, newConfiguration.Name, resp.Name)
	err = c.DeleteCapsManConfiguration(resp.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestDeleteCapsManConfiguration(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newConfiguration := new(CapsManConfiguration)
	newConfiguration.Name = "CapsManConfiguration_TEST"
	res, err := c.CreateCapsManConfiguration(newConfiguration)
	assert.Nil(t, err, "expecting a nil error")
	err = c.DeleteCapsManConfiguration(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}
