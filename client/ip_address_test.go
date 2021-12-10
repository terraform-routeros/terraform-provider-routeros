package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIpAddress(t *testing.T) {
	c := NewClient("BLAH", "BLAH", "BLAH", true)
	ctx := context.Background()
	res, err := c.GetIPAddresses(ctx, "*4")
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}
