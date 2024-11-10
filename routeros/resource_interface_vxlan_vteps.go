package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*3",
    "interface": "vxlan1",
    "port": "8472",
    "remote-ip": "::"
  }
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/100007937/VXLAN#VXLAN-Forwardingtable
func ResourceInterfaceVxlanVteps() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/vxlan/vteps"),
		MetaId:           PropId(Id),

		KeyComment:   PropCommentRw,
		KeyInterface: PropInterfaceRw,
		"port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Used UDP port number.",
			ValidateFunc:     Validation64k,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"remote_ip": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "The IPv4 or IPv6 destination address of remote VTEP.",
			ValidateFunc: validation.IsIPAddress,
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
