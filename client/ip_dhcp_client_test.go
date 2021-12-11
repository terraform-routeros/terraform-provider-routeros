package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateTestObjects() (*DhcpClient, error) {
	c := NewClient(GetCredentialsFromEnvVar())
	newClient := new(DhcpClient)
	newClient.Disabled = "true"
	newClient.Interface = "bridge"
	newClient.UsePeerDNS = "true"
	newClient.UsePeerNtp = "true"
	res, err := c.CreateDhcpClient(newClient)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func TestCreateDhcpClient(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newClient := new(DhcpClient)
	newClient.Disabled = "true"
	newClient.Interface = "bridge"
	newClient.UsePeerDNS = "true"
	newClient.UsePeerNtp = "true"
	res, err := c.CreateDhcpClient(newClient)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the to have an id")
	assert.Equal(t, newClient.UsePeerDNS, res.UsePeerDNS)
	assert.Equal(t, newClient.UsePeerNtp, res.UsePeerNtp)
	assert.Equal(t, newClient.Disabled, res.Disabled)
	assert.Equal(t, newClient.Interface, res.Interface)
	err = c.DeleteDhcpClient(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadDhcpClient(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	dhcp_client, err := CreateTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	res, err := c.ReadDhcpClient(dhcp_client.ID)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.ID, dhcp_client.ID)
	assert.Equal(t, res.Disabled, dhcp_client.Disabled)
	assert.Equal(t, res.Interface, dhcp_client.Interface)
	assert.Equal(t, res.UsePeerDNS, dhcp_client.UsePeerDNS)
	assert.Equal(t, res.UsePeerNtp, dhcp_client.UsePeerNtp)
	err = c.DeleteDhcpClient(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestUpdateDhcpClient(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	dhcp_client, err := CreateTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	dhcp_client_id := dhcp_client.ID
	new_client := DhcpClient{}
	new_client.UsePeerDNS = "false"
	res, err := c.UpdateDhcpClient(dhcp_client_id, &new_client)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.UsePeerDNS, new_client.UsePeerDNS)
	err = c.DeleteDhcpClient(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}
