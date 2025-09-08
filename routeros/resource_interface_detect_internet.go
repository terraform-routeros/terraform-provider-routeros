package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "detect-interface-list": "none",
  "internet-interface-list": "none",
  "lan-interface-list": "none",
  "wan-interface-list": "none"
}
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/8323187/Detect+Internet
func ResourceInterfaceDetectInternet() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/detect-internet"),
		MetaId:           PropId(Id),

		"detect_interface_list": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "All interfaces in the list will be monitored by Detect Internet.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"internet_interface_list": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Interfaces with state Internet will be dynamically added to this list.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"lan_interface_list": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Interfaces with state Lan will be dynamically added to this list.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"wan_interface_list": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Interfaces with state Wan will be dynamically added to this list.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
