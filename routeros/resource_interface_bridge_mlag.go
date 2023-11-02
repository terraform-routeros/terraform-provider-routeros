package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
	"bridge":"bridge1",
	"peer-port":"stack-link"
}
*/

// https://help.mikrotik.com/docs/display/ROS/Multi-chassis+Link+Aggregation+Group
func ResourceInterfaceBridgeMlag() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/bridge/mlag"),
		MetaId:           PropId(Id),

		"bridge": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The bridge interface where MLAG is being created.",
		},
		"peer_port": {
			Type:     schema.TypeString,
			Required: true,
			Description: "An interface that will be used as a peer port. Both peer devices are using inter-chassis " +
				"communication over these peer ports to establish MLAG and update the host table. Peer port should be " +
				"isolated on a different untagged VLAN using a pvid setting. Peer port can be configured as a bonding " +
				"interface.",
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
