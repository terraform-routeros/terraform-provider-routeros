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
		//  * Bridge Port
		//  * Host Name
		//  * Status (this may just be a meta field)

		// TODO: Investigate the CLI command parameters.
		"address": {
			Type:        schema.TypeString,
			Optional:    false,
			Description: "IP address to be mapped.",
		},
		"interface": {
			Type:        schema.TypeString,
			Optional:    false,
			Description: "Interface name the IP address is assigned to.",
		},
		"mac_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The MAC address that the IP will be mapped to. Defaults to 00:00:00:00:00:00.",
		},
		"published": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Static proxy-arp entry for individual IP addresses. When an ARP query is received for the specific IP address, the device will respond with its own MAC address. No need to set proxy-arp on the interface itself for all the MAC addresses to be proxied. The interface will respond to an ARP request only when the device has an active route towards the destination.",
		},
		// Read Only Properties
		"complete": {
			Type:        schema.TypeBool,
			Optional:    false,
			Description: "This flag is included in ARP entries when the ARP status is: permanent, reachable, stale, probe, or delay.",
		},
		"dhcp": {
			Type:        schema.TypeBool,
			Optional:    false,
			Description: "Whether the ARP entry is added by a DHCP server.",
		},
		"disabled": {
			Type:        schema.TypeBool,
			Optional:    false,
			Description: "Whether the ARP entry is disabled",
		},
		"dynamic": {
			Type:        schema.TypeBool,
			Optional:    false,
			Description: "Whether the entry is dynamically created",
		},
		"invalid": {
			Type:        schema.TypeBool,
			Optional:    false,
			Description: "Whether the entry is not valid",
		},
		"status": {
			Type:        schema.TypeString,
			Optional:    false,
			Description: "Shows the state of the ARP entry. This can be: delay, failed, incomplete, permanent, permanent, reachable, or stale.",
		},
		"VRF": {
			Type:        schema.TypeString,
			Optional:    false,
			Description: "\tIndicates which VRF this ARP entry is associated with.",
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
