package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadIpAddresses(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	res, err := c.ReadIPAddresses()
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}
