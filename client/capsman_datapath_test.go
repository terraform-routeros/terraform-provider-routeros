package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCapsManDatapath(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newDatapath := new(CapsManDatapath)
	newDatapath.Name = "CapsManDatapath_TEST"
	res, err := c.CreateCapsManDatapath(newDatapath)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the CapsManDatapath to have an id")
	assert.Equal(t, newDatapath.Name, res.Name)
	err = c.DeleteCapsManDatapath(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadCapsManDatapath(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newDatapath := new(CapsManDatapath)
	newDatapath.Name = "CapsManDatapath_TEST"
	res, err := c.CreateCapsManDatapath(newDatapath)
	assert.Nil(t, err, "expecting a nil error")
	datapath_get, err := c.ReadCapsManDatapath(res.ID)
	assert.Equal(t, newDatapath.Name, datapath_get.Name)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	err = c.DeleteCapsManDatapath(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}
func TestUpdateCapsManDatapath(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newDatapath := new(CapsManDatapath)
	newDatapath.Name = "CapsManDatapath_TEST"
	res, err := c.CreateCapsManDatapath(newDatapath)
	assert.Nil(t, err, "expecting a nil error")
	newDatapath.Name = "CapsManDatapath_NEW"
	resp, err := c.UpdateCapsManDatapath(res.ID, newDatapath)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, newDatapath.Name, resp.Name)
	err = c.DeleteCapsManDatapath(resp.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestDeleteCapsManDatapath(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newDatapath := new(CapsManDatapath)
	newDatapath.Name = "CapsManDatapath_TEST"
	res, err := c.CreateCapsManDatapath(newDatapath)
	assert.Nil(t, err, "expecting a nil error")
	err = c.DeleteCapsManDatapath(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}
