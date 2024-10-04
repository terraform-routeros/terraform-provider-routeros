package routeros

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-routeros/routeros"
)

type ApiClient struct {
	ctx       context.Context
	HostURL   string
	Username  string
	Password  string
	Transport TransportType
	extra     *ExtraParams
	*routeros.Client
}

var (
	apiMethodName = map[crudMethod]string{
		crudCreate:      "/add",
		crudRead:        "/print",
		crudUpdate:      "/set",
		crudDelete:      "/remove",
		crudPost:        "/set",
		crudImport:      "/import",
		crudSign:        "/sign",
		crudSignViaScep: "/add-scep",
		crudRemove:      "/remove",
		crudRevoke:      "/issued-revoke",
		crudMove:        "/move",
		crudStart:       "/start",
		crudStop:        "/stop",
		crudGenerateKey: "/generate-key",
	}
)

func (c *ApiClient) GetExtraParams() *ExtraParams {
	return c.extra
}

func (c *ApiClient) GetTransport() TransportType {
	return c.Transport
}

func (c *ApiClient) SendRequest(method crudMethod, url *URL, item MikrotikItem, result interface{}) error {

	// https://help.mikrotik.com/docs/display/ROS/API
	// /interface/vlan/print + '?.id=*39' + '?type=vlan'
	cmd := url.GetApiCmd()
	// The first element is the Path
	cmd[0] += apiMethodName[method]

	// Marshal
	for fieldName, fieldValue := range item {
		cmd = append(cmd, fmt.Sprintf("=%s=%s", fieldName, fieldValue))
	}
	ColorizedDebug(c.ctx, "request body:  "+strings.Join(cmd, " "))

	resp, err := c.RunArgs(cmd)
	if err != nil {
		return err
	}

	ColorizedDebug(c.ctx, "response body: "+resp.String())

	if result == nil {
		return nil
	}

	// Unmarshal

	switch r := result.(type) {
	case *MikrotikItem:
		// Only ID returned.
		// !done @ [{`ret` `*7F`}]
		if len(resp.Re) == 0 {
			for k, v := range resp.Done.Map {
				(*r)[k] = v
			}
			break
		}

		// Fill in only one item.
		for k, v := range resp.Re[0].Map {
			(*r)[k] = v
		}
	case *[]MikrotikItem:
		// !re
		for _, sentence := range resp.Re {
			m := MikrotikItem{}
			for k, v := range sentence.Map {
				m[k] = v
			}
			*r = append(*r, m)
		}

		// !done @ [] is empty...
	default:
		panic("[SendRequest] type " + reflect.TypeOf(result).String() + " is not supported for API response unmarshaling.")
	}

	return nil
}
