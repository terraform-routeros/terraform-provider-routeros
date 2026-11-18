package routeros

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func sshKeyDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	oldFields := strings.Fields(old)
	newFields := strings.Fields(new)
	if len(oldFields) >= 2 && len(newFields) >= 2 {
		return oldFields[0] == newFields[0] && oldFields[1] == newFields[1]
	}
	return old == new
}

/*
  {
    ".id": "*1",
    "RSA":"true",
    "bits":"2048",
	"fingerprint": "SHA256:ABCD..."
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
			Type:             schema.TypeString,
			Required:         true,
			ForceNew:         true,
			Sensitive:        true,
			DiffSuppressFunc: sshKeyDiffSuppress,
			Description:      "key",
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
			Description: "SSH key owner (RouterOS < 7.25)",
		},
		"info": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "SSH key comment/info (RouterOS >= 7.25)",
		},
		"fingerprint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "SSH key fingerprint",
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
