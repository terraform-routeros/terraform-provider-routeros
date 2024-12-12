package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*A",
    "default": "true",
    "dh-group": "modp2048,modp1024",
    "dpd-interval": "2m",
    "dpd-maximum-failures": "5",
    "enc-algorithm": "aes-128,3des",
    "hash-algorithm": "sha1",
    "lifetime": "1d",
    "name": "default",
    "nat-traversal": "true",
    "proposal-check": "obey"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/IPsec#IPsec-Profiles
func ResourceIpIpsecProfile() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/ipsec/profile"),
		MetaId:           PropId(Id),

		"dh_group": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Diffie-Hellman group (cipher strength).",
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"modp768", "modp1024 ", "modp1536", "modp2048",
					"modp3072", "modp4096", "modp6144", "modp8192", "ecp256", "ecp384", "ecp521"}, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dpd_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Dead peer detection interval. If set to disable-dpd, dead peer detection will not be used.",
			DiffSuppressFunc: TimeEquall,
		},
		"dpd_maximum_failures": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum count of failures until peer is considered to be dead. Applicable if DPD is enabled.",
			ValidateFunc:     validation.IntBetween(1, 100),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"enc_algorithm": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "List of encryption algorithms that will be used by the peer.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"3des", "aes-128", "aes-192", "aes-256", "blowfish",
					"camellia-128", "camellia-192", "camellia-256", "des"}, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"hash_algorithm": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Hashing algorithm. SHA (Secure Hash Algorithm) is stronger, but slower. MD5 uses 128-bit key, " +
				"sha1-160bit key.",
			ValidateFunc:     validation.StringInSlice([]string{"md5", "sha1", "sha256", "sha512"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"lifebytes": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Phase 1 lifebytes is used only as administrative value which is added to proposal. Used in " +
				"cases if remote peer requires specific lifebytes value to establish phase 1.",
		},
		"lifetime": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Phase 1 lifetime: specifies how long the SA will be valid.",
			DiffSuppressFunc: TimeEquall,
		},
		KeyName: PropName(""),
		"nat_traversal": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Use Linux NAT-T mechanism to solve IPsec incompatibility with NAT routers between IPsec peers. " +
				"This can only be used with ESP protocol (AH is not supported by design, as it signs the complete packet, " +
				"including the IP header, which is changed by NAT, rendering AH signature invalid). The method encapsulates " +
				"IPsec ESP traffic into UDP streams in order to overcome some minor issues that made ESP incompatible " +
				"with NAT.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"prf_algorithm": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "",
			ValidateFunc: validation.StringInSlice([]string{"auto", "sha1", "sha256", "sha384", "sha512"}, false),
		},
		"proposal_check": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Phase 2 lifetime check logic:\n  * claim - take shortest of proposed and configured lifetimes and " +
				"notify initiator about it\n  * exact - require lifetimes to be the same\n  * obey - accept whatever is sent by an " +
				"initiator\n  * strict - if the proposed lifetime is longer than the default then reject the proposal otherwise " +
				"accept a proposed lifetime.",
			ValidateFunc:     validation.StringInSlice([]string{"claim", "exact", "obey", "strict"}, false),
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
