package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*0",
    "auth-algorithms": "sha1",
    "default": "true",
    "disabled": "false",
    "enc-algorithms": "aes-256-cbc,aes-192-cbc,aes-128-cbc",
    "lifetime": "30m",
    "name": "default",
    "pfs-group": "modp1024"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/IPsec#IPsec-Proposals
func ResourceIpIpsecProposal() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/ipsec/proposal"),
		MetaId:           PropId(Id),

		"auth_algorithms": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Allowed algorithms for authorization. SHA (Secure Hash Algorithm) is stronger but slower. " +
				"MD5 uses a 128-bit key, sha1-160bit key.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"md5", "null", "sha1", "sha256", "sha512"}, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDefault:  PropDefaultRo,
		KeyDisabled: PropDisabledRw,
		"enc_algorithms": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Allowed algorithms and key lengths to use for SAs.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"null", "des", "3des", "aes-128-cbc", "aes-128-cbc",
					"aes-128gcm", "aes-192-cbc", "aes-192-ctr", "aes-192-gcm", "aes-256-cbc", "aes-256-ctr", "aes-256-gcm",
					"blowfish", "camellia-128", "camellia-192", "camellia-256", "twofish"}, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"lifetime": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "How long to use SA before throwing it out.",
			DiffSuppressFunc: TimeEquall,
		},
		KeyName: PropName(""),
		"pfs_group": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The diffie-Helman group used for Perfect Forward Secrecy.",
			ValidateFunc: validation.StringInSlice([]string{"ecp256", "ecp384", "ecp521", "modp768", "modp1024",
				"modp1536", "modp2048", "modp3072", "modp4096", "modp6144", "modp8192", "none"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
