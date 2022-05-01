package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type InterfaceVrrp struct {
	ID                     string `json:".id,omitempty"`
	Arp                    string `json:"arp,omitempty"`
	V3Protocol             string `json:"v3-protocol,omitempty"`
	Authentication         string `json:"authentication,omitempty"`
	Interface              string `json:"interface,omitempty"`
	Mtu                    string `json:"mtu,omitempty"`
	Password               string `json:"password,omitempty"`
	Priority               string `json:"priority,omitempty"`
	OnFail                 string `json:"on-fail,omitempty"`
	OnMaster               string `json:"on-master,omitempty"`
	OnBackup               string `json:"on-backup,omitempty"`
	SyncConnectionTracking string `json:"sync-connection-tracking,omitempty"`
	Vrid                   string `json:"vrid,omitempty"`
	Interval               string `json:"interval,omitempty"`
	Name                   string `json:"name,omitempty"`
	PreemptionMode         string `json:"preemption-mode ,omitempty"`
	Version                string `json:"version,omitempty"`
	ArpTimeout             string `json:"arp-timeout,omitempty"`
}

func (c *Client) CreateInterfaceVrrp(vrrp *InterfaceVrrp) (*InterfaceVrrp, error) {
	reqBody, err := json.Marshal(vrrp)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/interface/vrrp", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	res := InterfaceVrrp{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadInterfaceVrrp(id string) (*InterfaceVrrp, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/interface/vrrp/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	res := InterfaceVrrp{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) UpdateInterfaceVrrp(id string, vrrp *InterfaceVrrp) (*InterfaceVrrp, error) {
	reqBody, err := json.Marshal(vrrp)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/interface/vrrp/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := InterfaceVrrp{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteInterfaceVrrp(vrrp *InterfaceVrrp) error {
	id := vrrp.ID
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/interface/vrrp/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := InterfaceVrrp{}
	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
