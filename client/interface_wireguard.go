package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type InterfaceWireguard struct {
	ID         string `json:".id,omitempty"`
	Disabled   string `json:"disabled,omitempty"`
	ListenPort string `json:"listen-port,omitempty"`
	Mtu        string `json:"mtu,omitempty"`
	Name       string `json:"name,omitempty"`
	PrivateKey string `json:"private-key,omitempty"`
	PublicKey  string `json:"public-key,omitempty"`
	Running    string `json:"running,omitempty"`
}

func (c *Client) CreateInterfaceWireguard(ip_pool *InterfaceWireguard) (*InterfaceWireguard, error) {
	reqBody, err := json.Marshal(ip_pool)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/interface/wireguard", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	res := InterfaceWireguard{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadInterfaceWireguard(id string) (*InterfaceWireguard, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/interface/wireguard/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	res := InterfaceWireguard{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) UpdateInterfaceWireguard(id string, ip_pool *InterfaceWireguard) (*InterfaceWireguard, error) {
	reqBody, err := json.Marshal(ip_pool)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/interface/wireguard/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := InterfaceWireguard{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteInterfaceWireguard(ip_pool *InterfaceWireguard) error {
	id := ip_pool.ID
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/interface/wireguard/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := InterfaceWireguard{}
	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
