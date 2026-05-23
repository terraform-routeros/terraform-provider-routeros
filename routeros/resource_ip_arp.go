package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceArp https://help.mikrotik.com/docs/spaces/ROS/pages/100892687/ARP
func ResourceArp() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		// This aligns with the ARP sub-menu - which is at "/ip arp".
		MetaResourcePath: PropResourcePath("/ip/arp"),
		MetaId:           PropId(Id),

		// TODO: Investigate which items in WinBox directly correspond to the API/CLI commands:
		//  * Enabled
		//  * Comment
		//  * IP Address
		//  * MAC Address
		//  * Interface
		//  * Bridge Port
		//  * Host Name
		//  * Published
		//  * Status (this may just be a meta field)

		// TODO: Investigate the CLI command parameters.
		"address": {
			Type:        schema.TypeString,
			Optional:    false,
			Description: "IP address to be mapped",
		},
		"interface": {
			Type:        schema.TypeString,
			Optional:    false,
			Description: "",
		},
		"mac_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The MAC address that the IP will be mapped to. Defaults to 00:00:00:00:00:00",
		},
		"published": {
			Type:        schema.TypeString,
			Optional:    false,
			Description: "Static proxy-arp entry for individual IP addresses. When an ARP query is received for the specific IP address, the device will respond with its own MAC address. No need to set proxy-arp on the interface itself for all the MAC addresses to be proxied. The interface will respond to an ARP request only when the device has an active route towards the destination.",
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),
		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		SchemaVersion: 1,

		Schema: resSchema,
	}
}
