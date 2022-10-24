package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCapsManSecurity(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newSecurity := new(CapsManSecurity)
	newSecurity.Name = "CapsManSecurity_TEST"
	res, err := c.CreateCapsManSecurity(newSecurity)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the CapsManSecurity to have an id")
	assert.Equal(t, newSecurity.Name, res.Name)
	err = c.DeleteCapsManSecurity(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadCapsManSecurity(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newSecurity := new(CapsManSecurity)
	newSecurity.Name = "CapsManSecurity_TEST"
	res, err := c.CreateCapsManSecurity(newSecurity)
	assert.Nil(t, err, "expecting a nil error")
	security_get, err := c.ReadCapsManSecurity(res.ID)
	assert.Equal(t, newSecurity.Name, security_get.Name)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	err = c.DeleteCapsManSecurity(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}
func TestUpdateCapsManSecurity(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newSecurity := new(CapsManSecurity)
	newSecurity.Name = "CapsManSecurity_TEST"
	res, err := c.CreateCapsManSecurity(newSecurity)
	assert.Nil(t, err, "expecting a nil error")
	newSecurity.Name = "CapsManSecurity_NEW"
	resp, err := c.UpdateCapsManSecurity(res.ID, newSecurity)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, newSecurity.Name, resp.Name)
	err = c.DeleteCapsManSecurity(resp.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestDeleteCapsManSecurity(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newSecurity := new(CapsManSecurity)
	newSecurity.Name = "CapsManSecurity_TEST"
	res, err := c.CreateCapsManSecurity(newSecurity)
	assert.Nil(t, err, "expecting a nil error")
	err = c.DeleteCapsManSecurity(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}
