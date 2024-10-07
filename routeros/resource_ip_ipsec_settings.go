package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
REST JSON
*/

// https://help.mikrotik.com/docs/display/ROS/IPsec#IPsec-Settings
func ResourceIpIpsecSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/ipsec/settings"),
		MetaId:           PropId(Id),

		"accounting": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to send RADIUS accounting requests to a RADIUS server. Applicable if EAP Radius " +
				"(`auth-method=eap-radius`) or pre-shared key with XAuth authentication method " +
				"(`auth-method=pre-shared-key-xauth`) is used.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"interim_update": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The interval between each consecutive RADIUS accounting Interim update. Accounting must be " +
				"enabled.",
			DiffSuppressFunc: TimeEquall,
		},
		"xauth_use_radius": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to use Radius client for XAuth users or not. Property is only applicable to peers " +
				"using the IKEv1 exchange mode.",
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
