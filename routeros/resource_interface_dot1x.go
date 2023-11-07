package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// https://help.mikrotik.com/docs/display/ROS/Dot1X#Dot1X-Client
func ResourceInterfaceDot1xClient() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/dot1x/client"),
		MetaId:           PropId(Id),

		"anon_identity": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Identity for outer layer EAP authentication. Used only with `eap-ttls` and `eap-peap` methods. If not set, the value from the identity parameter will be used for outer layer EAP authentication.",
		},
		"certificate": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "none",
			Description: "Name of a certificate. Required when the `eap-tls` method is used.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"eap_methods": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "A list of EAP methods used for authentication: `eap-tls`, `eap-ttls`, `eap-peap`, `eap-mschapv2`.",
		},
		"identity": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The supplicant identity that is used for EAP authentication.",
		},
		KeyInterface: PropInterfaceRw,
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Cleartext password for the supplicant.",
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext: DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
