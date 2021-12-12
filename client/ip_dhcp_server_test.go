package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateDhcpServerTestObjects() (*DhcpServer, error) {
	c := NewClient(GetCredentialsFromEnvVar())
	dhcp_server := new(DhcpServer)
	dhcp_server.Disabled = "true"
	dhcp_server.Interface = "bridge"
	dhcp_server.AddressPool = "dhcp"
	dhcp_server.Name = "test_dhcp"
	res, err := c.CreateDhcpServer(dhcp_server)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func TestCreateDhcpServer(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	dhcp_server := new(DhcpServer)
	dhcp_server.Disabled = "true"
	dhcp_server.Interface = "bridge"
	dhcp_server.AddressPool = "dhcp"
	dhcp_server.Name = "test_dhcp"
	res, err := c.CreateDhcpServer(dhcp_server)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the to have an id")
	assert.Equal(t, dhcp_server.AddressPool, res.AddressPool)
	assert.Equal(t, dhcp_server.Name, res.Name)
	assert.Equal(t, dhcp_server.Disabled, res.Disabled)
	assert.Equal(t, dhcp_server.Interface, res.Interface)
	err = c.DeleteDhcpServer(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadDhcpServer(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	dhcp_server, err := CreateDhcpServerTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	res, err := c.ReadDhcpServer(dhcp_server.ID)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.ID, dhcp_server.ID)
	assert.Equal(t, res.Disabled, dhcp_server.Disabled)
	assert.Equal(t, res.Interface, dhcp_server.Interface)
	assert.Equal(t, res.AddressPool, dhcp_server.AddressPool)
	assert.Equal(t, res.Name, dhcp_server.Name)
	err = c.DeleteDhcpServer(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestUpdateDhcpServer(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	dhcp_server, err := CreateDhcpServerTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	dhcp_server_id := dhcp_server.ID
	new_server := DhcpServer{}
	new_server.Name = "false"
	res, err := c.UpdateDhcpServer(dhcp_server_id, &new_server)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.Name, new_server.Name)
	err = c.DeleteDhcpServer(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}
