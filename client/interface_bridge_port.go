package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type InterfaceBridgePort struct {
	ID                    string `json:".id,omitempty"`
	Nextid                string `json:".nextid,omitempty"`
	AutoIsolate           string `json:"auto-isolate,omitempty"`
	BpduGuard             string `json:"bpdu-guard,omitempty"`
	Bridge                string `json:"bridge,omitempty"`
	BroadcastFlood        string `json:"broadcast-flood,omitempty"`
	DebugInfo             string `json:"debug-info,omitempty"`
	Disabled              string `json:"disabled,omitempty"`
	Dynamic               string `json:"dynamic,omitempty"`
	Edge                  string `json:"edge,omitempty"`
	EdgePort              string `json:"edge-port,omitempty"`
	EdgePortDiscovery     string `json:"edge-port-discovery,omitempty"`
	ExternalFdbStatus     string `json:"external-fdb-status,omitempty"`
	FastLeave             string `json:"fast-leave,omitempty"`
	Forwarding            string `json:"forwarding,omitempty"`
	FrameTypes            string `json:"frame-types,omitempty"`
	Horizon               string `json:"horizon,omitempty"`
	Hw                    string `json:"hw,omitempty"`
	HwOffload             string `json:"hw-offload,omitempty"`
	HwOffloadGroup        string `json:"hw-offload-group,omitempty"`
	Inactive              string `json:"inactive,omitempty"`
	IngressFiltering      string `json:"ingress-filtering,omitempty"`
	Interface             string `json:"interface,omitempty"`
	InternalPathCost      string `json:"internal-path-cost,omitempty"`
	Learn                 string `json:"learn,omitempty"`
	Learning              string `json:"learning,omitempty"`
	MulticastRouter       string `json:"multicast-router,omitempty"`
	PathCost              string `json:"path-cost,omitempty"`
	PointToPoint          string `json:"point-to-point,omitempty"`
	PointToPointPort      string `json:"point-to-point-port,omitempty"`
	PortNumber            string `json:"port-number,omitempty"`
	Priority              string `json:"priority,omitempty"`
	Pvid                  string `json:"pvid,omitempty"`
	RestrictedRole        string `json:"restricted-role,omitempty"`
	RestrictedTcn         string `json:"restricted-tcn,omitempty"`
	Role                  string `json:"role,omitempty"`
	SendingRstp           string `json:"sending-rstp,omitempty"`
	Status                string `json:"status,omitempty"`
	TagStacking           string `json:"tag-stacking,omitempty"`
	Trusted               string `json:"trusted,omitempty"`
	UnknownMulticastFlood string `json:"unknown-multicast-flood,omitempty"`
	UnknownUnicastFlood   string `json:"unknown-unicast-flood,omitempty"`
}

func (c *Client) CreateInterfaceBridgePort(ip_route *InterfaceBridgePort) (*InterfaceBridgePort, error) {
	reqBody, err := json.Marshal(ip_route)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/interface/bridge/port", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	res := InterfaceBridgePort{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadInterfaceBridgePort(id string) (*InterfaceBridgePort, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/interface/bridge/port/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	res := InterfaceBridgePort{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) UpdateInterfaceBridgePort(id string, ip_route *InterfaceBridgePort) (*InterfaceBridgePort, error) {
	reqBody, err := json.Marshal(ip_route)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/interface/bridge/port/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := InterfaceBridgePort{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteInterfaceBridgePort(ip_route *InterfaceBridgePort) error {
	id := ip_route.ID
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/interface/bridge/port/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := InterfaceBridgePort{}
	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
