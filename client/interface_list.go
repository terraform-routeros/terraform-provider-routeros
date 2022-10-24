package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type InterfaceList struct {
	ID      string `json:".id,omitempty"`
	Exclude string `json:"exclude,omitempty"`
	Include string `json:"include,omitempty"`
	Dynamic string `json:"dynamic,omitempty"`
	Name    string `json:"name,omitempty"`
	Comment string `json:"comment,omitempty"`
}

func (c *Client) CreateInterfaceList(list *InterfaceList) (*InterfaceList, error) {
	reqBody, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/interface/list", c.HostURL), bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, err
	}

	res := InterfaceList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadInterfaceList(id string) (*InterfaceList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/interface/list/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	res := InterfaceList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateInterfaceList(id string, list *InterfaceList) (*InterfaceList, error) {
	reqBody, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/interface/list/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := InterfaceList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteInterfaceList(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/interface/list/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := InterfaceList{}

	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
