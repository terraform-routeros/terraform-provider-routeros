package client

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCapsManManager(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	manager_get, err := c.ReadCapsManManager()
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, manager_get, "expecting non-nil result")
}

func TestUpdateCapsManManager(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	manager, err := c.ReadCapsManManager()
	assert.Nil(t, err, "expecting a nil error")
	t.Log(manager.Enabled)
	newManager := new(CapsManManager)
	opp, err := strconv.ParseBool(manager.Enabled)
	newManager.Enabled = strconv.FormatBool(!opp)
	c.UpdateCapsManManager(newManager)
	manager_get, err := c.ReadCapsManManager()
	assert.Nil(t, err, "expecting a nil error on get")
	assert.Equal(t, newManager.Enabled, manager_get.Enabled)
	assert.NotEqual(t, manager.Enabled, manager_get.Enabled)
	newManager.Enabled = strconv.FormatBool(opp)
	c.UpdateCapsManManager(newManager)
}
