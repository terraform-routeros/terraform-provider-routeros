package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CapsManProvisioning struct {
	ID                  string `json:".id,omitempty"`
	Comment             string `json:"comment,omitempty"`
	Action              string `json:"action,omitempty"`
	CommonNameRegexp    string `json:"common-name-regexp,omitempty"`
	HwSupportedModes    string `json:"hw-supported-modes,omitempty"`
	IdentityRegexp      string `json:"identity-regexp,omitempty"`
	IpAddressRanges     string `json:"ip-address-ranges,omitempty"`
	MasterConfiguration string `json:"master-configuration,omitempty"`
	NameFormat          string `json:"name-format,omitempty"`
	NamePrefix          string `json:"name-prefix,omitempty"`
	RadioMAC            string `json:"radio-mac,omitempty"`
	SlaveConfigurations string `json:"slave-conigurations,omitempty"`
}

func (c *Client) CreateCapsManProvisioning(provisioning *CapsManProvisioning) (*CapsManProvisioning, error) {
	reqBody, err := json.Marshal(provisioning)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/caps-man/provisioning", c.HostURL), bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, err
	}

	res := CapsManProvisioning{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadCapsManProvisioning(id string) (*CapsManProvisioning, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/caps-man/provisioning/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	res := CapsManProvisioning{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateCapsManProvisioning(id string, provisioning *CapsManProvisioning) (*CapsManProvisioning, error) {
	reqBody, err := json.Marshal(provisioning)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/caps-man/provisioning/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := CapsManProvisioning{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteCapsManProvisioning(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/caps-man/provisioning/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := CapsManProvisioning{}

	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
