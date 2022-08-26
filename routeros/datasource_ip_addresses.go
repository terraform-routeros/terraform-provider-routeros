package routeros

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceIPAddresses() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceIPAddressesRead,
		Schema: map[string]*schema.Schema{
			"addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						MetaResourcePath: PropResourcePath("/ip/address"),
						MetaId:           PropId(Id),

						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"actual_interface": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"disabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"dynamic": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						KeyInterface: PropInterfaceRw,
						"invalid": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"network": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceIPAddressesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	s := DatasourceIPAddresses().Schema
	path := s["addresses"].Elem.(*schema.Resource).Schema[MetaResourcePath].Default.(string)

	res, err := ReadItems(nil, path, m.(Client))
	if err != nil {
		return diag.FromErr(err)
	}

	return MikrotikResourceDataToTerraformDatasource(res, "addresses", s, d)
}
