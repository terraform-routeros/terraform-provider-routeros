package routeros

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
    ".id": "*1",
    "name": "read",
    "policy": "local,telnet,ssh,reboot,read,test,winbox,password,web,sniff,sensitive,api,romon,rest-api,!ftp,!write,!policy",
    "skin": "default"
}
*/

// https://help.mikrotik.com/docs/display/ROS/User#User-UserGroups
func ResourceUserGroup() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath:   PropResourcePath("/user/group"),
		MetaId:             PropId(Id),
		MetaSetUnsetFields: PropSetUnsetFields("policy"),

		KeyComment: PropCommentRw,
		KeyName:    PropName("The name of the user group"),
		"policy": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateDiagFunc: ValidationValInSlice([]string{
					"api", "dude", "ftp", "local", "password", "policy", "read", "reboot", "rest-api", "romon", "sensitive", "sniff",
					"ssh", "telnet", "test", "tikapp", "web", "winbox", "write",
				}, false, true),
			},
			Description: "A set of allowed policies.",
			DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
				// cty.SetVal([]cty.Value{cty.StringVal("!api"), cty.StringVal("!ftp"), ... })
				if len(oldValue) > 0 && oldValue[0] == '!' && newValue == "" {
					return true
				}
				if strings.HasSuffix(k, ".#") || d.GetRawConfig().GetAttr("policy").IsNull() {
					return true
				}
				return false
			},
		},
		"skin": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "default",
			Description: "The name of the skin that will be used for WebFig.",
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
