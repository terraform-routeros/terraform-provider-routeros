package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    "caps-man-addresses": "192.168.88.1",
    "caps-man-certificate-common-names": "CAPsMAN-0000000",
    "caps-man-names": "router",
    "certificate": "request",
    "current-caps-man-address": "192.168.88.1",
    "current-caps-man-identity": "router",
    "discovery-interfaces": "lan",
    "enabled": "no",
    "lock-to-caps-man": "true",
    "slaves-datapath": "lan",
    "slaves-static": "true"
}
*/

// https://help.mikrotik.com/docs/display/ROS/WiFi#WiFi-CAPconfiguration
func ResourceWifiCap() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wifi/cap"),
		MetaId:           PropId(Name),

		"caps_man_addresses": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPAddress,
			},
			Description: "List of Manager IP addresses that CAP will attempt to contact during discovery.",
		},
		"caps_man_certificate_common_names": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of manager certificate common names that CAP will connect to.",
		},
		"caps_man_names": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "An ordered list of CAPs Manager names that the CAP will connect to.",
		},
		"certificate": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Certificate to use for authentication.",
		},
		"current_caps_man_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Currently used CAPsMAN address.",
		},
		"current_caps_man_identity": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Currently used CAPsMAN identity.",
		},
		"discovery_interfaces": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of interfaces over which CAP should attempt to discover CAPs Manager.",
		},
		KeyEnabled: PropEnabled("Disable or enable the CAP functionality."),
		"lock_to_caps_man": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Lock CAP to the first CAPsMAN it connects to.",
		},
		"locked_caps_man_common_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Common name of the CAPsMAN that the CAP is locked to.",
		},
		"requested_certificate": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Requested certificate.",
		},
		"slaves_datapath": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the bridge interface the CAP will be added to.",
		},
		"slaves_static": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option that creates static virtual interfaces.",
		},
	}

	return &schema.Resource{
		Description:   `*<span style="color:red">This resource requires a minimum version of RouterOS 7.13.</span>*`,
		CreateContext: DefaultSystemCreate(resSchema),
		ReadContext:   DefaultSystemRead(resSchema),
		UpdateContext: DefaultSystemUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
