package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type InterfaceListMember struct {
	ID        string `json:".id,omitempty"`
	List      string `json:"list,omitempty"`
	Interface string `json:"interface,omitempty"`
}

func (c *Client) CreateInterfaceListMember(member *InterfaceListMember) (*InterfaceListMember, error) {
	reqBody, err := json.Marshal(member)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/interface/list/member", c.HostURL), bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, err
	}

	res := InterfaceListMember{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadInterfaceListMember(id string) (*InterfaceListMember, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/interface/list/member/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	res := InterfaceListMember{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateInterfaceListMember(id string, member *InterfaceListMember) (*InterfaceListMember, error) {
	reqBody, err := json.Marshal(member)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/interface/list/member/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := InterfaceListMember{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteInterfaceListMember(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/interface/list/member/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := InterfaceListMember{}

	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
