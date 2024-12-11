package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "enabled": "true",
  "freq-drift": "3.109",
  "mode": "unicast",
  "servers": "168.119.4.163",
  "status": "synchronized",
  "synced-server": "168.119.4.163",
  "synced-stratum": "1",
  "system-offset": "1.703",
  "vrf": "main"
}
*/

// https://help.mikrotik.com/docs/display/ROS/NTP#NTP-NTPClientproperties:
func ResourceSystemNtpClient() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/ntp/client"),
		MetaId:           PropId(Id),

		KeyEnabled: PropEnabled("Enable NTP client."),
		"freq_drift": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The fractional frequency drift per unit time.",
		},
		"mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Mode that the NTP client will operate in",
			ValidateFunc:     validation.StringInSlice([]string{"broadcast", "manycast", "multicast", "unicast"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"servers": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The list of NTP servers. It is possible to add static entries." +
				"The following formats are accepted:" +
				"\n  * FQDN (\"Resolved Address\" will appear in the \"Servers\"- window in an appropriate column if the address is " +
				"resolved) or IP address can be used. If DHCP-Client property `use-peer-ntp=yes` - the dynamic entries " +
				"advertised by DHCP" +
				"\n  * ipv4" +
				"\n  * ipv4@vrf" +
				"\n  * ipv6" +
				"\n  * ipv6@vrf" +
				"\n  * ipv6-linklocal%interface",

			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Current status of the NTP client.",
		},
		"synced_server": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The IP address of the NTP Server.",
		},
		"synced_stratum": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "The accuracy of each server is defined by a number called the stratum, with the topmost level " +
				"(primary servers) assigned as one and each level downwards (secondary servers) in the hierarchy " +
				"assigned as one greater than the preceding level.",
		},
		"system_offset": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "This is a signed, fixed-point number indicating the offset of the NTP server's clock " +
				"relative to the local clock, in seconds.",
		},
		KeyVrf: PropVrfRw,
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
