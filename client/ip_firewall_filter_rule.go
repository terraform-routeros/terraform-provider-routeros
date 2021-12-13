package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type IPFirewallFilterRule struct {
	ID                 string `json:".id,omitempty"`
	Action             string `json:"action,omitempty"`
	Bytes              string `json:"bytes,omitempty"`
	Chain              string `json:"chain,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Dynamic            string `json:"dynamic,omitempty"`
	Packets            string `json:"packets,omitempty"`
	ConnectionState    string `json:"connection-state,omitempty"`
	Invalid            string `json:"invalid,omitempty"`
	Disabled           string `json:"disabled,omitempty"`
	Log                string `json:"log,omitempty"`
	LogPrefix          string `json:"log-prefix,omitempty"`
	DstAddress         string `json:"dst-address,omitempty"`
	DstPort            string `json:"dst-port,omitempty"`
	Protocol           string `json:"protocol,omitempty"`
	SrcAddress         string `json:"src-address,omitempty"`
	InInterfaceList    string `json:"in-interface-list,omitempty"`
	SrcAddressList     string `json:"src-address-list,omitempty"`
	DstAddressList     string `json:"dst-address-list,omitempty"`
	OutInterfaceList   string `json:"out-interface-list,omitempty"`
	ConnectionNatState string `json:"connection-nat-state,omitempty"`
	HwOffload          string `json:"hw-offload,omitempty"`
}

func (c *Client) CreateIPFirewallFilterRule(ip_pool *IPFirewallFilterRule) (*IPFirewallFilterRule, error) {
	reqBody, err := json.Marshal(ip_pool)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/ip/firewall/filter", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	res := IPFirewallFilterRule{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadIPFirewallFilterRule(id string) (*IPFirewallFilterRule, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/ip/firewall/filter/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	res := IPFirewallFilterRule{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) UpdateIPFirewallFilterRule(id string, ip_pool *IPFirewallFilterRule) (*IPFirewallFilterRule, error) {
	reqBody, err := json.Marshal(ip_pool)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/ip/firewall/filter/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := IPFirewallFilterRule{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteIPFirewallFilterRule(ip_pool *IPFirewallFilterRule) error {
	id := ip_pool.ID
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/ip/firewall/filter/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := IPFirewallFilterRule{}
	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
