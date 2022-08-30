package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GRE struct {
	ID            string `json:".id,omitempty"`
	ActualMtu     int    `json:"actual-mtu,string,omitempty"`
	AllowFastPath string `json:"allow-fast-path,omitempty"`
	ClampTcpMss   string `json:"clamp-tcp-mss,omitempty"`
	Comment       string `json:"comment,omitempty"`
	Disabled      bool   `json:"disabled,string"`
	DontFragment  string `json:"dont-fragment,omitempty"`
	Dscp          string `json:"dscp,omitempty"`
	IpsecSecret   string `json:"ipsec-secret,omitempty"`
	Keepalive     string `json:"keepalive,omitempty"`
	L2Mtu         int    `json:"l2mtu,string,omitempty"`
	LocalAddress  string `json:"local-address,omitempty"`
	Mtu           int    `json:"mtu,string,omitempty"`
	Name          string `json:"name"`
	RemoteAddress string `json:"remote-address"`
	Running       bool   `json:"running,string,omitempty"`
}

func (c *Client) CreateGRE(iface *GRE) (*GRE, error) {
	reqBody, err := json.Marshal(iface)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/interface/gre", c.HostURL), bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, err
	}

	res := GRE{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadGRE(id string) (*GRE, error) {
	// Prevent 404 - Not Found error. Error occurs when a resource has been deleted outside terraform control (winbox,
	// ssh, etc).
	// In this implementation, we have an empty [] or non-empty array [{...}].
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/interface/gre?.id=%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	var res []GRE
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	if len(res) > 0 {
		return &res[0], nil
	}
	return &GRE{}, nil
}

func (c *Client) UpdateGRE(id string, iface *GRE) (*GRE, error) {
	reqBody, err := json.Marshal(iface)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/interface/gre/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := GRE{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteGRE(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/interface/gre/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := GRE{}

	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
