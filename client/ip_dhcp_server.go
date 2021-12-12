package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DhcpServer struct {
	ID            string `json:".id,omitempty"`
	AddressPool   string `json:"address-pool,omitempty"`
	Authoritative string `json:"authoritative,omitempty"`
	Disabled      string `json:"disabled,omitempty"`
	Dynamic       string `json:"dynamic,omitempty"`
	Interface     string `json:"interface,omitempty"`
	Invalid       string `json:"invalid,omitempty"`
	LeaseScript   string `json:"lease-script,omitempty"`
	LeaseTime     string `json:"lease-time,omitempty"`
	Name          string `json:"name,omitempty"`
	UseRadius     string `json:"use-radius,omitempty"`
}

func (c *Client) CreateDhcpServer(dhcp_server *DhcpServer) (*DhcpServer, error) {
	reqBody, err := json.Marshal(dhcp_server)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/ip/dhcp-server", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	res := DhcpServer{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadDhcpServer(id string) (*DhcpServer, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/ip/dhcp-server/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	res := DhcpServer{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) UpdateDhcpServer(id string, dhcp_server *DhcpServer) (*DhcpServer, error) {
	reqBody, err := json.Marshal(dhcp_server)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/ip/dhcp-server/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := DhcpServer{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteDhcpServer(dhcp_server *DhcpServer) error {
	id := dhcp_server.ID
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/ip/dhcp-server/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := DhcpServer{}
	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
