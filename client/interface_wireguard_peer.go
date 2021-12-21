package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type InterfaceWireguardPeer struct {
	ID                     string `json:".id,omitempty"`
	AllowedAddress         string `json:"allowed-address,omitempty"`
	CurrentEndpointAddress string `json:"current-endpoint-address,omitempty"`
	CurrentEndpointPort    string `json:"current-endpoint-port,omitempty"`
	Disabled               string `json:"disabled,omitempty"`
	EndpointAddress        string `json:"endpoint-address,omitempty"`
	EndpointPort           string `json:"endpoint-port,omitempty"`
	Interface              string `json:"interface,omitempty"`
	LastHandshake          string `json:"last-handshake,omitempty"`
	PublicKey              string `json:"public-key,omitempty"`
	Rx                     string `json:"rx,omitempty"`
	Tx                     string `json:"tx,omitempty"`
}

func (c *Client) CreateInterfaceWireguardPeer(ip_pool *InterfaceWireguardPeer) (*InterfaceWireguardPeer, error) {
	reqBody, err := json.Marshal(ip_pool)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/interface/wireguard/peers", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	res := InterfaceWireguardPeer{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadInterfaceWireguardPeer(id string) (*InterfaceWireguardPeer, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/interface/wireguard/peers/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	res := InterfaceWireguardPeer{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) UpdateInterfaceWireguardPeer(id string, ip_pool *InterfaceWireguardPeer) (*InterfaceWireguardPeer, error) {
	reqBody, err := json.Marshal(ip_pool)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/interface/wireguard/peers/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := InterfaceWireguardPeer{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteInterfaceWireguardPeer(ip_pool *InterfaceWireguardPeer) error {
	id := ip_pool.ID
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/interface/wireguard/peers/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := InterfaceWireguardPeer{}
	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
