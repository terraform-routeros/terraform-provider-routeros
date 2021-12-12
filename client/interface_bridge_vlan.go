package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type InterfaceBridgeVlan struct {
	ID              string `json:".id,omitempty"`
	Bridge          string `json:"bridge,omitempty"`
	CurrentTagged   string `json:"current-tagged,omitempty"`
	CurrentUntagged string `json:"current-untagged,omitempty"`
	Disabled        string `json:"disabled,omitempty"`
	Dynamic         string `json:"dynamic,omitempty"`
	Tagged          string `json:"tagged,omitempty"`
	Untagged        string `json:"untagged,omitempty"`
	VlanIds         string `json:"vlan-ids,omitempty"`
}

func (c *Client) CreateInterfaceBridgeVlan(ip_route *InterfaceBridgeVlan) (*InterfaceBridgeVlan, error) {
	reqBody, err := json.Marshal(ip_route)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/interface/bridge/vlan", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	res := InterfaceBridgeVlan{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadInterfaceBridgeVlan(id string) (*InterfaceBridgeVlan, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/interface/bridge/vlan/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	res := InterfaceBridgeVlan{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) UpdateInterfaceBridgeVlan(id string, ip_route *InterfaceBridgeVlan) (*InterfaceBridgeVlan, error) {
	reqBody, err := json.Marshal(ip_route)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/interface/bridge/vlan/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := InterfaceBridgeVlan{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteInterfaceBridgeVlan(ip_route *InterfaceBridgeVlan) error {
	id := ip_route.ID
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/interface/bridge/vlan/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := InterfaceBridgeVlan{}
	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
