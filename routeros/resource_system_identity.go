package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* {"name":"MikroTik"} */

func ResourceSystemIdentity() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/identity"),
		MetaId:           PropId(Name),

		KeyName: PropNameRw,
	}

	resRead := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		metadata := GetMetadata(resSchema)

		res := MikrotikItem{}
		err := m.(Client).SendRequest(crudRead, &URL{Path: metadata.Path}, nil, &res)
		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId(d.Get("name").(string))

		return MikrotikResourceDataToTerraform(res, resSchema, d)
	}

	resUpdate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		item, metadata := TerraformResourceDataToMikrotik(resSchema, d)

		var resUrl string
		if m.(Client).GetTransport() == TransportREST {
			// https://router/rest/system/identity/set
			resUrl = "/set"
		}

		err := m.(Client).SendRequest(crudPost, &URL{Path: metadata.Path + resUrl}, item, nil)
		if err != nil {
			return diag.FromErr(err)
		}

		return resRead(ctx, d, m)
	}

	return &schema.Resource{
		CreateContext: resUpdate,
		ReadContext:   resRead,
		UpdateContext: resUpdate,
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			// No delete functionality provided by API for System Identity.
			// Delete function will remove the object from the Terraform state
			d.SetId("")
			return DeleteSystemObject
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
