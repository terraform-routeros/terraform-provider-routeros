package client

import (
	"fmt"
	"net/http"
)

type ROSInterface struct {
	ID               string `json:".id"`
	ActualMtu        string `json:"actual-mtu"`
	DefaultName      string `json:"default-name"`
	Disabled         string `json:"disabled"`
	FpRxByte         string `json:"fp-rx-byte"`
	FpRxPacket       string `json:"fp-rx-packet"`
	FpTxByte         string `json:"fp-tx-byte"`
	FpTxPacket       string `json:"fp-tx-packet"`
	L2Mtu            string `json:"l2mtu"`
	LastLinkDownTime string `json:"last-link-down-time"`
	LastLinkUpTime   string `json:"last-link-up-time"`
	LinkDowns        string `json:"link-downs"`
	MacAddress       string `json:"mac-address"`
	MaxL2Mtu         string `json:"max-l2mtu"`
	Mtu              string `json:"mtu"`
	Name             string `json:"name"`
	Running          string `json:"running"`
	RxByte           string `json:"rx-byte"`
	RxPacket         string `json:"rx-packet"`
	Slave            string `json:"slave"`
	TxByte           string `json:"tx-byte"`
	TxPacket         string `json:"tx-packet"`
	TxQueueDrop      string `json:"tx-queue-drop"`
	Type             string `json:"type"`
}

func (c *Client) ReadInterfaces() ([]ROSInterface, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/interface", c.HostURL), nil)
	if err != nil {
		return nil, err
	}
	res := []ROSInterface{}
	if c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return res, nil
}
