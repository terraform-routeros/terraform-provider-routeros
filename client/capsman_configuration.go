package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CapsManConfiguration struct {
	ID                               string `json:".id,omitempty"`
	Name                             string `json:"name"`
	Comment                          string `json:"comment,omitempty"`
	Channel                          string `json:"channel,omitempty"`
	ChannelSaveSelected              string `json:"channel.save-selected,omitempty"`
	ChannelWidth                     string `json:"channel.width,omitempty"`
	ChannelBand                      string `json:"channel.band,omitempty"`
	ChannelExtensionChannel          string `json:"channel.extension-channel,omitempty"`
	ChannelFrequency                 string `json:"channel.frequency,omitempty"`
	ChannelTXPower                   string `json:"channel.tx-power,omitempty"`
	ChannelControlChannelWidth       string `json:"channel.control-channel-width,omitempty"`
	ChannelReselectInterval          string `json:"channel.reselect-interval,omitempty"`
	ChannelSecondaryFrequency        string `json:"channel.secondary-frequency,omitempty"`
	ChannelSkipDfsChannels           string `json:"channel.skip-dfs-channels,omitempty"`
	Country                          string `json:"country,omitempty"`
	DatapathBridge                   string `json:"datapath.bridge,omitempty"`
	DatapathBridgeCost               string `json:"datapath.bridge-cost,omitempty"`
	DatapathBridgeHorizon            string `json:"datapath.bridge-horizon,omitempty"`
	DatapathClientToClientForwarding string `json:"datapath.client-to-client-forwarding,omitempty"`
	DatapathInterfaceList            string `json:"datapath.interface-list,omitempty"`
	DatapathL2MTU                    string `json:"datapath.l2mtu,omitempty"`
	DatapathLocalForwarding          string `json:"datapath.local-forwarding,omitempty"`
	DatapathMTU                      string `json:"datapath.mtu,omitempty"`
	DatapathOpenFlowSwitch           string `json:"datapath.openflow-switch,omitempty"`
	DatapathVlanID                   string `json:"datapath.vlan-id,omitempty"`
	DatapathVlanMode                 string `json:"datapath.vlan-mode,omitempty"`
	DisconnectTimeout                string `json:"disconnect-timeout,omitempty"`
	Distance                         string `json:"distance,omitempty"`
	FrameLifetime                    string `json:"frame-lifetime,omitempty"`
	GuardInterval                    string `json:"guard-interval,omitempty"`
	HideSsid                         string `json:"hide-ssid,omitempty"`
	HwProtectionMode                 string `json:"hw-protection-mode,omitempty"`
	HwRetries                        string `json:"hw-retries,omitempty"`
	Installation                     string `json:"installation,omitempty"`
	KeepAliveFrames                  string `json:"keepalive-frames,omitempty"`
	LoadBalancingGroup               string `json:"load-balancing-group,omitempty"`
	MaxStaCount                      string `json:"max-sta-count,omitempty"`
	Mode                             string `json:"mode,omitempty"`
	MulticastHelper                  string `json:"multicast-helper,omitempty"`
	Rates                            string `json:"rates,omitempty"`
	RatesBasic                       string `json:"rates.basic,omitempty"`
	RatesSupported                   string `json:"rates.supported,omitempty"`
	RatesHtBasicMcs                  string `json:"rates.ht-basic-mcs,omitempty"`
	RatesHtSupportedMcs              string `json:"rates.ht-supported-mcs,omitempty"`
	RatesVhtBasicMcs                 string `json:"rates.vht-basic-mcs,omitempty"`
	RatesVhtSupportedMcs             string `json:"rates.vht-supported-mcs,omitempty"`
	RxChains                         string `json:"rx-chains,omitempty"`
	Security                         string `json:"security,omitempty"`
	SecurityGroupEncryption          string `json:"security.group-encryption,omitempty"`
	SecurityAuthenticationTypes      string `json:"security.authentication-types,omitempty"`
	SecurityEapMethods               string `json:"security.eap-methods,omitempty"`
	SecurityEapRadiusAccounting      string `json:"security.eap-radius-accounting,omitempty"`
	SecurityEncryption               string `json:"security.encryption,omitempty"`
	SecurityGroupKeyUpdate           string `json:"security.group-key-update,omitempty"`
	SecurityPassphrase               string `json:"security.passphrase,omitempty"`
	SecurityTlsCertificate           string `json:"security.tls-certificate,omitempty"`
	SecurityTlsMode                  string `json:"security.tls-mode,omitempty"`
	Ssid                             string `json:"ssid,omitempty"`
	TxChains                         string `json:"tx-chains,omitempty"`
}

func (c *Client) CreateCapsManConfiguration(configuration *CapsManConfiguration) (*CapsManConfiguration, error) {
	reqBody, err := json.Marshal(configuration)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/caps-man/configuration", c.HostURL), bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, err
	}

	res := CapsManConfiguration{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadCapsManConfiguration(id string) (*CapsManConfiguration, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/caps-man/configuration/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	res := CapsManConfiguration{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateCapsManConfiguration(id string, configuration *CapsManConfiguration) (*CapsManConfiguration, error) {
	reqBody, err := json.Marshal(configuration)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/caps-man/configuration/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := CapsManConfiguration{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteCapsManConfiguration(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/caps-man/configuration/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	res := CapsManConfiguration{}

	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
