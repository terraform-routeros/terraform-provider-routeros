package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*1",
  "address": "0.0.0.1",
  "bytes-in": "0",
  "bytes-out": "0",
  "disabled": "false",
  "dynamic": "false",
  "email": "mail@g.com",
  "limit-bytes-in": "100",
  "limit-bytes-out": "200",
  "limit-bytes-total": "500",
  "limit-uptime": "1m",
  "mac-address": "11:00:00:00:00:00",
  "name": "user1",
  "packets-in": "0",
  "packets-out": "0",
  "password": "123",
  "profile": "default",
  "routes": "10.0.0.0/24",
  "uptime": "0s"
}
*/

// https://wiki.mikrotik.com/wiki/Manual:IP/Hotspot/User#Users
func ResourceIpHotspotUser() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/hotspot/user"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("bytes_in", "bytes_out", "packets_in", "packets_out", "uptime"),

		"address": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "IP address, when specified client will get the address from the HotSpot one-to-one NAT translations. " +
				"Address does not restrict HotSpot login only from this address.",
		},
		KeyComment:  PropCommentRw,
		KeyDefault:  PropDefaultRo,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"email": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "HotSpot client's e-mail, informational value for the HotSpot user.",
		},
		"limit_bytes_in": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Maximal amount of bytes that can be received from the user. User is disconnected from HotSpot " +
				"after the limit is reached.",
		},
		"limit_bytes_out": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Maximal amount of bytes that can be transmitted from the user. User is disconnected from HotSpot " +
				"after the limit is reached.",
		},
		"limit_bytes_total": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "(limit-bytes-in+limit-bytes-out). User is disconnected from HotSpot after the limit is reached.",
		},
		"limit_uptime": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Uptime limit for the HotSpot client, user is disconnected from HotSpot as soon as uptime is " +
				"reached.",
		},
		"mac_address": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Client is allowed to login only from the specified MAC-address. If value is 00:00:00:00:00:00, " +
				"any mac address is allowed.",
		},
		KeyName: PropName("HotSpot login page username, when MAC-address authentication is used name is configured as " +
			"client's MAC-address."),
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "User password.",
			Sensitive:   true,
		},
		"profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "User profile configured in `/ip hotspot user profile`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"routes": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Routes added to HotSpot gateway when client is connected. The route format dst-address gateway " +
				"metric (for example, `192.168.1.0/24 192.168.0.1 1`).",
		},
		"server": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "HotSpot server's name to which user is allowed login.",
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
