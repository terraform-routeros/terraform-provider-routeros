package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CapsManChannel struct {
	ID                  string `json:".id,omitempty"`
	Comment             string `json:"comment,omitempty"`
	SaveSelected        string `json:"save-selected,omitempty"`
	Width               string `json:"width,omitempty"`
	Band                string `json:"band,omitempty"`
	ExtensionChannel    string `json:"extension-channel,omitempty"`
	Frequency           string `json:"frequency,omitempty"`
	TXPower             string `json:"tx-power,omitempty"`
	ControlChannelWidth string `json:"control-channel-width,omitempty"`
	ReselectInterval    string `json:"reselect-interval,omitempty"`
	SecondaryFrequency  string `json:"secondary-frequency,omitempty"`
	SkipDfsChannels     string `json:"skip-dfs-channels,omitempty"`
	Name                string `json:"name"`
}

func (c *Client) CreateCapsManChannel(channel *CapsManChannel) (*CapsManChannel, error) {
	reqBody, err := json.Marshal(channel)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/caps-man/channel", c.HostURL), bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, err
	}

	res := CapsManChannel{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadCapsManChannel(id string) (*CapsManChannel, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/caps-man/channel/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	res := CapsManChannel{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateCapsManChannel(id string, channel *CapsManChannel) (*CapsManChannel, error) {
	reqBody, err := json.Marshal(channel)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/caps-man/channel/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := CapsManChannel{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteCapsManChannel(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/caps-man/channel/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := CapsManChannel{}

	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
