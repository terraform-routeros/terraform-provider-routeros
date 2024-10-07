package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*2",
    "default": "true",
    "name": "default"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/IPsec#IPsec-Groups
func ResourceIpIpsecPolicyGroup() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/ipsec/policy/group"),
		MetaId:           PropId(Id),

		KeyDefault: PropDefaultRo,
		KeyName:    PropName(""),
		KeyComment: PropCommentRw,
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
