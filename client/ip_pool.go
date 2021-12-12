package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type IPPool struct {
	ID     string `json:".id,omitempty"`
	Name   string `json:"name,omitempty"`
	Ranges string `json:"ranges,omitempty"`
}

func (c *Client) CreateIPPool(dhcp_server *IPPool) (*IPPool, error) {
	reqBody, err := json.Marshal(dhcp_server)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/ip/pool", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	res := IPPool{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadIPPool(id string) (*IPPool, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/ip/pool/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	res := IPPool{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) UpdateIPPool(id string, dhcp_server *IPPool) (*IPPool, error) {
	reqBody, err := json.Marshal(dhcp_server)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/ip/pool/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := IPPool{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteIPPool(dhcp_server *IPPool) error {
	id := dhcp_server.ID
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/ip/pool/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := IPPool{}
	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
