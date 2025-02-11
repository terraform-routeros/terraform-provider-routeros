package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "allow-signal-out-of-range": "10s",
    "comment": "Laptop",
    "disabled": "false",
    "mac-address": "00:00:00:00:00:00",
    "signal-range": "-120..120",
    "ssid-regexp": "",
    "time": "0s-1d,sun,mon,tue,wed,thu,fri,sat",
    "vlan-id": "1",
    "vlan-mode": "use-tag"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/CAPsMAN
func ResourceCapsManAccessList() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/caps-man/access-list"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"action": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "An action to take when a client matches.",
			ValidateFunc: validation.StringInSlice([]string{"accept", "reject", "query-radius"}, false),
		},
		"allow_signal_out_of_range": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "10s",
			Description:      "An option that permits the client's signal to be out of the range always or for some time interval.",
			DiffSuppressFunc: TimeEqual,
		},
		"ap_tx_limit": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Transmission speed limit in the direction of the client..",
		},
		"client_to_client_forwarding": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option that specifies whether to allow forwarding data between clients connected to the same interface.",
		},
		"client_tx_limit": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Transmission speed limit in the direction of the access point.",
		},
		"mac_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "MAC address of the client.",
		},
		"mac_mask": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "MAC address mask to apply when comparing clients' addresses.",
		},
		"interface": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Interface name to compare with an interface to which the client actually connects to.",
		},
		KeyPlaceBefore: PropPlaceBefore,
		"private_passphrase": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "PSK passphrase for the client if some PSK authentication algorithm is used.",
		},
		"radius_accounting": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option that specifies if RADIUS traffic accounting should be used in case of RADIUS authentication of the client.",
		},
		"signal_range": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "-120..120",
			Description: "The range in which the client signal must fall.",
		},
		"ssid_regexp": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The regular expression to compare the actual SSID the client connects to.",
		},
		"time": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "0s-1d,sun,mon,tue,wed,thu,fri,sat",
			Description: "Time of the day and days of the week when the rule is applicable.",
		},
		"vlan_id": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "VLAN ID to use if vlan-mode enables use of VLAN tagging.",
			ValidateFunc: validation.IntBetween(1, 4094),
		},
		"vlan_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "VLAN tagging mode specifies if traffic coming from a client should get tagged and untagged when it goes back to the client.",
			ValidateFunc: validation.StringInSlice([]string{"no-tag", "use-service-tag", "use-tag"}, false),
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			resSchema[MetaSkipFields].Default = `"place_before"`
			defer func() {
				resSchema[MetaSkipFields].Default = ``
			}()

			return ResourceUpdate(ctx, resSchema, d, m)
		},
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
