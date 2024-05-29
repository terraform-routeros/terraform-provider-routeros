package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "auth-key": "none",
  "broadcast": "false",
  "broadcast-addresses": "",
  "enabled": "false",
  "local-clock-stratum": "5",
  "manycast": "false",
  "multicast": "false",
  "use-local-clock": "false",
  "vrf": "main"
}
*/

// https://help.mikrotik.com/docs/display/ROS/NTP#NTP-NTPServersettings:.1
func ResourceSystemNtpServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/ntp/server"),
		MetaId:           PropId(Id),

		"auth_key": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "NTP symmetric key, used for authentication between the NTP client and server. Key Identifier " +
				"(Key ID) - an integer identifying the cryptographic key used to generate the message-authentication code.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"broadcast": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enable certain NTP server mode, for this mode to work you have to set up broadcast-addresses field.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"broadcast_addresses": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Set broadcast address to use for NTP server broadcast mode.",
		},
		KeyEnabled: PropEnabled("Enable NTP server."),
		"local_clock_stratum": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Manually set stratum if ```use_local_clock = true```.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.IntBetween(1, 16),
		},
		"manycast": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enable certain NTP server mode.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"multicast": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enable certain NTP server mode.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyVrf: PropVrfRw,
		"use_local_clock": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "The server will supply its local system time as valid if others are not available.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
