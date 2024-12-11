package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "allow-signal-out-of-range": "10s",
    "ap-tx-limit": "0",
    "authentication": "true",
    "client-tx-limit": "0",
    "disabled": "false",
    "forwarding": "true",
    "interface": "any",
    "mac-address": "00:00:00:00:00:00",
    "management-protection-key": "",
    "private-algo": "none",
    "private-key": "",
    "private-pre-shared-key": "",
    "signal-range": "-120..120",
    "time": "3h3m-5h,mon,tue,wed,thu,fri",
    "vlan-id": "1",
    "vlan-mode": "default"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/
func ResourceInterfaceWirelessAccessList() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wireless/access-list"),
		MetaId:           PropId(Id),

		"allow_signal_out_of_range": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Option which permits client's signal to be out of the range always or for some time interval.",
			DiffSuppressFunc: TimeEquall,
		},
		"ap_tx_limit": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Limit rate of data transmission to this client. Value 0 means no limit. Value is in bits per " +
				"second.",
		},
		"authentication": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "No - Client association will always fail.yes - Use authentication procedure that is specified " +
				"in the security-profile of the interface.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"client_tx_limit": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Ask client to limit rate of data transmission. Value 0 means no limit.This is a proprietary " +
				"extension that is supported by RouterOS clients.Value is in bits per second.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"forwarding": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "\n  * false - Client cannot send frames to other station that are connected to same access point." +
				"\n  *true - Client can send frames to other stations on the same access point.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"interface": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Rules with interface=any are used for any wireless interface and the `interface = all` defines interface-list `all` " +
				"name. To make rule that applies only to one wireless interface, specify that interface as a value of " +
				"this property.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyMacAddress: PropMacAddressRw("Rule matches client with the specified MAC address. Value 00:00:00:00:00:00 matches always.", false),
		"management_protection_key": {
			Type:        schema.TypeString,
			Description: "Management protection shared secret.",
			Optional:    true,
		},
		"private_algo": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Only for `WEP` modes.",
			ValidateFunc:     validation.StringInSlice([]string{"104bit-wep", "40bit-wep", "aes-ccm", "none", "tkip"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"private_key": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Only for `WEP` modes (HEX).",
		},
		"private_pre_shared_key": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Used in `WPA PSK` mode.",
		},
		"signal_range": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Rule matches if signal strength of the station is within the range.If signal strength of the " +
				"station will go out of the range that is specified in the rule, access point will disconnect that station.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"time": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Rule will match only during specified time.Station will be disconnected after specified time " +
				"ends. Both start and end time is expressed as time since midnight, 00:00. Rule will match only during " +
				"specified days of the week. Ex: \"3h3m-5h,mon,tue,wed,thu,fri\"",
		},
		"vlan_id": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "VLAN ID to use if doing VLAN tagging.",
			ValidateFunc:     validation.IntBetween(0, 4094),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"vlan_mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "VLAN tagging mode specifies if traffic coming from client should get tagged (and untagged when going to client).",
			ValidateFunc:     validation.StringInSlice([]string{"default", "no-tag", "use-service-tag", "use-tag"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
