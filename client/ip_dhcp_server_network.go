package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DhcpServerNetwork struct {
	ID            string `json:".id,omitempty"`
	Address       string `json:"address,omitempty"`
	BootFileName  string `json:"boot-file-name,omitempty"`
	CapsManager   string `json:"caps-manager,omitempty"`
	DhcpOption    string `json:"dhcp-option,omitempty"`
	DhcpOptionSet string `json:"dhcp-option-set,omitempty"`
	DnsNone       string `json:"dns-none,omitempty"`
	DnsServer     string `json:"dns-server,omitempty"`
	Domain        string `json:"domain,omitempty"`
	Gateway       string `json:"gateway,omitempty"`
	Netmask       string `json:"netmask,omitempty"`
	NextServer    string `json:"next-server,omitempty"`
	NtpServer     string `json:"ntp-server,omitempty"`
	WinsServer    string `json:"wins-server,omitempty"`
}

func (c *Client) CreateDhcpServerNetwork(dhcp_server_network *DhcpServerNetwork) (*DhcpServerNetwork, error) {
	reqBody, err := json.Marshal(dhcp_server_network)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/ip/dhcp-server/network", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	res := DhcpServerNetwork{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadDhcpServerNetwork(id string) (*DhcpServerNetwork, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/ip/dhcp-server/network/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	res := DhcpServerNetwork{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) UpdateDhcpServerNetwork(id string, dhcp_server_network *DhcpServerNetwork) (*DhcpServerNetwork, error) {
	reqBody, err := json.Marshal(dhcp_server_network)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/ip/dhcp-server/network/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := DhcpServerNetwork{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteDhcpServerNetwork(dhcp_server_network *DhcpServerNetwork) error {
	id := dhcp_server_network.ID
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/ip/dhcp-server/network/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := DhcpServerNetwork{}
	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
