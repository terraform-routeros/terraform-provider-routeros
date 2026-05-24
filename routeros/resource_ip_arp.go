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
		"address": {
			Type:        schema.TypeString,
			Optional:    false,
			Description: "IP address to be mapped.",
		},
		// TODO: Switch to KeyInterface/PropInterfaceRw
		"interface": {
			Type:        schema.TypeString,
			Optional:    false,
			Description: "Interface name the IP address is assigned to.",
		},
		// TODO: Switch to KeyMacAddress/PropMacAddressRW
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
		// Read-only properties
		"complete": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "This flag is included in ARP entries when the ARP status is: permanent, reachable, stale, probe, or delay.",
		},
		// TODO: If an ARP entry was added by the DHCP server, then it shouldn't exist within Terraform - since this could result in a conflict. If this is the case, should we error in some way?
		"dhcp": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the ARP entry is added by a DHCP server.",
		},
		KeyDisabled: PropDisabledRo,
		KeyDynamic: {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the entry is dynamically created",
		},
		KeyInvalid: {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the entry is not valid",
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Shows the state of the ARP entry. This can be: delay, failed, incomplete, permanent, permanent, reachable, or stale.",
		},
		// NOTE: The documentation shows this as capitalised - so I opted against using the KeyVrf constant (which is lowercase).
		"VRF": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Indicates which VRF this ARP entry is associated with.",
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
