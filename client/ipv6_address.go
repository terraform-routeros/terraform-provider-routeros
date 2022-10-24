package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type IPv6Address struct {
	ID              string `json:".id,omitempty"`
	ActualInterface string `json:"actual-interface,omitempty"`
	Address         string `json:"address,omitempty"`
	Advertise       string `json:"advertise,omitempty"`
	Comment         string `json:"comment,omitempty"`
	Disabled        string `json:"disabled,omitempty"`
	Dynamic         string `json:"dynamic,omitempty"`
	Eui64           string `json:"eui-64,omitempty"`
	FromPool        string `json:"from-pool,omitempty"`
	Interface       string `json:"interface,omitempty"`
	Invalid         string `json:"invalid,omitempty"`
	LinkLocal       string `json:"link-local,omitempty"`
	NoDad           string `json:"no-dad,omitempty"`
}

func (c *Client) GetIPv6Address(id string) (*IPv6Address, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/ipv6/address/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	res := IPv6Address{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateIPv6Address(ip_address *IPv6Address) (*IPv6Address, error) {
	reqBody, err := json.Marshal(ip_address)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/ipv6/address", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := IPv6Address{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateIPv6Address(id string, ip_address *IPv6Address) (*IPv6Address, error) {
	reqBody, err := json.Marshal(ip_address)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/ipv6/address/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := IPv6Address{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteIPv6Address(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/ipv6/address/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}
	res := IPv6Address{}
	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
