package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CapsManDatapath struct {
	ID                       string `json:".id,omitempty"`
	Comment                  string `json:"comment,omitempty"`
	Bridge                   string `json:"bridge,omitempty"`
	BridgeCost               string `json:"bridge-cost,omitempty"`
	BridgeHorizon            string `json:"bridge-horizon,omitempty"`
	ClientToClientForwarding string `json:"client-to-client-forwarding,omitempty"`
	LocalForwarding          string `json:"local-forwarding,omitempty"`
	OpenFlowSwitch           string `json:"openflow-switch,omitempty"`
	VlanID                   string `json:"vlan-id,omitempty"`
	VlanMode                 string `json:"vlan-mode,omitempty"`
	Name                     string `json:"name"`
}

func (c *Client) CreateCapsManDatapath(datapath *CapsManDatapath) (*CapsManDatapath, error) {
	reqBody, err := json.Marshal(datapath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/caps-man/datapath", c.HostURL), bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, err
	}

	res := CapsManDatapath{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadCapsManDatapath(id string) (*CapsManDatapath, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/caps-man/datapath/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	res := CapsManDatapath{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateCapsManDatapath(id string, datapath *CapsManDatapath) (*CapsManDatapath, error) {
	reqBody, err := json.Marshal(datapath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/caps-man/datapath/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := CapsManDatapath{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteCapsManDatapath(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/caps-man/datapath/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := CapsManDatapath{}

	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
