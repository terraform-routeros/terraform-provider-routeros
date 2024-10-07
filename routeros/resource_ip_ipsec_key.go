package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "key-size": "1024", <<<< !!! /ip/ipsec/key/generate-key name=new-key key-size=   2048     4096     8192
    "name": "new-key",
    "private-key": "true",
    "rsa": "true"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/IPsec#IPsec-Keys
func ResourceIpIpsecKey() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/ipsec/key"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("private_key", "rsa"),

		"key_size": {
			Type:             schema.TypeInt,
			Required:         true,
			ForceNew:         true,
			Description:      "Size of this key.",
			ValidateFunc:     validation.IntInSlice([]int{1024, 2048, 4096}),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName(""),
	}

	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return ResourceCreateAndWait(ctxSetCrudMethod(ctx, crudGenerateKey), resSchema, d, m, d.Timeout(schema.TimeoutCreate))
		},
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
