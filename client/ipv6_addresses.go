package client

import (
	"fmt"
	"net/http"
)

func (c *Client) ReadIPv6Addresses() ([]IPv6Address, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/ipv6/address", c.HostURL), nil)
	if err != nil {
		return nil, err
	}
	res := []IPv6Address{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
