package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
 {
	".id": "*1D",
	"comment": "comment",
    "client-address": "172.18.0.2",
    "disabled": "false",
    "encoding": "BF-128-CBC/SHA256",
    "mtu": "1500",
    "name": "ovpn-in1",
    "running": "true",
    "uptime": "1m25s",
    "user": "user1"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/???
func ResourceInterfaceOpenVPNServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ovpn-server"),
		MetaId:           PropId(Id),

		"client_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The address of the remote side.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"encoding": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Encryption characteristics.",
		},
		KeyMtu:     PropL2MtuRo,
		KeyName:    PropName("Interface name (Example: ovpn-in1)."),
		KeyRunning: PropRunningRo,
		"uptime": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Connection uptime.",
		},
		"user": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "User name used for authentication.",
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

		Schema: resSchema,
	}
}
