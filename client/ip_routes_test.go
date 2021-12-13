package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadIpRoutes(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	res, err := c.ReadIPRoutes()
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}
