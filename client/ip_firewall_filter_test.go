package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateIPFirewallFilterTestObjects() (*IPFirewallFilter, error) {
	c := NewClient(GetCredentialsFromEnvVar())
	filter_rule := new(IPFirewallFilter)
	filter_rule.Chain = "forward"
	filter_rule.SrcAddress = "10.0.0.1"
	filter_rule.DstAddress = "10.0.1.1"
	filter_rule.DstPort = "443"
	filter_rule.Protocol = "udp"
	res, err := c.CreateIPFirewallFilter(filter_rule)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func TestCreateIPFirewallFilter(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	filter_rule := new(IPFirewallFilter)
	filter_rule.Chain = "forward"
	filter_rule.SrcAddress = "10.0.0.1"
	filter_rule.DstAddress = "10.0.1.1"
	filter_rule.DstPort = "443"
	filter_rule.Protocol = "udp"
	res, err := c.CreateIPFirewallFilter(filter_rule)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the to have an id")
	assert.Equal(t, filter_rule.Chain, res.Chain)
	assert.Equal(t, filter_rule.SrcAddress, res.SrcAddress)
	err = c.DeleteIPFirewallFilter(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadIPFirewallFilter(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	filter_rule, err := CreateIPFirewallFilterTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	res, err := c.ReadIPFirewallFilter(filter_rule.ID)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.ID, filter_rule.ID)
	assert.Equal(t, filter_rule.SrcAddress, res.SrcAddress)
	assert.Equal(t, filter_rule.Chain, res.Chain)
	err = c.DeleteIPFirewallFilter(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestUpdateIPFirewallFilter(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	filter_rule, err := CreateIPFirewallFilterTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	filter_rule_id := filter_rule.ID
	new_filter_rule := IPFirewallFilter{}
	new_filter_rule.DstPort = "8443"
	res, err := c.UpdateIPFirewallFilter(filter_rule_id, &new_filter_rule)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.DstPort, new_filter_rule.DstPort)
	err = c.DeleteIPFirewallFilter(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}
