package client

import (
	"context"
	"fmt"
	"net/http"
)

type IPAddress struct {
	ID              string `json:".id"`
	ActualInterface string `json:"actual-interface"`
	Address         string `json:"address"`
	Disabled        string `json:"disabled"`
	Dynamic         string `json:"dynamic"`
	Interface       string `json:"interface"`
	Invalid         string `json:"invalid"`
	Network         string `json:"network"`
}

func (c *Client) GetIPAddresses(ctx context.Context, id string) (*IPAddress, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/ip/address/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	res := IPAddress{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateIPAddress(ctx context.Context, ip_address *IPAddress) (*IPAddress, error) {
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/ip/address", c.HostURL), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	res := IPAddress{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
