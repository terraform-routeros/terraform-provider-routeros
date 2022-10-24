package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadIpv6Addresses(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	res, err := c.ReadIPv6Addresses()
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}
