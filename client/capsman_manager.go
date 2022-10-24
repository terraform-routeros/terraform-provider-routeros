package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CapsManManager struct {
	Certificate            string `json:"certificate,omitempty"`
	Enabled                string `json:"enabled,omitempty"`
	CaCertificate          string `json:"ca-certificate,omitempty"`
	RequirePeerCertificate string `json:"require-peer-certificate,omitempty"`
	Frequency              string `json:"frequency,omitempty"`
	PackagePath            string `json:"package-path,omitempty"`
	UpgradePolicy          string `json:"upgrade-policy,omitempty"`
}

func (c *Client) ReadCapsManManager() (*CapsManManager, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/caps-man/manager", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	res := CapsManManager{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateCapsManManager(manager *CapsManManager) (*CapsManManager, error) {
	reqBody, err := json.Marshal(manager)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/rest/caps-man/manager/set", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	//POST request in this case returns empty list
	var data []string
	if err := c.sendRequest(req, &data); err != nil {
		return nil, err
	}
	// Get the data
	req, err = http.NewRequest("GET", fmt.Sprintf("%s/rest/caps-man/manager", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	res := CapsManManager{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
