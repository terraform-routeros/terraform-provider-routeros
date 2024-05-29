package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    ".about": "",
    "bridge": "bridge1",
    "caps-man-addresses": "",
    "caps-man-certificate-common-names": "",
    "caps-man-names": "",
    "certificate": "CAP-000000000000",
    "discovery-interfaces": "bridge1",
    "enabled": "true",
    "interfaces": "wlan1,wlan2",
    "lock-to-caps-man": "false",
    "requested-certificate": "CAP-000000000000",
    "static-virtual": "false"
}
*/

// https://help.mikrotik.com/docs/pages/viewpage.action?pageId=1409149#APController(CAPsMAN)-CAPtoCAPsMANConnection
func ResourceInterfaceWirelessCap() *schema.Resource {

	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wireless/cap"),
		MetaId:           PropId(Name),
		MetaSkipFields:   PropSkipFields(".about"),

		"bridge": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "none",
			Description: "Bridge interface to add the interface as a bridge port.",
		},
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
			Default:     "none",
			Description: "Certificate to use for authentication.",
		},
		"discovery_interfaces": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of interfaces over which CAP should attempt to discover CAPs Manager.",
		},
		KeyEnabled: PropEnabled("Disable or enable the CAP functionality."),
		"interfaces": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of interfaces managed by CAPs Manager.",
		},
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
		"static_virtual": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option that creates static virtual interfaces.",
		},
	}

	return &schema.Resource{
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
