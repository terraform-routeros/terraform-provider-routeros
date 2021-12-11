package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DhcpClient struct {
	ID                   string `json:".id,omitempty"`
	AddDefaultRoute      string `json:"add-default-route,omitempty"`
	Address              string `json:"address,omitempty"`
	DefaultRouteDistance string `json:"default-route-distance,omitempty"`
	DhcpOptions          string `json:"dhcp-options,omitempty"`
	DhcpServer           string `json:"dhcp-server,omitempty"`
	Disabled             string `json:"disabled,omitempty"`
	Dynamic              string `json:"dynamic,omitempty"`
	ExpiresAfter         string `json:"expires-after,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Interface            string `json:"interface,omitempty"`
	Invalid              string `json:"invalid,omitempty"`
	PrimaryDNS           string `json:"primary-dns,omitempty"`
	SecondaryDNS         string `json:"secondary-dns,omitempty"`
	Status               string `json:"status,omitempty"`
	UsePeerDNS           string `json:"use-peer-dns,omitempty"`
	UsePeerNtp           string `json:"use-peer-ntp,omitempty"`
}

func (c *Client) CreateDhcpClient(dhcp_client *DhcpClient) (*DhcpClient, error) {
	if dhcp_client.AddDefaultRoute == "true" {
		dhcp_client.AddDefaultRoute = "yes"
	} else {
		dhcp_client.AddDefaultRoute = "no"
	}
	reqBody, err := json.Marshal(dhcp_client)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/ip/dhcp-client", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	res := DhcpClient{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadDhcpClient(id string) (*DhcpClient, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/ip/dhcp-client/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	res := DhcpClient{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) UpdateDhcpClient(id string, dhcp_client *DhcpClient) (*DhcpClient, error) {
	reqBody, err := json.Marshal(dhcp_client)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/ip/dhcp-client/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := DhcpClient{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteDhcpClient(dhcp_client *DhcpClient) error {
	id := dhcp_client.ID
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/ip/dhcp-client/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := DhcpClient{}
	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
