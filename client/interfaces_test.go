package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInterfaces(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	res, err := c.ReadInterfaces()
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}
