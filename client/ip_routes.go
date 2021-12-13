package client

import (
	"fmt"
	"net/http"
)

func (c *Client) ReadIPRoutes() ([]IPRoute, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/ip/route", c.HostURL), nil)
	if err != nil {
		return nil, err
	}
	res := []IPRoute{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return res, nil
}
