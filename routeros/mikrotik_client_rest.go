package routeros

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
)

type RestClient struct {
	ctx       context.Context
	HostURL   string
	Username  string
	Password  string
	Transport TransportType
	extra     *ExtraParams
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
		crudGenerateKey: "POST",
	}
)

func (c *RestClient) GetExtraParams() *ExtraParams {
	return c.extra
}

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

	// Cast single return values to an array.
	if !isJsonArray(body) {
		body = append([]byte{'['}, append(body, []byte{']'}...)...)
	}

	// If the requested value is a slice, then parse the JSON directly into it.
	var slice = new([]MikrotikItem)
	if result != nil && isSlice(result) {
		slice = result.(*[]MikrotikItem)
	}

	if len(body) != 0 && result != nil {
		if err = json.Unmarshal(body, &slice); err != nil {

			if e, ok := err.(*json.SyntaxError); ok {
				ColorizedDebug(c.ctx, fmt.Sprintf("json.Unmarshal(response body): syntax error at byte offset %d", e.Offset))

				if err = json.Unmarshal(EscapeChars(body), &slice); err != nil {
					return fmt.Errorf("json.Unmarshal(response body): %v", err)
				}
			} else {
				return err
			}
		}
	}

	if result != nil && !isSlice(result) && len(*slice) > 0 {
		// result.(*MikrotikItem).replace(&(*slice)[0])
		for k, v := range (*slice)[0] {
			(*result.(*MikrotikItem))[k] = v
		}
	}

	return nil
}

// isSlice The function returns information whether the passed parameter is a slice.
// The incoming type is a variable or pointer.
func isSlice(i any) bool {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t.Kind() == reflect.Slice
}

// isJsonArray The function returns information about the type of JSON response.
// Based on the response, we can cast MT's response to an array of values.
// After some time, we can say that it is easier to operate with an array of values,
// since MT can return '[]' which is not obvious in the process of creating a single resource.
func isJsonArray(b []byte) bool {
	b = bytes.TrimLeft(b, " \t\r\n")
	return len(b) > 0 && b[0] == '['
}
