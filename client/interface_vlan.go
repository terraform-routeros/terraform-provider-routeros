package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type VLAN struct {
	ID                      string `json:".id,omitempty"`
	Arp                     string `json:"arp,omitempty"`
	ArpTimeout              string `json:"arp-timeout,omitempty"`
	Disabled                string `json:"disabled"`
	Interface               string `json:"interface"`
	L2Mtu                   string `json:"l2mtu,omitempty"`
	LoopProtect             string `json:"loop-protect,omitempty"`
	LoopProtectDisableTime  string `json:"loop-protect-disable-time,omitempty"`
	LoopProtectSendInterval string `json:"loop-protect-send-interval,omitempty"`
	LoopProtectStatus       string `json:"loop-protect-status,omitempty"`
	MacAddress              string `json:"mac-address,omitempty"`
	Mtu                     string `json:"mtu,omitempty"`
	Name                    string `json:"name"`
	Running                 string `json:"running,omitempty"`
	UseServiceTag           string `json:"use-service-tag,omitempty"`
	VlanID                  string `json:"vlan-id"`
}

func (c *Client) CreateVLAN(vlan *VLAN) (*VLAN, error) {
	reqBody, err := json.Marshal(vlan)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/interface/vlan", c.HostURL), bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, err
	}

	res := VLAN{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadVLAN(name string) (*VLAN, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/interface/vlan/%s", c.HostURL, name), nil)
	if err != nil {
		return nil, err
	}

	res := VLAN{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateVLAN(name string, vlan *VLAN) (*VLAN, error) {
	reqBody, err := json.Marshal(vlan)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/interface/vlan/%s", c.HostURL, name), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := VLAN{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteVLAN(name string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/interface/vlan/%s", c.HostURL, name), nil)
	if err != nil {
		return err
	}

	res := VLAN{}

	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
