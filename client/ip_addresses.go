package client

import (
	"fmt"
	"net/http"
)

func (c *Client) ReadIPAddresses() ([]IPAddress, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/ip/address", c.HostURL), nil)
	if err != nil {
		return nil, err
	}
	res := []IPAddress{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
