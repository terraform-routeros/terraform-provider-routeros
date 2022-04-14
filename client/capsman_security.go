package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CapsManSecurity struct {
	ID                  string `json:".id,omitempty"`
	Comment             string `json:"comment,omitempty"`
	GroupEncryption     string `json:"group-encryption,omitempty"`
	EapRadiusAccounting string `json:"eap-radius-accounting,omitempty"`
	Encryption          string `json:"encryption,omitempty"`
	GroupKeyUpdate      string `json:"group-key-update,omitempty"`
	Passphrase          string `json:"passphrase,omitempty"`
	TlsCertificate      string `json:"tls-certificate,omitempty"`
	TlsMode             string `json:"tls-mode,omitempty"`
	Name                string `json:"name"`
}

func (c *Client) CreateCapsManSecurity(security *CapsManSecurity) (*CapsManSecurity, error) {
	reqBody, err := json.Marshal(security)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/caps-man/security", c.HostURL), bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, err
	}

	res := CapsManSecurity{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadCapsManSecurity(id string) (*CapsManSecurity, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/caps-man/security/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	res := CapsManSecurity{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateCapsManSecurity(id string, security *CapsManSecurity) (*CapsManSecurity, error) {
	reqBody, err := json.Marshal(security)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/caps-man/security/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := CapsManSecurity{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteCapsManSecurity(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/caps-man/security/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := CapsManSecurity{}

	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
