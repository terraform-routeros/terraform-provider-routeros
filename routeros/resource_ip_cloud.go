package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "ddns-enabled": "true",
  "ddns-update-interval": "1m",
  "dns-name": "ad8a0be701ea.sn.mynetname.net",
  "public-address": "31.173.86.120",
  "status": "updated",
  "update-time": "true",
  "warning": "Router is behind a NAT. Remote connection might not work."
}
*/

// https://wiki.mikrotik.com/wiki/Manual:IP/Cloud
func ResourceIpCloud() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/cloud"),
		MetaId:           PropId(Id),

		"ddns_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "If set to yes, then the device will send an encrypted message to the MikroTik's Cloud " +
				"server. The server will then decrypt the message and verify that the sender is an " +
				"authentic MikroTik device. If all is OK, then the MikroTik's Cloud server will create a " +
				"DDNS record for this device and send a response to the device. Every minute the IP/Cloud " +
				"service on the router will check if WAN IP address matches the one sent to MikroTik's " +
				"Cloud server and will send encrypted update to cloud server if IP address changes.",
		},
		"ddns_update_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "none",
			Description: "If set DDNS will attempt to connect IP Cloud servers at the set interval. If set to none " +
				"it will continue to internally check IP address update and connect to IP Cloud servers " +
				"as needed. Useful if IP address used is not on the router itself and thus, cannot be " +
				"checked as a value internal to the router.",
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == new {
					return true
				}

				if old == "none" || new == "none" {
					return false
				}

				return TimeEquall(k, old, new, d)
			},
		},
		"dns_name": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Shows DNS name assigned to the rdevice. Name consists of 12 character serial number " +
				"appended by .sn.mynetname.net. This field is visible only after at least one " +
				"ddns-request is successfully completed.",
		},
		"public_address": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Shows device's IPv4 address that was sent to cloud server. This field is visible only " +
				"after at least one IP Cloud request was successfully completed.",
		},
		"public_address_ivp6": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Shows device's IPv6 address that was sent to cloud server. This field is visible only " +
				"after at least one IP Cloud request was successfully completed.",
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Contains text string that describes current dns-service state. The messages are self " +
				"explanatory  updating... updated Error: no Internet connection Error: request timed out " +
				"Error: REJECTED. Contact MikroTik support Error: internal error - should not happen. One " +
				"possible cause is if router runs out of memory.",
		},
		"update_time": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
			Description: "If set to yes then router clock will be set to time, provided by cloud server IF there " +
				"is no NTP or SNTP client enabled. If set to no, then IP/Cloud service will never update " +
				"the device's clock. If update-time is set to yes, Clock will be updated even when " +
				"ddns-enabled is set to no.",
		},
		"warning": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Shows a warning message if IP address sent by the device differs from the IP address in " +
				"UDP packet header as visible by the MikroTik's Cloud server. Typically this happens if " +
				"the device is behind NAT. Example: 'DDNS server received request from IP 123.123.123.123 " +
				"but your local IP was 192.168.88.23; DDNS service might not work'",
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
