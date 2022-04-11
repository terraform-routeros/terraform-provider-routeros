package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCapsManChannel(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newChannel := new(CapsManChannel)
	newChannel.Name = "CapsManChannel_TEST"
	res, err := c.CreateCapsManChannel(newChannel)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the CapsManChannel to have an id")
	assert.Equal(t, newChannel.Name, res.Name)
	err = c.DeleteCapsManChannel(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadCapsManChannel(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newChannel := new(CapsManChannel)
	newChannel.Name = "CapsManChannel_TEST"
	res, err := c.CreateCapsManChannel(newChannel)
	assert.Nil(t, err, "expecting a nil error")
	channel_get, err := c.ReadCapsManChannel(res.ID)
	assert.Equal(t, newChannel.Name, channel_get.Name)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	err = c.DeleteCapsManChannel(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}
func TestUpdateCapsManChannel(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newChannel := new(CapsManChannel)
	newChannel.Name = "CapsManChannel_TEST"
	res, err := c.CreateCapsManChannel(newChannel)
	assert.Nil(t, err, "expecting a nil error")
	newChannel.Name = "CapsManChannel_NEW"
	resp, err := c.UpdateCapsManChannel(res.ID, newChannel)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, newChannel.Name, resp.Name)
	err = c.DeleteCapsManChannel(resp.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestDeleteCapsManChannel(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newChannel := new(CapsManChannel)
	newChannel.Name = "CapsManChannel_TEST"
	res, err := c.CreateCapsManChannel(newChannel)
	assert.Nil(t, err, "expecting a nil error")
	err = c.DeleteCapsManChannel(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}
