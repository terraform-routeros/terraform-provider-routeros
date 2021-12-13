package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type InterfaceBridge struct {
	ID                string `json:".id,omitempty"`
	ActualMtu         string `json:"actual-mtu,omitempty"`
	AdminMac          string `json:"admin-mac,omitempty"`
	AgeingTime        string `json:"ageing-time,omitempty"`
	Arp               string `json:"arp,omitempty"`
	ArpTimeout        string `json:"arp-timeout,omitempty"`
	AutoMac           string `json:"auto-mac,omitempty"`
	Comment           string `json:"comment,omitempty"`
	DhcpSnooping      string `json:"dhcp-snooping,omitempty"`
	Disabled          string `json:"disabled,omitempty"`
	EtherType         string `json:"ether-type,omitempty"`
	FastForward       string `json:"fast-forward,omitempty"`
	ForwardDelay      string `json:"forward-delay,omitempty"`
	FrameTypes        string `json:"frame-types,omitempty"`
	IgmpSnooping      string `json:"igmp-snooping,omitempty"`
	IngressFiltering  string `json:"ingress-filtering,omitempty"`
	L2Mtu             string `json:"l2mtu,omitempty"`
	MacAddress        string `json:"mac-address,omitempty"`
	MaxMessageAge     string `json:"max-message-age,omitempty"`
	Mtu               string `json:"mtu,omitempty"`
	Name              string `json:"name,omitempty"`
	Priority          string `json:"priority,omitempty"`
	ProtocolMode      string `json:"protocol-mode,omitempty"`
	Pvid              string `json:"pvid,omitempty"`
	Running           string `json:"running,omitempty"`
	TransmitHoldCount string `json:"transmit-hold-count,omitempty"`
	VlanFiltering     string `json:"vlan-filtering,omitempty"`
}

func (c *Client) CreateInterfaceBridge(ip_route *InterfaceBridge) (*InterfaceBridge, error) {
	reqBody, err := json.Marshal(ip_route)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/interface/bridge", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	res := InterfaceBridge{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadInterfaceBridge(id string) (*InterfaceBridge, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/interface/bridge/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	res := InterfaceBridge{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) UpdateInterfaceBridge(id string, ip_route *InterfaceBridge) (*InterfaceBridge, error) {
	reqBody, err := json.Marshal(ip_route)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/interface/bridge/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := InterfaceBridge{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteInterfaceBridge(ip_route *InterfaceBridge) error {
	id := ip_route.ID
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/interface/bridge/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := InterfaceBridge{}
	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
