package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    ".id": "*1",
    "attributes": "Mikrotik-Wireless-Comment:Test Group,Mikrotik-Wireless-VLANID:100",
    "default": "false",
    "default-name": "",
    "inner-auths": "ttls-pap,ttls-chap,ttls-mschap1,ttls-mschap2,peap-mschap2",
    "name": "test",
    "outer-auths": "pap,chap,mschap1,mschap2,eap-tls,eap-ttls,eap-peap,eap-mschap2"
}
*/

// https://help.mikrotik.com/docs/display/ROS/User+Manager#UserManager-UserGroups
func ResourceUserManagerUserGroup() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/user-manager/user/group"),
		MetaId:           PropId(Id),

		"attributes": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "A custom set of colon-separated attributes with their values will be added to `Access-Accept` messages for users in this group.",
		},
		KeyDefault:     PropDefaultRo,
		KeyDefaultName: PropDefaultNameRo("The default name of the group."),
		"inner_auths": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ttls-pap", "ttls-chap", "ttls-mschap1", "ttls-mschap2", "peap-mschap2"}, false),
			},
			Description: "A set of allowed authentication methods for tunneled authentication methods (`ttls-pap`, `ttls-chap`, `ttls-mschap1`, `ttls-mschap2`, `peap-mschap2`).",
		},
		KeyName: PropName("Unique name of the group."),
		"outer_auths": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"pap", "chap", "mschap1", "mschap2", "eap-tls", "eap-ttls", "eap-peap", "eap-mschap2"}, false),
			},
			Description: "A set of allowed authentication methods (`pap`, `chap`, `mschap1`, `mschap2`, `eap-tls`, `eap-ttls`, `eap-peap`, `eap-mschap2`).",
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
