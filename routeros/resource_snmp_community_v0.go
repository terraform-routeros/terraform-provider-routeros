package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*2",
    "addresses": "::/0",
    "authentication-password": "",
    "authentication-protocol": "MD5",
    "comment": "Comment",
    "default": "false",
    "disabled": "true",
    "encryption-password": "",
    "encryption-protocol": "DES",
    "name": "private",
    "read-access": "true",
    "security": "none",
    "write-access": "false"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/SNMP#SNMP-CommunityProperties
func ResourceSNMPCommunityV0() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/snmp/community"),
		MetaId:           PropId(Id),

		"addresses": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Set of IP (v4 or v6) addresses or CIDR networks from which connections to SNMP server are allowed.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"authentication_password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Password used to authenticate the connection to the server (SNMPv3).",
		},
		"authentication_protocol": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "MD5",
			Description:  "The protocol used for authentication (SNMPv3).",
			ValidateFunc: validation.StringInSlice([]string{"MD5", "SHA1"}, false),
		},
		KeyComment: PropCommentRw,
		"default": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "It's a default community.",
		},
		KeyDisabled: PropDisabledRw,
		"encryption_password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "The password used for encryption (SNMPv3).",
		},
		"encryption_protocol": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "DES",
			Description: "encryption protocol to be used to encrypt the communication (SNMPv3). AES (see rfc3826) " +
				"available since v6.16.",
			ValidateFunc: validation.StringInSlice([]string{"DES", "AES"}, false),
		},

		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Community Name.",
		},
		"read_access": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Whether read access is enabled for this community.",
		},
		"security": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "none",
			Description:  "Security features.",
			ValidateFunc: validation.StringInSlice([]string{"authorized", "none", "private"}, false),
		},
		"write_access": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether write access is enabled for this community.",
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
