package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateInterfaceWireguardPeerTestObjects() (*InterfaceWireguardPeer, error) {
	c := NewClient(GetCredentialsFromEnvVar())
	interface_wireguard_peer := new(InterfaceWireguardPeer)
	interface_wireguard_peer.Interface = "wg1"
	interface_wireguard_peer.PublicKey = "QxC+CTcrDdU5+ny0+2ChUH3NegTrwoVCv53TllI5T0I="
	res, err := c.CreateInterfaceWireguardPeer(interface_wireguard_peer)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func TestCreateInterfaceWireguardPeer(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	interface_wireguard_peer := new(InterfaceWireguardPeer)
	interface_wireguard_peer.Interface = "wg1"
	interface_wireguard_peer.PublicKey = "QxC+CTcrDdU5+ny0+2ChUH3NegTrwoVCv53TllI5T0I="
	res, err := c.CreateInterfaceWireguardPeer(interface_wireguard_peer)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the to have an id")
	assert.Equal(t, interface_wireguard_peer.Interface, res.Interface)
	assert.Equal(t, interface_wireguard_peer.PublicKey, res.PublicKey)
	err = c.DeleteInterfaceWireguardPeer(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadInterfaceWireguardPeer(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	interface_wireguard_peer, err := CreateInterfaceWireguardPeerTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	res, err := c.ReadInterfaceWireguardPeer(interface_wireguard_peer.ID)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.ID, interface_wireguard_peer.ID)
	assert.Equal(t, interface_wireguard_peer.Interface, res.Interface)
	assert.Equal(t, interface_wireguard_peer.PublicKey, res.PublicKey)
	err = c.DeleteInterfaceWireguardPeer(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestUpdateInterfaceWireguardPeer(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	interface_wireguard_peer, err := CreateInterfaceWireguardPeerTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	interface_wireguard_peer_id := interface_wireguard_peer.ID
	new_interface_wireguard_peer := InterfaceWireguardPeer{}
	new_interface_wireguard_peer.PublicKey = "s2iOmDMxsjeZZ3UreOmMmL830QMyYOJEjZnpbJCG5yo="
	res, err := c.UpdateInterfaceWireguardPeer(interface_wireguard_peer_id, &new_interface_wireguard_peer)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, res.PublicKey, new_interface_wireguard_peer.PublicKey)
	err = c.DeleteInterfaceWireguardPeer(res)
	assert.Nil(t, err, "expecting a nil error on delete")
}
