package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCapsManProvisioning(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newProvisioning := new(CapsManProvisioning)
	newProvisioning.NameFormat = "prefix"
	res, err := c.CreateCapsManProvisioning(newProvisioning)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the CapsManProvisioning to have an id")
	assert.Equal(t, newProvisioning.NameFormat, res.NameFormat)
	err = c.DeleteCapsManProvisioning(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadCapsManProvisioning(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newProvisioning := new(CapsManProvisioning)
	newProvisioning.NameFormat = "prefix"
	res, err := c.CreateCapsManProvisioning(newProvisioning)
	assert.Nil(t, err, "expecting a nil error")
	provisioning_get, err := c.ReadCapsManProvisioning(res.ID)
	assert.Equal(t, newProvisioning.NameFormat, provisioning_get.NameFormat)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	err = c.DeleteCapsManProvisioning(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}
func TestUpdateCapsManProvisioning(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newProvisioning := new(CapsManProvisioning)
	newProvisioning.NameFormat = "prefix"
	res, err := c.CreateCapsManProvisioning(newProvisioning)
	assert.Nil(t, err, "expecting a nil error")
	newProvisioning.NameFormat = "cap"
	resp, err := c.UpdateCapsManProvisioning(res.ID, newProvisioning)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, newProvisioning.NameFormat, resp.NameFormat)
	err = c.DeleteCapsManProvisioning(resp.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestDeleteCapsManProvisioning(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newProvisioning := new(CapsManProvisioning)
	newProvisioning.NameFormat = "prefix"
	res, err := c.CreateCapsManProvisioning(newProvisioning)
	assert.Nil(t, err, "expecting a nil error")
	err = c.DeleteCapsManProvisioning(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}
