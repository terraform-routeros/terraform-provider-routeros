package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIpAddress(t *testing.T) {
	c := NewClient("BLAH", "BLAH", "BLAH", true)
	res, err := c.GetIPAddresses("*4")
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}
