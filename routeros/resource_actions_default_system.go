package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DefaultSystemCreate(s map[string]*schema.Schema) schema.CreateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return SystemResourceCreateUpdate(ctx, s, d, m)
	}
}

func DefaultSystemRead(s map[string]*schema.Schema) schema.ReadContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return SystemResourceRead(ctx, s, d, m)
	}
}

func DefaultSystemUpdate(s map[string]*schema.Schema) schema.UpdateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return SystemResourceCreateUpdate(ctx, s, d, m)
	}
}

func DefaultSystemDelete(s map[string]*schema.Schema) schema.DeleteContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return SystemResourceDelete(ctx, s, d, m)
	}
}

func DefaultSystemDatasourceRead(s map[string]*schema.Schema) schema.ReadContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		res := MikrotikItem{}
		path := s[MetaResourcePath].Default.(string)

		err := m.(Client).SendRequest(crudRead, &URL{Path: path}, nil, &res)
		if err != nil {
			return diag.FromErr(err)
		}

		return MikrotikResourceDataToTerraformDatasource(&[]MikrotikItem{res}, "", s, d)
	}
}
