package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateDhcpServerNetworkTestObjects() (*DhcpServerNetwork, error) {
	c := NewClient(GetCredentialsFromEnvVar())
	dhcp_server_network := new(DhcpServerNetwork)
	dhcp_server_network.Address = "192.168.1.0/24"
	res, err := c.CreateDhcpServerNetwork(dhcp_server_network)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func TestCreateDhcpServerNetwork(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	dhcp_server_network := new(DhcpServerNetwork)
	dhcp_server_network.Address = "192.168.1.0/24"
	res, err := c.CreateDhcpServerNetwork(dhcp_server_network)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the to have an id")
	assert.Equal(t, dhcp_server_network.Address, res.Address)
	err = c.DeleteDhcpServerNetwork(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadDhcpServerNetwork(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	dhcp_server_network, err := CreateDhcpServerNetworkTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	res, err := c.ReadDhcpServerNetwork(dhcp_server_network.ID)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.ID, dhcp_server_network.ID)
	assert.Equal(t, res.Address, dhcp_server_network.Address)
	err = c.DeleteDhcpServerNetwork(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestUpdateDhcpServerNetwork(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	dhcp_server_network, err := CreateDhcpServerNetworkTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	dhcp_server_network_id := dhcp_server_network.ID
	new_server := DhcpServerNetwork{}
	new_server.Address = "192.168.2.0/24"
	res, err := c.UpdateDhcpServerNetwork(dhcp_server_network_id, &new_server)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.Address, new_server.Address)
	err = c.DeleteDhcpServerNetwork(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}
