package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceUPNPInterfaces https://help.mikrotik.com/docs/display/ROS/UPnP
func ResourceUPNPInterfaces() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath:   PropResourcePath("/ip/upnp/interfaces"),
		MetaId:             PropId(Id),
		MetaSetUnsetFields: PropSetUnsetFields("type"),

		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"forced_ip": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Allow specifying what public IP to use if the external interface has more than one IP available.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"interface": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Interface name on which uPnP will be running.",
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "UPnP interface type:" +
				"\n  * external - the interface a global IP address is assigned to" +
				"\n  * internal - router's local interface the clients are connected to",
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
