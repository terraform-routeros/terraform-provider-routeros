package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type IPFirewallFilterRule struct {
	ID                      string `json:".id,omitempty"`
	Action                  string `json:"action,omitempty"`
	AddressListTimeout      string `json:"address-list-timeout,omitempty"`
	Bytes                   string `json:"bytes,omitempty"`
	Chain                   string `json:"chain,omitempty"`
	Comment                 string `json:"comment,omitempty"`
	ConnectionBytes         string `json:"connection-bytes,omitempty"`
	ConnectionLimit         string `json:"connection-limit,omitempty"`
	ConnectionMark          string `json:"connection-mark,omitempty"`
	ConnectionNatState      string `json:"connection-nat-state,omitempty"`
	ConnectionRate          string `json:"connection-rate,omitempty"`
	ConnectionState         string `json:"connection-state,omitempty"`
	ConnectionType          string `json:"connection-type,omitempty"`
	Content                 string `json:"content,omitempty"`
	Disabled                string `json:"disabled,omitempty"`
	Dscp                    string `json:"dscp,omitempty"`
	DstAddress              string `json:"dst-address,omitempty"`
	DstAddressList          string `json:"dst-address-list,omitempty"`
	DstAddressType          string `json:"dst-address-type,omitempty"`
	DstLimit                string `json:"dst-limit,omitempty"`
	DstPort                 string `json:"dst-port,omitempty"`
	Dynamic                 string `json:"dynamic,omitempty"`
	Fragment                string `json:"fragment,omitempty"`
	HotSpot                 string `json:"hotspot,omitempty"`
	HwOffload               string `json:"hw-offload,omitempty"`
	IcmpOptions             string `json:"icmp-options,omitempty"`
	InBridgePort            string `json:"in-bridge-port,omitempty"`
	InBridgePortList        string `json:"in-bridge-port-list,omitempty"`
	InInterface             string `json:"in-interface,omitempty"`
	InInterfaceList         string `json:"in-interface-list,omitempty"`
	IngressPriority         string `json:"ingress-priority,omitempty"`
	Invalid                 string `json:"invalid,omitempty"`
	IpsecPolicy             string `json:"ipsec-policy,omitempty"`
	Ipv4Options             string `json:"ipv4-options,omitempty"`
	JumpTarget              string `json:"jump-target,omitempty"`
	Layer7Protocol          string `json:"layer7-protocol,omitempty"`
	Limit                   string `json:"limit,omitempty"`
	Log                     string `json:"log,omitempty"`
	LogPrefix               string `json:"log-prefix,omitempty"`
	Nth                     string `json:"nth,omitempty"`
	OutBridgePort           string `json:"out-bridge-port,omitempty"`
	OutBridgePortList       string `json:"out-bridge-port-list,omitempty"`
	OutInterface            string `json:"out-interface,omitempty"`
	OutInterfaceList        string `json:"out-interface-list,omitempty"`
	Packets                 string `json:"packets,omitempty"`
	PacketMark              string `json:"packet-mark,omitempty"`
	PacketSize              string `json:"packet-size,omitempty"`
	PerConnectionClassifier string `json:"per-connection-classifier,omitempty"`
	Port                    string `json:"port,omitempty"`
	Priority                string `json:"priority,omitempty"`
	Protocol                string `json:"protocol,omitempty"`
	Psd                     string `json:"psd,omitempty"`
	Random                  string `json:"random,omitempty"`
	RejectWith              string `json:"reject-with,omitempty"`
	RoutingTable            string `json:"routing-table,omitempty"`
	RoutingMark             string `json:"routing-mark,omitempty"`
	SrcAddress              string `json:"src-address,omitempty"`
	SrcAddressList          string `json:"src-address-list,omitempty"`
	SrcAddressType          string `json:"src-address-type,omitempty"`
	SrcPort                 string `json:"src-port,omitempty"`
	SrcMacAddress           string `json:"src-mac-address,omitempty"`
	TcpFlags                string `json:"tcp-flags,omitempty"`
	TcpMss                  string `json:"tcp-mss,omitempty"`
	Time                    string `json:"time,omitempty"`
	TlsHost                 string `json:"tls-host,omitempty"`
	Ttl                     string `json:"ttl,omitempty"`
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
