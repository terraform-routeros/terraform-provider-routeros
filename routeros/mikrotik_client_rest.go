package routeros

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"io"
	"net/http"
)

type RestClient struct {
	ctx       context.Context
	HostURL   string
	Username  string
	Password  string
	Transport TransportType
	*http.Client
}

type errorResponse struct {
	Detail  string `json:"detail"`
	Error   int    `json:"error"`
	Message string `json:"message"`
}

var (
	restMethodName = map[crudMethod]string{
		crudCreate: "PUT",
		crudRead:   "GET",
		crudUpdate: "PATCH",
		crudDelete: "DELETE",
	}
)

func (c *RestClient) GetTransport() TransportType {
	return c.Transport
}

func (c *RestClient) SendRequest(method crudMethod, url *URL, item MikrotikItem, result interface{}) error {
	var data io.Reader

	if item != nil {
		b, err := json.Marshal(&item)
		if err != nil {
			return err
		}

		tflog.Debug(c.ctx, "request body:  "+string(b))
		data = bytes.NewBuffer(b)
	}

	// https://mikrotik + /rest + /interface/vlan + ? + .id=*39
	requestUrl := c.HostURL + "/rest" + url.GetRestURL()
	tflog.Debug(c.ctx, restMethodName[method]+" request URL:  "+requestUrl)

	req, err := http.NewRequest(restMethodName[method], requestUrl, data)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)

	res, err := c.Do(req)
	if err != nil {
		return err
	}

	defer func() { _ = res.Body.Close() }()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return fmt.Errorf("%v %v returned response - %v: '%v' (%v)",
				restMethodName[method], requestUrl, res.StatusCode, errRes.Message, errRes.Detail)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}
	body, _ := io.ReadAll(res.Body)

	tflog.Debug(c.ctx, "response body: "+string(body))

	if len(body) != 0 {
		if err = json.Unmarshal(body, &result); err != nil {
			return err
		}
	}
	return nil
}
