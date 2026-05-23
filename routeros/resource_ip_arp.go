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
