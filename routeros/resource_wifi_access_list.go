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
    "action": "accept",
    "allow-signal-out-of-range": "always",
    "client-isolation": "true",
    "disabled": "false",
    "interface": "2ghz",
    "mac-address": "00:00:00:00:00:00",
    "mac-address-mask": "00:00:00:00:00:00",
    "passphrase": "password",
    "radius-accounting": "true",
    "signal-range": "-120..-85",
    "ssid-regexp": "something",
    "time": "6m-20m,sun,sat",
    "vlan-id": "none"
}
*/

// https://help.mikrotik.com/docs/display/ROS/WiFi#WiFi-AccessList.1
func ResourceWifiAccessList() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wifi/access-list"),
		MetaId:           PropId(Id),

		"action": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "accept",
			Description:  "An action to take when a client matches.",
			ValidateFunc: validation.StringInSlice([]string{"accept", "reject", "query-radius"}, false),
		},
		"allow_signal_out_of_range": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "An option that permits the client's signal to be out of the range always or for some time interval.",
		},
		"client_isolation": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option that specifies whether to deny forwarding data between clients connected to the same interface.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"interface": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Interface name to compare with an interface to which the client actually connects to.",
		},
		"last_logged_in": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Last time this client logged in.",
		},
		"last_logged_out": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Last time this client logged out.",
		},
		"mac_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "MAC address of the client.",
			ValidateFunc: ValidationMacAddress,
		},
		"mac_address_mask": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "MAC address mask to apply when comparing clients' addresses.",
		},
		"match_count": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Number of times this entry was matched.",
		},
		KeyPlaceBefore: PropPlaceBefore,
		"passphrase": {
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
			Description: "Time of the day and days of the week when the rule is applicable.",
		},
		KeyVlanId: PropVlanIdRw("VLAN ID to use for VLAN tagging or `none`.", false),
	}

	return &schema.Resource{
		Description:   `*<span style="color:red">This resource requires a minimum version of RouterOS 7.13.</span>*`,
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
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
