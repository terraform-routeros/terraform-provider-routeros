package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "RSA":"true",
    "bits":"2048",
    "key-owner":"admin",
    "key-type":"rsa",
    "user":"admin"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/User#User-SSHKeys
func ResourceUserSshKeys() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/user/ssh-keys"),
		MetaId:           PropId(Id),
		MetaTransformSet: PropTransformSet("rsa: RSA"),

		"user": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "username to which the SSH key is assigned.",
		},
		KeyComment: PropCommentRw,
		"key": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "key",
		},
		"rsa": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "key type is rsa",
		},
		"bits": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "key length",
		},
		"key_type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "key type",
		},
		"key_owner": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "SSH key owner",
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
