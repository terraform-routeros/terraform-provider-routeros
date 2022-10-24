package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type IPv6NeighborDiscovery struct {
	ID                          string `json:".id,omitempty"`
	AdvertiseDNS                string `json:"advertise-dns,omitempty"`
	AdvertiseMACAddress         string `json:"advertise-mac-address,omitempty"`
	Disabled                    string `json:"disabled,omitempty"`
	DNS                         string `json:"dns,omitempty"`
	HopLimit                    string `json:"hop-limit,omitempty"`
	Interface                   string `json:"interface,omitempty"`
	Invalid                     string `json:"invalid,omitempty"`
	ManagedAddressConfiguration string `json:"managed-address-configuration,omitempty"`
	MTU                         string `json:"mtu,omitempty"`
	OtherConfiguration          string `json:"other-configuration,omitempty"`
	RADelay                     string `json:"ra-delay,omitempty"`
	RAInterval                  string `json:"ra-interval,omitempty"`
	RALifetime                  string `json:"ra-lifetime,omitempty"`
	ReachableTime               string `json:"reachable-time,omitempty"`
	RetransmitInterval          string `json:"retransmit-interval,omitempty"`
}

func (c *Client) GetIPv6NeighborDiscovery(id string) (*IPv6NeighborDiscovery, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/ipv6/nd/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	res := IPv6NeighborDiscovery{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateIPv6NeighborDiscovery(nd *IPv6NeighborDiscovery) (*IPv6NeighborDiscovery, error) {
	reqBody, err := json.Marshal(nd)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/ipv6/nd", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := IPv6NeighborDiscovery{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateIPv6NeighborDiscovery(id string, nd *IPv6NeighborDiscovery) (*IPv6NeighborDiscovery, error) {
	reqBody, err := json.Marshal(nd)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/ipv6/nd/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := IPv6NeighborDiscovery{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteIPv6NeighborDiscovery(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/ipv6/nd/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}
	res := IPv6NeighborDiscovery{}
	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
