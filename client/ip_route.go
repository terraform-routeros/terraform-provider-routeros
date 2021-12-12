package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type IPRoute struct {
	ID                string `json:".id,omitempty"`
	Active            string `json:"active,omitempty"`
	Dhcp              string `json:"dhcp,omitempty"`
	Distance          string `json:"distance,omitempty"`
	DstAddress        string `json:"dst-address,omitempty"`
	Dynamic           string `json:"dynamic,omitempty"`
	Ecmp              string `json:"ecmp,omitempty"`
	Gateway           string `json:"gateway,omitempty"`
	HwOffloaded       string `json:"hw-offloaded,omitempty"`
	ImmediateGw       string `json:"immediate-gw,omitempty"`
	Inactive          string `json:"inactive,omitempty"`
	PrefSrc           string `json:"pref-src,omitempty"`
	RoutingTable      string `json:"routing-table,omitempty"`
	Scope             string `json:"scope,omitempty"`
	SuppressHwOffload string `json:"suppress-hw-offload,omitempty"`
	TargetScope       string `json:"target-scope,omitempty"`
	VrfInterface      string `json:"vrf-interface,omitempty"`
}

func (c *Client) CreateIPRoute(ip_route *IPRoute) (*IPRoute, error) {
	reqBody, err := json.Marshal(ip_route)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/ip/route", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	res := IPRoute{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadIPRoute(id string) (*IPRoute, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/ip/route/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	res := IPRoute{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) UpdateIPRoute(id string, ip_route *IPRoute) (*IPRoute, error) {
	reqBody, err := json.Marshal(ip_route)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/ip/route/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := IPRoute{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteIPRoute(ip_route *IPRoute) error {
	id := ip_route.ID
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/ip/route/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := IPRoute{}
	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
