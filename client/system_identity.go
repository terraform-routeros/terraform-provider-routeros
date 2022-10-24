package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SystemIdentity struct {
	ID   string `json:".id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (c *Client) CreateSystemIdentity(system_identity *SystemIdentity) (*SystemIdentity, error) {
	reqBody, err := json.Marshal(system_identity)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/rest/system/identity/set", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

    var res interface{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}

    return c.ReadSystemIdentity()
}

func (c *Client) ReadSystemIdentity() (*SystemIdentity, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/system/identity", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	res := SystemIdentity{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateSystemIdentity(system_identity *SystemIdentity) (*SystemIdentity, error) {
	reqBody, err := json.Marshal(system_identity)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/rest/system/identity/set", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

    var res interface{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}

    return c.ReadSystemIdentity()
}

func (c *Client) DeleteSystemIdentity(system_identity *SystemIdentity) error {
    default_system_identity := new(SystemIdentity)
    default_system_identity.Name = "MikroTik"

    reqBody, err := json.Marshal(default_system_identity)
    if err != nil {
        return err
    }

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/rest/system/identity/set", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

    var res interface{}
	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
