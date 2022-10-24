package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateSystemIdentityTestObjects() (*SystemIdentity, error) {
	c := NewClient(GetCredentialsFromEnvVar())
	system_identity := new(SystemIdentity)
	system_identity.Name = "myself"
	res, err := c.CreateSystemIdentity(system_identity)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func TestCreateSystemIdentity(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	system_identity := new(SystemIdentity)
	system_identity.Name = "myself"
	res, err := c.CreateSystemIdentity(system_identity)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting result to have an id")
	assert.Equal(t, system_identity.Name, res.Name)
	err = c.DeleteSystemIdentity(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadSystemIdentity(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())

	system_identity, err := CreateSystemIdentityTestObjects()
	assert.Nil(t, err, "expecting a nil error")

	res, err := c.ReadSystemIdentity()
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.ID, system_identity.ID)
	assert.Equal(t, system_identity.Name, res.Name)

	err = c.DeleteSystemIdentity(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestUpdateSystemIdentity(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())

	new_system_identity := SystemIdentity{}
	new_system_identity.Name = "yourself"

	res, err := c.UpdateSystemIdentity(&new_system_identity)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.Name, new_system_identity.Name)

	err = c.DeleteSystemIdentity(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}
