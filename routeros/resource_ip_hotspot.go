package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*5",
    "HTTPS": "false",
    "addresses-per-mac": "unlimited",
    "disabled": "false",
    "idle-timeout": "5m",
    "interface": "ether4",
    "invalid": "false",
    "keepalive-timeout": "none",
    "login-timeout": "none",
    "name": "server1",
    "profile": "default",
    "proxy-status": "running"
  }
*/

// https://help.mikrotik.com/docs/pages/viewpage.action?pageId=56459266#HotSpot(Captiveportal)-IPHotSpot
func ResourceIpHotspot() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/hotspot"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("HTTPS", "keepalive_timeout", "proxy_status"),

		"address_pool": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Address space used to change HotSpot client any IP address to a valid address. Useful for " +
				"providing public network access to mobile clients that are not willing to change their networking settings.",
		},
		"addresses_per_mac": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Number of IP addresses allowed to be bind with the MAC address, when multiple HotSpot clients " +
				"connected with one MAC-address.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		"idle_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Period of inactivity for unauthorized clients. When there is no traffic from this client (literally " +
				"client computer should be switched off), once the timeout is reached, a user is dropped from the HotSpot " +
				"host list, its used address becomes available.",
			DiffSuppressFunc: TimeEquall,
		},
		"interface": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Interface to run HotSpot on.",
		},
		KeyInvalid: PropInvalidRo,
		"keepalive_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The exact value of the keepalive-timeout, that is applied to the user. Value shows how long " +
				"the host can stay out of reach to be removed from the HotSpot.",
			DiffSuppressFunc: TimeEquall,
		},
		"login_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Period of time after which if a host hasn't been authorized itself with a system the host " +
				"entry gets deleted from host table. Loop repeats until the host logs in the system. Enable if there " +
				"are situations where a host cannot log in after being too long in the host table unauthorized.",
			DiffSuppressFunc: TimeEquall,
		},
		KeyName: PropName("HotSpot server's name or identifier."),
		"profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "HotSpot server default HotSpot profile, which is located in `/ip/hotspot/profile`.",
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
