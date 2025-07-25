package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {"max-block-size": "4096"}
*/

// ResourceIPTFTPSettings https://wiki.mikrotik.com/Manual:IP/TFTP
func ResourceIpTFTPSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		"max_block_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Maximum accepted block size value. During transfer " +
				"negotiation phase, RouterOS device will not negotiate larger value " +
				"than this.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
