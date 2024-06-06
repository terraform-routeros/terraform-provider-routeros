package routeros

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
		crudCreate:      "PUT",
		crudRead:        "GET",
		crudUpdate:      "PATCH",
		crudDelete:      "DELETE",
		crudPost:        "POST",
		crudImport:      "POST",
		crudSign:        "POST",
		crudSignViaScep: "POST",
		crudRemove:      "POST",
		crudRevoke:      "POST",
		crudMove:        "POST",
		crudStart:       "POST",
		crudStop:        "POST",
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

		ColorizedDebug(c.ctx, "request body:  "+string(b))
		data = bytes.NewBuffer(b)
	}

	// https://mikrotik + /rest + /interface/vlan + ? + .id=*39
	// Escaping spaces!
	requestUrl := c.HostURL + "/rest" + strings.Replace(url.GetRestURL(), " ", "%20", -1)
	ColorizedDebug(c.ctx, restMethodName[method]+" request URL:  "+requestUrl)

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

	body, _ := io.ReadAll(res.Body)

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse

		ColorizedDebug(c.ctx, fmt.Sprintf("error response body:\n%s", body))

		if err = json.Unmarshal(body, &errRes); err != nil {
			return fmt.Errorf("json.Unmarshal - %v", err)
		} else {
			return fmt.Errorf("%v '%v' returned response code: %v, message: '%v', details: '%v'",
				restMethodName[method], requestUrl, res.StatusCode, errRes.Message, errRes.Detail)
		}
	}

	ColorizedDebug(c.ctx, "response body: "+string(body))

	if len(body) != 0 && result != nil {
		if err = json.Unmarshal(body, &result); err != nil {

			if e, ok := err.(*json.SyntaxError); ok {
				ColorizedDebug(c.ctx, fmt.Sprintf("json.Unmarshal(response body): syntax error at byte offset %d", e.Offset))

				if err = json.Unmarshal(EscapeChars(body), &result); err != nil {
					return fmt.Errorf("json.Unmarshal(response body): %v", err)
				}
			} else {
				return err
			}
		}
	}
	return nil
}
