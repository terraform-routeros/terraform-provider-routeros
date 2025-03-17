package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "allowed-interface-list": "LAN"
}
*/

// https://help.mikrotik.com/docs/display/ROS/MAC+server
func ResourceToolMacServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/mac-server"),
		MetaId:           PropId(Id),

		"allowed_interface_list": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "List of interfaces for MAC Telnet access.",
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

// https://help.mikrotik.com/docs/display/ROS/MAC+server
func ResourceToolMacServerWinBox() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/mac-server/mac-winbox"),
		MetaId:           PropId(Id),

		"allowed_interface_list": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "List of interfaces for MAC WinBox access.",
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

// https://help.mikrotik.com/docs/spaces/ROS/pages/98795539/MAC+server#MACserver-MACPingServer
func ResourceToolMacServerPing() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/mac-server/ping"),
		MetaId:           PropId(Id),

		KeyEnabled: PropEnabled("Whether to enable the MAC Ping server."),
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
